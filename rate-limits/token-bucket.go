package main

import (
	"fmt"
	"time"
)

type TokenBucket struct {
	capacity int
	fillRate int
	tokens   int
	lastTime time.Time
}

func NewTokenBucket(capacity, fillRate int) *TokenBucket {
	return &TokenBucket{
		capacity: capacity,
		fillRate: fillRate,
		tokens:   capacity,
		lastTime: time.Now(),
	}
}

func (tb *TokenBucket) AllowRequest(tokens int) bool {
	now := time.Now()
	time_passed := int(now.Sub(tb.lastTime).Seconds()) // the difference in seconds between lastTime and now
	// the min between the difference in seconds * fillRate and the total capacity (aka max per bucket)
	tb.tokens = min(tb.capacity, (tb.tokens + time_passed*tb.fillRate))
	tb.lastTime = now

	// checks if requested toke ns are bigger than the available tokens
	if tb.tokens >= tokens {
		tb.tokens -= tokens
		return true
	}

	return false
}

func main() {
	bucket := NewTokenBucket(10, 1)

	// simulating calls each 0.5 second
	for i := 0; i < 15; i++ {
		fmt.Println(bucket.AllowRequest(1))
		time.Sleep(500 * time.Millisecond)
	}

	// sleeps for the bucket to be filled
	time.Sleep(5 * time.Second)

	fmt.Println(bucket.AllowRequest(1))
}
