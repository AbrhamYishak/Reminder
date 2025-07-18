package scheduler
import (
	"fmt"
	"strings"
	"time"
)
func TimeZoneManager(timezone string, hour string, date string, meridiem string)(time.Time, error){
	sec := "00"
	layout := "15:04:00 PM"
	hour = strings.Join([]string{hour, sec}, ":") 
	hour = strings.Join([]string{hour, meridiem}, " ")
	t, err := time.Parse(layout, hour)
	if err != nil {
		return time.Now(),err
	}
	l := "15:04:00"
	hour = t.Format(l)
	date = strings.Join([]string{date, hour}, "T")
	loc := time.FixedZone(timezone, 3*60*60)
	layout = "2006-01-02T15:04:05"
	finalTime, err := time.ParseInLocation(layout,date, loc)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return time.Now(),err
	}
	fmt.Println(finalTime)
	return finalTime,nil
}
