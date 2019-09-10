package cache

import "testing"

func TestInitRedis(t *testing.T) {
	err := InitRedis()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}
