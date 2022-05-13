package main

import "github.com/hasanbakirci/mongodb-kafka-connector/cmd"

func main() {
	//cmd.StartApi()
	forever := make(chan bool)

	l := cmd.NewListener()
	l.StartListener()

	<-forever

}
