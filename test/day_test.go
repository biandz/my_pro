package test

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
)

func Test_01(t *testing.T) {
	for i := 0; i < 10000; i++ {
		fmt.Println(i)
	}
}

var pk1 = []int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 9, 10, 10, 10, 10}

var pk = make([]int, len(pk1))

func InitPk() {
	m := make(map[int]int)
	for i, i2 := range pk1 {
		m[i] = i2
	}
	for _, i := range m {
		pk = append(pk, i)
	}
}

func Test_02(t *testing.T) {
	InitPk()
	rst, err := deal(5)
	if err != nil {
		return
	}

	for i, ints := range rst {
		fmt.Println(ints)
		fmt.Println(fmt.Sprintf("第%d个玩家的牌点：%d<->%d;分数:%d", i, pk[ints[0]], pk[ints[1]], compute(ints)))
	}
}

//发牌的索引
func deal(num int) ([][]int, error) {
	var rst = make([][]int, 0, num)
	if num*2 > len(pk) {
		return rst, errors.New("num is to much!!")
	}
	var hasUsed = make(map[int]struct{})
	for i := 0; i < num; i++ {
		var n1, n2 int
		//取第一个数
		for {
			n1 = getRand()
			if _, ok := hasUsed[n1]; !ok {
				hasUsed[n1] = struct{}{}
				break
			} else {
				hasUsed[n1] = struct{}{}
			}
		}

		//取第二个数
		for {
			n2 = getRand()
			if _, ok := hasUsed[n2]; !ok {
				hasUsed[n2] = struct{}{}
				break
			} else {
				hasUsed[n2] = struct{}{}
			}
		}
		rst = append(rst, []int{n1, n2})
	}
	return rst, nil
}

//获取随机数
func getRand() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(pk))
}

//计算牌力值
func compute(rst []int) int {
	var score = 0
	first := rst[0]
	second := rst[1]
	if pk[first] == pk[second] { //豹子
		score = 10 + pk[first]*2
	} else {
		score = (pk[first] + pk[second]) % 10
	}
	return score
}

//LRU算法实践
func Test_03(t *testing.T) {
	l := NewLRU()
	for i := 0; i < 20; i++ {
		l.Set(fmt.Sprintf("key%d", i), i)
		fmt.Println(l.getAllKey())
	}
	get, _ := l.Get("key11")
	fmt.Println(get)
	fmt.Println(l.getAllKey())
	fmt.Println("===============================")
	get1, _ := l.Get("key10")
	fmt.Println(get1)
	fmt.Println(l.getAllKey())
}

const MemoryList = 10

type LRU struct {
	List []string
	Map  map[string]any
	sync.RWMutex
}

func NewLRU() *LRU {
	return &LRU{
		Map: make(map[string]any),
	}
}

func (l *LRU) Get(key string) (any, bool) {
	l.Lock()
	defer l.Unlock()
	val, ok := l.Map[key]
	if ok {
		if len(l.List) >= MemoryList {
			//淘汰第一个key，并将新ky放入队尾
			firstKey := l.List[0]
			delete(l.Map, firstKey)
			l.List = append(l.List[1:], key)
		} else {
			l.List = append(l.List, key)
		}
	}
	return val, ok
}

func (l *LRU) Set(key string, val any) {
	l.Lock()
	defer l.Unlock()
	l.Map[key] = val
	if len(l.List) >= MemoryList {
		l.List = append(l.List[1:], key)
	} else {
		l.List = append(l.List, key)
	}
}

func (l *LRU) getAllKey() []string {
	return l.List
}

func Test_55(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
}
func canJump(nums []int) bool {
	var maxMove = 02
	for i, num := range nums {
		if i == 0 {
			maxMove = num //最远可达
		} else {
			if maxMove >= i { //说明能到达当前下标
				maxMove = max(maxMove, i+num)
			}
		}
	}
	return maxMove >= len(nums)-1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_06(t *testing.T) {
	timeout, err := net.DialTimeout("tcp", "1.14.59.249:9998", 3*time.Second)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	timeout.Close()
}

func Test_07(t *testing.T) {
	a := []int{1, 2, 3}
	test7(a)
	fmt.Println()
	fmt.Printf("%p", a)
}

func test7(b []int) {
	fmt.Printf("%p", b)
}

func Test_08(t *testing.T) {
	fmt.Println("bdz\nwy")
}

func Test_09(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	for _, i := range a {
		go func(i int) {
			print(i)
		}(i)
	}
}

func Test_10(t *testing.T) {
	var a any = 3.2
	fmt.Println()
	switch a.(type) {
	case int:
		fmt.Println("int:", a)
	case float64:
		fmt.Println("float64:", a)
	}
}

func Test_11(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "bdz")
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)
	fmt.Println("------开始子协程------")
	go test1101(ctx)
	go test1102(ctx)
	//模拟业务场景
	time.Sleep(10 * time.Second)
	fmt.Println("------结束子协程------")
	//cancelFunc()
	time.Sleep(5 * time.Second)
	fmt.Println("------end-------")
}

func test1101(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("test1101:", ctx.Value("name"))
			time.Sleep(2 * time.Second)
		}
	}

}
func test1102(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("test1102", ctx.Value("name"))
			time.Sleep(2 * time.Second)
		}
	}

}

func Test_12(t *testing.T) {
	fmt.Println(strconv.FormatInt(time.Now().UnixNano(), 16))
}

func Test_13(t *testing.T) {
	type s struct {
		name string
	}

	s1 := s{name: "bdz"}
	s2 := s{name: "wy"}
	fmt.Println(s1 == s2)
}

var l sync.Mutex

func Test_14(t *testing.T) {
	l.Lock()
	fmt.Println(123)
	l.Unlock()
}

var wg sync.WaitGroup

func Test_15(t *testing.T) {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("end")
	fmt.Println(runtime.Version())
}

func Test_16(t *testing.T) {
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for {
			<-ticker.C
			test16()
		}
	}()
	select {}
}

func test16() {
	fmt.Println("111")
	return
}
