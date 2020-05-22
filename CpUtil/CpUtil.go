package CpUtil

// 求解器常用函数和常量

const BITSIZE = 64
const DIVBIT = 6
const MODMASK = 0x3f
const INDEXOVERFLOW = -1
const ALLONELONG = 0xFFFFFFFFFFFFFFFF
const INTMAXINF = 0x3f3f3f3f
const INTMININF = -0x3f3f3f3f
const LONGMAXINF = 0x3f3f3f3f3f3f3f3f
const LONGMININF = -0x3f3f3f3f3f3f3f3f

type Int2d struct {
	X int
	Y int
}

//
func GetInt2d(a int) Int2d {
	return Int2d{a >> DIVBIT, a & MODMASK}
}

func GetInt2(a int) (int, int) {
	return a >> DIVBIT, a & MODMASK
}

func GetIndex(x, y int) int {
	return (x << DIVBIT) + y
}

// 十亿级前缀标识常量
// 30位值部，2位标识部
const VALUEPARTBITLENGTH = 30
const PERFIXMASK = 0x40000000 // 十亿级
const SUFFIXMASK = 0x3fffffff // 十亿级

// 给数据添加前缀
func markValue(a uint) uint {
	return a | PERFIXMASK
}

// 去掉数据据前缀
func demarkValue(a uint) uint {
	return a & SUFFIXMASK
}

// 解析数据: (数据类型,数据值)
func resolveMark(a uint) (uint, uint) {
	return a >> VALUEPARTBITLENGTH, a & SUFFIXMASK
}

// 解析数据: (数据bool类型,数据值)
func resolveMarkBoolean(a uint) (bool, uint) {
	return a >= PERFIXMASK, a & SUFFIXMASK
}
