package thermoprinter

import (
	"os"
	"testing"
)

func MainTest(m *testing.M) {
	retCode := m.Run()

	os.Exit(retCode)
}

func TestThermoPrinter(t *testing.T) {
	printer := NewPrinter()

	printer.Write("hello world.")
}
