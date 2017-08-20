package main

import (
	"fmt"
	"github.com/megawubs/remote"
	"strconv"
	"bufio"
	"os"
)

func main(){
	done := make(chan int)
	fmt.Print("Choose action:\n[0:Up, 1:Down, 2:Power, 3:Volume Up, 4: Volume Down]\n")
	go forever()
	<-done
}

func forever() {
	for{
		reader := bufio.NewReader(os.Stdin)

		text, _ := reader.ReadString('\n')
		res, _ := strconv.ParseInt(text[:1], 10, 64)
		action := remote.Action(res)
		var command = remote.Command{Action:action}
		remote.Send(command)
	}
}