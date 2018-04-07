package thermoprinter

func (p *Printer) UpsideDownOn() error {
	return p.setPrintMode(UpsideDownMask)
}

func (p *Printer) setPrintMode(mask byte) error {
	p.printMode |= mask
	err := p.writePrintMode()

	return err
}

func (p *Printer) unsetPrintMode(mask byte) error {
	p.printMode &= ^mask
	err := p.writePrintMode()

	return err
}

func (p *Printer) writePrintMode() error {
	return p.writeBytes([]byte{ASCII_ESC, byte('!'), p.printMode})
}
