package c_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
	"unsafe"
)

func IsLittleEndian() bool {
	var n int32 = 0x01020304

	// 讲int 转换为 void*
	u := unsafe.Pointer(&n)
	// 强转 byte*
	pb := (*byte)(u)
	// 获取第一位
	b := *pb

	// 小端 : 04
	// 大端 : 01
	return b == 0x04
}

func IntToBytes(n uint32) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})

	var order binary.ByteOrder
	if IsLittleEndian() {
		order = binary.LittleEndian
	} else {
		order = binary.BigEndian
	}
	binary.Write(bytesBuffer, order, x)
	return bytesBuffer.Bytes()
}

func TestMemoryC(t *testing.T) {
	data := Malloc(4)
	fmt.Printf(" data % +v, %T\n", data, data)

	myData := (*uint32)(data)
	*myData = 5
	fmt.Printf(" data % +v, %T\n", *myData, *myData)

	var a uint32 = 100
	MemCopy(data, IntToBytes(a), 4)
	fmt.Printf(" data % +v, %T\n", *myData, *myData)
	Free(data)
}
