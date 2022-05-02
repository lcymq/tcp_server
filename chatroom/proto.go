package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

func Encode(message string) ([]byte, error) {
	length := len(message)
	pkg := new(bytes.Buffer)
    err := binary.Write(pkg, binary.LittleEndian, length)
    if err != nil {
        return nil, err
    }
    err = binary.Write(pkg, binary.LittleEndian, []byte(message))
    if err != nil {
        return nil, err
    }
    err = binary.Write(pkg, binary.LittleEndian, []byte(message))
    if err != nil {
        return nil, err
    }
    fmt.Println(pkg)
    return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
    lengthHeader, _ := reader.Peek(4) // read 4 bytes data
    lengthBuff := bytes.NewBuffer(lengthHeader)
    var length int32
    err := binary.Read(lengthBuff, binary.LittleEndian, &length)
    if err != nil {
        return "", err
    }
    if int32(reader.Buffered()) < length + 4 {
        return "", err
    }
    // read the real msg
    pack := make([]byte, int(length + 4))
    _, err = reader.Read(pack)
    if err != nil {
        return "", err
    }
    return string(pack[4:]), nil
}