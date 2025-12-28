package main

import (
	"fmt"
	"obuchenie/feature2"
	"obuchenie/feature3"
	simpleconection "obuchenie/feauture4/simple_conection"
	"obuchenie/ffeature"
)

func main() {
	fmt.Println("Hello git")
	ffeature.Featture1()
	feature2.Feature2()
	feature3.Feature3("привет как дела", 5)
	simpleconection.CheckConection()
}
