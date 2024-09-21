package my_math


import (
	"fmt"
	"strconv"
)


type Sign int
const (
	Positive Sign = iota
	Negative
	None
)


type BigInt struct {
	sign Sign
	len uint64
	digits []uint64
}


func NewUnsignedBigInt[V uint|uint8|uint32|uint64](number V) BigInt {
	len := uint64(0)
	digits := make([]uint64, 0)

	if number == V(0) {
		digits = append(digits, 0)
		len++
	} else {
		for {
			if number == V(0) {
				break
			}
			digits = append(digits, uint64(number % 10))
			number = V(number/V(10))
			len++
		}
	}

	return BigInt{
		sign: None,
		len: len,
		digits: digits,
	}
}


func (b *BigInt) Len() uint64 {
	return b.len
}


func (b *BigInt) Copy() BigInt {
	digits := make([]uint64, 0)
	digits = append(digits, b.digits...)

	return BigInt{
		sign: b.sign,
		digits: digits,
		len: b.len,
	}
}


func (b1 *BigInt) Equal(b2 *BigInt) bool {
	if b1.len != b2.len {
		return false
	}

	for i := uint64(0); i < b1.len; i++ {
		if b1.digits[i] != b2.digits[i] {
			return false
		}
	}

	return true
}


func (b *BigInt) Get(index uint64) (uint64, error) {
	if b.len <= index {
		return uint64(0), fmt.Errorf("index cannot be exceed the number length\n")
	}
	return b.digits[index], nil
}


func (b BigInt) String() string {
	str := ""
	if b.len > 0 {
		for i := b.len-1; i != 0; i-- {
			str += strconv.FormatUint(b.digits[i], 10)
		}
	}
	str += strconv.FormatUint(b.digits[0], 10)
	return str
}


func (b1 *BigInt) Add(b2 *BigInt) {
	i := uint64(0)
	moveNextVal := uint64(0)

	for i < b1.len || i < b2.len {
		val1 := uint64(0)
		val2 := uint64(0)

		if i < b1.len {
			val1 = b1.digits[i]
		}

		if i < b2.len {
			val2 = b2.digits[i]
		}

		val3 := val1 + val2 + moveNextVal
		if val3 < 10 {
			if i >= b1.len {
				b1.digits = append(b1.digits, val3)
				b1.len++
			} else {
				b1.digits[i] = val3
			}
			moveNextVal = 0
			i++
			continue
		}

		if i >= b1.len {
			b1.digits = append(b1.digits, val3 % 10)
			b1.len++
		} else {
			b1.digits[i] = val3 % 10
		}
		moveNextVal = uint64(val3 / uint64(10))
		i++
	}

	if moveNextVal != 0 {
		b1.digits = append(b1.digits, moveNextVal)
		b1.len++
	}
}


func (b1 *BigInt) Mult(b2 *BigInt) {
	temp1 := NewUnsignedBigInt(uint(0))
	multParts := make([]BigInt, 0)
	tenthPower := uint64(0)

	if temp1.Equal(b2) || temp1.Equal(b1) {
		*b1 = temp1 
		return
	}

	for i := uint64(0); i < b2.len; i++ {
		n1, _ := b2.Get(i)
		temp := uint64(0)
		moveNextVal := uint64(0)

		b1Copy := b1.Copy()
		if n1 == uint64(0) {
			tenthPower += 1
			continue
		}

		for j := uint64(0); j < b1Copy.len; j++ {
			n2, _ := b1Copy.Get(j)
			temp = n1 * n2 + moveNextVal
			b1Copy.digits[j] = temp % 10 
			moveNextVal = uint64(temp / uint64(10))
		}

		if moveNextVal != uint64(0) {
			b1Copy.digits = append(b1Copy.digits, moveNextVal)
			b1Copy.len++
		}

		if (tenthPower > 0) {
			tempSlice := make([]uint64, 0)
			for k := uint64(0); k < tenthPower; k++ {
				tempSlice = append(tempSlice, 0)
				b1Copy.len++
			}
			tempSlice = append(tempSlice, b1Copy.digits...)
			b1Copy.digits = tempSlice
		}

		multParts = append(multParts, b1Copy)
		tenthPower += 1
	}

	for _, m := range multParts {
		temp1.Add(&m)
	}

	*b1 = temp1
}


func (b1 *BigInt) Pow(b2 *BigInt) {
	temp1 := NewUnsignedBigInt(uint(1))
	temp2 := NewUnsignedBigInt(uint(0)) 

	if temp2.Equal(b2) {
		*b1 = temp1 
		return
	}
	temp2.Add(&temp1)

	if temp2.Equal(b2) {
		return
	}

	b1Copy := b1.Copy()
	for {
		if temp2.Equal(b2) {
			break
		}
		temp2.Add(&temp1)
		b1.Mult(&b1Copy)
	}
}
