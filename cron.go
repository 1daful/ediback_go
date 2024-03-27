package main

import (
	"time"
	"github.com/go-co-op/gocron"
)

func StartSchedule(interval int, timeUnit TimeUnit, startDelay int,  task interface{}, params ...interface{}) {
	s := gocron.NewScheduler(time.UTC)
	
	switch timeUnit {
		case "second": // for a seocnd interval
			s.Every(interval).Second().At(startDelay).Do(task, params...)
			break
		case "minute": // for a minute interval
			s.Every(interval).Minute().At(startDelay)
			break
		case "hour": // for a hour interval
			s.Every(interval).Hour().At(startDelay)
			break
		case "day": // for a daily interval
			s.Every(interval).Day().At(startDelay)
			break
		case "week": // for a weekly interval
			s.Every(interval).Week().At(startDelay)
			break
		case "month": // for a monthly interval
			s.Every(interval).Month().At(startDelay)
			break
		// Add more cases for different intervals as needed
		default:
			s.Every(1).Day() // Default to daily interval
			break
		}

	//s.Every(interval).timeUnit().At(timeToStart).Do(sendMessage, user)
	s.StartAsync()
}


type TimeUnit string

const (
    TimeUnitSecond    TimeUnit = "second"
    TimeUnitMinute TimeUnit = "minute"
    TimeUnitHour   TimeUnit = "hour"
    TimeUnitDay    TimeUnit = "day"
    TimeUnitWeek TimeUnit = "week"
    TimeUnitMonth   TimeUnit = "month"
)