package main

import (
	"goim/libs/define"
	"strconv"
	log "github.com/thinkboy/log4go"
	"goim/libs/proto"
	"encoding/json"
)

// developer could implement "Auth" interface for decide how get userId, or roomId
type Auther interface {
	Auth(token string) (userId int64, roomId int32)
}

type DefaultAuther struct {
}

func NewDefaultAuther() *DefaultAuther {
	return &DefaultAuther{}
}

func (a *DefaultAuther) Auth(token string) (userId int64, roomId int32) {
	var err error
	info := new(proto.Info)
	err = json. Unmarshal ( []byte(token) , info )
	if err != nil{
		log.Warn("json unmarshal is  %v ", err)
	}
	if userId, err = strconv.ParseInt(info.Uid, 10, 64); err != nil {
		log.Info("token err is  %v ", err)
		userId = 0
		roomId = define.NoRoom
	} else {
		//roomId = 1 // only for debug
		if info.Rid == 0{
			roomId = define.NoRoom
		}else{
			roomId = info.Rid
		}
	}
	return
}
