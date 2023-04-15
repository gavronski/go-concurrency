package main

import (
	"app/waitGroup"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	waitGroup.Wg.Add(1)

	go waitGroup.UpdateMessage("test", &waitGroup.Wg)

	waitGroup.Wg.Wait()

	if waitGroup.Msg != "test" {
		t.Error("missing expected 'test'")
	}
}
