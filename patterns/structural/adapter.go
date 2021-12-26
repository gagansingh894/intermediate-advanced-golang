package structural

import "fmt"

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct {}

func (l *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg =fmt.Sprintf("Legacy Printer: %s\n", s)
	fmt.Println(newMsg)
	return
}

type NewPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg string
}

func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
		fmt.Println("Not using legacy")
	}
	return
}


func AdapterExample() {
	msg := "Hi, I am printing"
	adapter := PrinterAdapter{
		OldPrinter: &MyLegacyPrinter{},
		Msg:        msg,
	}
	adapter.PrintStored()
	adapter2 := PrinterAdapter{
		OldPrinter: nil,
		Msg:        msg,
	}
	adapter2.PrintStored()
}