package gribtest

import (
	"os"
	"testing"

	"github.com/nilsmagnus/grib/griblib"
)

func BenchmarkReadMessages(b *testing.B) {
	// run the Fib function b.N times

	f, err := os.Open("../integrationtestdata/template5_3.grib2")
	if err != nil {
		b.Fatalf("Could not open test-file %v", err)
	}

	for n := 0; n < b.N; n++ {
		griblib.ReadMessages(f)
	}
}