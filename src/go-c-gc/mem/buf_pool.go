package mem

import (
	"errors"
	"fmt"
	"sync"
)

// 内存池类型
type Pool map[int]*Buf

type BufPool struct {
	Pool     Pool
	PoolLock sync.RWMutex

	// 总buffer池大小 单位KB
	TotalMem uint64
}

var bufPoolInstance *BufPool
var once sync.Once

// 一个池有多个Buflist的大小组
const (
	m4K   int = 4096
	m16K  int = 16384
	m64K  int = 65535
	m256K int = 262144
	m1M   int = 1048576
	m4M   int = 4194304
	m8M   int = 8388608
)

const (
	EXTRA_MEM_LIMIT int = 5 * 1024 * 1024 // 池最大大小 5GB
)

func MemPool() *BufPool {
	once.Do(func() {
		bufPoolInstance = new(BufPool)
		bufPoolInstance.Pool = make(Pool)
		bufPoolInstance.TotalMem = 0
		bufPoolInstance.initPool()
	})
	return bufPoolInstance
}

func (bp *BufPool) initPool() {
	// ----> 4KB
	// 5000个Buf 约20M
	bp.makeBufList(m4K, 5000)
	// ----> 16KB
	// 1000个Buf 约16M
	bp.makeBufList(m16K, 1000)
	// ----> 64KB
	// 500个Buf 约32M
	bp.makeBufList(m64K, 500)
	// ----> 256KB
	// 200个Buf 约50M
	bp.makeBufList(m256K, 200)
	// ----> 1MB
	// 50个Buf 约50M
	bp.makeBufList(m1M, 50)
	// ----> 4MB
	// 20个Buf 约80M
	bp.makeBufList(m4M, 20)
	// ----> 8MB
	// 10个Buf 约80M
	bp.makeBufList(m8M, 10)
}

func (bp *BufPool) makeBufList(cap int, num int) {
	bp.Pool[cap] = NewBuf(cap)
	var prev = bp.Pool[cap]
	for i := 0; i < num; i++ {
		prev.Next = NewBuf(cap)
		prev = prev.Next
	}
	bp.TotalMem += (uint64(cap) / 1024) * uint64(num)
}

// 从池种申请一块内存
func (bp *BufPool) Alloc(N int) (*Buf, error) {
	//如果上层需要N字节大小的空间，找到与N最接近的Buf链表集合，从当前Buf集合取出。
	var index int
	if N <= m4K {
		index = m4K
	} else if N <= m16K {
		index = m16K
	} else if N <= m64K {
		index = m64K
	} else if N <= m256K {
		index = m256K
	} else if N <= m1M {
		index = m1M
	} else if N <= m4M {
		index = m4M
	} else if N <= m8M {
		index = m8M
	} else {
		return nil, errors.New("Alloc size Too Large!")
	}

	//如果该组已经没有节点可供使用，则可以额外申请总申请长度不能够超过最大的限制大小EXTRA_MEM_LIMIT。
	bp.PoolLock.Lock()
	if bp.Pool[index] == nil {
		if (bp.TotalMem + uint64(index/1024)) >= uint64(EXTRA_MEM_LIMIT) {
			errStr := fmt.Sprintf("already use too many memory!\n!")
			return nil, errors.New(errStr)
		}
		newBuf := NewBuf(index)
		bp.TotalMem += uint64(index / 1024)
		bp.PoolLock.Unlock()
		fmt.Printf("Alloc Mem Size: %d KB\n", newBuf.Capacity/1024)
		return newBuf, nil
	}

	//如果有该节点需要的内存块，则直接取出，并且将该内存块从BufPool移除
	targetBuf := bp.Pool[index]
	bp.Pool[index] = targetBuf.Next
	bp.TotalMem -= uint64(index / 1024)
	bp.PoolLock.Unlock()
	targetBuf.Next = nil
	fmt.Printf("Alloc Mem Size: %d KB\n", targetBuf.Capacity/1024)
	return targetBuf, nil
}

// 归还一块内存
func (bp *BufPool) Revert(buf *Buf) error {
	index := buf.Capacity
	buf.Clear()

	bp.PoolLock.Lock()

	if _, ok := bp.Pool[index]; !ok {
		errStr := fmt.Sprintf("Index %d not in BufPool!\n", index)
		return errors.New(errStr)
	}

	buf.Next = bp.Pool[index]
	bp.Pool[index] = buf
	bp.TotalMem += uint64(index / 1024)
	bp.PoolLock.Unlock()
	fmt.Printf("Revert Mem Size: %d KB\n", bp.TotalMem/1024)
	return nil
}
