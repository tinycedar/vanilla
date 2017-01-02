package runtime

type thread struct {
	pc    int // length of int is platform dependent
	stack *Stack
}
