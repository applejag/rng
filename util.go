package main

import "strconv"

func parseUint32(value string, base int) (uint32, error) {
	var result, err = strconv.ParseUint(value, base, 32)
	if err != nil {
		return 0, err
	}
	return uint32(result), nil
}

func parseUint16(value string, base int) (uint16, error) {
	var result, err = strconv.ParseUint(value, base, 16)
	if err != nil {
		return 0, err
	}
	return uint16(result), nil
}
