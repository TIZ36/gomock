package base

import (
	"encoding/json"
	"fmt"
)

type Service[T any] struct {
	funcMap map[string]func()
}

type ProxyService[T any] struct {
	Cache  Cache
	Method func() (T, error)
}

type Cache struct {
	Enable   bool
	Cache    CommonCache[string, []byte]
	Cascade  any
	QueryKey string
	label    string
	k        string
	others   string
}

type ServiceContext struct {
}

func (service *Service[T]) Proxy(f func() (T, error)) *ProxyService[T] {
	return &ProxyService[T]{
		Method: f,
	}
}

func (proxyService *ProxyService[T]) Cached(keyCompos ...string) *ProxyService[T] {
	apiKey := buildKey(proxyService.Cache.label, proxyService.Cache.k)

	keyCompos = append(keyCompos, apiKey)

	qk := buildKey(keyCompos...)

	proxyService.Cache.Enable = true
	proxyService.Cache.QueryKey = qk
	proxyService.Cache.Cache = CommonCache[string, []byte]{}
	return proxyService
}

func (proxyService *ProxyService[T]) CancelCached(keyCompos ...string) *ProxyService[T] {
	proxyService.Cache.Enable = false
	return proxyService
}

func (proxyService *ProxyService[T]) Exec() (*T, error) {
	var reBin []byte
	var re T
	if proxyService.Cache.Enable {
		reAny, err := proxyService.Cache.Cache.Get(proxyService.Cache.QueryKey)

		if err == nil {
			reBin, _ := reAny.([]byte)

			json.Unmarshal(reBin, &re)

			return &re, nil
		}

	}

	// 真正执行
	re, err := proxyService.Method()

	if proxyService.Cache.Enable {
		reBin, _ = json.Marshal(re)
		proxyService.Cache.Cache.Put(proxyService.Cache.QueryKey, reBin)
	}

	if err != nil {
		return nil, err
	}

	return &re, nil
}

func buildKey(compo ...string) string {
	k := ""

	for _, c := range compo {
		k += fmt.Sprintf("%s-", c)
	}

	return k[:len(k)-1]
}

func (proxyService *ProxyService[T]) DoCascade(reBin []byte) (*T, error) {
	var re T
	_ = json.Unmarshal(reBin, &re)
	return &re, nil
}
