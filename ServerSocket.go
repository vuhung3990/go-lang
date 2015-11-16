package main
import (
	"fmt"
	"net"
	"bufio"
	"strings"
)
func main() {
	fmt.Println("start server.")

	// listen on all interface
	ln, _ := net.Listen("tcp", ":8881")

	// accept connection on port
	conn, _ := ln.Accept()

	for  {
		// will listen for message to process ending newline \n.
		message,_ := bufio.NewReader(conn).ReadString('\n')

		// output received mesage
		fmt.Println("received: ",message)

		// send mesage to client
		// convert string received to title then parse to bytes array
		conn.Write([]byte(strings.ToTitle(message+"\n")))
	}
}
