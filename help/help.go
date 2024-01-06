package help

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
