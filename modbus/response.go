package modbus

import "fmt"

// ResponseCheck response check
func ResponseCheck(req, resp []byte) error {
	if !CRC16Check(resp) {
		return fmt.Errorf("response %X check failed", resp)
	}

	if resp[0] != req[0] || resp[1] != req[1] {
		return fmt.Errorf("response %X can't match request %X", resp, req)
	}
	return nil
}
