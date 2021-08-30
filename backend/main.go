package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
)

type inputData struct {
	Expression string
}
type expression struct {
	firstOperand  float64
	operator      string
	secondOperand float64
}

var validExpression = regexp.MustCompile(`(\d+\.?d*)([+|-|\*|\/])(\d+\.?d*)`)

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/evaluate", PostHandler).Methods("POST")
	log.Println("listing 5000")
	http.ListenAndServe(":5000", r)
}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	var input inputData
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(input)
	var expressionToEvaluate expression
	m := validExpression.FindStringSubmatch(input.Expression)
	log.Println(m)
	if m == nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Incorrect Expression")
		return
	} else {
		firstOperand, _ := strconv.ParseFloat(m[1], 64)
		secondOperand, _ := strconv.ParseFloat(m[3], 64)
		operator := m[2]
		expressionToEvaluate = expression{
			firstOperand:  firstOperand,
			operator:      operator,
			secondOperand: secondOperand,
		}
	}

	answer, err := evaluate(expressionToEvaluate)

	w.Header().Set("Content-Type", "application/text")

	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%g", answer)))
	}
}

func evaluate(exp expression) (float64, error) {
	var answer float64
	var err error = nil
	switch exp.operator {
	case "+":
		answer = float64(exp.firstOperand) + float64(exp.secondOperand)

	case "-":
		answer = float64(exp.firstOperand) - float64(exp.secondOperand)

	case "*":
		answer = float64(exp.firstOperand) * float64(exp.secondOperand)

	case "/":
		if exp.secondOperand == 0 {
			err = fmt.Errorf("Can not divide to 0")
			answer = 0
		} else {
			answer = float64(exp.firstOperand) / float64(exp.secondOperand)
		}
	}

	return answer, err
}
