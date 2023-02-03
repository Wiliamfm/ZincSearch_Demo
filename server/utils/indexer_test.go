package indexer

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestDocumentCreation(t *testing.T) {
	data := struct {
		Name string
		Age  int
	}{
		"asdfasdf",
		23,
	}
	body, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	loginData := struct {
		_id      string
		password string
	}{
		"admin",
		"Complexpass#123",
	}
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:4080/api/test/_doc", bytes.NewBuffer(body))
	t.Log(bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth(loginData._id, loginData.password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatal("Error in response", resp.Status)
	}
}
