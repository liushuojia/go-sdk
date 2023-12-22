package token

import (
	"fmt"
	"testing"
	"time"
)

// 目录下 go test
func TestWJT(t *testing.T) {
	t.Log("wjt test")
	two()
}

type d struct {
	OID int    `json:"oid,omitempty"`
	A   string `json:"a,omitempty"`
	B   string `json:"b,omitempty"`
}

func single() {
	type d struct {
		OID int    `json:"oid,omitempty"`
		A   string `json:"a,omitempty"`
		B   string `json:"b,omitempty"`
	}

	o := New().SetID(10).SetData(d{
		OID: 10086,
		A:   "a",
		B:   "b",
	})

	token, err := o.Generate()

	fmt.Println("+================================+")
	fmt.Println(token, err)
	fmt.Println(o.Payload(token))

	rToken, err := o.Parse(token)
	fmt.Println("+================================+")
	fmt.Println(rToken, err)

	var DD d
	err = o.LoadData(&DD)
	fmt.Println("data", err, DD)

}

func two() {

	o := NewDouble().SetData(d{
		OID: 10086,
		A:   "a",
		B:   "b",
	})
	o.GetAccess().SetID(10)

	accessToken, refreshToken, expireTime, err := o.Generate()

	fmt.Println("+================================+")
	fmt.Println("accessToken", accessToken)
	fmt.Println("refreshToken", refreshToken)
	fmt.Println("expireTime", expireTime)
	fmt.Println("err", err)
	fmt.Println("")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("+================================+")
	fmt.Println("accessToken Payload")
	fmt.Println(o.GetAccess().Payload(accessToken))
	fmt.Println("")
	fmt.Println("refreshToken Payload")
	fmt.Println(o.GetRefresh().Payload(refreshToken))

	fmt.Println("+================================+")
	rToken, err := o.GetAccess().Parse(accessToken)
	fmt.Println(rToken, err)

	var DD d
	err = o.GetAccess().LoadData(&DD)
	fmt.Println("data", err, DD)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	time.Sleep(time.Second * 5)

	tt, tr, te, terr := o.Refresh(refreshToken)
	fmt.Println("+================================+")
	fmt.Println("accessToken", tt)
	fmt.Println("refreshToken", tr)
	fmt.Println("expireTime", te)
	fmt.Println("err", terr)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("+================================+")
	fmt.Println("accessToken Payload")
	fmt.Println(o.GetAccess().Payload(accessToken))
	fmt.Println("")
	fmt.Println("refreshToken Payload")
	fmt.Println(o.GetRefresh().Payload(refreshToken))

	fmt.Println("+================================+")
	rToken, err = o.GetAccess().Parse(accessToken)
	fmt.Println(rToken, err)

	err = o.GetAccess().LoadData(&DD)
	fmt.Println("data", err, DD)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 2)
		tt, tr, te, terr = o.Refresh(refreshToken)
		fmt.Println("")
		fmt.Println("accessToken", tt)
		fmt.Println("refreshToken", tr)
		fmt.Println("expireTime", te)
		fmt.Println("err", terr)
	}

}
