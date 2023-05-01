package main

import "fmt"

func main() {
	m := map[string]int{"partha": 3, "homi": 2, "soumya": 1}
	n := make(map[string]int)
	fmt.Println(m)
	arr := []string{}
	for key := range m {
		fmt.Println(key)

		arr = append(arr, key)
	}
	for _, k := range arr {
		n[k] = m[k]
	}
	fmt.Println(n)
}
