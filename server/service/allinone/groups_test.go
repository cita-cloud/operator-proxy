package allinone

import (
	"context"
	"testing"
	"time"
)

func TestG(t *testing.T) {
	g, _ := NewGroup(context.Background())

	var args []interface{}
	args = append(args, 2*time.Second)

	for i := 0; i < 3; i++ {
		g.Go(WrappedFunc(sleep, args...))
	}
	err := g.Wait()
	if err != nil {
		t.Fatal(err)
	}

}

func sleep(d time.Duration) error {
	time.Sleep(d)
	return nil
}
