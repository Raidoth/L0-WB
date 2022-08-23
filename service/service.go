package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"test/service/config"
	"test/service/dataModel"
	"test/service/view"

	"github.com/nats-io/stan.go"
)

func StartServer() {

	log.Println("Server started")
	http.HandleFunc("/", view.MainPage)
	http.ListenAndServe(config.ServerPort, nil)
}

func UnwrapJsonFromNats(data *stan.Msg) *dataModel.Order_t {

	tmp := dataModel.Order_t{}
	err := json.Unmarshal(data.Data, &tmp)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &tmp
}

func ConvertJsonToByte(data *dataModel.Order_t) *[]byte {
	exampleBytes, err := json.Marshal(data)
	if err != nil {
		log.Println("Data no convert to byte ", err)
	}
	return &exampleBytes
}
