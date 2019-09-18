package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/k8s-api/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	namespace = "monitoring"
)

var err error

func PutSvcProm(w http.ResponseWriter, req *http.Request) {

	clientset := utils.GetKubeClientset()

	svcs, err := clientset.CoreV1().Services(namespace).List(metav1.ListOptions{
		LabelSelector: "app=grafana",
	})
	if err != nil {
		log.Fatalln("failed to get svc: ", err)
	}
	// print svc
	for i, svc := range svcs.Items {
		service := svc.Status.LoadBalancer.Ingress
		fmt.Fprintf(w, "Access the Grafana on the following url\n")
		json, err := json.Marshal(service)
		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest(http.MethodPut, "", bytes.NewBuffer(json))
		/*
			if err != nil {
				panic(err)
			}
		*/
		req.Header.Set("Content-Type", "application/json; charset=utf-8")

		log.Printf("Grafana Service Name: %d Service Name: %+v\n", i, service)
	}

}

func PutSvcAnc(w http.ResponseWriter, req *http.Request) {

	clientset := utils.GetKubeClientset()

	svcs, err := clientset.CoreV1().Services("default").List(metav1.ListOptions{
		LabelSelector: "app=anchore",
	})
	if err != nil {
		log.Fatalln("failed to get svc: ", err)
	}
	// print svc
	for i, svc := range svcs.Items {
		service := svc.Status.LoadBalancer.Ingress
		fmt.Fprintf(w, "Access the Anchore on the following url\n")
		json.NewEncoder(w).Encode(service)
		json, err := json.Marshal(service)
		if err != nil {
			panic(err)
		}
		req, err := http.NewRequest(http.MethodPut, "", bytes.NewBuffer(json))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		log.Printf("Anchore Service Name: %d Service Name: %+v\n", i, service)
	}

	//json.NewEncoder(w).Encode(svc)

}
