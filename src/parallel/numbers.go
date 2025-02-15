package main

func SumFromSlice(sl []int) int64 {
	var sum int64 = 0
	for i := 0; i < len(sl); i++ {
		sum += int64(sl[i])
	}
	return sum
}
