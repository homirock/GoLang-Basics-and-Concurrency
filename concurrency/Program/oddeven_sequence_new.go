package main

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	status := true
	for i := 0; i < 10; i++ {
		val <- ch
		if val%2 == 0 && status {
			fmt.Println("even:", val)
			status = false
		} else {
			fmt.Println("odd:", val)
			status = true
		}
	}
}
