package main

import (
	"fmt"
	"time"
)

type logWriter struct{}

func (w logWriter) Write(b []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05") + " " + string(b))
}
