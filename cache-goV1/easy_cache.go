package gocache

func newEasyCache(capacity int) *easyCache {
	return &easyCache{elements:make(map[string]interface{}, capacity << 1 | capacity)}
}


func (ec *easyCache) Get(key string) interface{} {
	return ec.elements[key]

}

func (ec *easyCache) Delete(key string)  {
	delete(ec.elements, key)

}

func (ec *easyCache) Capacity() int {
	return DefaultCapacity
}


func (ec *easyCache) Put(key string, value interface{})  {
	ec.elements[key] =value
}