package utils

import "strconv"

func ParseStr(str string) (int, error) {
	if len(str) == 0 {
		return 0, nil
	}

	res, err := strconv.Atoi(str)
	if err != nil {
		return -1, err
	}

	return res, nil
}
