package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	for _, v := range input {
		output = string(v) + output
	}
	return output
}

//
//func main() {
//	t3 := "Hello world!"
//	res := ReverseString(t3)
//	fmt.Println(res)
//}

//
//input: Hello world!
//
//output: !dlrow olleH
