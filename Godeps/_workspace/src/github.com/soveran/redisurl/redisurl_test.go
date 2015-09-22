package redisurl_test

import (
	"nyanpass/Godeps/_workspace/src/github.com/garyburd/redigo/redis"
	"nyanpass/Godeps/_workspace/src/github.com/soveran/redisurl"
	"testing"
)

func TestConnect(t *testing.T) {
	c, err := redisurl.Connect()

	if err != nil {
		t.Errorf("Error returned")
	}

	pong, err := redis.String(c.Do("PING"))

	if err != nil {
		t.Errorf("Call to PING returned an error: %v", err)
	}

	if pong != "PONG" {
		t.Errorf("Wanted PONG, got %v\n", pong)
	}
}