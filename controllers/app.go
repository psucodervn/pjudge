package controllers

import (
	"github.com/psufighter/pjudge/models"
	"sync"
)

var stats struct {
	lock *sync.Mutex
	all  []models.PStat
	All  func() *[]models.PStat
}

func init() {
	stats.All = func() *[]models.PStat {
		stats.lock.Lock()
		defer stats.lock.Unlock()
		return &stats.all
	}
	models.PStatList(nil, &stats.all)
}
