package core

type JVM struct {
	heap       *heap
	methodArea *methodArea
}

// class instances and arrays
type heap struct {
}

// per-class structures: constant pool & fileds & methods & codes
type methodArea struct {
}
