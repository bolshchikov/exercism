package reverse

//String should reverse given string
func String(str string) (res string) {
	for _, v := range str {
		res = string(v) + res
	}
	return
}
