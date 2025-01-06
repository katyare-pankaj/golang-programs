package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gosuri/uiprogress"
	"github.com/muesli/reflow/format"
	"github.com/xuri/excelize"
)

const numLoops = 10000000

func main() {
	// Comparing fmt.Sprintf with some popular third-party libraries
	t1 := time.Now()
	formatStr := "Hello, my name is %s and I am %d years old."
	fmt.Sprintf(formatStr, "Alice", 30)
	t2 := time.Now()
	durationFmt := t2.Sub(t1)

	t3 := time.Now()
	for i := 0; i < numLoops; i++ {
		format.Sprintf(formatStr, "Bob", i%100)
	}
	t4 := time.Now()
	durationFormatter := t4.Sub(t3)

	t5 := time.Now()
	for i := 0; i < numLoops; i++ {
		strconv.Itoa(i)
	}
	t6 := time.Now()
	durationStrconv := t6.Sub(t5)

	t7 := time.Now()
	wb := excelize.NewFile()
	_ = wb.SetCellValue("Sheet1", "A1", "Formatted String")
	for i := 0; i < numLoops; i++ {
		wb.SetCellValue("Sheet1", "A2", fmt.Sprintf(formatStr, "Charlie", i))
	}
	_ = wb.SaveAs("book.xlsx")
	t8 := time.Now()
	durationExcelize := t8.Sub(t7)

	// Progress bar to show the comparison
	p := uiprogress.New()
	p.Start()
	bar1 := p.AddBar(int(durationFmt.Seconds())).AppendCompleted().PrependFunc(func(b *uiprogress.Bar) string {
		return fmt.Sprintf("fmt.Sprintf:   %s", b.String())
	})
	bar1.Incr()

	bar2 := p.AddBar(int(durationFormatter.Seconds())).AppendCompleted().PrependFunc(func(b *uiprogress.Bar) string {
		return fmt.Sprintf("muesli/reflow: %s", b.String())
	})
	bar2.Incr()

	bar3 := p.AddBar(int(durationStrconv.Seconds())).AppendCompleted().PrependFunc(func(b *uiprogress.Bar) string {
		return fmt.Sprintf("strconv:      %s", b.String())
	})
	bar3.Incr()

	bar4 := p.AddBar(int(durationExcelize.Seconds())).AppendCompleted().PrependFunc(func(b *uiprogress.Bar) string {
		return fmt.Sprintf("excelize:      %s", b.String())
	})
	bar4.Incr()
	p.Wait()

	fmt.Printf("\n")
	fmt.Printf("Elapsed Time Comparison:\n")
	fmt.Printf("-----------------------\n")
	fmt.Printf("fmt.Sprintf:   \t%v\n", durationFmt)
	fmt.Printf("muesli/reflow: \t%v\n", durationFormatter)
	fmt.Printf("strconv:       \t%v\n", durationStrconv)
	fmt.Printf("excelize:       \t%v\n", durationExcelize)
}
