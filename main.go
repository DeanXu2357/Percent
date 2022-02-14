package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			now := time.Now()
			year := now.Year()
			startOfNextYearTimestamp := time.Date(year+1, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
			startOfYearTimestamp := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).Unix()

			percent := float64(now.Unix()-startOfYearTimestamp) / float64(startOfNextYearTimestamp-startOfYearTimestamp) * 100
			fmt.Printf("\rNow: %s, Percent of %d: %.2f%%", now.Format("2006-01-02 15:04:05"), year, percent)
		}
	}
}
