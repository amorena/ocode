package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	time.Sleep(3 * time.Second)
	fmt.Println(os.Stdin)
	os.Stderr.WriteString("PROCESSING CAR <ID>")
}
