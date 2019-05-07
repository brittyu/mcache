package cache

import "log"

func New(typ string) Cache {
	var c Cache
	if typ == "inmemory" {
		c = newInMemoryCache()
	}

	if c == nil {
		panic("unknow cache type " + typ)
	}

	log.Println(typ, "ready to serve")
	return c
}
