package thermoprinter

import "time"

const (
	DefaultBaudRate           = 19200
	DefaultSerialPortName     = "/dev/serial0"
	DefaultWaitTimeAfterWrite = time.Millisecond * 10

	MaxColumn = 32

	// ASCII codes
	ASCII_TAB = "\t" // horizontal tab
	ASCII_LF  = "\n" // line feed
	ASCII_FF  = "\f" // form feed
	ASCII_CR  = "\r" // carriage return
	ASCII_DC2 = 18   // device control 2
	ASCII_ESC = 27   // escape
	ASCII_FS  = 28   // field separator
	ASCII_GS  = 29   // group separator

	// print mode bit masks
	UpsideDownMask = (1 << 2)
)
