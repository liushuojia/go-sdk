package utils

import (
	"errors"
	"sync"
)

func WaitError(funcList ...func() error) error {
	return NewWait().AddFuncErr(funcList...).Exec().Error()
}

func Wait(funcList ...func()) {
	NewWait().Add(funcList...).Exec()
}

func NewWait() *waitGroup {
	return &waitGroup{
		i: 0,
	}
}

type waitGroup struct {
	i    int64
	lock sync.Mutex
	w    sync.WaitGroup
	m    sync.Map
	err  sync.Map
}

func (w *waitGroup) index() int64 {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.i += 1
	return w.i
}
func (w *waitGroup) Add(ll ...func()) *waitGroup {
	for _, fn := range ll {
		w.w.Add(1)
		w.m.Store(w.index(), fn)
	}
	return w
}
func (w *waitGroup) AddFuncErr(ll ...func() error) *waitGroup {
	for _, fn := range ll {
		w.w.Add(1)
		w.m.Store(w.index(), fn)
	}
	return w
}
func (w *waitGroup) Exec() *waitGroup {
	w.m.Range(func(key, value interface{}) bool {
		if fn, ok := value.(func() error); ok {
			go func() {
				defer w.w.Done()
				w.err.Store(key, fn())
			}()
			goto END
		}
		if fn, ok := value.(func()); ok {
			go func() {
				defer w.w.Done()
				fn()
			}()
			goto END
		}
		defer w.w.Done()
	END:
		return true
	})
	w.w.Wait()
	return w
}
func (w *waitGroup) Error() error {
	var err error
	w.err.Range(func(key, value interface{}) bool {
		errValue, ok := value.(error)
		if !ok || errValue == nil {
			goto END
		}
		err = errors.Join(err, errValue)
	END:
		return true
	})

	return err
}
