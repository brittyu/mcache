package cache

import "sync"

type inMemoryCache struct {
	c     map[string][]byte
	mutex sync.RWMutex
	Stat
}

func (this *inMemoryCache) Set(key string, value []byte) error {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	tmp, exist := this.c[key]
	if exist {
		this.del(key, tmp)
	}
	this.c[key] = value
	this.add(key, value)
	return nil
}

func (this *inMemoryCache) Get(key string) ([]byte, error) {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.c[key], nil
}

func (this *inMemoryCache) Del(key string) error {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	value, exist := this.c[key]
	if exist {
		delete(this.c, key)
		this.del(key, value)
	}

	return nil
}

func (this *inMemoryCache) GetStat() Stat {
	return this.Stat
}

func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{make(map[string][]byte), sync.RWMutex{}, Stat{}}
}
