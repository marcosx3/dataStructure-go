package array

func InvertArray(array []int) []int {
	start := 0
	end := len(array) - 1
	for start < end {
		array[start], array[end] = array[end], array[start]
		start++
		end--
	}
	return array
}

// func main() {
// 	arr := []int{10, 20, 30, 40, 55}

// 	resultado := InvertArray(arr)
// 	fmt.Println(resultado) // [55 40 30 20 10]
// }
