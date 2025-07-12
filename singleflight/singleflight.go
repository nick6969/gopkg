package singleflight

import (
	"errors"
	"time"

	"golang.org/x/sync/singleflight"
)

type Group interface {
	Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool)
	DoChan(key string, fn func() (interface{}, error)) <-chan singleflight.Result
}

type Job[T any] struct {
	WorkIdentify string
	CacheGetter  func() (*T, error)
	CacheSetter  func(*T) error
	OnceGetter   func() (*T, error)
}

func (job Job[T]) DoWith(engine Group) (*T, error) {
	v, e := job.CacheGetter()
	if e == nil {
		return v, nil
	}

	m, err, _ := engine.Do(job.WorkIdentify, func() (any, error) {
		v, e = job.OnceGetter()
		if e != nil {
			return nil, e
		}

		e = job.CacheSetter(v)
		if e != nil {
			return nil, e
		}

		return v, nil
	})

	if err != nil {
		return nil, err
	}

	value, ok := m.(*T)
	if !ok {
		return nil, errors.New("type assertion failed")
	}

	return value, nil
}

func (job Job[T]) DoWithTimeout(engine Group, timeout time.Duration) (*T, error) {
	v, e := job.CacheGetter()
	if e == nil {
		return v, nil
	}

	ch := engine.DoChan(job.WorkIdentify, func() (any, error) {
		v, e := job.OnceGetter()
		if e != nil {
			return nil, e
		}

		e = job.CacheSetter(v)
		if e != nil {
			return nil, e
		}

		return v, nil
	})

	select {
	case res := <-ch:
		if res.Err != nil {
			return nil, res.Err
		}

		value, ok := res.Val.(*T)
		if !ok {
			return nil, errors.New("type assertion failed")
		}

		return value, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}
