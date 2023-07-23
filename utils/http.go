package utils

import "errors"

func CheckValidPortNumber(port uint16) error {
	const MAX_PORT_RANGE = 65535
	if port > MAX_PORT_RANGE {
		return errors.New("invalid port number")
	}
	return nil
}