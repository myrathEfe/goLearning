// channels
package main

import (
	"fmt"
	"time"
)

/*func printMessage(msg string) {
	for i := 0; i <= 10; i++ {
		fmt.Println(msg, " ", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func sayHello(ch chan string) {
	ch <- "Merhaba"
}*/

func main() {

	for i := 0; i < 100; i++ {
		ch1 := make(chan string)
		ch2 := make(chan string)
		ch3 := make(chan string)

		go func() {
			time.Sleep(time.Second / 1000)
			ch1 <- "Channel 1"
		}()

		go func() {
			time.Sleep(time.Second / 1000)
			ch2 <- "Channel 2"
		}()

		go func() {
			time.Sleep(time.Second / 1000)
			ch3 <- "Channel 3"
		}()

		select {
		case w1 := <-ch1:
			fmt.Println("Kazanan : ", w1)
		case w2 := <-ch2:
			fmt.Println("Kazanan : ", w2)
		case w3 := <-ch3:
			fmt.Println("Kazanan : ", w3)
		}
		/*go printMessage("Go Routined Message")

		for i := 0; i <= 10; i++ {
			fmt.Println("Normal Message ", i)
			time.Sleep(time.Millisecond * 500)
		}

		channel := make(chan string)
		go sayHello(channel)
		msg := <-channel
		fmt.Println(msg)*/

	}
}
