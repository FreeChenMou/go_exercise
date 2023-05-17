package main

const (
	_ = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	println(KB, MB, GB, TB, PB, EB, ZB, YB)
}
