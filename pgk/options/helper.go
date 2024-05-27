package options

// 定义单位常量
const (
	_   = iota
	KiB = 1 << (10 * iota)
	MiB
	GiB
	TiB
)
