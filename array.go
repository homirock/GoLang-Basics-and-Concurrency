package main
import "fmt"
func main() {
	//1D array
	var arr [5]int
	fmt.Println("array:",arr)
	fmt.Println("length of array:",len(arr))
	arr[4],arr[3] = 5 ,4
	fmt.Println("insert:",arr)
	fmt.Println("length of array:",len(arr))
	b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

	//2D array
	var twod [2][3]int
	for i:=0;i<2;i++{
		for j:=0; j<3;j++{
			twod[i][j] = i+j
		}
	}
	fmt.Println("two dimension:",twod)
}