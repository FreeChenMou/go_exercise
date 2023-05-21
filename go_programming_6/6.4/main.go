package main

type IntSet struct {
	words []uint64
}

//exercise6.4 返回slice
func (s *IntSet) Elems() []uint64 {
	var elems []uint64
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint64(j)) != 0 {
				elems = append(elems, uint64(i*64+j))
			}
		}
	}
	return elems
}

func main() {

}
