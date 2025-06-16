package mem

import "C"
import (
	"fmt"
	"goland-study/src/go-c-gc/c"
	"unsafe"
)

type Buf struct {
	Next     *Buf           // 链表
	Capacity int            // 此Buf容量
	length   int            // 当前数据长度
	head     int            // 数据头地址
	data     unsafe.Pointer // 数据地址
}

func NewBuf(capacity int) *Buf {
	return &Buf{
		Capacity: capacity,
		length:   0,
		head:     0,
		data:     c.Malloc(capacity),
	}
}

func (b *Buf) SetBytes(src []byte) {
	c.MemCopy(unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), src, len(src))
	b.length = len(src)
}

func (b *Buf) GetBytes() []byte {
	data := C.GoBytes(unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), C.int(b.length))
	return data
}

func (b *Buf) Copy(other *Buf) {
	c.MemCopy(b.data, other.GetBytes(), other.length)
	b.head = 0
	b.length = other.length
}

func (b *Buf) Pop(len int) {
	if b.data == nil {
		fmt.Printf("buf pop data is nil\n")
		return
	}
	if len > b.length {
		fmt.Printf("buf pop len > length\n")
	}
	b.length -= len
	b.head += len
}

func (b *Buf) Adjust() {
	if b.head != 0 {
		if b.length != 0 {
			c.MemMove(b.data, unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), b.length)
		}
		b.head = 0
	}
}

func (b *Buf) Clear() {
	b.head = 0
	b.length = 0
}

func (b *Buf) Head() int {
	return b.head
}

func (b *Buf) Length() int {
	return b.length
}
