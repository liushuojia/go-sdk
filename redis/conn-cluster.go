package redisConn

import (
	"context"
	"github.com/liushuojia/go-sdk/ssh"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"net"
	"time"
)

func NewCluster() *ConnCluster {
	return (&ConnCluster{}).SetContext(context.Background())
}

type ConnCluster struct {
	addrs    []string
	password string

	SSH ssh.SSH

	ctx    context.Context
	cancel context.CancelFunc
	*redis.ClusterClient
}

func (c *ConnCluster) Connect() *ConnCluster {
	ops := &redis.ClusterOptions{
		Addrs:    c.addrs,
		Password: c.password,

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
	if c.SSH.Host != "" {
		sshClient, err := c.SSH.Dial()
		if err == nil {
			ops.Dialer = func(ctx context.Context, network, addr string) (net.Conn, error) {
				return sshClient.Dial(network, addr)
			}
		}
	}
	c.ClusterClient = redis.NewClusterClient(ops)
	return c
}
func (c *ConnCluster) Close() {
	c.cancel()
}

func (c *ConnCluster) SetSSH(s ssh.SSH) *ConnCluster {
	c.SSH = s
	return c
}
func (c *ConnCluster) SetAddr(addrs ...string) *ConnCluster {
	c.addrs = addrs
	return c
}
func (c *ConnCluster) SetPassword(password string) *ConnCluster {
	c.password = password
	return c
}
func (c *ConnCluster) SetContext(ctx context.Context) *ConnCluster {
	c.ctx, c.cancel = context.WithCancel(ctx)
	return c
}

func (c *ConnCluster) GetRedis() *redis.ClusterClient {
	return c.ClusterClient
}

func (c *ConnCluster) GetByKey(key string) ([]string, error) {
	return c.ClusterClient.Do(c.ctx, "KEYS", key).StringSlice() // *prefix*为要查找的Key的前缀部分
}
func (c *ConnCluster) DeleteByKey(prefix_key string) error {
	list, err := c.GetByKey(prefix_key)
	if err != nil {
		return err
	}
	return c.ClusterClient.Del(c.ctx, list...).Err()
}

func (c *ConnCluster) Reader(cb Callback, topicList ...string) {
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
