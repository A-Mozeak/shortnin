package main

import (
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// GenerateShort creates a 6-character alphanumeric shortlink using a random number generator.
func GenerateShort() string {
	var sb strings.Builder
	alphabet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	seed := rand.NewSource(time.Now().UnixNano())
	randHex := rand.New(seed)
	for i := 0; i < 6; i++ {
		pick := randHex.Intn(62)
		sb.WriteByte(alphabet[pick])
	}
	return sb.String()
}

// UpdateStats updates the statistics for a given shortlink.
func UpdateStats(s *Shorty) {
	if time.Since(s.LastVis) > (86400 * time.Second) {
		s.DailyVis = append(s.DailyVis, s.TodayVis)
		s.TodayVis = 0
	}
	s.TodayVis++
	s.TotalVis++
	s.LastVis = time.Now()
}

// GetQueryParams parses a URL and returns only the query as a set of key:value pairs.
func GetQueryParams(r *http.Request) url.Values {
	u, err := url.Parse(r.URL.RequestURI())
	if err != nil {
		panic(err)
	}
	return u.Query()
}
