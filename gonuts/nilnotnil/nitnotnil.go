//
// Package nilnotnil demonstrates that nil can have a underlying type.
//
package nilnotnil

type Worker interface {
	Work()
}

type WorkerA struct{}

func (a *WorkerA) Work() {}

type WorkerB struct{}

func (b *WorkerB) Work() {}

// NewNilWorker returns nil pointer of specific type
// All return types satisfies Worker interface
func NewNilWorker(workerType string) Worker {
	switch workerType {
	case "A":
		var w *WorkerA // return type is nil pointer to *WorkerA type
		return w
	case "B":
		var w *WorkerB // return type is nil pointer to *WorkerB type
		return w
	default:
		var w Worker // return type is nil interface w/o concrete type
		return w
	}
}
