package gocache


func newCommons(fn *func(*commonsData), data *commonsData) *commons {
	return &commons{
		fn:   fn,
		data: data,
	}
}

func newCommonsData(key string, val interface{}, c *Cache, out chan interface{}) *commonsData {
	return &commonsData{
		key:    &key,
		val:    val,
		master: c,
		out:    out,
	}
}

