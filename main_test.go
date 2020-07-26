package main

import (
	"testing"

	"github.com/alanslko/golab/pkg/formatter"
)

func TestAll(t *testing.T) {
	if !(formatter.GetName() == "formatter") {
		t.Error("testall fail")
	}
}
