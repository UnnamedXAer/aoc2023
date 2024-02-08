package help

import "slices"

func IfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadNumValue(line []byte, lineSize, pos int) (numValue int, numLen int) {
	for ; pos < lineSize; pos++ {
		if !(line[pos] >= '0' && line[pos] <= '9') {
			pos--
			break
		}
	}

	if pos == lineSize {
		pos--
	}

	for ; pos >= 0; pos-- {
		if !(line[pos] >= '0' && line[pos] <= '9') {
			pos++
			break
		}

		numLen++
		continue
	}
	if pos == -1 {
		// pos eq -1 means that the start was at pos eq 0
		pos++
	}

	multiplier := 1
	for i := 0; i < numLen; i++ {
		tmp := int(line[pos+numLen-1-i] - '0')
		numValue += tmp * multiplier
		multiplier *= 10
	}

	return numValue, numLen
}

func ReadNumValueFromEnd(line []byte, lastDigitPos int) (v int, vLen int) {
	multiplier := 1
	for ; lastDigitPos > -1 && IsNumber(line[lastDigitPos]); lastDigitPos-- {
		v += int(line[lastDigitPos]-'0') * multiplier
		vLen++
		multiplier *= 10
	}

	if lastDigitPos > -1 {
		if line[lastDigitPos] == '-' {
			v = -v
			vLen++
		}
	}

	return v, vLen
}

func IsNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

type Point struct {
	X, Y int
}

// queue
// type QElementAny = any
type QueueAny[T any] []T

func NewQAny[T any](n ...int) QueueAny[T] {
	size := 10
	if len(n) > 0 {
		size = n[0]
	}

	return make(QueueAny[T], size)
}

func (q QueueAny[T]) IsEmpty() bool {
	return len(q) == 0
}

func (q QueueAny[T]) Len() int {
	return len(q)
}

func (q *QueueAny[T]) Push(element T) {

	*q = append(*q, element)
}

func (q *QueueAny[T]) Pop() T {
	if len(*q) == 0 {
		panic("popping from empty queue")
	}

	first := (*q)[0]
	// q[0] = nil
	*q = slices.Delete(*q, 0, 0+1)

	return first
}
