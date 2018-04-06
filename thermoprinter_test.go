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

	err := printer.Write("hello world.")
	if err != nil {
		t.Fatal(err)
	}
}
