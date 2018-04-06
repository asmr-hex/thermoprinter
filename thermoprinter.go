package thermoprinter

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

type Printer struct {
	// symbols per second. Defaults to 19200.
	BaudRate int
	// the number of microseconds it takes to issue one byte to the printer.
	// since the printer depends on physical components, we need to wait a
	// reasonable time after each write is issued s.t. the printer can physically
	// carry out the operations. The default is 300 milliseconds.
	WaitTimeAfterWrite time.Time
	// serial port device name. On Raspberry Pi 3, this is typically /dev/serial1
	// but on older models it could be /dev/ttyS0 or /dev/AMA0
	SerialPortName string
	// serial port used for UART comminucation, acts as a stream for reading
	// and writing bytes.
	Stream *serial.Port
}

type PrinterOptions struct {
	BaudRate       int
	SerialPortName string
}

func NewPrinter(options ...PrinterOptions) *Printer {
	var (
		err error
	)

	if len(options) > 1 {
		log.Fatal("Printer can only accept one set of options!")
	}

	printer := &Printer{
		BaudRate:       DefaultBaudRate,
		SerialPortName: DefaultSerialPortName,
	}

	// apply overrides if options are provided
	if len(options) == 1 {
		override := options[0]
		printer.BaudRate = override.BaudRate
		printer.SerialPortName = override.SerialPortName
	}

	// open stream
	printer.Stream, err = serial.OpenPort(
		&serial.Config{
			Name: printer.SerialPortName,
			Baud: printer.BaudRate,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return printer
}
