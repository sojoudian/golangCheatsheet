package main

import (
	"fmt"
)

// d:=<-chanelection  // reading the data from the channel
// channel <- a // write the data to the channel
//send and receive to the channel is blocked by default,
// the data is sent to a channel the control is blocked in the
//same statement,
//untill some other go routine reads from that channel
//similiarly when data is read from a channelt the read is blocked
//blocked until some go routines write data to that channel
func fOne(d chan bool) {
	fmt.Println("funcOne")
	d <- true
}

func main() {
	d := make(chan bool)
	go fOne(d)
	<-d // the control will not move to the next line untill the write process is done
	//so we can remove sleep
	// time.Sleep(1 * time.Second)
	fmt.Println("main()")
}
