package json

import (
	"testing"

	"github.com/Schinti95/storm/codec/internal"
)

func TestJSON(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}
