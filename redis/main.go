package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var rdb *redis.Client

func initRedisClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:8888",
		Password: "tomjack",
		DB:       0,
		PoolSize: 100,
	})

	_, err = rdb.Ping().Result()
	return err
}

// 字符串以及一些键值对的基本操作
func stringEx() {
	// 添加/修改一个键
	err := rdb.Set("name", "rzzy", 0).Err()
	if err != nil {
		fmt.Printf("set err,err:%v\n", err)
		return
	}

	// 查看所有键
	fmt.Println("所有键:", rdb.Keys("*").Val())

	// 添加或修改多个键
	err = rdb.MSet("age", "20", "gender", "男", "length", "190").Err()
	if err != nil {
		fmt.Printf("mset err,err:%v\n", err)
		return
	}

	// 删除一个键
	err = rdb.Del("gender").Err()
	if err != nil {
		fmt.Printf("del err,err:%v\n", err)
		return
	}

	// 获取一个键
	name, err := rdb.Get("name").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Printf("this key not found, err:%v\n", err)
		} else {
			fmt.Printf("get error,err:%v\n", err)
			return
		}
	}
	fmt.Println("name:", name)

	// 获取多个键
	mulkey, err := rdb.MGet("name", "age").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Printf("this key not found, err:%v\n", err)
		} else {
			fmt.Printf("get error,err:%v\n", err)
			return
		}
	}
	fmt.Println(mulkey)

	// 判断键是否存在
	fmt.Println(rdb.Exists("gender"))

	// 查看键类型
	fmt.Println(rdb.Type("name"))

	// 设置键的过期时间
	rdb.Expire("length", time.Second*20) // 第二个时间参数是以纳秒为单位，所以直接用了time.Second

	// 查看键的过期时间
	fmt.Println(rdb.TTL("name"))

	// 创建有过期时间的键值对
	rdb.Set("gender", "男", time.Second*20)

	// 获取原值并设置新值
	fmt.Println(rdb.GetSet("name", "rzxy"))

	// 给值后追加内容
	rdb.Append("name", "is rzzy")

	// 获取值的长度
	fmt.Println(rdb.StrLen("name"))

	// 只有当键值对不存在时才能设置键值对
	rdb.SetNX("location", "China", 0)

	// 给纯数字键值对的值自增1
	fmt.Println(rdb.Incr("age"))

	// 给纯数字键值对的值自减1
	fmt.Println(rdb.Decr("age"))

	// 给纯数字键值对的值自增n
	fmt.Println(rdb.IncrBy("age", 100))

	// 给纯数字键值对的值自减n
	fmt.Println(rdb.DecrBy("age", 100))

	// 获取值的0~3长度内的内容
	fmt.Println(rdb.GetRange("name", 0, 3))

	// 从开始位置开始覆写值
	rdb.SetRange("name", 9, "eihei")

	// 展示下现在所有的键值对
	for _, s := range rdb.Keys("*").Val() {
		fmt.Println(s, ":", rdb.Get(s).Val())
	}
}

// List示例
func listEx() {
	// 从列表左边插入一个或多个值
	rdb.LPush("userName", "rzzy", "rzxy")

	// 从列表右边插入一个或多个值
	rdb.RPush("nickName", "SukiMegumi", "dtmyx")

	// 从列表左边弹出一个值，右边的话把L改成R
	fmt.Println(rdb.LPop("userName"))

	// 从第一个列表的右边弹出一个值，将这个弹出的值添加到第二个列表的左边
	rdb.RPopLPush("userName", "nickName")

	// 按照索引下标获取元素
	rdb.LIndex("nickName", 0)

	// 获取列表长度
	rdb.LLen("nickName")

	// 插入查找到的第一个Value后（插入其之前的话，把After改成Before）
	rdb.LInsertAfter("userName", "rzzy", "eihei")

	// 删除 1 个value值
	rdb.LRem("userName", 1, "eihei")

	// 将列表key下标为index的值替换为value
	rdb.LSet("useName", 0, "rzxy")

}

// set示例
func setEx() {
	// 添加一个或多个元素到集合key中，已存在的忽略掉
	rdb.SAdd("name", "rzzy", "rzxy", "dtmyx", "SukiMegumi", "eihei")
	rdb.SAdd("nickName", "a", "rzxy", "b", "v", "eihei")

	// 取出该集合所有元素
	fmt.Println(rdb.SMembers("name"))

	// 判断集合中是否含有某个元素
	fmt.Println(rdb.SIsMember("name", "rzzy"))

	// 返回该集合元素个数
	fmt.Println(rdb.SCard("name"))

	// 删除集合中的某个值
	fmt.Println(rdb.SRem("name", "rzxy"))

	// 随机从集合中弹出一个值
	fmt.Println(rdb.SPop("name"))

	// 随机从集合中弹出n个值
	fmt.Println(rdb.SPopN("name", 1))

	// 把集合中的一个值从一个集合移动到另一个集合
	fmt.Println(rdb.SMove("name", "nickName", "SukiMegumi"))

	// 返回两个集合的交集元素,并集为SUnion,差集为SDiff
	fmt.Println(rdb.SInter("name", "nickName"))

}

// zset示例
func zsetEx() {
	languages := []redis.Z{
		{Score: 64.96, Member: "JavaScript"},
		{Score: 56.07, Member: "HTML/CSS"},
		{Score: 48.24, Member: "Python"},
		{Score: 47.08, Member: "SQL"},
		{Score: 35.35, Member: "Java"},
	}
	// 将一个或多个member及其score添加到有序集合中
	rdb.ZAdd("programRank", languages...)

	// 返回有序集中，下标在start和stop之间的元素
	fmt.Println(rdb.ZRange("programRank", 0, 3))

	// 返回有序集中，所有score介于min和max之间的成员，有序成员按从小到大排序(从大到小用rdb.ZRevRangeByScoreWithScores())
	fmt.Println(rdb.ZRangeByScoreWithScores("programRank", redis.ZRangeBy{Min: "50", Max: "100"}))

	// 让元素的score加上指定值
	fmt.Println(rdb.ZIncrBy("programRank", 3, "Java"))
	// 自减
	fmt.Println(rdb.ZIncrBy("programRank", -3, "Java"))

	// 删除该集合下指定的元素
	fmt.Println(rdb.ZRem("programRank", "Java"))

	// 统计该集合,min和max区间内元素个数
	fmt.Println(rdb.ZCount("programRank", "30", "60"))

	// 返回该值的在集合中的排名，排名从0开始(按score从小到大排序，倒序使用ZRevRank)
	fmt.Println(rdb.ZRank("programRank", "JavaScript"))
}

// hash 示例
func hashEx() {
	// 给哈希表中filed赋值value
	rdb.HSet("user10001", "name", "rzzy")

	// 批量设置filed值
	user10001 := map[string]interface{}{
		"name":   "rzzy",
		"age":    "100",
		"height": "180",
	}
	rdb.HMSet("user10001", user10001)

	// 设置哈希表中的filed的值，当且仅当该field不存在
	rdb.HSetNX("user10001", "weight", "80")

	// 从哈希表中去除filed的值
	fmt.Println(rdb.HGet("user10001", "age"))

	// 判断哈希表中，给定filed是否存在
	fmt.Println(rdb.HExists("user10001", "location"))

	// 列出该哈希表中所有的filed
	fmt.Println(rdb.HKeys("user10001"))

	// 列出该哈希表中所有的value
	fmt.Println(rdb.HVals("user10001"))

	// 列出该哈希表所有内容
	fmt.Println(rdb.HGetAll("user10001"))

	// 为哈希表field的纯数字值加上增量（负数也行）
	fmt.Println(rdb.HIncrBy("user10001", "age", -1))
}

func main() {
	if err := initRedisClient(); err != nil {
		fmt.Printf("init redis client failed, err:%v\n", err)
		return
	}
	fmt.Println("connect redis success...")

	// 释放相关资源
	defer rdb.Close()

	// 字符串以及一些键值对的基本操作
	fmt.Println("************************* sting类型 **********************************")
	stringEx()

	rdb.FlushDB() // 清空当前库

	// list示例
	fmt.Println("************************* list类型 **********************************")
	listEx()

	rdb.FlushDB() // 清空当前库

	// set 示例
	fmt.Println("************************* set类型 **********************************")
	setEx()

	rdb.FlushDB() // 清空当前库

	// zset 示例
	fmt.Println("************************* zset类型 **********************************")
	zsetEx()

	rdb.FlushDB() // 清空当前库

	// hash 示例
	fmt.Println("************************* hash类型 **********************************")
	hashEx()

	//rdb.FlushDB() // 清空当前库
}
