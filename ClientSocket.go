package main
import (
	"net"
	"bufio"
	"os"
	"fmt"
)
func main() {
	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8881")
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		// send to socket server
		conn.Write([]byte(text+"\n"))

		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}