package stdlib

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestLocation(t *testing.T) {
	zones := []string{"Asia/Shanghai", "Asia/Tokyo", "Asia/Seoul", "UTC", "CST-8", "KST-9"}
	now := time.Now()
	for _, zone := range zones {
		location, err := time.LoadLocation(zone)
		if err != nil {
			fmt.Println("load location err: ", err.Error())
			fmt.Println()
			continue
		}
		fmt.Println("input: ", zone, ", output: ", location.String())
		fmt.Println("local to location: ", now.In(location))
		fmt.Println("local: ", now.Local())
		fmt.Println("location: ", now.Location().String())
		fmt.Println("location: ", now.UTC())
		zone, offset := now.Zone()
		fmt.Println("local zone: ", zone, " ,offset: ", offset)
		lzone, loffset := now.In(location).Zone()
		fmt.Println("location zone: ", lzone, " ,offset: ", loffset)
		fmt.Println()
	}
}

func TestTime(t *testing.T) {
	now := time.Now()
	fmt.Println("now str:, ", now.String())
	fmt.Println("now utc:, ", now.UTC().String())
	hh, mm, ss := now.Clock()
	yy, MM, dd := now.Date()
	fmt.Printf("%4d-%2d-%2d %2d:%2d:%2d\n",
		yy, MM, dd, hh, mm, ss)
	fmt.Printf("%4d-%s-%2d %2d:%2d:%2d\n",
		now.Year(), now.Month().String(), now.Day(), now.Hour(), now.Minute(), now.Second())
	weekday := now.Weekday()
	yearday := now.YearDay()
	month := now.Month()
	fmt.Println("weekday type: ", reflect.TypeOf(weekday).Name(), ", kind: ", reflect.TypeOf(weekday).Kind().String())
	fmt.Println("yearday type: ", reflect.TypeOf(yearday).Name(), ", kind: ", reflect.TypeOf(yearday).Kind().String())
	fmt.Println("month type: ", reflect.TypeOf(month).Name(), ", kind: ", reflect.TypeOf(month).Kind().String())
	fmt.Println("time type: ", reflect.TypeOf(now).Name(), ", kind: ", reflect.TypeOf(now).Kind().String())
	fmt.Println("weekday: ", weekday.String(), int(weekday))
	fmt.Println("yearday: ", yearday)
	fmt.Println("month: ", month.String(), int(month))
	future := now.Add(1*time.Hour + 23*time.Minute)
	past := now.Add(-(1*time.Hour + 23*time.Minute))
	fmt.Println("future: ", future)
	fmt.Println("past: ", past)
	fmt.Println("after future: ", now.After(future), ", before past: ", now.Before(past))
	fmt.Println("add -10 years 2 months and 3 days: ", now.AddDate(-10, 3, 25))
	isoyear, isoweek := now.ISOWeek()
	fmt.Println("iso year: ", isoyear, " ,isoweek: ", isoweek)
	fmt.Println("unix: ", now.Unix(), " , unix nano: ", now.UnixNano())
	fmt.Println("format: ", "2006-01-02 15:04:05", ", formatted: ", now.Format("2006-01-02 3:04:05 PM"))
	fmt.Println("format: ", "2006-01-02 3:04:05 PM", ", formatted: ", now.Format("2006-01-02 15:04:05"))
	fmt.Println("format: ", time.RFC3339, ", formatted: ", now.Format(time.RFC3339))
	date := time.Date(2019, time.October, 31, 4, 41, 7, 123456789, time.Now().Location())
	fmt.Println("date: ", date, date.Format("2006-01-02 3:04:05 PM"))
	fmt.Println("date unix: ", date.Unix(), ", nano: ", date.UnixNano())
	unix := time.Unix(1572468067, 987654321)
	fmt.Println("unix: ", unix, unix.Format("2006-01-02 3:04:05 PM"))
	fmt.Println("unix unix: ", unix.Unix(), ", nano: ", unix.UnixNano())
}

func TestDuration(t *testing.T) {
	now := time.Now()
	ano, _ := time.Parse("2006-01-02 15:04:05", "2019-11-30 14:41:57")
	fmt.Println("now: ", now)
	fmt.Println("ano: ", ano)
	dur := ano.Sub(now)
	fmt.Println(dur, int64(dur))
	rev := now.Sub(ano)
	fmt.Println(rev, int64(rev))
	fmt.Printf("%fh %fm %fs %dns\n", dur.Hours(), dur.Minutes(), dur.Seconds(), dur.Nanoseconds())
	fmt.Printf("%fh %fm %fs %dns\n", rev.Hours(), rev.Minutes(), rev.Seconds(), rev.Nanoseconds())
	dur1 := 6*time.Minute + 84*time.Second + 13*time.Millisecond
	fmt.Println("dur1: ", dur1)
	fmt.Printf("%fh %fm %fs %dns\n", dur1.Hours(), dur1.Minutes(), dur1.Seconds(), dur1.Nanoseconds())
	fmt.Println(reflect.TypeOf(dur1).String(), reflect.TypeOf(dur1).Kind().String())
	fmt.Println("round hour: ", dur.Round(time.Hour).Hours())
	fmt.Println("round minute: ", dur.Round(time.Minute).Minutes())
	fmt.Println("round second: ", dur.Round(time.Second).Seconds())
	fmt.Println("truncate hour: ", dur.Truncate(time.Hour))
	fmt.Println("truncate minute: ", dur.Truncate(time.Minute))
	fmt.Println("truncate second: ", dur.Truncate(time.Second))
	fmt.Println(int(dur.Round(time.Hour)), int(dur.Truncate(time.Hour)))
}

func TestSleep(t *testing.T) {
	fmt.Println(time.Now())
	time.Sleep(3*time.Second)
	fmt.Println(time.Now())
}

func TestAfter(t *testing.T) {
	fmt.Println(time.Now().String())
	ch := time.After(5*time.Second)
	time.Sleep(3 * time.Second)
	fmt.Println(time.Now())
	ti := <- ch
	fmt.Println(ti)
}

func TestTick(t *testing.T) {
	fmt.Println(time.Now())
	ch := time.Tick(2*time.Second)
	for {
		ti := <- ch
		fmt.Println(ti)
	}
}

func TestTimer(t *testing.T) {
	fmt.Println(time.Now())
	timer := time.NewTimer(6*time.Second)
	go func() {
		time.Sleep(3*time.Second)
		timer.Reset(1*time.Second)  // wait for 4 seconds
	}()
	ti := <- timer.C
	fmt.Println(ti)
}

func TestTimer1(t *testing.T) {
	fmt.Println(time.Now())
	timer := time.NewTimer(6*time.Second)
	go func() {
		time.Sleep(3*time.Second)
		timer.Reset(0)  // use timer.Reset(0) to stop waiting
		// can not use timer.Stop() concurrently
	}()
	ti := <- timer.C
	fmt.Println(ti)
}

func TestTicker(t *testing.T) {
	fmt.Println(time.Now())
	ticker := time.NewTicker(2*time.Second)
	for i:=0; i<3; i++ {
		ti := <- ticker.C
		fmt.Println(ti)
	}
	time.Sleep(3*time.Second) // if time.Sleep(3*time.Second), the last receive will block
	ticker.Stop()
	fmt.Println(time.Now())
	fmt.Println("last: ", <-ticker.C)
}