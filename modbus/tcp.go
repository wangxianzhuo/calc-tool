package modbus

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

// TCPResponseParseReadHoldingRegister pares modbus tcp read holding register response,
// It can only parse the data to 1 number(1,2,4,8 bytes number)
func TCPResponseParseReadHoldingRegister(resp []byte, log io.Writer) (float64, error) {
	var size int16
	if len(resp) > 6 {
		size = int16(binary.BigEndian.Uint16(resp[4:6]))
		if len(resp) != (6 + int(size)) {
			return 0, fmt.Errorf("modbus-tcp response %X is illegal, pdu size is %d", resp, size)
		}
	} else {
		return 0, fmt.Errorf("the response %X is not modbus-tcp response", resp)
	}

	protocl := resp[2:4]
	pdu := resp[6 : 6+int(size)]

	if bytes.Compare(protocl, []byte{0, 0}) != 0 {
		return 0, fmt.Errorf("do not support protocol %X for modbus", protocl)
	}

	slaveAddr, funCode, dataSize, value, err := parsePDU(pdu)
	if err != nil {
		return 0, fmt.Errorf("parse modbus tcp PDU %X error: %v", pdu, err)
	}

	if log == nil {
		log = ioutil.Discard
	}
	fmt.Fprintf(log, "parse tcp PDU [%X]: slave address %d, function code %d, data size %d, value %f", pdu, slaveAddr, funCode, dataSize, value)

	return value, nil
}

// TCPRequestGenerateReadHoldingRegister generate read holding regisers request
func TCPRequestGenerateReadHoldingRegister(slaveAddr, registerAddr, registerNum int, taskHi, taskLo byte) ([]byte, error) {
	registerAddrHi := byte(int32(registerAddr) >> 8)
	registerAddrLo := byte(int32(registerAddr) & 0xFF)

	registerNumHi := byte(int32(registerNum) >> 8)
	registerNumLo := byte(int32(registerNum) & 0xFF)

	pdu := []byte{byte(slaveAddr), 0x03, registerAddrHi, registerAddrLo, registerNumHi, registerNumLo}

	request := []byte{taskHi, taskLo, 0x0, 0x0, 0x0, 0x0, 0x6, pdu[0], pdu[1], pdu[2], pdu[3], pdu[4], pdu[5]}

	return request, nil
}

func parsePDU(pdu []byte) (slaveAddr, funCode, size int, data float64, err error) {
	if len(pdu) < 6 {
		return 0, 0, 0, 0, fmt.Errorf("pdu %X illegal", pdu)
	}

	slaveAddr = int(pdu[0])
	funCode = int(pdu[1])
	size = int(pdu[2])

	switch size {
	case 1:
		data = float64(pdu[3])
	case 2:
		data = float64(binary.BigEndian.Uint16(pdu[3 : 3+size]))
	case 4:
		data = float64(math.Float32frombits(binary.BigEndian.Uint32(pdu[3 : 3+size])))
	case 8:
		data = float64(math.Float64frombits(binary.BigEndian.Uint64(pdu[3 : 3+size])))
	default:
		return 0, 0, 0, 0, fmt.Errorf("can't support more than 8 bytes value convert")
	}
	return
}
