package servicepods

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

func PodHealthCheck(name string, namespace string) string {
	client, err := api.NewClient(api.Config{
		Address: "http://192.168.10.200:32001",
	})
	if err != nil {
		log.Fatal(err)
	}

	prometheusClient := v1.NewAPI(client)

	// Check if the "nginx2" pod in the "default" namespace is ready
	query := fmt.Sprintf("kube_pod_container_status_ready{namespace='%s', pod='%s'}", namespace, name)
	value, warnings, err := prometheusClient.Query(context.Background(), query, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	if len(warnings) > 0 {
		log.Println("Warnings:", warnings)
	}
	vector, ok := value.(model.Vector)
	if !ok {
		log.Fatal("Unexpected query result type")
	}
	if len(vector) == 0 {
		return "Unhealthy"
	}

	return "Healthy"
}
