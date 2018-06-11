package modbus

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

// InstructionParse parse modbus-rtu instruction
func InstructionParse(ins string) ([]byte, error) {
	instruction := []byte{}
	if ins == "" {
		return nil, fmt.Errorf("parse instruction error: instruction is empty")
	} else if len(ins) < 6 {
		return nil, fmt.Errorf("parse instruction %v error: instruction is illegal", ins)
	}

	for i := 2; i <= len(ins); i += 2 {
		t, err := parseByte(ins[i-2 : i])
		if err != nil {
			return nil, fmt.Errorf("parse instruction %v error: %v", ins, err)
		}
		instruction = append(instruction, t)
	}
	return instruction, nil
}

func parseByte(byteStr string) (byte, error) {
	bufIns, err := strconv.ParseUint(byteStr, 16, 8)
	if err != nil {
		return 0, fmt.Errorf("parse byte %v error: %v", byteStr, err)
	}
	i := bytes.NewBuffer([]byte{})
	binary.Write(i, binary.BigEndian, &bufIns)
	return i.Bytes()[7], nil
}

// InstructionCheck check modbus-rtu instruction check code
func InstructionCheck(ins string) error {
	instruction, err := InstructionParse(ins)
	if err != nil {
		return fmt.Errorf("instruction %v check error: %v", ins, err)
	}

	if !CRC16Check(instruction) {
		return fmt.Errorf("instruction %v check code illegal", ins)
	}
	return nil
}

// InstructionWithCheckCode generate modbus-rtu instruction with check code
func InstructionWithCheckCode(instructionWithoutCheckCode string) ([]byte, error) {
	instruction, err := InstructionParse(instructionWithoutCheckCode)
	if err != nil {
		return nil, fmt.Errorf("instruction %v check error: %v", instructionWithoutCheckCode, err)
	}
	hi, lo, _ := CRC16CheckCode(instruction)
	instruction = append(instruction, hi)
	instruction = append(instruction, lo)
	return instruction, nil
}
