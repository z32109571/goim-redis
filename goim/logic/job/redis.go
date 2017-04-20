package main

import (
	log "github.com/thinkboy/log4go"
	"gopkg.in/redis.v3"
	"time"
)

func InitRedis() error {
	log.Info("redis is host", Conf.RedisHost)
	log.Info("redis topic is :%s", Conf.RedisTopic)
	client := redis.NewClient(&redis.Options{
		Addr:     Conf.RedisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	go func() {
		for {
			result, err := client.BLPop(1*time.Second, Conf.RedisTopic).Result()
			if err != redis.Nil && err != nil{
				log.Warn("redis blpop err:%s", err.Error())
			}
			if len(result) != 0 {
				msg := result[1]
				log.Info("deal with redis list is %s,msg is %s ", Conf.RedisTopic, msg)
				pushR([]byte(msg))
			}
		}
	}()
	return nil
}
