package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisDb *redis.Client

func main() {
	err := initRedis()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("连接redis成功")

	redisExample()
	redisExample2()
}

func initRedis() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	result, err := redisDb.Ping().Result()
	if err != nil {
		return err
	}

	fmt.Println(result)
	return
}

func redisExample() {
	err := redisDb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	val1, err := redisDb.Get("score").Result()
	if err != nil {
		return
	}

	fmt.Println("scoer", val1)

	val2, err := redisDb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(val2)
	}
}

func redisExample2() {
	//zset
	key := "rank"
	items := []redis.Z{
		redis.Z{
			Score:  99,
			Member: "php",
		},
		redis.Z{
			Score:  96,
			Member: "golang",
		},
		redis.Z{
			Score:  97,
			Member: "python",
		},
		redis.Z{
			Score:  99,
			Member: "java",
		},
	}
	fmt.Println(items)
	//把元素都追加到key中
	num, err := redisDb.ZAdd(key, items...).Result()
	if err != nil {
		return
	}
	fmt.Println(num)

	//加分数
	newScore, err := redisDb.ZIncrBy(key, -1, "java").Result()
	if err != nil {
		return
	}
	fmt.Println(newScore)

	//取分数最高的
	scoreList, err := redisDb.ZRevRangeWithScores(key, 0, 3).Result()
	if err != nil {
		return
	}
	for _, z := range scoreList {
		fmt.Println(z.Member, z.Score)
	}

	//取95到100分的
	option := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err := redisDb.ZRangeByScoreWithScores(key, *option).Result()
	if err != nil {
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}
