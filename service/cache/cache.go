package cache

import (
	"test/service/dataModel"
)

var Cache = make(map[string]dataModel.Order_t, 5)

func AddCache(data dataModel.Order_t) {

	if data.OrderUid != "" {
		Cache[data.OrderUid] = data
	}

}

func IsInCache(find string) bool {

	if _, ok := Cache[find]; ok {
		return true
	}
	return false

}

func GetCache(find string) (*dataModel.Order_t, bool) {

	if val, ok := Cache[find]; ok {
		return &val, true
	}
	return nil, false
}

func UpdateCache(data *dataModel.Order_t) {

	if _, ok := Cache[data.OrderUid]; ok {
		Cache[data.OrderUid] = *data
	}

}

// func CheckCacheKey() {
// 	for i := range Cache {
// 		fmt.Println(i)

// 	}
// }
