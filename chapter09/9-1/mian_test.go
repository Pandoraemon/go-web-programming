package main

import (
	"testing"
	"time"
)

func TestPrint1(t *testing.T) {
	print1()
	time.Sleep(1 * time.Millisecond)
}

func TestGoPrint1(t *testing.T) {
	goPrint1()
	time.Sleep(1 *  time.Millisecond)
}
