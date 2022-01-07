package main

import (
	"fmt"
	"net/http"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func test_rate_limit() {
	rate := vegeta.Rate{Freq: 20, Per: time.Second}
	duration := 60 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:8080/notes/5",
		Header: http.Header{"Authorization": {"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDE1Nzg3MzYsImlzX2FkbWluIjpmYWxzZSwidXNlcl9pZCI6MTB9.VqS1mFd33RmaeNAFCwQIC9_ZySJha9_vhpIOLpQzj0U"}},
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}
