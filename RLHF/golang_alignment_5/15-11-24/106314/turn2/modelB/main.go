package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxConcurrentTransactions = 10
	transactionBatchSize      = 5
	transactionInterval       = 10 * time.Millisecond
)

type Payment struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}

type PaymentProcessor struct {
	wg    sync.WaitGroup
	sem   chan struct{}
	queue []Payment
	db    *mockDatabase
}

func NewPaymentProcessor(db *mockDatabase) *PaymentProcessor {
	sem := make(chan struct{}, maxConcurrentTransactions)
	return &PaymentProcessor{
		wg:    sync.WaitGroup{},
		sem:   sem,
		queue: make([]Payment, 0),
		db:    db,
	}
}

func (pp *PaymentProcessor) Start() {
	go pp.processQueue()
}

func (pp *PaymentProcessor) Stop() {
	close(pp.queue)
	pp.wg.Wait()
}

func (pp *PaymentProcessor) AddPayment(payment Payment) {
	pp.queue = append(pp.queue, payment)
	if len(pp.queue) >= transactionBatchSize {
		pp.processBatch()
	}
}

func (pp *PaymentProcessor) processBatch() {
	batch := pp.queue
	pp.queue = pp.queue[:0]

	pp.wg.Add(len(batch))
	for _, payment := range batch {
		go pp.processPayment(payment)
	}
}

func (pp *PaymentProcessor) processPayment(payment Payment) {
	<-pp.sem
	defer func() {
		pp.sem <- struct{}{}
		pp.wg.Done()
	}()

	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	if err := pp.db.InsertPayment(payment); err != nil {
		fmt.Println("Error inserting payment:", err)
		return
	}

	fmt.Println("Payment processed:", payment.ID)
}

func (pp *PaymentProcessor) processQueue() {
	for payment := range pp.queue {
		pp.processPayment(payment)
	}
}

type mockDatabase struct{}

func (db *mockDatabase) InsertPayment(payment Payment) error {
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	fmt.Println("Payment inserted:", payment.ID)
	return nil
}

func main() {
	ctx := context.Background()

	db := &mockDatabase{}
	pp := NewPaymentProcessor(db)
	pp.Start()

	defer pp.Stop()

	ticker := time.NewTicker(transactionInterval)
	for range ticker.C {
		payment := Payment{ID: fmt.Sprintf("payment-%d", atomic.AddUint64(&counter, 1)), Amount: 100.0}
		pp.AddPayment(payment)
	}

	ticker.Stop()
}

var counter uint64
