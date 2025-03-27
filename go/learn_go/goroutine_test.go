package learn_go

import (
	"testing"
)

func TestMain(t *testing.T) {
	go func(t *testing.T) {
		t.Log("goroutine 1")
	}(t)
}
