package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

type Deployment struct {
	Deployment string `json:"deployment,omitempty"`
}

var deployment []Deployment

func CreateOperators(w http.ResponseWriter, req *http.Request) {
	var (
		cmdOut     []byte
		err        error
		deployment Deployment
	)

	err = json.NewDecoder(req.Body).Decode(&deployment)
	if err != nil {
		panic(err)
	}

	if deployment.Deployment == "prometheus" {
		cmdName := "kubectl"
		cmdArgs := []string{"apply", "-f", "00namespace-namespace.yaml"}
		cmdExec := exec.Command(cmdName, cmdArgs...)
		cmdExec.Dir = "./manifests"
		if cmdOut, err = cmdExec.Output(); err != nil {
			fmt.Fprintf(w, string("There was an error running the kubectl create: "+err.Error()))
		}
		log.Println(deployment.Deployment, "operator created")
		fmt.Fprintf(w, string(cmdOut))

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "deployment with this name does'nt exist, valid ones are prometheus and anchor")
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(deployment); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("HTTP server listening 8081")
	router := mux.NewRouter()
	router.HandleFunc("/deploy", CreateOperators).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))

}
