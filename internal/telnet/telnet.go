package telnet

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/Komilov31/telnet/internal/flags"
)

type Telnet struct {
	flags  *flags.Flags
	dialer net.Dialer
}

func New(flags *flags.Flags) *Telnet {
	return &Telnet{
		flags: flags,
		dialer: net.Dialer{
			Timeout: time.Duration(flags.Timeout) * time.Second,
		},
	}
}

func (t *Telnet) ProcessProgram() {
	address := t.flags.Host + ":" + t.flags.Port
	fmt.Println(address)
	conn, err := t.dialer.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal("could not close connection: ", err)
		}
	}()

	//запускаем горутину для чтения из STDIN из записи в канал
	lines := make(chan []byte)
	go readInput(lines, conn)

	// запускаем горутину которая читает из канала и отправки в сокет
	// также выводит результат в STDOUT
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go writeOutput(wg, conn, lines)

	wg.Wait()
}
