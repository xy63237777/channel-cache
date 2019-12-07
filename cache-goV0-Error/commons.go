package gocache

func newCommons(fn func()) commons {
	return commons{
		fn:  fn,
	}
}

