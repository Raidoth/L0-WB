package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"test/service"
	"test/service/cache"
	"test/service/config"
	"test/service/dataModel"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

var cn *sql.DB
var Validator = validator.New()

func ConnectionDB(wg *sync.WaitGroup) {
	log.Println("Connecting database...")
	PgInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", config.HostPg, config.PortPg, config.UserPg, config.PasswordPg, config.NameDbPg)
	connection, err := sql.Open("postgres", PgInfo)
	if err != nil {
		log.Println("Dattabase no connected ", err)
		panic(err)
	} else {
		log.Println("Database connected")
	}
	if err = connection.Ping(); err != nil {
		log.Println("Database no answer ", err)
	} else {
		log.Println("Database answer")
	}
	cn = connection
	wg.Done()

}

func PushDB(data *dataModel.Order_t) {

	log.Println("Pushing data in database")

	if err := Validator.Struct(data); err != nil {
		log.Println("Accepted json no valid", err)
		return
	}

	if !cache.IsInCache(data.OrderUid) {
		cache.AddCache(*data)

		bytesStruct := service.ConvertJsonToByte(data)
		_, err := cn.Exec("INSERT INTO orders (id, data) VALUES ($1,$2)", data.OrderUid, bytesStruct)
		if err != nil {
			log.Println("Database insert violated ", err)
		} else {
			log.Println("Data in database")
		}
	} else {
		UpdateDataDB(data)
	}

}

func UpdateDataDB(data *dataModel.Order_t) {

	log.Println("Updating data in database")

	if err := Validator.Struct(data); err != nil {
		log.Println("Accepted json no valid", err)
		return
	}

	if !cache.IsInCache(data.OrderUid) {
		cache.AddCache(*data)
	}
	bytesStruct := service.ConvertJsonToByte(data)
	_, err := cn.Exec("UPDATE orders set data=$1 where id=$2", bytesStruct, data.OrderUid)

	if err != nil {
		log.Println("Database update violated")
		panic(err)
	}
	cache.UpdateCache(data)
	log.Println("Data updated")
}

func ShowDB() {

	rows, err := cn.Query("SELECT id FROM orders")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var tmp string
	for rows.Next() {
		rows.Scan(&tmp)
		fmt.Println(tmp)
	}
}

func IsHaveIdDB(data *dataModel.Order_t) bool {

	log.Println("Check sent json")

	if err := Validator.Struct(data); err != nil {
		log.Println("Accepted json no valid", err)
		return false
	}

	rows, err := cn.Query("SELECT id FROM orders WHERE id=$1", data.OrderUid)
	if err != nil {
		log.Println("Query database don't send")
		panic(err)
	}
	defer rows.Close()
	t := rows.Next()
	return t

}

// func CheckDB() error {

// 	for {
// 		time.Sleep(3 * time.Second)
// 		if err := cn.Ping(); err != nil {
// 			return err
// 		}
// 		//fmt.Println("Checking db...")
// 		//cache.CheckCacheKey()
// 	}

// }

func DisconnectionDB() {
	cn.Close()
	log.Println("Database close")
}
