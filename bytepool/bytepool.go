package bytepool

type BytePoolCap struct {
	pool chan []byte // []byte pool，通过 maxSize 来指定池中可存放 []byte 对象的数量，但即使池中空了，也可以获取
	w    int         // 每个 []byte 对象的初始 width，及默认起始位置。建议为 0，即从 0 开始使用
	wcap int         // 每个 []byte 对象的容量。参考值，但必须大于 w，对象容量不足时，runtime 会自动扩容
}

// maxSize 过高时，会浪费内存
// maxSize 过低时，会频繁创建 []byte，失去 pool 的意义
// width 建议为 0，表示从 0 开始使用，回收时也能置零
// capwidth 建议根据需要指定，避免 runtime 频繁的扩容
func NewBytePoolCap(maxSize int, width int, capwidth int) (bp *BytePoolCap) {
	return &BytePoolCap{
		pool: make(chan []byte, maxSize),
		w:    width,
		wcap: capwidth,
	}
}

// 获取对象
func (bp *BytePoolCap) Get() (b []byte) {
	select {
	case b = <-bp.pool: // 池中有对象时，复用池中的 []byte
	default: // 池中没有时，新建一个 []byte
		if bp.wcap > 0 {
			b = make([]byte, bp.w, bp.wcap)
		} else {
			b = make([]byte, bp.w)
		}
	}
	return
}

// 回收对象
func (bp *BytePoolCap) Put(b []byte) {
	// 如果传入的 []byte width 小于初始 width，就放弃回收此对象
	if cap(b) < bp.w {
		return
	}

	// 如果传入的 []byte 容量远超初始容量，就放弃回收此对象
	// 避免偶尔的 []byte 使用时容量扩的极大，回收后会导致 pool 中对象逐渐都是大容量，需要在时间和空间上做好平衡
	// if cap(b) > 10*bp.wcap {
	// 	return
	// }

	select {
	case bp.pool <- b[:bp.w]: // 如果传入的 []byte width 大于等于初始 width，就回收，并指定 width 为默认起始位置
	default: // 除上述情况以外时，不回收
	}
}

// Width returns the width of the byte arrays in this pool.
func (bp *BytePoolCap) Width() (n int) {
	return bp.w
}

// 返回 pool 中可用对象的数量
func (bp *BytePoolCap) Len() (n int) {
	return len(bp.pool)
}
