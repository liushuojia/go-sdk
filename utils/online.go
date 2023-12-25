package utils

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
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

type Online struct {
	idKey   string // 在线用户ID集合
	id2uuid string // 某一个用户ID的所有uuid
	uuid2id string // 对象中uuid对应的在线用户id
	prefix  string

	client        *redis.Client
	clusterClient *redis.ClusterClient
}

func (o *Online) SetClusterClient(r *redis.ClusterClient) *Online {
	o.clusterClient = r
	return o
}
func (o *Online) SetClient(r *redis.Client) *Online {
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
	}
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
}
func (o *Online) CleanAll() {
	idKey, uuid2id := o.prefix+o.idKey, o.prefix+o.uuid2id
	id2uuidList := []string{
		idKey, uuid2id,
	}
	ctx := context.Background()
	if r := o.client; r != nil || r.Ping(ctx).Err() == nil {
		l, err := r.SMembers(ctx, idKey).Result()
		if err == nil {
			for _, id := range l {
				id2uuidList = append(id2uuidList, o.prefix+o.id2uuid+":"+id)
			}
		}
		r.Del(ctx, id2uuidList...)
	}
	if r := o.clusterClient; r != nil || r.Ping(ctx).Err() == nil {
		l, err := r.SMembers(ctx, idKey).Result()
		if err == nil {
			for _, id := range l {
				id2uuidList = append(id2uuidList, o.prefix+o.id2uuid+":"+id)
			}
		}
		r.Del(ctx, id2uuidList...)
	}
}
func (o *Online) GetAllID() ([]string, error) {
	idKey := o.prefix + o.idKey
	ctx := context.Background()
	if r := o.client; r != nil || r.Ping(ctx).Err() == nil {
		return r.SMembers(ctx, idKey).Result()
	}
	if r := o.clusterClient; r != nil || r.Ping(ctx).Err() == nil {
		return r.SMembers(ctx, idKey).Result()
	}
	return nil, errors.New("redis 未设置")
}
func (o *Online) GetIDByUUID(id string) (string, error) {
	uuid2id := o.prefix + o.uuid2id

	ctx := context.Background()
	if r := o.client; r != nil || r.Ping(ctx).Err() == nil {
		return r.HGet(ctx, uuid2id, id).Result()
	}
	if r := o.clusterClient; r != nil || r.Ping(ctx).Err() == nil {
		return r.HGet(ctx, uuid2id, id).Result()
	}
	return "", errors.New("redis 未设置")
}
func (o *Online) GetUUIDByID(id string) ([]string, error) {
	id2uuid := o.prefix + o.id2uuid + ":" + id

	ctx := context.Background()
	if r := o.client; r != nil || r.Ping(ctx).Err() == nil {
		return r.SMembers(ctx, id2uuid).Result()
	}
	if r := o.clusterClient; r != nil || r.Ping(ctx).Err() == nil {
		return r.SMembers(ctx, id2uuid).Result()
	}
	return nil, errors.New("redis 未设置")
}
