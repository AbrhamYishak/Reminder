package scheduler
import (
	"strings"
	"time"
	"fmt"
	"strconv"
)
func parseUTCOffset(offsetStr string) (int, error) {
    if !strings.HasPrefix(offsetStr, "UTC") {
        return 0, fmt.Errorf("invalid offset format")
    }

    sign := offsetStr[3]
    parts := strings.Split(offsetStr[4:], ":")
    if len(parts) != 2 {
        return 0, fmt.Errorf("invalid offset format")
    }

    hour, err := strconv.Atoi(parts[0])
    if err != nil {
        return 0, err
    }

    min, err := strconv.Atoi(parts[1])
    if err != nil {
        return 0, err
    }

    offset := hour*3600 + min*60
    if sign == '-' {
        offset = -offset
    }

    return offset, nil
}
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
    offset, err := parseUTCOffset(timezone)
    if err != nil {
    	return time.Now(), err
	}
	loc := time.FixedZone(timezone, offset)
	layout = "2006-01-02T15:04:05"
	finalTime, err := time.ParseInLocation(layout,date, loc)
	if err != nil {
		return time.Now(),err
	}
	return finalTime,nil
}
