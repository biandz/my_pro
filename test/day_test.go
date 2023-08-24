package test

import (
	"container/heap"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
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

var sli []int

func Test_17(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go appendTo()
	}
	wg.Wait()
	fmt.Println(sli, len(sli))
	fmt.Println("执行完毕！！")
}

func appendTo() {
	l.Lock()
	defer l.Unlock()
	sli = append(sli, 1)
	fmt.Printf("%p", sli)
	fmt.Println()
	wg.Done()
}

//实现最大堆
func Test_18(t *testing.T) {
	var q = &Queue{}
	var sli = []int{1, 8, 1, 85, 1, 3, 7, 1, 6}
	for i, v := range sli {
		var p = Person{
			Name: fmt.Sprintf("name%d", i),
			Age:  v,
		}
		heap.Push(q, p)
	}
	for i := 0; i < len(sli); i++ {
		fmt.Println(heap.Pop(q).(Person))
	}
}

type Person struct {
	Name string
	Age  int
}

type Queue []Person

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	return q[i].Age > q[j].Age
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x any) {
	*q = append(*q, x.(Person))
}

func (q *Queue) Pop() any {
	old, n := *q, len(*q)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

func Test_19(t *testing.T) {
	var a int8
	var b int16
	var c int32
	var d int64
	var e uint8
	var f uint16
	var g uint32
	var h uint64
	var i bool
	var j string

	fmt.Println("字节数int8:", unsafe.Sizeof(a))
	fmt.Println("字节数int16:", unsafe.Sizeof(b))
	fmt.Println("字节数int32:", unsafe.Sizeof(c))
	fmt.Println("字节数int64:", unsafe.Sizeof(d))
	fmt.Println("字节数uint8:", unsafe.Sizeof(e))
	fmt.Println("字节数uint16:", unsafe.Sizeof(f))
	fmt.Println("字节数uint32:", unsafe.Sizeof(g))
	fmt.Println("字节数uint64:", unsafe.Sizeof(h))
	fmt.Println("字节数bool:", unsafe.Sizeof(i))
	fmt.Println("字节数string:", unsafe.Sizeof(j))
}

func Test_20(t *testing.T) {
	deviceID, err := getWindowsDeviceID()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Device ID:", deviceID)
}
func getWindowsDeviceID() (string, error) {
	cmd := exec.Command("wmic", "cpu", "get", "ProcessorID")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	deviceID := strings.TrimSpace(string(output))
	return deviceID, nil
}

func Test_21(t *testing.T) {
	interInfos, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, info := range interInfos {
		fmt.Println(info.Name)
		fmt.Println(info.HardwareAddr.String())
	}
}

func Test_22(t *testing.T) {
	nums := []int{2, 3, 6, 7}
	fmt.Println(zuhe22(nums))
}

func zuhe22(nums []int) [][]int {
	var rst [][]int
	var tempSli []int
	var dfs func(nums, tempSli []int, start int)
	dfs = func(nums, tempSli []int, start int) {
		if len(tempSli) > 0 {
			rst = append(rst, tempSli)
		}
		for i := start; i < len(nums); i++ {
			tempSli = append(tempSli, nums[i])
			//进入下一层
			dfs(nums, tempSli, i+1)
			tempSli = removeLast(tempSli)
		}
	}
	dfs(nums, tempSli, 0)
	return rst
}

func removeLast(nums []int) []int {
	s := []int{}
	for k, v := range nums {
		if k != len(nums)-1 {
			s = append(s, v)
		}
	}
	return s
}

//无重复元素，可复选
func Test_23(t *testing.T) {
	candidates := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(candidates, target))
}

func combinationSum(candidates []int, target int) [][]int {
	var rst [][]int
	var tempSum int
	var tempSli []int
	var start int
	var dfs func(candidates, tempSli []int, tempSum, start int)
	dfs = func(candidates, tempSli []int, tempSum, start int) {
		if tempSum == target {
			rst = append(rst, tempSli)
			return
		}
		if tempSum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			//选择
			tempSli = append(tempSli, candidates[i])
			tempSum += candidates[i]
			//进入
			dfs(candidates, tempSli, tempSum, i)
			//撤销
			tempSum -= candidates[i]
			tempSli = removeLast(tempSli)
		}
	}
	dfs(candidates, tempSli, tempSum, start)
	return rst
}

//可重复元素，不能重复选的剪枝
func Test_24(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}

func combinationSum2(candidates []int, target int) [][]int {
	var rst [][]int
	var tempSum int
	var tempSli []int
	var start int
	//减枝所需
	var isUsed = make([]bool, len(candidates))
	var dfs func(candidates, tempSli []int, tempSum, start int)
	dfs = func(candidates, tempSli []int, tempSum, start int) {
		if tempSum == target {
			rst = append(rst, tempSli)
			return
		}
		if tempSum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			//减枝所需
			if i > 0 && candidates[i] == candidates[i-1] && !isUsed[i-1] {
				continue
			}
			//选择
			tempSli = append(tempSli, candidates[i])
			tempSum += candidates[i]
			isUsed[i] = true
			//进入
			dfs(candidates, tempSli, tempSum, i+1)
			//撤销
			isUsed[i] = false
			tempSum -= candidates[i]
			tempSli = removeLast(tempSli)
		}
	}
	//减枝所需
	sort.Ints(candidates)
	dfs(candidates, tempSli, tempSum, start)
	return rst
}

func Test_25(t *testing.T) {
	nums := []int{3, 3, 0, 3}
	fmt.Println(permuteUnique(nums))
}

func permuteUnique(nums []int) [][]int {
	var rst [][]int
	var isUsed = make([]bool, len(nums))
	var dfs func(tempSli []int)
	var tempSli []int
	dfs = func(tempSli []int) {
		if len(tempSli) == len(nums) {
			rst = append(rst, tempSli)
		}

		for i := 0; i < len(nums); i++ {
			//重复元素，不可复选的减枝
			if i > 0 && nums[i] == nums[i-1] && !isUsed[i-1] {
				continue
			}
			if isUsed[i] {
				continue
			}
			//选择
			tempSli = append(tempSli, nums[i])
			isUsed[i] = true
			//进入
			dfs(tempSli)
			//撤销
			tempSli = removeLast(tempSli)
			isUsed[i] = false
		}
	}
	sort.Ints(nums)
	dfs(tempSli)
	return rst
}

func Test_26(t *testing.T) {
	n, k := 1, 1
	fmt.Println(combine(n, k))
}

func combine(n int, k int) [][]int {
	nums := genSli(n)
	var rst [][]int
	var tempSli []int
	var start int
	var dfs func(tempSli []int, start int)
	dfs = func(tempSli []int, start int) {
		if len(tempSli) == k {
			rst = append(rst, tempSli)
			return
		}
		for i := start; i < n; i++ {
			//选择
			tempSli = append(tempSli, nums[i])
			//进入
			dfs(tempSli, i+1)
			//撤销
			tempSli = removeLast(tempSli)
		}
	}
	dfs(tempSli, start)
	return rst
}

func genSli(n int) []int {
	sli := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		sli = append(sli, i)
	}
	return sli
}

func Test_27(t *testing.T) {
	nums := []int{0}
	fmt.Println(subsets(nums))
}

func subsets(nums []int) [][]int {
	var rst [][]int
	var start int
	var tempSli []int
	var dfs func(tempSli []int, start int)
	dfs = func(tempSli []int, start int) {
		rst = append(rst, tempSli)
		for i := start; i < len(nums); i++ {
			//选择
			tempSli = append(tempSli, nums[i])
			//进入
			dfs(tempSli, i+1)
			//撤销
			tempSli = removeLast(tempSli)
		}
	}
	dfs(tempSli, start)
	return rst
}

func Test_28(t *testing.T) {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "SEE"
	fmt.Println(exist(board, word))

}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var isUsed = make([][]bool, m)
	for i := 0; i < m; i++ {
		isUsed[i] = make([]bool, n)
	}
	var dfs func(x, y, index int) bool
	dfs = func(x, y, index int) bool {
		if board[x][y] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		//选择
		isUsed[x][y] = true
		//进入
		//下
		if 0 <= x+1 && x+1 < m && !isUsed[x+1][y] {
			if dfs(x+1, y, index+1) {
				return true
			}
		}
		//上
		if 0 <= x-1 && x-1 < m && !isUsed[x-1][y] {
			if dfs(x-1, y, index+1) {
				return true
			}
		}
		//左
		if 0 <= y-1 && y-1 < n && !isUsed[x][y-1] {
			if dfs(x, y-1, index+1) {
				return true
			}
		}
		//右
		if 0 <= y+1 && y+1 < n && !isUsed[x][y+1] {
			if dfs(x, y+1, index+1) {
				return true
			}
		}
		//撤销
		isUsed[x][y] = false
		return false
	}

	for x, bytes := range board {
		for y, _ := range bytes {
			if dfs(x, y, 0) {
				return true
			}
		}
	}
	return false
}

func Test_29(t *testing.T) {
	nums := []int{0}
	fmt.Println(subsetsWithDup(nums))
}

func subsetsWithDup(nums []int) [][]int {
	var rst [][]int
	var tempSli []int
	var isUsed = make([]bool, len(nums))
	var start int
	var dfs func(tempSli []int, start int)
	dfs = func(tempSli []int, start int) {
		rst = append(rst, tempSli)
		for i := start; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && !isUsed[i-1] {
				continue
			}
			if isUsed[i] {
				continue
			}
			//选择
			isUsed[i] = true
			tempSli = append(tempSli, nums[i])
			//进入
			dfs(tempSli, i+1)
			//撤销
			isUsed[i] = false
			tempSli = removeLast(tempSli)
		}
	}
	sort.Ints(nums)
	dfs(tempSli, start)
	return rst
}

func Test_30(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(findLadders(beginWord, endWord, wordList))
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var rst [][]string

	var queue = []string{beginWord}
	for len(queue) != 0 {
		for i := 0; i < len(queue); i++ {
			//取出第一个并更新queue
			queue := queue[1:]
			first := queue[0]
			queue = append(queue, rex(first, wordList)...)
		}
	}

	return rst
}

//返回与自身只有一位字符不同的字符串
func rex(str string, strs []string) []string {
	var rst []string
	for i := 0; i < len(str); i++ {
		//修改第i个字符从‘a’到‘z’
		for j := 0; j < 26; j++ {
			newStr := genNewStr(i, m[j], str)
			if inArray(newStr, strs) {
				rst = append(rst, newStr)
			}
		}
	}
	return rst
}

func genNewStr(index int, char byte, str string) string {
	b := []byte(str)
	b[index] = char
	return string(b)
}

var m = map[int]byte{
	0: 'a', 1: 'b', 2: 'c', 3: 'd', 4: 'e',
	5: 'f', 6: 'g', 7: 'h', 8: 'i', 9: 'j',
	10: 'k', 11: 'l', 12: 'm', 13: 'n', 14: 'o',
	15: 'p', 16: 'q', 17: 'r', 18: 's', 19: 't',
	20: 'u', 21: 'v', 22: 'w', 23: 'x', 24: 'y', 25: 'z',
}

func inArray(str string, strs []string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

func Test_31(t *testing.T) {
	sli := []int{1}
	fmt.Println(sli[0])
	fmt.Println(sli[1:])
}

func Test_32(t *testing.T) {
	var ch = make(chan string, 100)

	go func() {
		for {
			ch <- "协程1"
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			ch <- "协程2"
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			ch <- "协程3"
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		select {
		case a := <-ch:
			fmt.Println(a)
		}
	}

}

type Man struct {
	Name string
	Age  int
}

func (m *Man) setName(name string) {
	m.Name = name
}

func (m *Man) getName() string {
	return m.Name
}

func (m *Man) setAge(age int) {
	m.Age = age
}

func (m *Man) getAge() int {
	return m.Age
}

func newMan(name string, age int) *Man {
	return &Man{name, age}
}

func Test_33(t *testing.T) {
	m := newMan("bdz", 18)
	fmt.Println(m.getAge())
	m.setAge(19)
	fmt.Println(m.getAge())
}

func Test_34(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	s2 := s1
	fmt.Println(s2)
	s2[3] = 777
	fmt.Println(s1)
	fmt.Println(s2)
}

var rst uint32 = 0

func Test_35(t *testing.T) {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("打印前结果：", rst)
			atomic.AddUint32(&rst, 2)
			fmt.Println("打印后结果：", rst)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(rst)
}
