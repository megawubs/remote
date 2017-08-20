package remote

import (
	"net/http"
	"encoding/json"
	"bytes"
)

func Send(command Command)  {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(command)

	http.Post("http://localhost:8000/", "application/json; charset=utf-8", b)
}