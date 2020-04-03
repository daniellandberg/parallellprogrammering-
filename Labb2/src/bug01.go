package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
//felet i tidigare kod var att det blev en deadlock. Dvs programmet visst inte hur den
//skulle skicka infon i channeln, den fick inget slut. Löste detta genom en go rutin som
//får tillbaka strängen hello world i channeln och då går den att printa.
func hej(ch chan string){
	ch <- "Hello world!"
}

func main() {
	ch := make(chan string)
	go hej(ch)
	fmt.Println(<-ch)

}
