package runtime

type Slot struct {
	num int32
	ref *Object
}

type Object struct {
	data interface{}
}
