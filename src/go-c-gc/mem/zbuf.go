package mem

import "fmt"

type ZBuf struct {
	b *Buf
}

func (zb *ZBuf) Clear() {
	if zb.b != nil {
		err := MemPool().Revert(zb.b)
		if err != nil {
			return
		}
		zb.b = nil
	}
}

func (zb *ZBuf) Pop(len int) {
	if zb.b == nil || len > zb.b.length {
		return
	}
	zb.b.Pop(len)

	if zb.b.Length() == 0 {
		err := MemPool().Revert(zb.b)
		if err != nil {
			return
		}
		zb.b = nil
	}
}

func (zb *ZBuf) Data() []byte {
	if zb.b == nil {
		return nil
	}
	return zb.b.GetBytes()
}

func (zb *ZBuf) Adjust() {
	if zb.b == nil {
		zb.b.Adjust()
	}
}

func (zb *ZBuf) Read(src []byte) (err error) {
	if zb.b == nil {
		zb.b, err = MemPool().Alloc(len(src))
		if err != nil {
			fmt.Println("pool Alloc Error ", err)
		}
	} else {
		if zb.b.Head() != 0 {
			return nil
		}
		if zb.b.Capacity-zb.b.Length() < len(src) {
			newBuf, err := MemPool().Alloc(len(src) + zb.b.Length())
			if err != nil {
				return err
			}

			newBuf.Copy(zb.b)
			MemPool().Revert(zb.b)
			zb.b = newBuf
		}
	}

	zb.b.SetBytes(src)

	return nil
}
