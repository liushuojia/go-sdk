package websocketConn

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"sync"
)

/*
在线用户
	客户端唯一id uuid 与 用户id 关系， 该功能主要在 websocket 或 mqtt回调后端
	默认使用redis存储在线用户

存储数据内容
 1. 在线用户ID集合
 2. 某一个用户的在线的所有uuid
 3. uuid对应的在线用户id
*/

func New() *Online {
	return &Online{}
}

type Online struct {
	idKey   string // 在线用户ID集合
	id2uuid string // 某一个用户ID的所有在线uuid
	uuid2id string // 在线uuid对应的在线用户id
	prefix  string

	kind          string
	client        *redis.Client
	clusterClient *redis.ClusterClient

	onlineMap   sync.Map
	key2uuidMap sync.Map
}

func (o *Online) SetClusterClient(r *redis.ClusterClient) *Online {
	o.kind = "ClusterClient"
	o.clusterClient = r
	return o
}
func (o *Online) SetClient(r *redis.Client) *Online {
	o.kind = "Client"
	o.client = r
	return o
}
func (o *Online) SetPrefix(prefix string) *Online {
	o.prefix = prefix
	return o
}
func (o *Online) SetIdKey(idKey string) *Online {
	o.idKey = idKey
	return o
}
func (o *Online) SetId2uuid(id2uuid string) *Online {
	o.id2uuid = id2uuid
	return o
}
func (o *Online) SetUuid2id(uuid2id string) *Online {
	o.uuid2id = uuid2id
	return o
}

func (o *Online) Create(id string, uuid string) error {
	idKey, id2uuid, uuid2id := o.prefix+o.idKey, o.prefix+o.id2uuid+":"+id, o.prefix+o.uuid2id
	//var rr redis.Client | redis.ClusterClient
	ctx := context.Background()
	switch o.kind {
	case "ClusterClient":
		if r := o.clusterClient; r != nil || r.Ping(ctx).Err() == nil {
			if err := r.SAdd(ctx, idKey, id).Err(); err != nil {
				return err
			}
			if err := r.SAdd(ctx, id2uuid, uuid).Err(); err != nil {
				return err
			}
			if err := r.HSet(ctx, uuid2id, uuid, id).Err(); err != nil {
				return err
			}
			return nil
		}
	case "Client":
		if r := o.client; r != nil || r.Ping(ctx).Err() == nil {
			if err := r.SAdd(ctx, idKey, id).Err(); err != nil {
				return err
			}
			if err := r.SAdd(ctx, id2uuid, uuid).Err(); err != nil {
				return err
			}
			if err := r.HSet(ctx, uuid2id, uuid, id).Err(); err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("redis 未设置")
}
func (o *Online) Clean(uuid string) {
	id, err := o.GetIDByUUID(uuid)
	idKey, id2uuid, uuid2id := o.prefix+o.idKey, o.prefix+o.id2uuid+":"+id, o.prefix+o.uuid2id
	if err != nil && id == "" {
		return
	}
	ctx := context.Background()
	switch o.kind {
	case "ClusterClient":
		if r := o.clusterClient; r != nil || r.Ping(ctx).Err() == nil {
			r.HDel(ctx, uuid2id, uuid)
			if err := r.SRem(ctx, id2uuid, uuid).Err(); err != nil {
				return
			}
			if num, err := r.SCard(ctx, id2uuid).Result(); err != nil || num <= 0 {
				r.SRem(ctx, idKey, id)
			}
			return
		}
	case "Client":
		if r := o.client; r != nil || r.Ping(ctx).Err() == nil {
			r.HDel(ctx, uuid2id, uuid)
			if err := r.SRem(ctx, id2uuid, uuid).Err(); err != nil {
				return
			}
			if num, err := r.SCard(ctx, id2uuid).Result(); err != nil || num <= 0 {
				r.SRem(ctx, idKey, id)
			}
			return
		}
	}
}
func (o *Online) CleanAll() {
	idKey, uuid2id := o.prefix+o.idKey, o.prefix+o.uuid2id
	id2uuidList := []string{
		idKey, uuid2id,
	}
	ctx := context.Background()
	switch o.kind {
	case "ClusterClient":
		if r := o.clusterClient; r != nil || r.Ping(ctx).Err() == nil {
			l, err := r.SMembers(ctx, idKey).Result()
			if err == nil {
				for _, id := range l {
					id2uuidList = append(id2uuidList, o.prefix+o.id2uuid+":"+id)
				}
			}
			r.Del(ctx, id2uuidList...)
		}
	case "Client":
		if r := o.client; r != nil || r.Ping(ctx).Err() == nil {
			l, err := r.SMembers(ctx, idKey).Result()
			if err == nil {
				for _, id := range l {
					id2uuidList = append(id2uuidList, o.prefix+o.id2uuid+":"+id)
				}
			}
			r.Del(ctx, id2uuidList...)
		}
	}
}

func (o *Online) GetAllID() ([]string, error) {
	idKey := o.prefix + o.idKey
	ctx := context.Background()
	switch o.kind {
	case "ClusterClient":
		if r := o.clusterClient; r != nil || r.Ping(ctx).Err() == nil {
			return r.SMembers(ctx, idKey).Result()
		}
	case "Client":
		if r := o.client; r != nil || r.Ping(ctx).Err() == nil {
			return r.SMembers(ctx, idKey).Result()
		}
	}
	return nil, errors.New("redis 未设置")
}
func (o *Online) GetIDByUUID(uuid string) (string, error) {
	uuid2id := o.prefix + o.uuid2id

	ctx := context.Background()
	switch o.kind {
	case "ClusterClient":
		if r := o.clusterClient; r != nil || r.Ping(ctx).Err() == nil {
			return r.HGet(ctx, uuid2id, uuid).Result()
		}
	case "Client":
		if r := o.client; r != nil || r.Ping(ctx).Err() == nil {
			return r.HGet(ctx, uuid2id, uuid).Result()
		}
	}
	return "", errors.New("redis 未设置")
}
func (o *Online) GetUUIDByID(id string) ([]string, error) {
	id2uuid := o.prefix + o.id2uuid + ":" + id

	ctx := context.Background()
	switch o.kind {
	case "ClusterClient":
		if r := o.clusterClient; r != nil || r.Ping(ctx).Err() == nil {
			return r.SMembers(ctx, id2uuid).Result()
		}
	case "Client":
		if r := o.client; r != nil || r.Ping(ctx).Err() == nil {
			return r.SMembers(ctx, id2uuid).Result()
		}
	}
	return nil, errors.New("redis 未设置")
}

func (o *Online) SendMessageById(id string, message []byte) error {
	uuidList, err := o.GetUUIDByID(id)
	if err != nil {
		return err
	}
	for _, uuid := range uuidList {
		o.SendMessageByUUID(uuid, message)
	}
	return nil
}
func (o *Online) SendMessageByUUID(uuid string, message []byte) error {
	err := o.Message(uuid, message)
	if err != nil {
		o.Clean(uuid)
	}
	return err
}

// Init 长连接操作 长连接的uuid与用户的uuid不一样
func (o *Online) Init(c *gin.Context) (*Conn, string, error) {
	conn, err := NewByGin(c)
	if err != nil {
		return nil, "", err
	}

	key := ""
	for i := 0; i <= 100; i++ {
		keyTmp := uuid.New().String()
		if _, ok := o.onlineMap.Load(keyTmp); !ok {
			key = keyTmp
			break
		}
	}
	if key == "" {
		return nil, "", errors.New("唯一id已存在")
	}

	o.onlineMap.Store(key, conn)
	return conn, key, nil
}
func (o *Online) GetConn(key string) (*Conn, error) {
	conn, ok := o.onlineMap.Load(key)
	if !ok {
		return nil, errors.New("连接不存在")
	}
	c, ok := conn.(*Conn)
	if !ok {
		o.onlineMap.Delete(key)
		return nil, errors.New("连接不存在")
	}
	return c, nil
}
func (o *Online) Message(key string, message []byte) error {
	conn, err := o.GetConn(key)
	if err != nil {
		return err
	}

	if err := conn.Write(message); err != nil {
		o.CloseConn(key)
	}
	return nil
}
func (o *Online) CleanConn() {
	o.onlineMap.Range(func(key, value any) bool {
		conn, ok := value.(Conn)
		if !ok {
			o.onlineMap.Delete(key)
			return true
		}
		if err := conn.Write([]byte("HeartBeat")); err != nil {
			conn.Close()
			o.onlineMap.Delete(key)
		}
		return true
	})
}
func (o *Online) CloseConn(key string) {
	value, ok := o.onlineMap.LoadAndDelete(key)
	if !ok {
		return
	}
	if conn, ok := value.(Conn); ok {
		conn.Close()
	}
}

func (o *Online) SetKey2Uuid(key string, uuid string) {
	o.key2uuidMap.Store(key, uuid)
}
func (o *Online) DeleteKey2Uuid(key string) {
	o.key2uuidMap.Delete(key)
}
func (o *Online) GetKey2Uuid(key string) (string, error) {
	v, ok := o.key2uuidMap.Load(key)
	if !ok {
		return "", errors.New("key不存在")
	}
	value, ok := v.(string)
	if !ok {
		o.DeleteKey2Uuid(key)
		return "", errors.New("key不存在")
	}
	return value, nil
}
