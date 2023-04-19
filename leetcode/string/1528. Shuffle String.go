package main

func restoreString(s string, indices []int) string {

	res := make([]byte, len(s))
	for i := 0; i < len(s); i++ {

		res[indices[i]] = s[i]

	}
	return string(res)
}

func main() {
	restoreString("codeleet", []int{4,5,6,7,0,2,1,3})
}
