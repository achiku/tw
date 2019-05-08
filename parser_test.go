package main

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	f, err := os.Open("./testdata/achiku.csv")
	if err != nil {
		t.Fatal(err)
	}
	sts, err := parser(f)
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range sts {
		t.Logf("%+v", s)
	}
}
