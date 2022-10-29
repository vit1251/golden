package commonfunc

import "strconv"

func ParseSize(value []byte) (int, error) {
	var str string = string(value)
	size, err1 := strconv.ParseUint(str, 10, 32)
	return int(size), err1
}

