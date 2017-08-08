package msgpack

import (
	"testing"

	"github.com/Schinti95/storm/codec/internal"
)

func TestMsgpack(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}
