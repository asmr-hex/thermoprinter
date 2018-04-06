package thermoprinter

import "time"

// set initial settings for thermoprinter.
func (p *Printer) init() {
	// we need to allow the printer at least 0.5 seconds of uptime
	// before it can get data.
	time.Sleep(time.Millisecond * 500)

}
