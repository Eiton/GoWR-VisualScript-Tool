package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

// type FileHeader struct {
// 	MagicNumber [4]byte
// 	Version     uint32
// 	Unknown1    uint32
// 	Unknown2    uint32
// 	NodeCount   uint32
// }

type VsFunctionHead struct {
	Padding      [4]uint64
	NumFunctions uint32
	Padding2     uint32
	Offset       uint32
	Padding3     uint32
	Padding4     uint64
}

type VsFunction struct {
	nameHash uint64
	name     string
	head     Node
}

type VisualScript struct {
	// Header FileHeader
	name string
	// functions []VsFunctionHead
	// Nodes     []*Node
}

var nodeOffsets []uint32

func ReadLengthAndOffset(file *os.File, fileOffset int64) (uint32, uint32, error) {
	_, err := file.Seek(fileOffset, 0)
	if err != nil {
		return 0, 0, err
	}
	var length uint32
	err = binary.Read(file, binary.LittleEndian, &length)
	if err != nil {
		return 0, 0, err
	}
	_, err = file.Seek(fileOffset+0x8, 0)
	if err != nil {
		return 0, 0, err
	}
	var offset uint32
	err = binary.Read(file, binary.LittleEndian, &offset)
	if err != nil {
		return 0, 0, err
	}
	return length, offset, nil
}

func ReadVisualScriptFromFile(filePath string) (*VisualScript, error) {
	vs := &VisualScript{}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	_, err = file.Seek(0x7c, 0)
	if err != nil {
		return nil, err
	}
	var length uint32
	var offset uint32
	length, offset, err = ReadLengthAndOffset(file, 0x78)
	if err != nil {
		return nil, err
	}
	_, err = file.Seek(int64(0x80+offset), 0)
	if err != nil {
		return nil, err
	}
	nameData := make([]byte, length)
	_, err = file.Read(nameData)
	if err != nil {
		return nil, err
	}
	vs.name = string(nameData)
	fmt.Printf("Visual Script Name: %s\n", vs.name)

	length, offset, err = ReadLengthAndOffset(file, 0xa8)
	if err != nil {
		return nil, err
	}
	nodeOffsets = make([]uint32, length)
	if err != nil {
		return nil, err
	}
	for i := uint32(0); i < length; i++ {
		_, err = file.Seek(int64(0xb0+offset+i*8), 0)
		if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &nodeOffsets[i])
		if err != nil {
			return nil, err
		}
		nodeOffsets[i] = nodeOffsets[i] + uint32(0xb0) + offset + i*8
	}

	length, offset, err = ReadLengthAndOffset(file, 0x1d8)
	if err != nil {
		return nil, err
	}
	functions := make([]VsFunctionHead, length)
	_, err = file.Seek(int64(0x1e0+offset), 0)
	if err != nil {
		return nil, err
	}
	err = binary.Read(file, binary.LittleEndian, &functions)
	if err != nil {
		return nil, err
	}
	for i, function := range functions {
		for j := uint32(0); j < function.NumFunctions; j++ {
			_, err = file.Seek(int64(0x1e0+offset+0x28+0x38*uint32(i)+function.Offset+uint32(j)*0x30), 0)
			if err != nil {
				return nil, err
			}
			f := VsFunction{}
			err = binary.Read(file, binary.LittleEndian, &f.nameHash)
			if err != nil {
				return nil, err
			}
			_, err = file.Seek(int64(0x1e0+offset+0x28+0x38*uint32(i)+function.Offset+uint32(j)*0x30+0x10), 0)
			if err != nil {
				return nil, err
			}
			var functionNameOffset uint32
			err = binary.Read(file, binary.LittleEndian, &functionNameOffset)
			if err != nil {
				return nil, err
			}
			_, err = file.Seek(int64(0x1e0+offset+0x28+0x38*uint32(i)+function.Offset+uint32(j)*0x30+0x10+functionNameOffset), 0)
			if err != nil {
				return nil, err
			}
			var nameLength uint32
			err = binary.Read(file, binary.LittleEndian, &nameLength)
			_, err = file.Seek(int64(0x1e0+offset+0x28+0x38*uint32(i)+function.Offset+uint32(j)*0x30+0x10+functionNameOffset+0x10), 0)
			if err != nil {
				return nil, err
			}
			nameData := make([]byte, nameLength)
			_, err = file.Read(nameData)
			if err != nil {
				return nil, err
			}
			f.name = string(nameData)

			_, err = file.Seek(int64(0x1e0+offset+0x28+0x38*uint32(i)+function.Offset+uint32(j)*0x30+0x20), 0)
			if err != nil {
				return nil, err
			}
			var functionHeadOffset uint32
			err = binary.Read(file, binary.LittleEndian, &functionHeadOffset)
			if err != nil {
				return nil, err
			}

			_, err = file.Seek(int64(0x1e0+offset+0x28+0x38*uint32(i)+function.Offset+uint32(j)*0x30+0x20+functionHeadOffset), 0)
			if err != nil {
				return nil, err
			}
			var nodeIndex uint16
			err = binary.Read(file, binary.LittleEndian, &nodeIndex)
			if err != nil {
				return nil, err
			}
			f.head, err = ReadNodeHead(file, nodeOffsets[uint32(nodeIndex)])
			if err != nil {
				return nil, err
			}
			fmt.Printf("function %s(){\n%s\n}\n", f.name, f.head.To_String())
		}
	}
	return vs, nil
}
