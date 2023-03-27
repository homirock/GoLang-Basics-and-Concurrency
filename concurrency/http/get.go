package main
import (
	"fmt"
	"sync"
	"io/ioutil"
	"net/http"
)

func getApiCall(v string, wg *sync.WaitGroup){
	resp,_:=http.Get(v)
	defer resp.Body.Close()
	body, _:= ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	url := []string{
		"https://httpbin.org/get",
		"http://api.github.com/users/hadley/orgs",
	}

	for _,v:=range url{
		wg.Add(1)
		go getApiCall(v,&wg)
	}
	wg.Wait()
}