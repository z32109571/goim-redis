package main

import (
	"gopkg.in/redis.v3"
	"goim/libs/proto"
	"goim/libs/define"
	"encoding/json"
	log "github.com/thinkboy/log4go"
)

var (
	client *redis.Client
	topicList string
)

func InitRedis(host string,topic string) error {
	client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	topicList = topic
	return nil
}

func mpushRedis(serverId int32, keys []string, msg []byte) (err error) {
	var (
		vBytes []byte
		v      = &proto.RedisMsg{OP: define.Redis_MESSAGE_MULTI, ServerId: serverId, SubKeys: keys, Msg: msg}
	)
	if vBytes, err = json.Marshal(v); err != nil {
		log.Warn("mpushRedis json err:%s", err.Error())
		return 
	}
	if err := client.RPush(topicList, string(vBytes)).Err(); err != nil {
		log.Warn("mpushRedis push err:%s", err.Error())
		return err
	}
	return nil
}

func broadcastRedis(msg []byte) (err error) {
	var (
		vBytes []byte
		v      = &proto.RedisMsg{OP: define.Redis_MESSAGE_BROADCAST, Msg: msg}
	)
	if vBytes, err = json.Marshal(v); err != nil {
		return
	}
	if err := client.RPush(topicList, string(vBytes)).Err(); err != nil {
		return err
	}
	return
}

func broadcastRoomRedis(rid int32, msg []byte, ensure bool) (err error) {
	var (
		vBytes   []byte
		v        = &proto.RedisMsg{OP: define.Redis_MESSAGE_BROADCAST_ROOM, RoomId: rid, Msg: msg, Ensure: ensure}
	)
	if vBytes, err = json.Marshal(v); err != nil {
		return
	}
	if err := client.RPush(topicList, string(vBytes)).Err(); err != nil {
		return err
	}
	return
}
