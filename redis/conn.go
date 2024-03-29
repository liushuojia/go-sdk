package redisConn

import (
	"context"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"time"
)

const (
	poolSize     = 30
	minIdleConns = 6
)

func New() *Conn {
	return (&Conn{}).SetContext(context.Background())
}

//
// type Callback func(message *redis.Message)

// Callback 使用string方便其他包快速使用
type Callback func(channel, message string) error

type Conn struct {
	addr     string
	password string
	db       int

	ctx    context.Context
	cancel context.CancelFunc
	*redis.Client
}

func (c *Conn) ConnectWithSSH(client *ssh.Client) *Conn {
	ops := &redis.Options{
		Addr:     c.addr,
		Password: c.password,
		DB:       c.db,

		Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return client.DialContext(ctx, network, addr)
		},
		ReadTimeout:  -2,
		WriteTimeout: -2,
	}

	c.Client = redis.NewClient(ops)
	return c
}
func (c *Conn) Connect() *Conn {
	ops := &redis.Options{
		Addr:     c.addr,
		Password: c.password,
		DB:       c.db,

		PoolSize:     poolSize,     // 连接池最大socket连接数，默认为5倍CPU数， 5 * runtime.NumCPU
		MinIdleConns: minIdleConns, //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。

		MaxRetries:      0,                      // 命令执行失败时，最多重试多少次，默认为0即不重试
		MinRetryBackoff: 8 * time.Millisecond,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		MaxRetryBackoff: 512 * time.Millisecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔

		//超时
		DialTimeout:  5 * time.Second, //连接建立超时时间，默认5秒。
		ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
		WriteTimeout: 3 * time.Second, //写超时，默认等于读超时，-1表示取消读超时
		PoolTimeout:  4 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。

		//------------------------------------------------------------------------------------------------------
		//ClusterClient管理着一组redis.Client,下面的参数和非集群模式下的redis.Options参数一致，但默认值有差别。
		//初始化时，ClusterClient会把下列参数传递给每一个redis.Client
		//钩子函数
		//仅当客户端执行命令需要从连接池获取连接时，如果连接池需要新建连接则会调用此钩子函数
		OnConnect: func(ctx context.Context, conn *redis.Conn) error {
			//log.Println("redis", "connect", conn)
			return nil
		},
	}

	c.Client = redis.NewClient(ops)
	return c
}

func (c *Conn) Close() {
	c.cancel()
}
func (c *Conn) SetAddr(addr string) *Conn {
	c.addr = addr
	return c
}
func (c *Conn) SetPassword(password string) *Conn {
	c.password = password
	return c
}
func (c *Conn) SetDB(db int) *Conn {
	c.db = db
	return c
}
func (c *Conn) SetContext(ctx context.Context) *Conn {
	c.ctx, c.cancel = context.WithCancel(ctx)
	return c
}

func (c *Conn) GetRedis() *redis.Client {
	return c.Client
}

func (c *Conn) GetByKey(key string) ([]string, error) {
	return c.Client.Do(c.ctx, "KEYS", key).StringSlice() // *prefix*为要查找的Key的前缀部分
}
func (c *Conn) DeleteByKey(prefix_key string) error {
	list, err := c.GetByKey(prefix_key)
	if err != nil {
		return err
	}
	return c.Client.Del(c.ctx, list...).Err()
}

func (c *Conn) Reader(cb Callback, topicList ...string) {
	// There is no error because go-redis automatically reconnects on error.
	pub := c.Subscribe(c.ctx, topicList...)
	defer pub.Close()
	log.Println("[subscribe]", "topic:", topicList, "start")

	for {
		select {
		case message := <-pub.Channel():
			if message != nil {
				log.Println("[subscribe]", "channel:", message.Channel, "payload", message.Payload)
				//cb(message)
				cb(message.Channel, message.Payload)
			}
		case <-c.ctx.Done():
			goto END
		}
	}
END:
	log.Println("[subscribe]", "topic:", topicList, "end")
}

// r.Publish(r.ctx, "topic", "topic a "+strconv.Itoa(i))
