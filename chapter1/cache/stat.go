package cache

type Stat struct {
	Count     int64
	KeySize   int64
	ValueSize int64
}

func (this *Stat) add(key string, value []byte) {
	this.Count += 1
	this.KeySize += int64(len(key))
	this.ValueSize += int64(len(value))
}

func (this *Stat) del(key string, value []byte) {
	this.Count -= 1
	this.KeySize -= int64(len(key))
	this.ValueSize -= int64(len(value))
}
