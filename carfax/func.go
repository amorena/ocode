package main

import (
	"os"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	os.Stderr.WriteString("PULLING CARFAX REPORT FOR CAR: <id>")
}
