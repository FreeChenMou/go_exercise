package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//exercise2.3 循环右移x，根据计数数组累加位值
func PopCountRange(x uint64) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

//exercise2.4 遍历64个位
func PopCount64(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int((x >> i) & 1)
	}
	return count
}

//exercise2.5 根据&运算从右到左消1并计数
func PopCountAnd(x uint64) int {
	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}
