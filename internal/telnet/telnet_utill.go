package telnet

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

// чтение из stdin и отправка в канал
func readInput(lines chan<- []byte, conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)

	for {
		nextLine, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			log.Fatal("could not read from stdin: ", err)
		}

		if err == io.EOF {
			close(lines)
			break
		}

		lines <- nextLine
	}
}

func writeOutput(wg *sync.WaitGroup, conn net.Conn, lines <-chan []byte) {
	defer wg.Done()

	for line := range lines {
		_, err := conn.Write(line)
		if err != nil {
			log.Println("could not send meesage to server:", err)
			return
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("could not read answer from server:", err)
			return
		}

		fmt.Print(string(buf[:n]))
	}
}
