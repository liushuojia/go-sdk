package utils

import (
	"fmt"
	"testing"
)

// go test -run TestSnowlake
func TestSnowlake(t *testing.T) {
	s, _ := NewSnowflake(1)

	m := make(map[int64]bool)
	for i := 0; i < 10; i++ {
		n := s.Generate()
		if _, ok := m[n]; ok {
			fmt.Println("存在")
			panic("")
		}
		fmt.Println(n)
	}

	//fmt.Println(len(strconv.Itoa(int(s.Generate()))))
	//
	//workeridBits := uint(10)                         // 机器id所占的位数
	//sequenceBits := uint(12)                         // 序列所占的位数
	//workeridMax := int64(-1 ^ (-1 << workeridBits))  // 支持的最大机器id数量
	//sequenceMask := int64(-1 ^ (-1 << sequenceBits)) // 支持的序列数量
	//
	//fmt.Println("支持的最大机器id数量", workeridMax)
	//fmt.Println("支持的序列数量", sequenceMask)
	//
	//fmt.Println(uint(99))
	//fmt.Println(uint(999))
	//
	//fmt.Println(fmt.Sprintf("%s", time.Now().Format("20060102")))
}

/*
淘宝订单号总共有18位。
淘宝订单号的前14位为序号。
淘宝订单号的第15-16位是买家ID的倒数1-2位。
淘宝订单号的17-18位买家ID的倒数3-4位。
*/
