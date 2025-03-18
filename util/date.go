package utils

import "time"

func GetRecentSevenDays() []string {
	var days []string
	now := time.Now()
	for i := 0; i < 7; i++ {
		day := now.AddDate(0, 0, -i)
		formattedDay := day.Format("2006-01-02")
		days = append(days, formattedDay)
	}
	return days
}
