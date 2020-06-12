package main

import (
	"fmt"
	"go-jingdong/base"
	"go-jingdong/config"
)

func main()  {
	conf:=config.Config{
		Url:    "https://router.jd.com/api",
		AppKey: "7215a4b0d**********ecdde786d90",
		Secretkey: "9a1d812d**********78ae3a1ea90905",
		V: "1.0",
	}

	jd:=NewDj(&conf)
	data,err := jd.Get(base.System{
		Method:     "jd.union.open.goods.jingfen.query",
		Param_json: `{"goodsReq":{"eliteId":"26"}}`,
	})
	if err!=nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data)
	data2,err := jd.Get(base.System{
		Method:     "jd.union.open.order.query",
		Param_json: `{"orderReq":{"time":"202006121645","pageNo":"1","pageSize":"10","type":"1"}}`,
	})
	if err!=nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data2)

}