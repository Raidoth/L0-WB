package main

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"test/database"
	"test/service/cache"
	"test/service/dataModel"
	"test/service/subscriber"
	"testing"

	"github.com/nats-io/stan.go"
)

func TestSubscribeNats(t *testing.T) {
	sc, err := stan.Connect(subscriber.ClusterID, "test2")
	if err != nil {
		t.Error("Nats not connection ", err)
	}
	defer sc.Close()
	sub, err := sc.Subscribe("WB", func(m *stan.Msg) {
	})
	if err != nil {
		t.Error("Nats not subscribe ", err)
	}
	defer sub.Close()
}

func TestDb(t *testing.T) {

	data, err := ioutil.ReadFile("model.json")
	if err != nil {
		t.Error("Cannot read file: ", err)
	}
	tmp := dataModel.Order_t{}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		t.Error("Unmarsahl error", err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	database.ConnectionDB(&wg)
	wg.Wait()
	database.PushDB(&tmp)
	if ok := database.IsHaveIdDB(&tmp); !ok {
		t.Error("No data in database")
	}
}

func TestCache(t *testing.T) {
	data, err := ioutil.ReadFile("model.json")
	if err != nil {
		t.Error("Cannot read file: ", err)
	}
	tmp := dataModel.Order_t{}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		t.Error("Unmarsahl error", err)
	}
	cache.AddCache(tmp)
	ord, ok := cache.GetCache("b563feb7b2b84b6test")
	if !ok {
		t.Error("No struct in cache")
	}
	if ord.Payment.Amount != tmp.Payment.Amount {
		t.Error("Fild struct error")
	}

}
