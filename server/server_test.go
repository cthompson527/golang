package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"dev.azure.com/rchi-texas/Golang/server"
)

var tests = []struct {
	args string
	sum  int
}{
	{"[1, 2, 3]", 6},
	{"[1, 2, 3, 4]", 10},
	{"[1, 2, 3, 4, 5]", 15},
	{"[1024, 2048, 4096, 8192]", 15360},
}

type response struct {
	Sum int `json:"sum"`
}

func TestHealthCheckHandler(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			sum, err := makeRequest(tt.args)
			if err != nil {
				t.Fatal(err)
			} else if sum != tt.sum {
				t.Fatalf("invalid sum results. expected %d but got %d", tt.sum, sum)
			}
		})
	}

}

func makeRequest(args string) (int, error) {
	body := []byte(fmt.Sprintf(`{ "args": %s }`, args))

	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
	if err != nil {
		return 0, err
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		return 0, fmt.Errorf("handler returned status code %d", status)
	}

	b, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading body: %v", err)
	}

	var resp response
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return 0, fmt.Errorf("error parsing body: %v", err)
	}

	return resp.Sum, nil
}
