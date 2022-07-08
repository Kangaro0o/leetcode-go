package iota

const (
	bit0, mask0 = 1 << iota, 1<<iota - 1
	bit1, mask1 = 1 << iota, 1<<iota - 1
	_, _
	bit3, mask3
	// 空行并不会使得 iota 增大
	bit4, mask4
)
