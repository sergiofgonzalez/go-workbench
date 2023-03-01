package mascot_test

import (
	"testing"

	"example.com/01_hello-go/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher!" {
		t.Fatal("Wrong mascot!")
	}
}