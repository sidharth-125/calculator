package app

import (
	"calculator/consts"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
Procedure

1. Perform basic calculations
2. Advanced calculations
3. Advanced arithmetic operations
4. Deploy to a host
5. Prepare a UI
*/
func Run() {
	log.Println("starting the server on port: " + consts.Host)
	m := http.NewServeMux()
	m.HandleFunc("/calc", ServeCalc)
	http.ListenAndServe(consts.Host, m)
}

func ServeCalc(w http.ResponseWriter, r *http.Request) {
	var p Payload

	byteArr, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something Went Wrong"))
		return

	}

	err = json.Unmarshal(byteArr, &p)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something Went Wrong"))
		return
	}

	fmt.Println(p.Expression)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

type Payload struct {
	Expression string `json:"exp"`
}
