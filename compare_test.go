package cmp

import (
	"reflect"
	"testing"
	"time"
)

type testPair struct {
	a      interface{}
	b      interface{}
	expEq  bool
	expNeq bool
	expLt  bool
	expGt  bool
	expLte bool
	expGte bool
}

type testInterface struct {
	Number int
}

func (t testInterface) Eq(b interface{}) (bool, error) {
	return t.Number == b.(testInterface).Number, nil
}

func (t testInterface) Lt(b interface{}) (bool, error) {
	return t.Number < b.(testInterface).Number, nil
}

var now = time.Now()
var testData = []testPair{
	{a: true, b: true, expEq: true, expNeq: false},
	{a: true, b: false, expEq: false, expNeq: true},

	{a: int(1), b: int(1), expEq: true, expNeq: false, expLt: false, expGt: false, expLte: true, expGte: true},
	{a: int(1), b: int(2), expEq: false, expNeq: true, expLt: true, expGt: false, expLte: true, expGte: false},
	{a: int(2), b: int(1), expEq: false, expNeq: true, expLt: false, expGt: true, expLte: false, expGte: true},

	{a: int8(1), b: int8(1), expEq: true, expNeq: false, expLt: false, expGt: false, expLte: true, expGte: true},
	{a: int8(1), b: int8(2), expEq: false, expNeq: true, expLt: true, expGt: false, expLte: true, expGte: false},
	{a: int8(2), b: int8(1), expEq: false, expNeq: true, expLt: false, expGt: true, expLte: false, expGte: true},

	{a: int16(1), b: int16(1), expEq: true, expNeq: false, expLt: false, expGt: false, expLte: true, expGte: true},
	{a: int16(1), b: int16(2), expEq: false, expNeq: true, expLt: true, expGt: false, expLte: true, expGte: false},
	{a: int16(2), b: int16(1), expEq: false, expNeq: true, expLt: false, expGt: true, expLte: false, expGte: true},

	{a: int32(1), b: int32(1), expEq: true, expNeq: false, expLt: false, expGt: false, expLte: true, expGte: true},
	{a: int32(1), b: int32(2), expEq: false, expNeq: true, expLt: true, expGt: false, expLte: true, expGte: false},
	{a: int32(2), b: int32(1), expEq: false, expNeq: true, expLt: false, expGt: true, expLte: false, expGte: true},

	{a: int64(1), b: int64(1), expEq: true, expNeq: false, expLt: false, expGt: false, expLte: true, expGte: true},
	{a: int64(1), b: int64(2), expEq: false, expNeq: true, expLt: true, expGt: false, expLte: true, expGte: false},
	{a: int64(2), b: int64(1), expEq: false, expNeq: true, expLt: false, expGt: true, expLte: false, expGte: true},

	{a: uint(1), b: uint(1), expEq: true, expNeq: false, expLt: false, expGt: false, expLte: true, expGte: true},
	{a: uint(1), b: uint(2), expEq: false, expNeq: true, expLt: true, expGt: false, expLte: true, expGte: false},
	{a: uint(2), b: uint(1), expEq: false, expNeq: true, expLt: false, expGt: true, expLte: false, expGte: true},

	{a: uint8(1), b: uint8(1), expEq: true, expNeq: false, expLt: false, expGt: false, expLte: true, expGte: true},
	{a: uint8(1), b: uint8(2), expEq: false, expNeq: true, expLt: true, expGt: false, expLte: true, expGte: false},
	{a: uint8(2), b: uint8(1), expEq: false, expNeq: true, expLt: false, expGt: true, expLte: false, expGte: true},

	{a: uint16(1), b: uint16(1), expEq: true, expNeq: false, expLt: false, expGt: false, expLte: true, expGte: true},
	{a: uint16(1), b: uint16(2), expEq: false, expNeq: true, expLt: true, expGt: false, expLte: true, expGte: false},
	{a: uint16(2), b: uint16(1), expEq: false, expNeq: true, expLt: false, expGt: true, expLte: false, expGte: true},

	{a: uint32(1), b: uint32(1), expEq: true, expNeq: false, expLt: false, expGt: false, expLte: true, expGte: true},
	{a: uint32(1), b: uint32(2), expEq: false, expNeq: true, expLt: true, expGt: false, expLte: true, expGte: false},
	{a: uint32(2), b: uint32(1), expEq: false, expNeq: true, expLt: false, expGt: true, expLte: false, expGte: true},

	{a: uint64(1), b: uint64(1), expEq: true, expNeq: false, expLt: false, expGt: false, expLte: true, expGte: true},
	{a: uint64(1), b: uint64(2), expEq: false, expNeq: true, expLt: true, expGt: false, expLte: true, expGte: false},
	{a: uint64(2), b: uint64(1), expEq: false, expNeq: true, expLt: false, expGt: true, expLte: false, expGte: true},

	{a: float32(1), b: float32(1), expEq: true, expNeq: false, expLt: false, expGt: false, expLte: true, expGte: true},
	{a: float32(1), b: float32(2), expEq: false, expNeq: true, expLt: true, expGt: false, expLte: true, expGte: false},
	{a: float32(2), b: float32(1), expEq: false, expNeq: true, expLt: false, expGt: true, expLte: false, expGte: true},

	{a: float64(1), b: float64(1), expEq: true, expNeq: false, expLt: false, expGt: false, expLte: true, expGte: true},
	{a: float64(1), b: float64(2), expEq: false, expNeq: true, expLt: true, expGt: false, expLte: true, expGte: false},
	{a: float64(2), b: float64(1), expEq: false, expNeq: true, expLt: false, expGt: true, expLte: false, expGte: true},

	{a: now, b: now, expEq: true, expNeq: false, expLt: false, expGt: false, expLte: true, expGte: true},
	{a: now, b: now.Add(time.Second), expEq: false, expNeq: true, expLt: true, expGt: false, expLte: true, expGte: false},
	{a: now.Add(time.Second), b: now, expEq: false, expNeq: true, expLt: false, expGt: true, expLte: false, expGte: true},

	{a: "abc123", b: "abc123", expEq: true, expNeq: false, expLt: false, expGt: false, expLte: true, expGte: true},
	{a: "abc123", b: "def456", expEq: false, expNeq: true, expLt: true, expGt: false, expLte: true, expGte: false},
	{a: "def456", b: "abc123", expEq: false, expNeq: true, expLt: false, expGt: true, expLte: false, expGte: true},

	{a: testInterface{1}, b: testInterface{1}, expEq: true, expNeq: false, expLt: false, expGt: false, expLte: true, expGte: true},
	{a: testInterface{1}, b: testInterface{2}, expEq: false, expNeq: true, expLt: true, expGt: false, expLte: true, expGte: false},
	{a: testInterface{2}, b: testInterface{1}, expEq: false, expNeq: true, expLt: false, expGt: true, expLte: false, expGte: true},
}

func TestCompare(t *testing.T) {
	_, err := Eq(1, "abc123")
	if err != ErrNotSameKind {
		t.Error(err)
	}

	_, err = Eq(testData, testData)
	if err != ErrKindNotSupported {
		t.Error(err)
	}

	_, err = Lt(1, "abc123")
	if err != ErrNotSameKind {
		t.Error(err)
	}

	_, err = Lt(testData, testData)
	if err != ErrKindNotSupported {
		t.Error(err)
	}

	for _, d := range testData {
		m, err := Eq(d.a, d.b)
		if err != nil {
			t.Error(err)
		}
		if m != d.expEq {
			t.Errorf("want:\n%v, got:\n%v", d.expEq, m)
		}

		m, err = Neq(d.a, d.b)
		if err != nil {
			t.Error(err)
		}
		if m != d.expNeq {
			t.Errorf("want:\n%v, got:\n%v", d.expNeq, m)
		}

		v := reflect.ValueOf(d.a)
		if v.Kind() == reflect.Bool {
			continue
		}

		m, err = Lt(d.a, d.b)
		if err != nil {
			t.Error(err)
		}
		if m != d.expLt {
			t.Errorf("want:\n%v, got:\n%v", d.expLt, m)
		}

		m, err = Gt(d.a, d.b)
		if err != nil {
			t.Error(err)
		}
		if m != d.expGt {
			t.Errorf("want:\n%v, got:\n%v", d.expGt, m)
		}

		m, err = Lte(d.a, d.b)
		if err != nil {
			t.Error(err)
		}
		if m != d.expLte {
			t.Errorf("want:\n%v, got:\n%v", d.expLte, m)
		}

		m, err = Gte(d.a, d.b)
		if err != nil {
			t.Error(err)
		}
		if m != d.expGte {
			t.Errorf("want:\n%v, got:\n%v", d.expGte, m)
		}
	}
}
