package thermoprinter

import "log"

type Printer struct {
	// symbols per second.
	BaudRate int64
	// the number of microseconds it takes to issue one byte to the printer.
	// there are 11 bits in this byte (rather than 8) because there are the
	// idle, start, and stop bits for the serial protocol.
	ByteTime int64
}

type PrinterOptions struct {
	BaudRate int64
}

func NewPrinter(options ...PrinterOptions) *Printer {
	if len(options) > 1 {
		log.Fatal("Printer can only accept one set of options!")
	}

	printer := &Printer{
		BaudRate: DefaultBaudRate,
	}

	// apply overrides if options are provided
	if len(options) == 1 {
		override := options[0]
		printer.BaudRate = override.BaudRate
	}

	return printer
}
