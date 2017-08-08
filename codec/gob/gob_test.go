package gob

import (
	"testing"

	"github.com/Schinti95/storm/codec/internal"
)

func TestGob(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}
