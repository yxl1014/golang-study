package c

/*
#include <string.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// 调用C函数 malloc
func Malloc(size int) unsafe.Pointer {
	return C.malloc(C.size_t(size))
}

func Free(data unsafe.Pointer) {
	C.free(data)
}

func MemMove(dest, src unsafe.Pointer, length int) {
	C.memmove(dest, src, C.size_t(length))
}

func MemCopy(dest unsafe.Pointer, src []byte, length int) {
	srcData := C.CBytes(src)
	C.memcpy(dest, srcData, C.size_t(length))
}
