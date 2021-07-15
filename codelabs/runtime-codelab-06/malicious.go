package main

import "fmt"
import "time"

func main() {
	fmt.Println("malicious app")
	time.Sleep(1 * time.Hour)
}
