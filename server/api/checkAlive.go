package api

import (
	"time"
)

// Need some bug fix
func checkAlive(name string) {
	for {
		k := Kittens[name]
		time.Sleep(1 * time.Second)
		if k.GetAlive() {
			t := time.Since(k.GetLastSeen())
			sleepTime := time.Duration(k.GetSleep()) * time.Second
			if t > sleepTime+5*time.Second {
				k.SetAlive(false)
			}
		}
		return
	}
}
