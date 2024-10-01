package app

import (
	"calculator/consts"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/Knetic/govaluate"
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
		w.Write([]byte(consts.StatusInternalServerError))
		return

	}

	err = json.Unmarshal(byteArr, &p)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(consts.StatusInternalServerError))
		return
	}

	_, err = Valuate(p.Expression)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(consts.StatusInternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

type Payload struct {
	Expression string `json:"exp"`
}

func Valuate(exp string) (float64, error) {
	var (
		answer float64
		ok     bool
	)

	arithExp, err := govaluate.NewEvaluableExpression(exp)
	if err != nil {
		log.Println("make err: ", err)
		return 0, err
	}

	result, err := arithExp.Evaluate(nil)
	if err != nil {
		log.Println("evaluate err: ", err)
		return 0, err
	}

	if answer, ok = result.(float64); !ok {
		log.Printf("cannot convert the answer for expression :%s", exp)
		return 0, errors.New("cannot convert the answer for expression :" + exp)
	}

	return answer, nil

}
