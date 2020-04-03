package main

import (
	"fmt"
	"sync"
	"time"
)

// This program should go to 11, but it seemingly only prints 1 to 10.
// felet i programmet var att den nedre funktionen inte han klart innan main blev klar i sista loopen vilket gjorde
// att sista numret inte skrevs ut. Jag löste detta genom att fördröja main funktionen med lika mycket tid som
// print funktionen men man skulle även kunnat använda waitgroup funktionen vilket jag i efterhand ser att det hade
//varit snyggare, men det nedan löser också problemet.
func main() {
	ch := make(chan int)
	wf := new (sync.WaitGroup)
	go Print(ch, wf)
	wf.Add(1)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
	wf.Wait()

	//flyttade upp en förskjutning av close så att den hinner köra
	//nedanstående for loop och printa 11 innan den closar.
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, wf *sync.WaitGroup ) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
	wf.Done()
}
