package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/k8s-api/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	namespace = "monitoring"
)

var err error

func GetSvc(w http.ResponseWriter, req *http.Request) {

	clientset := utils.GetKubeClientset()

	svc, err := clientset.CoreV1().Services(namespace).List(metav1.ListOptions{
		LabelSelector: "app=grafana",
	})
	if err != nil {
		log.Fatalln("failed to get svc: ", err)
	}
	// print svc
	json.NewEncoder(w).Encode(svc)

}
