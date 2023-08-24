package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"reflect"
)

type Person struct {
	Name   string        `json:"name"`
	Age    int           `json:"age"`
	Gental bool          `json:"gental"`
	Hoppy  []interface{} `json:"hoppy"`
}

var (
	Index  = "index1"
	Url    = "http://1.14.59.249:9200"
	Ctx    = context.TODO()
	client *elastic.Client
)

func main() {
	c, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(Url))
	if err != nil {
		log.Fatal("创建链接失败:", err.Error())
		return
	}
	//也测试es链接状态
	_, _, err = c.Ping(Url).Do(Ctx)
	if err != nil {
		log.Fatal("链接es失败：", err.Error())
		return
	}
	//查看当前es版本
	version, err := c.ElasticsearchVersion(Url)
	if err != nil {
		log.Fatal("获取es版本失败：", err.Error())
		return
	} else {
		fmt.Println("当前es版本：", version)
	}
	client = c
	//set(client)
	get()
	get1()
}

func set() {

}

func get() {
	var p = make([]Person, 0)
	res, err := client.Search(Index).Query(elastic.NewQueryStringQuery("name:wy")).Do(Ctx)
	if err != nil {
		log.Fatal("获取数据失败：", err.Error())
		return
	}
	for _, i2 := range res.Each(reflect.TypeOf(Person{})) {
		t := i2.(Person)
		p = append(p, t)
	}
	fmt.Println(fmt.Sprintf("%v", p), len(p))
}

func get1() {
	res, err := client.Get().Index(Index).Id("EplZJYoBq_0wfSFy0jA5").Do(Ctx)
	if err != nil {
		log.Fatal("获取数据失败:", err.Error())
		return
	}
	if res.Found {
		var p = &Person{}
		_ = json.Unmarshal(res.Source, p)
		fmt.Println(fmt.Sprintf("%v", p))
	}
}
