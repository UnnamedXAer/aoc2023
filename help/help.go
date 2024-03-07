package help

import (
	"slices"

	"golang.org/x/exp/constraints"
)

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

func Gcd[T constraints.Integer](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Lcm[T constraints.Integer](a, b T, integers ...T) T {
	result := a * b / Gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = Lcm(result, integers[i])
	}

	return result
}

// queue
// type QElementAny = any
type QueueAny[T any] []T

func NewQAny[T any](n ...int) QueueAny[T] {
	size := 10
	if len(n) > 0 {
		size = n[0]
	}

	return make(QueueAny[T], 0, size)
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
	// (*q)[0] = *new(T) // delete does it for us
	*q = slices.Delete(*q, 0, 0+1)

	return first
}

// ////////////////////////////////
type Stack[T any] []T

func NewStack[T any](n ...int) Stack[T] {
	size := 10
	if len(n) > 0 {
		size = n[0]
	}

	return make(Stack[T], size)
}

func (s Stack[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Stack[T]) Len() int {
	return len(s)
}

func (s *Stack[T]) Push(element T) {

	*s = append(*s, element)
}

func (s *Stack[T]) Pop() T {
	size := len(*s)
	if size == 0 {
		panic("popping from empty stack")
	}

	top := (*s)[size-1]
	(*s)[size-1] = *new(T)
	*s = (*s)[:size-1]

	return top
}
