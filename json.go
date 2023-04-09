// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Emp struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func main() {
	b := &Emp{Name: "homi", Id: "1"}
	c, err := json.Marshal(b)
	if err != nil {
		errors.New("error found")
	}
	fmt.Println(string(c))
	f := string(c)
	var d Emp
	err = json.Unmarshal([]byte(f), &d)
	if err != nil {
		errors.New("error found")
	}
	fmt.Println(d)

	ma := map[string]string{"partha": "1", "sarathi": "2"}
	jm, err := json.Marshal(ma)
	if err != nil {
		errors.New("error found")
	}
	fmt.Println(string(jm))
}
