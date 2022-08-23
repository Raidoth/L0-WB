package subscriber

import (
	"log"
	"test/database"
	"test/service"
	"test/service/cache"

	"github.com/nats-io/stan.go"
)

const (
	ClusterID = "test-cluster"
)

func Subscribe() (*stan.Conn, *stan.Subscription) {
	log.Println("Nats subscribe")
	sc, err := stan.Connect(ClusterID, "test2")
	if err != nil {
		log.Println("Nats not connection ", err)
	}
	sub, err := sc.Subscribe("WB", getMsg, stan.DeliverAllAvailable())
	if err != nil {
		log.Println("Nats not subscribe ", err)
	}
	return &sc, &sub
}

func getMsg(m *stan.Msg) {
	data := service.UnwrapJsonFromNats(m)

	if !database.IsHaveIdDB(data) {
		database.PushDB(data)

	} else {
		database.UpdateDataDB(data)
	}

}
func NatsClose(sc stan.Conn, sub stan.Subscription) {

	err := sub.Unsubscribe()
	if err != nil {
		log.Println("Nats no unsubscribe")
	}
	err = sc.Close()
	if err != nil {
		log.Println("Nats no close")
	}
	log.Println("Nats close")
}

func GetsCache(id string) bool {
	_, s := cache.GetCache(id)
	return s
}
