package main
import(
	"fmt"
)

func main(){
	map1 := make(map[string]string)
	map1["name"] = "partha"
	map1["roll no"] = "4"
	map1["email"]= "pgiri@gmail.com"
	fmt.Printf("map1:%v",map1)
	fmt.Println(len(map1))
	delete(map1,"name")
	fmt.Printf("map2:%v",map1)
//check the key exists in map1 
if value,ok:=map1["email"];ok{
	fmt.Println("value-->",value)
}
}