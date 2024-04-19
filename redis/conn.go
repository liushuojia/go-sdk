package redisConn

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"net"
	"strconv"
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
		ReadTimeout:  -1,
		WriteTimeout: -1,
	}

	c.Client = redis.NewClient(ops)
	return c
}
func (c *Conn) Connect() *Conn {
	log.Println("[redis]", "connect redis", "host", c.addr, "db", c.db)

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
	log.Println("[subscribe]", "redis", "channel:", topicList, "start")

	for {
		select {
		case message := <-pub.Channel():
			if message != nil {
				log.Println("[subscribe]", "redis", "channel:", message.Channel, "payload", message.Payload)
				//cb(message)
				cb(message.Channel, message.Payload)
			}
		case <-c.ctx.Done():
			goto END
		}
	}
END:
	log.Println("[subscribe]", "redis", "channel:", topicList, "end")
}

/*
sequenceNumString,
workerIDString[len(workerIDString)-2:],
userIDString[len(userIDString)-2:],
dataIDString[len(dataIDString)-2:],

	自增数		workerIDBits	userIDBits		dataIDBits
	sequence	10				7				7
*/
const (
	workerIDBits = uint(10) // 机器id所占的位数
	userIDBits   = uint(7)  // 机器id所占的位数
	dataIDBits   = uint(7)  // 机器id所占的位数

	workerIDMax = int64(-1 ^ (-1 << workerIDBits)) // 支持的最大机器id数量
	userIDMax   = int64(-1 ^ (-1 << userIDBits))   // 支持的最大机器id数量
	dataIDMax   = int64(-1 ^ (-1 << dataIDBits))   // 支持的最大机器id数量

	//左移位数
	userIDShift   = dataIDBits
	workerIDShift = dataIDBits + userIDBits
	sequenceShift = dataIDBits + userIDBits + workerIDBits
)

func (c *Conn) ID(code string, dateTime time.Time, workerID, userID int64, dataID int64) (int64, error) {

	//2024
	//1231
	fmt.Println(int64(-1 ^ (-1 << 12)))
	fmt.Println(strconv.FormatInt(2024, 2), len(strconv.FormatInt(2024, 2)))
	fmt.Println(strconv.FormatInt(1231, 2), len(strconv.FormatInt(1231, 2)))
	panic("")

	key := fmt.Sprintf("%s:%s", code, dateTime.Format("20060102"))
	n, err := c.Exists(c.ctx, key).Result()
	if err != nil {
		return 0, err
	}
	sequenceNum, err := c.Incr(c.ctx, key).Result()
	if err != nil {
		return 0, err
	}
	if n <= 0 {
		//首次添加设置超时时间
		c.Expire(c.ctx, key, time.Hour*24)
	}

	//fmt.Println("============================================================")
	//fmt.Println("workerIDMax", workerIDMax)
	//fmt.Println("userIDMax", userIDMax)
	//fmt.Println("dataIDMax", dataIDMax)

	userID = userID & userIDMax
	dataID = dataID & dataIDMax
	workerID = workerID & workerIDMax

	r := sequenceNum<<sequenceShift | workerID<<workerIDShift | userID<<userIDShift | dataID

	fmt.Println("============================================================")
	fmt.Println("userID", userID)
	fmt.Println("dataID", dataID)
	fmt.Println("workerID", workerID)
	fmt.Println("sequenceNum", sequenceNum)
	fmt.Println("result", r, len(fmt.Sprintf("%d", r)))

	return r, nil

	userIDString := fmt.Sprintf("000000000000%d", userID)
	dataIDString := fmt.Sprintf("000000000000%d", dataID)
	workerIDString := fmt.Sprintf("000000000000%d", workerID)
	sequenceNumString := fmt.Sprintf("%d", sequenceNum)
	if len(sequenceNumString) < 6 {
		sequenceNumString = "000000000000000000000000" + sequenceNumString
		sequenceNumString = sequenceNumString[len(sequenceNumString)-6:]
	}
	numString := fmt.Sprintf("%s%s%s%s%s",
		time.Now().Format("2006"),
		sequenceNumString,
		workerIDString[len(workerIDString)-2:],
		userIDString[len(userIDString)-2:],
		dataIDString[len(dataIDString)-2:],
	)

	fmt.Println(numString, len(numString))
	num, err := strconv.Atoi(numString)
	if err != nil {
		return 0, err
	}

	return int64(num), nil
}
