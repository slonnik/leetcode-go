package concurrent

import (
	"fmt"
	"sync"
	"testing"
)

func TestCaseOne(t *testing.T) {
	var gr = sync.WaitGroup{}
	var data = "Hello"
	gr.Add(1)
	go func() {
		defer gr.Done()
		data = "Bye"
	}()
	gr.Wait()
	fmt.Println(data)
}
