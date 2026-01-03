package util

import "strconv"

var Convert ConvertManager

type ConvertManager struct{}

func (ConvertManager) Str2Bool(data string) (bool, error) {
	boolValue, err := strconv.ParseBool(data)
	if err != nil {
		return false, err
	}
	return boolValue, nil
}

func (ConvertManager) Str2Int(data string) (int, error) {
	i, err := strconv.Atoi(data)
	if err != nil {
		return -1, err
	}
	return i, nil
}

func (ConvertManager) Int2Str(data int) string {
	return strconv.Itoa(data)
}
