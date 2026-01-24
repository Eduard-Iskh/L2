package main

import (
	"fmt"
	"time"

	myntp "L2.8/pkg"
)

func main() {
	currentTime, err := myntp.CurrentTime("")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	fmt.Println("Задержка между серверами:", currentTime.Precision)
	fmt.Printf("NTP время: %v\n", currentTime.Time.Format(time.RFC3339))
	fmt.Printf("Локальное время: %v\n", time.Now().Format(time.RFC3339))
}
