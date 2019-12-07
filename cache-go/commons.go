package gocache

func newCommons(fn func()) commons {
	return commons{
		fn:  fn,
	}
}

func newCommonsFunc(fn func( string,  *item , *Cache, chan interface{}) interface{}) commons {
	return commons{
		fn:    nil,
		setFn: fn,
	}
}
