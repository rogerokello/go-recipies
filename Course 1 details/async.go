package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(8)

	// make a channel. NB: Channels take in a type of
	// object. In this case this channel will only
	// take in strings
	ch := make(chan string)

	donech := make(chan bool)

	go abcGen(ch)
	go printer(ch, donech)

	// Channels stop execution of the function that they are in
	<-donech
}

func abcGen(ch chan string) {
	for l := byte('a'); l <= byte('z'); l++ {
		ch <- string(l)
	}
	close(ch)
}

func printer(ch chan string, donech chan bool) {
	for l := range ch {
		println(l)
	}

	donech <- true
}
