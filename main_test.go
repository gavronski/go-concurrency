package main

import "testing"

func Test_updateMessage(t *testing.T) {
	wg.Add(1)

	go updateMessage("test", &wg)

	wg.Wait()

	if msg != "test" {
		t.Error("missing expected 'test'")
	}
}
