package main

// pgm à tester avec l'utilitaire "netcat" que j'ai aliasé avec nc
// donc : nc localhost 4000 si la foncion de Fprintln est utilisé
// en utilisant la libairie "io", tu peux aussi utiliser io.Copy(c,c) pour faire un echo !
import (
	"fmt"
	"log"
	"net"
)

const listenAddr = "localhost:4000"

func main() {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(c, "Hello World")
		c.Close()
	}
}
