package main

import (
	"fmt"
	"io"
	"log"

	sia "github.com/pietern/go-sia"
)

func Handle(r *sia.Reader, w *sia.Writer) {
	for {
		// Read whatever
		block, err := r.Read()
		if err != nil {
			if err == io.EOF {
				log.Printf("Connection closed")
				return
			}
			log.Fatalf("Error reading block: %s", err)
		}

		log.Printf(
			"Read block with function 0x%x, data \"%v\"",
			block.Function,
			string(block.Data))

		// Send acknowledgement
		ack := sia.Block{0x38, nil}
		err = w.Write(ack)
		if err != nil {
			log.Fatalf("Error writing block: %s", err)
		}
	}
}

func main() {
	port := 10002
	log.Printf("Listening on port %d", port)
	sia.Listen("tcp", fmt.Sprintf(":%d", port), Handle)
}
