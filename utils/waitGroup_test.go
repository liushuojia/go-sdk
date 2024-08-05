package utils

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	var err error

	err = WaitError()
	assert.Nil(t, err)

	err = WaitError(func() error {
		time.Sleep(2 * time.Second)
		fmt.Println("1")
		return nil
	}, func() error {
		time.Sleep(3 * time.Second)
		fmt.Println("2")
		return nil
	})
	assert.Nil(t, err)

	err = WaitError(func() error {
		fmt.Println("a001")
		return errors.New("aaa")
	}, func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("a")
		return errors.New("aaa")
	}, func() error {
		time.Sleep(2 * time.Second)
		fmt.Println("b")
		return nil
	})
	assert.NotNil(t, err)

	Wait(func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1")
	}, func() {
		time.Sleep(3 * time.Second)
		fmt.Println("2")
	})

	Wait(func() {
		fmt.Println("a001")
	}, func() {
		time.Sleep(1 * time.Second)
		fmt.Println("a")
	}, func() {
		time.Sleep(2 * time.Second)
		fmt.Println("b")
	})
	assert.NotNil(t, err)

}
