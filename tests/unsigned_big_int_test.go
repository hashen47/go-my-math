package my_math_test


import (
	"testing"
	"my_math"
)


type TestCase struct {
	A my_math.BigInt
	B my_math.BigInt
	Expect my_math.BigInt
}


func TestAdd(t *testing.T) {
	testcases := []TestCase{
		{
			my_math.NewUnsignedBigInt(uint64(10)),
			my_math.NewUnsignedBigInt(uint64(12)),
			my_math.NewUnsignedBigInt(uint64(22)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(2)),
			my_math.NewUnsignedBigInt(uint64(8)),
			my_math.NewUnsignedBigInt(uint64(10)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(20)),
			my_math.NewUnsignedBigInt(uint64(10)),
			my_math.NewUnsignedBigInt(uint64(30)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(14)),
			my_math.NewUnsignedBigInt(uint64(16)),
			my_math.NewUnsignedBigInt(uint64(30)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(2)),
			my_math.NewUnsignedBigInt(uint64(105)),
			my_math.NewUnsignedBigInt(uint64(107)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(1)),
			my_math.NewUnsignedBigInt(uint64(102)),
			my_math.NewUnsignedBigInt(uint64(103)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(11)),
			my_math.NewUnsignedBigInt(uint64(10000)),
			my_math.NewUnsignedBigInt(uint64(10011)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(100000001)),
			my_math.NewUnsignedBigInt(uint64(10000)),
			my_math.NewUnsignedBigInt(uint64(100010001)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(12)),
			my_math.NewUnsignedBigInt(uint64(0)),
			my_math.NewUnsignedBigInt(uint64(12)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(0)),
			my_math.NewUnsignedBigInt(uint64(12)),
			my_math.NewUnsignedBigInt(uint64(12)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(0)),
			my_math.NewUnsignedBigInt(uint64(0)),
			my_math.NewUnsignedBigInt(uint64(0)),
		},
	}

	for _, testcase := range testcases {
		oneCopy := testcase.A.Copy()
		oneCopy.Add(&testcase.B)
		if ! testcase.Expect.Equal(&oneCopy) {
			t.Errorf("A: %s, B: %s, Expect: %s, Real: %s\n", testcase.A, testcase.B, testcase.Expect, oneCopy)
		}
	}
}


func TestMult(t *testing.T) {
	testcases := []TestCase{
		{
			my_math.NewUnsignedBigInt(uint64(10)),
			my_math.NewUnsignedBigInt(uint64(12)),
			my_math.NewUnsignedBigInt(uint64(120)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(2)),
			my_math.NewUnsignedBigInt(uint64(8)),
			my_math.NewUnsignedBigInt(uint64(16)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(20)),
			my_math.NewUnsignedBigInt(uint64(10)),
			my_math.NewUnsignedBigInt(uint64(200)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(10)),
			my_math.NewUnsignedBigInt(uint64(10)),
			my_math.NewUnsignedBigInt(uint64(100)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(2)),
			my_math.NewUnsignedBigInt(uint64(105)),
			my_math.NewUnsignedBigInt(uint64(210)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(1)),
			my_math.NewUnsignedBigInt(uint64(102)),
			my_math.NewUnsignedBigInt(uint64(102)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(11)),
			my_math.NewUnsignedBigInt(uint64(10000)),
			my_math.NewUnsignedBigInt(uint64(110000)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(12)),
			my_math.NewUnsignedBigInt(uint64(0)),
			my_math.NewUnsignedBigInt(uint64(0)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(0)),
			my_math.NewUnsignedBigInt(uint64(12)),
			my_math.NewUnsignedBigInt(uint64(0)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(0)),
			my_math.NewUnsignedBigInt(uint64(0)),
			my_math.NewUnsignedBigInt(uint64(0)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(10)),
			my_math.NewUnsignedBigInt(uint64(2)),
			my_math.NewUnsignedBigInt(uint64(20)),
		},
	}

	for _, testcase := range testcases {
		oneCopy := testcase.A.Copy()
		oneCopy.Mult(&testcase.B)
		if ! testcase.Expect.Equal(&oneCopy) {
			t.Errorf("A: %s, B: %s, Expect: %s, Real: %s\n", testcase.A, testcase.B, testcase.Expect, oneCopy)
		}
	}
}


func TestPow(t *testing.T) {
	testcases := []TestCase{
		{
			my_math.NewUnsignedBigInt(uint64(10)),
			my_math.NewUnsignedBigInt(uint64(2)),
			my_math.NewUnsignedBigInt(uint64(100)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(2)),
			my_math.NewUnsignedBigInt(uint64(10)),
			my_math.NewUnsignedBigInt(uint64(1024)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(2)),
			my_math.NewUnsignedBigInt(uint64(10)),
			my_math.NewUnsignedBigInt(uint64(1024)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(3)),
			my_math.NewUnsignedBigInt(uint64(4)),
			my_math.NewUnsignedBigInt(uint64(81)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(5)),
			my_math.NewUnsignedBigInt(uint64(3)),
			my_math.NewUnsignedBigInt(uint64(125)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(5)),
			my_math.NewUnsignedBigInt(uint64(4)),
			my_math.NewUnsignedBigInt(uint64(625)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(2)),
			my_math.NewUnsignedBigInt(uint64(1)),
			my_math.NewUnsignedBigInt(uint64(2)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(200)),
			my_math.NewUnsignedBigInt(uint64(1)),
			my_math.NewUnsignedBigInt(uint64(200)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(1)),
			my_math.NewUnsignedBigInt(uint64(201)),
			my_math.NewUnsignedBigInt(uint64(1)),
		},
		{
			my_math.NewUnsignedBigInt(uint64(201)),
			my_math.NewUnsignedBigInt(uint64(0)),
			my_math.NewUnsignedBigInt(uint64(1)),
		},
	}

	for _, testcase := range testcases {
		oneCopy := testcase.A.Copy()
		oneCopy.Pow(&testcase.B)
		if ! testcase.Expect.Equal(&oneCopy) {
			t.Errorf("A: %s, B: %s, Expect: %s, Real: %s\n", testcase.A, testcase.B, testcase.Expect, oneCopy)
		}
	}
}
