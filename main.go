package main

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"github.com/rothink/imersaofsfc2-simulator/infra/kafka"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file")
	}
}

func main() {

	producer := kafka.NewKafkaProducer()
	kafka.Publish("Ola", "readTest", producer)

	//route := route2.Route{
	//	ID:       "1",
	//	ClientID: "1",
	//}
	//
	//route.LoadPositions()
	//stringjson, _ := route.ExportJsonPositions()
	//
	//fmt.Println(stringjson[0])
}
