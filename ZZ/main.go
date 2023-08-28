package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
)

// 请用本机 MySQL 实现下面这 4 个 test case (MySQL 版本 8.0), 注意操作都必须是存盘操作, 不能全部都在内存中返回数据. 数据一定是要经过了硬盘的
// testCase0 : 实现高性能的大量数据的序列化写入
// testCase1 : 在上述基础上可以读取某批 device 的 TimeZone
// testCase2 : 在上述基础上可以读取某批 device 的具体数据, 并且单个 device 的消息是顺序的
// testCase3 : 删除某批 device 写入的数据
// 请注意标记有 FIXME 的地方是一定需要实现代码的地方, 标记有 TODO 的地方是起到提示作用 . 79 行之前的代码结构都可以修改, 只要代码结构和数据总量不变, 表达的意思是一样的就行
// TODO 评分标准为 : testCase 时间尽可能短(时间长短会决定最终分数), 并且保证正确性
func main() {
	startTime := time.Now()
	globalDb = mustGetMysqlDb()
	defer globalDb.Close()
	err := globalDb.Ping()
	if err != nil {
		panic(err)
	}
	testCase0()
	testCase1()
	testCase2()
	testCase3()
	endTime := time.Now()
	fmt.Println("total delay", endTime.Sub(startTime).String())
}

var deviceIdIndex = "deviceId"
var deviceTimeZoneIndex = "deviceTimeZone"
var tableName = "zz"

//删除索引
func deleteIndex() {
	dropIndexSql1 := fmt.Sprintf("alter table  %s drop index %s", tableName, deviceIdIndex)
	globalDb.Exec(dropIndexSql1)
	//if err != nil {
	//	log.Fatal("删除索引deviceId失败：", err.Error())
	//	return
	//}

	dropIndexSql2 := fmt.Sprintf("alter table  %s drop index %s", tableName, deviceTimeZoneIndex)
	globalDb.Exec(dropIndexSql2)
	//if err != nil {
	//	log.Fatal("删除索引deviceTimeZone失败：", err.Error())
	//	return
	//}
}

func deleteDeviceTimeZoneIndex() {
	dropIndexSql2 := fmt.Sprintf("alter table  %s drop index %s", tableName, deviceTimeZoneIndex)
	globalDb.Exec(dropIndexSql2)
}

//创建索引
func createIndex() {
	createIndexSql1 := fmt.Sprintf("create index %s on %s (%s)", deviceIdIndex, tableName, "deviceId")
	_, err := globalDb.Exec(createIndexSql1)
	if err != nil {
		log.Fatal("创建索引deviceId失败：", err.Error())
		return
	}

	createIndexSql2 := fmt.Sprintf("create index %s on %s (%s)", deviceTimeZoneIndex, tableName, "deviceTimeZone")
	_, err = globalDb.Exec(createIndexSql2)
	if err != nil {
		log.Fatal("创建索引deviceTimeZone失败：", err.Error())
		return
	}
}

func create() {
	createSql := "CREATE TABLE `zz`  (\n  `time` bigint NOT NULL,\n  `eventRandId` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',\n  `deviceId` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',\n  `deviceTimeZone` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',\n  `clientVersionDetail` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',\n  `platform` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',\n  INDEX `deviceId`(`deviceId` ASC) USING BTREE\n) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;"
	_, err := globalDb.Exec(createSql)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// FIXME 请完善这个函数, 函数需要把 genStatisticsDataCb 返回的数据写入到数据库里面, 要求写入时间尽可能的短 (使用 MySQL 8.0实现, 一定要存盘)
func testCase0() {
	create()
	platformList := []string{
		"ios",
		"android",
		"mac",
		"amazon",
	}
	const deviceNum = 1 << 18
	startTime := time.Now()
	// FIXME 下面部分都可能会修改到 (不能修改数据量 platformList,deviceNum)
	// FIXME 需要实现把下面的数据写入到 MySQL 8.0 并且存盘, 这个函数执行完之后保证所有数据全部存盘
	//插入前，先删除索引（提高效率）
	deleteIndex()
	tx, err := globalDb.Begin()
	if err != nil {
		log.Fatal("开启事务失败：", err.Error())
	}
	//定义一个步长，每次插入这么多条数据
	skip := 1 << 11
	for _, platform := range platformList {
		insertSQL := "INSERT INTO zz (`time`, eventRandId,deviceId,deviceTimeZone,clientVersionDetail,platform)"
		for i := 0; i < deviceNum; i += skip {
			// FIXME 这里会有大量数据返回, 需要处理最终存盘 (MySQL 8.0)
			values := make([]interface{}, 0)
			for j := 0; j < skip; j++ {
				deviceId := platform + "_" + RandStringBytesMaskImprSrcUnsafe(32)
				genStatisticsDataCb(deviceId, platform, 5, func(msg *StatisticsMessage) {
					values = append(values, msg.Time.UnixNano(), msg.EventRandId, msg.DeviceId, msg.DeviceTimeZone, msg.ClientVersionDetail, msg.Platform)
				})
			}
			// 执行批量插入
			placeholders := make([]string, 0, len(values)/6)
			for j := 0; j < len(values)/6; j++ {
				placeholders = append(placeholders, "(?,?,?,?,?,?)")
			}
			query := insertSQL + " VALUES " + strings.Join(placeholders, ", ")
			_, err = tx.Exec(query, values...)
			if err != nil {
				tx.Rollback()
				log.Fatal(err)
			}
		}
	}
	tx.Commit()
	endTime := time.Now()
	// 需要保证下面这个 log 执行时所有数据均写入 MySQL (本地硬盘)
	fmt.Println("delay", "testCase0", endTime.Sub(startTime).String())
}

// FIXME 通过 DB 获取某批 DeviceId 的 TimeZone, 要求读取时间尽可能的短, 禁止读取内存里面已有的那个 map (MySQL 8.0 从磁盘里面获取数据)
func placeholders(tem []string) string {
	var rst string
	for i, s := range tem {
		if i != len(tem)-1 {
			rst += "\"" + s + "\"" + ","
		} else {
			rst += "\"" + s + "\""
		}
	}
	return rst
}
func getDeviceIdTimeTimeZoneMap(deviceIdList []string) map[string]string {
	//创建索引，增加查询效率
	createIndex()
	m := make(map[string]string)
	skip := 1 << 11
	for i := 0; i < len(deviceIdList); i += skip {
		tmp := deviceIdList[i : i+skip]
		args := make([]interface{}, len(tmp))
		for i2, s := range tmp {
			args[i2] = s
		}
		query := `select deviceId,deviceTimeZone from zz where deviceId in (` + deal(len(tmp)) + `) group by deviceId,deviceTimeZone`
		//fmt.Println("query:", query)
		rows, err := globalDb.Query(query, args...)
		if err != nil {
			log.Fatal("查询数据失败getDeviceIdTimeTimeZoneMap：", err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			var deviceId, timeZone string
			err := rows.Scan(&deviceId, &timeZone)
			if err != nil {
				log.Fatal("获取数据失败getDeviceIdTimeTimeZoneMap：", err.Error())
			}
			m[deviceId] = timeZone
		}
	}
	return m
}

// FIXME 通过 DB 获取某批 DeviceId 的 StatisticsMessage List 的 map,并且内部的 List 是顺序的, 要求读取时间尽可能的短 (MySQL 8.0 从磁盘里面获取数据)
func getDeviceIdStatisticDataListMap(deviceIdList []string) (deviceIdMsgMap map[string][]*StatisticsMessage) {
	m := make(map[string][]*StatisticsMessage)
	skip := 1 << 11
	for i := 0; i < len(deviceIdList); i += skip {
		tmp := deviceIdList[i : i+skip]
		args := make([]interface{}, len(tmp))
		for i2, s := range tmp {
			args[i2] = s
		}
		query := `select * from zz where deviceId in (` + deal(len(tmp)) + `)`
		//fmt.Println("query:", query)
		rows, err := globalDb.Query(query, args...)
		if err != nil {
			log.Fatal("查询数据失败getDeviceIdTimeTimeZoneMap：", err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			var eventRandId, deviceId, timeZone, clientVersionDetail, platform string
			var currTime int64
			err := rows.Scan(&currTime, &eventRandId, &deviceId, &timeZone, &clientVersionDetail, &platform)
			if err != nil {
				log.Fatal("获取数据失败getDeviceIdTimeTimeZoneMap：", err.Error())
			}
			t := time.Unix(0, currTime)
			tmp := &StatisticsMessage{
				Time:                t,
				EventRandId:         eventRandId,
				DeviceId:            deviceId,
				DeviceTimeZone:      timeZone,
				ClientVersionDetail: clientVersionDetail,
				Platform:            platform,
			}
			m[deviceId] = append(m[deviceId], tmp)
		}
	}
	return m
}

// FIXME 通过 DB 删除某批 DeviceId 的 StatisticsMessage (MySQL 8.0 从磁盘里面删除)
func deal(n int) string {
	var rst string
	for i := 0; i < n; i++ {
		if i != n-1 {
			rst += "?,"
		} else {
			rst += "?"
		}
	}
	return rst
}
func deleteDeviceIdStatisticData(deviceIdList []string) {
	deleteDeviceTimeZoneIndex()
	tx, _ := globalDb.Begin()
	skip := 1 << 11
	for i := 0; i < len(deviceIdList); i += skip {
		tmp := deviceIdList[i : i+skip]
		args := make([]interface{}, len(tmp))
		for i2, s := range tmp {
			args[i2] = s
		}
		query := `delete from zz where deviceId in (` + deal(len(tmp)) + `)`
		//fmt.Println("query:", query)
		//fmt.Println(fmt.Sprintf("args:%v", args))
		_, err := tx.Exec(query, args...)
		if err != nil {
			tx.Rollback()
			log.Fatal("删除数据失败deleteDeviceIdStatisticData：", err.Error())
		}
	}
	tx.Commit()
	//重新加上索引，方便下次查询
	//createIndex()
}

/* ================================================================================== 代码请写到上面,重要的事情说 3 遍 ================================================================================== */
/* ================================================================================== 代码请写到上面,重要的事情说 3 遍 ================================================================================== */
/* ================================================================================== 代码请写到上面,重要的事情说 3 遍 ================================================================================== */
/* ================================================================================== 下面都是不用修改的代码,重要的事情说 3 遍 ================================================================================== */
/* ================================================================================== 下面都是不用修改的代码,重要的事情说 3 遍 ================================================================================== */
/* ================================================================================== 下面都是不用修改的代码,重要的事情说 3 遍 ================================================================================== */

//////////////////////////////////////////////
///////// 接下来为测试的检测代码,不需要修改 ////////
//////////////////////////////////////////////

// TODO 注意后面 3 个 testCase 都不需要修改, 需要保证响应速度和正确性, 检测磁盘返回的数据是否正常
// TODO testCase1 的 TODO 都不用写代码, 只是正确性检测(MySQL 8.0 从磁盘读取出来的数据)
func testCase1() {
	startTime := time.Now()
	deviceIdList := getRandHalfDeviceIdList()

	m := getDeviceIdTimeTimeZoneMap(deviceIdList)
	if len(m) != len(deviceIdList) { // TODO 检测数据是否正常
		panic("tz data len error")
	}
	originDeviceIdTimeZoneMap := getDeviceIdTimeZoneMap()
	for deviceId, timeZone := range m {
		if originDeviceIdTimeZoneMap[deviceId] != timeZone { // TODO 检测数据是否正常
			panic("check deviceId timeZone failed")
		}
	}
	endTime := time.Now()
	fmt.Println("delay", "testCase1", endTime.Sub(startTime).String())
}

// TODO testCase2 的 TODO 都不用写代码, 只是正确性检测(MySQL 8.0 从磁盘读取出来的数据)
func testCase2() {
	startTime := time.Now()
	// 基于 testCase0 , 把写入的某批设备的 StatisticsMessageList 取出来
	deviceIdList := getRandHalfDeviceIdList()
	m := getDeviceIdStatisticDataListMap(deviceIdList)
	if len(m) != len(deviceIdList) { // TODO 检测数据是否正常
		panic("map data len error")
	}
	var tagDeviceId string
	for k := range m {
		tagDeviceId = k
		break
	}
	list := m[tagDeviceId]
	if len(list) != 5 { // TODO 检测数据是否正常
		panic("data len error")
	}
	getLastInt := func(s string) int {
		lastValue := s[len(s)-1:]
		out, err := strconv.Atoi(lastValue)
		if err != nil {
			panic(err)
		}
		return out
	}
	originDeviceIdTimeZoneMap := getDeviceIdTimeZoneMap()
	var lastMsg *StatisticsMessage
	for i, msg := range list {
		if msg.DeviceId != tagDeviceId { // TODO 检测数据是否正常
			panic("deviceId error")
		}
		if msg.DeviceTimeZone != originDeviceIdTimeZoneMap[msg.DeviceId] { // TODO 检测数据是否正常
			panic("DeviceTimeZone error")
		}
		if msg.ClientVersionDetail != "1.1.1" { // TODO 检测数据是否正常
			panic("ClientVersionDetail error")
		}
		if msg.Platform != strings.Split(msg.DeviceId, "_")[0] { // TODO 检测数据是否正常
			panic("Platform error")
		}
		if i > 0 {
			if msg.Time.Sub(lastMsg.Time) != 1 { // TODO 检测数据是否正常
				panic("data Time sort error")
			}
			if getLastInt(msg.EventRandId)-getLastInt(lastMsg.EventRandId) != 1 { // TODO 检测数据是否正常
				panic("data EventRandId sort error")
			}
		}
		lastMsg = msg
	}
	endTime := time.Now()
	fmt.Println("delay", "testCase2", endTime.Sub(startTime).String())
}

// TODO testCase3 的 TODO 都不用写代码, 只是正确性检测(MySQL 8.0 从磁盘读取出来的数据)
func testCase3() {
	startTime := time.Now()
	// TODO 把 testCase0 写入的数全部删除, 需要在磁盘上全部抹掉
	deviceIdList := getRandHalfDeviceIdList()
	deleteDeviceIdStatisticData(deviceIdList)
	list := getDeviceIdStatisticDataListMap(deviceIdList)
	if len(list) != 0 { // TODO 检测数据是否从磁盘上删除
		panic("not delete datasource")
	}
	endTime := time.Now()
	fmt.Println("delay", "testCase3", endTime.Sub(startTime).String())
}

//////////////////////////////////////////////
///////// 接下来的代码为模版代码，请勿修改 /////////
//////////////////////////////////////////////

var globalDb *sql.DB

func mustGetMysqlDb() *sql.DB {
	// 连接数据库
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/?parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	databaseName := "test"
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + databaseName)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("USE " + databaseName)
	if err != nil {
		panic(err)
	}
	return db
}

type StatisticsMessage struct {
	Time                time.Time `json:",omitempty"`
	EventRandId         string    `json:",omitempty"`
	DeviceId            string    `json:",omitempty"`
	DeviceTimeZone      string    `json:",omitempty"`
	ClientVersionDetail string    `json:",omitempty"`
	Platform            string    `json:",omitempty"`
}

var timeZoneList = []string{
	"Asia/Shanghai",
	"Asia/Urumqi",
	"Asia/Hong_Kong",
	"Asia/Taipei",
	"Asia/Singapore",
	"Asia/Qatar",
	"America/Cayman",
	"America/New_York",
	"America/Sao_Paulo",
}

var gTotalDeviceIdList []string
var gTotalDeviceIdListLock sync.Mutex
var gDeviceIdTimeZoneMap = map[string]string{}
var gDeviceIdTimeZoneMapLock sync.Mutex

func totalDeviceIdAddOne(deviceId string) {
	gTotalDeviceIdListLock.Lock()
	gTotalDeviceIdList = append(gTotalDeviceIdList, deviceId)
	gTotalDeviceIdListLock.Unlock()
}

// TODO 随机获取一半的数据
func getRandHalfDeviceIdList() []string {
	gTotalDeviceIdListLock.Lock()
	defer gTotalDeviceIdListLock.Unlock()
	srcDeviceIdList := gTotalDeviceIdList
	dest := make([]string, len(srcDeviceIdList))
	perm := rand.Perm(len(srcDeviceIdList))
	for i, v := range perm {
		dest[v] = srcDeviceIdList[i]
	}
	if len(dest) <= 1 {
		return dest
	}
	return dest[:len(dest)/2]
}

func deviceIdSetTimeZone(deviceId, timeZone string) {
	gDeviceIdTimeZoneMapLock.Lock()
	gDeviceIdTimeZoneMap[deviceId] = timeZone
	gDeviceIdTimeZoneMapLock.Unlock()
}

// TODO 此处只是用来做正确性的检查, 笔试禁止使用这个 map 来读取数据
func getDeviceIdTimeZoneMap() map[string]string {
	gDeviceIdTimeZoneMapLock.Lock()
	out := gDeviceIdTimeZoneMap
	gDeviceIdTimeZoneMapLock.Unlock()
	return out
}

func genStatisticsDataCb(deviceId, platform string, count int, f func(msg *StatisticsMessage)) {
	totalDeviceIdAddOne(deviceId)
	randId := RandStringBytesMaskImprSrcUnsafe(6)
	timeZone := timeZoneList[int(srcInt63())%len(timeZoneList)]
	deviceIdSetTimeZone(deviceId, timeZone)
	t := time.Now()
	for i := 0; i < count; i++ {
		msg := newStatisticsMessage()
		msg.Time = t.Add(time.Duration(i))
		msg.EventRandId = randId + strconv.Itoa(i)
		msg.DeviceId = deviceId
		msg.DeviceTimeZone = timeZone
		msg.ClientVersionDetail = "1.1.1"
		msg.Platform = platform
		f(msg)
		freeStatisticsMessage(msg)
	}
}

var structPool sync.Pool

func newStatisticsMessage() *StatisticsMessage {
	msg := structPool.Get()
	if msg != nil {
		return msg.(*StatisticsMessage)
	}
	return &StatisticsMessage{}
}

func freeStatisticsMessage(msg *StatisticsMessage) {
	msg.reset()
	structPool.Put(msg)
}

func (msg *StatisticsMessage) reset() {
	msg.Time = time.Time{}
	msg.EventRandId = ""
	msg.DeviceId = ""
	msg.DeviceTimeZone = ""
	msg.ClientVersionDetail = ""
	msg.Platform = ""
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var src = rand.NewSource(time.Now().UnixNano())
var randLock sync.Mutex

func srcInt63() int64 {
	randLock.Lock()
	out := src.Int63()
	randLock.Unlock()
	return out
}

func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, srcInt63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = srcInt63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
