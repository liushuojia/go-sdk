package utils

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

/*
* 1                                          34              44                 54
* +------------------------------------------+----------------+-------------------+
* | timestamp(s)                             | workerid       | sequence          |
* +------------------------------------------+----------------+-------------------+
* | 0000000000 0000000000 0000000000 000     | 0000000000     | 0000000000        |
* +------------------------------------------+----------------+-------------------+
*
* 1. 31位时间截，注意这是时间截的差值（当前时间截 - 开始时间截)。可以使用约272年: fmt.Println((1 << 33) / (60 * 60 * 24 * 365))
* 2. 10位数据机器位，可以部署在1024节点
* 3. 12位序列，精度时间内（毫秒内或秒）的计数，同一机器，同一时间截并发4096个序号
	now := time.Now().Unix()
	nowLast, _ := time.Parse("2006-01-02", "2200-12-30")

	tmpTT := nowLast.Unix() - now
	fmt.Println(tmpTT)

	时间戳长度33理论上够用

	缺点： 服务器时间回调可能会存在相同的id
*/

const (
	twe = 1000 // 时间戳精度 可以控制id的长度 10整数 js 最长53位  毫秒 / twe

	/*
		var t, _ = time.Parse("2006-01-02 15:03:04", "2024-01-01 00:00:00")
		var twepoch = t.UnixNano() / int64(time.Millisecond) / twe
	*/
	twepoch        = 1704067200000000000              // 开始时间截 (2021-01-01 00:00:00)  毫秒 1609430400000 秒 1609430400
	workerIDBits   = uint(10)                         // 机器id所占的位数
	sequenceBits   = uint(12)                         // 序列所占的位数
	workerIDMax    = int64(-1 ^ (-1 << workerIDBits)) // 支持的最大机器id数量
	sequenceMax    = int64(-1 ^ (-1 << sequenceBits)) // 支持的序列数量
	workerIDShift  = sequenceBits                     // 机器id左移位数
	timestampShift = sequenceBits + workerIDBits      // 时间戳左移位数
)

// NewSnowflake
// NewNode returns a new snowflake worker that can be used to generate snowflake IDs
func NewSnowflake(workerID int64) (*Snowflake, error) {
	if workerID < 0 || workerID > workerIDMax {
		return nil, errors.New(fmt.Sprintf("workerid must be between 0 and %d", workerIDMax))
	}

	return &Snowflake{
		timestamp: 0,
		workerID:  workerID,
		sequence:  0,
	}, nil
}

// A Snowflake struct holds the basic information needed for a snowflake generator worker
type Snowflake struct {
	sync.Mutex
	timestamp int64
	workerID  int64
	sequence  int64
}

// Generate creates and returns a unique snowflake ID
func (s *Snowflake) Generate() int64 {
	s.Lock()
	defer s.Unlock()

	now := s.timeGen()
	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & sequenceMax

		if s.sequence == 0 {
			for now <= s.timestamp {
				now = s.timeGen()
			}
		}
	} else {
		s.sequence = 0
	}

	s.timestamp = now

	r := (now-s.timePoch())<<timestampShift | (s.workerID << workerIDShift) | (s.sequence)
	return r
}
func (s *Snowflake) timeGen() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond) / twe
}
func (s *Snowflake) timePoch() int64 {
	return twepoch / int64(time.Millisecond) / twe
}
