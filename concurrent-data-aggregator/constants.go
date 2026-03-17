package main

import "time"

const (
	MAX_RETRIES    = 3
	MAX_BACKOFF    = 30 * time.Second
	GLOBAL_TIMEOUT = 5 * time.Second
)

type Source struct {
	Name       string
	Latency    time.Duration
	ShouldFail bool
}

type Result struct {
	Source   string
	Err      error
	Data     string
	Duration time.Duration
}

var SourceList = []Source{
	{"AlphaStream", 750 * time.Millisecond, false},
	{"BetaNode", 2100 * time.Millisecond, false},
	{"GammaFeed", 500 * time.Millisecond, true},
	{"DeltaPulse", 3000 * time.Millisecond, false},
	{"EpsilonCache", 1200 * time.Millisecond, false},
}
