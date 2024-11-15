package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	maxConcurrentTransactions = 10
	batchSize                 = 5
	paymentProcessingTime     = time.Duration(rand.Intn(500)) * time.Millisecond
)

type Payment struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}

type TransactionProcessor struct {
	semaphore chan struct{}
	db        *PaymentDB
}

func NewTransactionProcessor(db *PaymentDB) *TransactionProcessor {
	semaphore := make(chan struct{}, maxConcurrentTransactions)
	return &TransactionProcessor{
		semaphore: semaphore,
		db:        db,
	}
}

func (tp *TransactionProcessor) ProcessTransactions(ctx context.Context, transactions []Payment) {
	for i := 0; i < len(transactions); i += batchSize {
		batch := transactions[i : i+batchSize]
		tp.processBatch(ctx, batch)
	}
}

func (tp *TransactionProcessor) processBatch(ctx context.Context, batch []Payment) {
	var wg sync.WaitGroup

	for _, transaction := range batch {
		wg.Add(1)
		go func(t Payment) {
			defer wg.Done()
			tp.processTransaction(ctx, t)
		}(transaction)
	}

	wg.Wait()
}

func (tp *TransactionProcessor) processTransaction(ctx context.Context, transaction Payment) {
	<-tp.semaphore
	defer func() {
		tp.semaphore <- struct{}{}
	}()

	time.Sleep(paymentProcessingTime)

	if err := tp.db.SavePayment(transaction); err != nil {
		fmt.Println("Error saving payment:", err)
	}

	fmt.Println("Processed transaction:", transaction.ID)
}

type PaymentDB struct {
	payments map[string]Payment
	mu       sync.Mutex
}

func NewPaymentDB() *PaymentDB {
	return &PaymentDB{
		payments: make(map[string]Payment),
	}
}

func (db *PaymentDB) SavePayment(payment Payment) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.payments[payment.ID] = payment
	return nil
}

func main() {
	ctx := context.Background()
	db := NewPaymentDB()
	processor := NewTransactionProcessor(db)

	payments := generatePayments(100)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		processor.ProcessTransactions(ctx, payments)
	}()

	wg.Wait()

	fmt.Println("Processed all payments. Total payments:", len(db.payments))
}

func generatePayments(count int) []Payment {
	payments := make([]Payment, count)
	for i := 0; i < count; i++ {
		payments[i] = Payment{
			ID:     fmt.Sprintf("payment_%d", i+1),
			Amount: rand.Float64() * 100.0,
		}
	}
	return payments
}
