package main

import(

	"time"
)

func main() {
    go sleepyGopher() // Начало горутины
    time.Sleep(4 * time.Second) // Ожидание
}

func sleepyGopher() {
	panic("unimplemented")
} // Здесь все горутины останавливаются 