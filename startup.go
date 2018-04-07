package thermoprinter

import (
	"log"
	"time"
)

// set initial settings for thermoprinter.
func (p *Printer) init() {
	var (
		err error
	)

	// we need to allow the printer at least 0.01 seconds of uptime
	// before it can get data.
	time.Sleep(time.Millisecond * 10)

	// flush the stream before using.
	err = p.stream.Flush()
	if err != nil {
		log.Fatal(err)
	}

	// we want to print upsidedown since its better
	err = p.UpsideDownOn()
	if err != nil {
		log.Fatal(err)
	}
}
