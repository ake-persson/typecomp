package cmp

import (
	"testing"
	"time"
)

type testPair struct {
	a   interface{}
	b   interface{}
	exp bool
}

var now = time.Now()
var testData = []testPair{
	testPair{a: true, b: true, exp: true},
	testPair{a: true, b: false, exp: false},

	testPair{a: int(1), b: int(1), exp: true},
	testPair{a: int(1), b: int(2), exp: false},

	testPair{a: int8(1), b: int8(1), exp: true},
	testPair{a: int8(1), b: int8(2), exp: false},

	testPair{a: int16(1), b: int16(1), exp: true},
	testPair{a: int16(1), b: int16(2), exp: false},

	testPair{a: int32(1), b: int32(1), exp: true},
	testPair{a: int32(1), b: int32(2), exp: false},

	testPair{a: int64(1), b: int64(1), exp: true},
	testPair{a: int64(1), b: int64(2), exp: false},

	testPair{a: uint(1), b: uint(1), exp: true},
	testPair{a: uint(1), b: uint(2), exp: false},

	testPair{a: uint8(1), b: uint8(1), exp: true},
	testPair{a: uint8(1), b: uint8(2), exp: false},

	testPair{a: uint16(1), b: uint16(1), exp: true},
	testPair{a: uint16(1), b: uint16(2), exp: false},

	testPair{a: uint32(1), b: uint32(1), exp: true},
	testPair{a: uint32(1), b: uint32(2), exp: false},

	testPair{a: uint64(1), b: uint64(1), exp: true},
	testPair{a: uint64(1), b: uint64(2), exp: false},

	testPair{a: float32(1), b: float32(1), exp: true},
	testPair{a: float32(1), b: float32(2), exp: false},

	testPair{a: float64(1), b: float64(1), exp: true},
	testPair{a: float64(1), b: float64(2), exp: false},

	testPair{a: now, b: now, exp: true},
	testPair{a: now, b: now.Add(time.Second), exp: false},
}

func TestEq(t *testing.T) {
	for _, d := range testData {
		m, err := Eq(d.a, d.b)
		if err != nil {
			t.Error(err)
		}
		if m != d.exp {
			t.Errorf("want:\n%v, got:\n%v", d.exp, m)
		}
	}
}
