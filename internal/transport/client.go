package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ReadRequest struct {
	Session string `json:"session"`
}

type NoteResponse struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Session string `json:"session"`
	EditID  string `json:"editId"`
}

func UpdateTitle(noteID string, sessionToken string, editID string, hexPaylod string) error {
	client := &http.Client{Timeout: 5 * time.Second}

	reqData := UpdateRequest{
		Title:   hexPaylod,
		Content: "",
		Session: sessionToken,
		EditID:  editID,
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return err
	}

	url := "https://hypernotepad.com/api/notes/" + noteID + "/update"

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server rejected the update, http status %d", resp.StatusCode)
	}
	return nil
}

func UpdateContent(noteID string, sessionToken string, editID string, hexContent string) error {
	client := &http.Client{Timeout: 5 * time.Second}

	reqData := UpdateRequest{
		Title:   "",
		Content: hexContent,
		Session: sessionToken,
		EditID:  editID,
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return err
	}

	url := "https://hypernotepad.com/api/notes/" + noteID + "/update"

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server rejected the update, http status %d", resp.StatusCode)
	}
	return nil
}

func FetchCommand(noteID string, sessionToken string) (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}

	reqData := ReadRequest{Session: sessionToken}
	jsonData, err := json.Marshal(reqData)

	if err != nil {
		return "", err
	}

	url := "https://hypernotepad.com/api/notes/share/" + noteID
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var note NoteResponse
	if err := json.Unmarshal(body, &note); err != nil {
		return "", err
	}
	return note.Title, nil
}

func FetchResponse(noteID string, sessionToken string) (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}

	reqData := ReadRequest{Session: sessionToken}
	jsonData, err := json.Marshal(reqData)

	if err != nil {
		return "", err
	}

	url := "https://hypernotepad.com/api/notes/share/" + noteID
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var note NoteResponse
	if err := json.Unmarshal(body, &note); err != nil {
		return "", err
	}
	return note.Content, nil
}
