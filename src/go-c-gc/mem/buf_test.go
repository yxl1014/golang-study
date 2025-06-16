package mem

import (
	"fmt"
	"testing"
)

func TestBufPoolSetGet(t *testing.T) {
	pool := MemPool()

	buffer, err := pool.Alloc(1)
	if err != nil {
		t.Fatalf("Error allocating buffer: %s", err)
	}

	buffer.SetBytes([]byte("AceId12345"))
	fmt.Printf("GetBytes = %+v, ToString = %s\n", buffer.GetBytes(), string(buffer.GetBytes()))
	buffer.Pop(4)
	fmt.Printf("GetBytes = %+v, ToString = %s\n", buffer.GetBytes(), string(buffer.GetBytes()))
}

func TestBufPoolCopy(t *testing.T) {
	pool := MemPool()

	buffer, err := pool.Alloc(1)
	if err != nil {
		t.Fatalf("Error allocating buffer: %s", err)
	}

	buffer.SetBytes([]byte("AceId12345"))
	fmt.Printf("GetBytes = %+v, ToString = %s\n", buffer.GetBytes(), string(buffer.GetBytes()))

	buffer2, err := pool.Alloc(1)
	if err != nil {
		t.Fatalf("Error allocating buffer: %s", err)
	}

	buffer2.Copy(buffer)
	fmt.Printf("buffer2 GetBytes = %+v, ToString = %s\n", buffer2.GetBytes(), string(buffer2.GetBytes()))
}

func TestBufPoolAdjust(t *testing.T) {

	pool := MemPool()

	buffer, err := pool.Alloc(1)
	if err != nil {
		t.Fatalf("Error allocating buffer: %s", err)
	}
	buffer.SetBytes([]byte("AceId12345"))
	fmt.Printf("GetBytes = %+v, Head = %d, Length = %d\n", buffer.GetBytes(), buffer.Head(), buffer.Length())
	buffer.Pop(4)
	fmt.Printf("GetBytes = %+v, Head = %d, Length = %d\n", buffer.GetBytes(), buffer.Head(), buffer.Length())
	buffer.Adjust()
	fmt.Printf("GetBytes = %+v, Head = %d, Length = %d\n", buffer.GetBytes(), buffer.Head(), buffer.Length())
}
