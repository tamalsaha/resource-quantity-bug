package main

import (
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/api/resource"
)

func main() {
	q1, err := resource.ParseQuantity("1Gi")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(q1.String())

	q2 := resource.NewQuantity(1024*1024*1024, resource.BinarySI)
	fmt.Println(q2.String())

	q3 := resource.NewBinaryQuantity(1, 9)
	fmt.Println(q3.String())
}
