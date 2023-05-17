package main

//exercise3.12  判断两个字符串是否异构
func main() {
	var str1, str2 string = "abcdf", "fdcba"

	record := make(map[byte]int)
	var temp bool = true
	for i := 0; i < len(str1); i++ {
		record[str1[i]]++
		record[str2[i]]--
	}

	for _, i := range record {
		if i > 0 {
			temp = false
		}
	}
	println(temp)
}
