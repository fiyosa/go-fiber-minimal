package util

import (
	"strconv"
	"time"
)

var Convert ConvertManager

type ConvertManager struct{}

// "true" => true
func (ConvertManager) Str2Bool(data string) (bool, error) {
	boolValue, err := strconv.ParseBool(data)
	if err != nil {
		return false, err
	}
	return boolValue, nil
}

// "1" => 1
func (ConvertManager) Str2Int(data string) (int, error) {
	i, err := strconv.Atoi(data)
	if err != nil {
		return -1, err
	}
	return i, nil
}

// 1 => "1"
func (ConvertManager) Int2Str(data int) string {
	return strconv.Itoa(data)
}

// time.Time => "2006-01-02 15:04:05"
func (ConvertManager) Datetime2Str(date time.Time) string {
	layout := "2006-01-02 15:04:05"
	return date.Format(layout) // yyyy-MM-dd hh:mm:ss
}

// time.Time => "2006-01-02"
func (ConvertManager) Date2Str(date time.Time) string {
	layout := "2006-01-02"
	return date.Format(layout) // yyyy-MM-dd
}
