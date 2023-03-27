package main
import(
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"bytes"
)
type data struct{
	name string `json:"name"`
	email string `json:"email"`
}

// func postFunc(){

// }

func main(){
	url:="https://postman-echo.com/post"
	//send this json data
	// {
	// "name":  "Toby",
    //   "email": "Toby@example.com",
	// }
	// data:= map[string]string{

	// }
	data1:=&data{
		name: "Partha",
		email:"parthsarathigiri@gmail.com",
	}

	data,_:=json.Marshal(data1)
	fmt.Println("jason marshall data:",data)
	resp,_:=http.Post(url,"application/json",bytes.NewBuffer(data))
	defer resp.Body.Close()
	respBody,_:= ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))

}