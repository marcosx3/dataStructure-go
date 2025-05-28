package palindromo

// func main() {

// 	res := isPalindromo("arara")
// 	fmt.Println(res)
// }

func isPalindromo(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		letterStart := s[start]
		letterEnd := s[end]
		if letterStart != letterEnd {
			return false
		}
		start++
		end--
	}
	return true
}
