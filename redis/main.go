package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var Client *redis.Client

//go:embed main.lua
var script3 string

var lua1Hash string
var lua2Hash string
var lua3Hash string

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "1.14.59.249:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
	//启动时，预先将脚本缓存到redis，并返回客户端sha1 hash值
	lua1Hash, _ = Client.ScriptLoad(Script1).Result()
	lua2Hash, _ = Client.ScriptLoad(Script2).Result()
	lua3Hash, _ = Client.ScriptLoad(Script3).Result()
}

func main() {
	//fmt.Println(script3)
	//fmt.Println(200000*6*7 + 5400000)
	//Client.FlushAll()
	//demo1()
	//demo3()
	//println(script3)
	demo4()
}

func demo1() {
	Client.Set("stock", 11, -1)
	n, err := Client.EvalSha(lua1Hash, []string{"stock", "6"}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("结果", n, err)
}

//判断map里是否存在age字段，不存在则赋予初始值18，存在则加1
func demo2() {
	Client.HSet("me", "name", "bdz")
	n, err := Client.EvalSha(lua2Hash, []string{"me", "age", "1", "18"}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("结果", n, err)
}

//判断map里是否存在age字段，不存在则赋予初始值18，存在则加1
func demo3() {
	Client.HSet("me", "name", "bdz")
	n, err := Client.EvalSha(lua3Hash, []string{"me", "age", "1", "18"}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("结果", n, err)
}

func demo4() {
	//集合s1(1,2,3)与s2(1,4,5)

	//差集
	diff, err := Client.SDiff("s1", "s2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("差集：", diff)

	//交集
	inter, err := Client.SInter("s1", "s2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("交集：", inter)

	//并集
	union, err := Client.SUnion("s1", "s2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("并集：", union)

	//查看集合元素
	eles1, err := Client.SMembers("s1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("s1元素:", eles1)

	//随机获取集合中的一个或多个元素
	result, err := Client.SRandMemberN("s1", 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("随机获取一个或多个元素：", result)

	//
	p := Person{
		Name:  "bdz",
		Age:   18,
		Hoppy: []string{"play", "game", "mahjong"},
	}
	marshal, _ := json.Marshal(p)
	addRst, err := Client.SAdd("s3", marshal).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("s3添加对象成功：", addRst)

	rst, err := Client.SRandMember("s3").Result()
	if err != nil {
		panic(err)
	}
	var prst = &Person{}
	json.Unmarshal([]byte(rst), prst)

	fmt.Println(fmt.Sprintf("name:%s,age:%d,hoppy:%+v", prst.Name, prst.Age, prst.Hoppy))
	fmt.Println(fmt.Sprintf("rst:%+v", prst))
}

type Person struct {
	Name  string
	Age   int
	Hoppy []string
}

//bitmap 二值状态统计
func demo5() {
	result, err := Client.GetBit("bdz", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", result)
	Client.BitCount("bdz", &redis.BitCount{}).Result()

}

func demo6() {

}

//缓存一致性问题
func demo7() {
	//更新db
	//删除redis 重试。依然失败，记录失败key value。等后台程序后续处理
	//并发情况下 著有考虑加锁（需要写redis的时候加锁，这样会降低系统性能，与redis用于做缓存提升性能相违背，视业务情况而定）
}

//缓存雪崩、缓存击穿和缓存穿透
func demo8() {
	//雪崩：指某一时刻大量的热key过期失效，恰好此时并发访问这些key。就会导致流量全部打到mysql，mysql扛不住崩掉，系统瘫痪
	//解决：1、设置key的国企时间随机，不要同时失效  2、考虑服务降级（降一些不重要的业务接口直接返回nil，不走mysql查询）3、开启服务熔断机制（确保mysql正常运行）或者开启限流机制，降低访问速度

	//缓存击穿：某热点key失效，大量访问该热点key的请求全部打到mysql
	//解决：1、不设置过期时间 2、go可以使用singleFlight三方框架处理（参考singleFlight案例）

	//缓存穿透：指访问缓存和db都不存在的key，大量请求的话也会给mysql造成很大压力
	//解决：1、将该key写入到缓存，给一个默认值。下次访问直接返回该默认值（可能会造成数据泄露，因为有可能该key可以作为我们的正常key）
	//     2、布隆过滤器bitmap（二值状态），判断key一定不存在或者可能存在（可能存在的原因是hash冲突）
	//     3、检测恶意攻击
}

//指定某个key在某一事件第过期
func demo9() {
	parse, err := time.Parse("2006-01-02 15:04:05", "2023-12-25 15:00:00")
	if err != nil {
		return
	}
	Client.ExpireAt("name", parse)
}
