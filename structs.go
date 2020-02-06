package main

import (
	"time"
)

type Shorty struct {
	Name     string    `json:"name"`
	Link     string    `json:"link"`
	Created  time.Time `json:"created"`
	LastVis  time.Time `json:"lastVisited"`
	TotalVis int       `json:"totalVisits"`
	TodayVis int       `json:"visitsToday"`
	DailyVis []int     `json:"dailyEngagement"`
}

type MockDB struct {
	longMap  map[string]string
	shortMap map[string]*Shorty
}

// Database Constructor
func NewDB() *MockDB {
	var m MockDB
	m.longMap = make(map[string]string)
	m.shortMap = make(map[string]*Shorty)
	return &m
}
