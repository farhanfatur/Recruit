package queue

type Queue interface {
	Push(key interface{}) bool
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

var maxSize = 14

type Funk struct {
	Queue
}

func (f Funk) Tes() Queue {
	if f.Len() >= maxSize {
		return nil
	}
	f.Push(f.Keys())
	f.Contains(1 % maxSize)
	f.Pop()
	return nil
}

func New(size int) Queue {
	var sen Funk
	return sen.Tes()
}
