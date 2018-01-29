package main

import (
	"os"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	os.Stderr.WriteString("CRIMINAL LOOKUP HISTORY FOR CAR: <id>")
}
