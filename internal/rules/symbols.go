package rules

func InsertSymbols(inputText string) string {
	// r := []rune(inputText)
	// for i := range r {
	// 	if r[i] == '!' || r[i] == '.' || r[i] == ',' {
	// 		return string(append(r[:i+1], append([]rune(" *UWU*"), r[i+1:]...)...))
	// 	}
	// }
	return inputText + " *UWU*"
}
