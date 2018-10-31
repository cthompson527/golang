package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Body struct {
	Args []int `json:"args"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, fmt.Sprintf("could not read body: %v", err))
		return
	}

	var body Body
	err = json.Unmarshal(b, &body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, fmt.Sprintf("could not parse body: %v", err))
		return
	}

	args := body.Args

	sum := Sum(args...)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, fmt.Sprintf(`{ "sum": %d }`, sum))
}

func Sum(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}
