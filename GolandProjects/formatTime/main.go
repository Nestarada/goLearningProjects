package main

import (
	"fmt"
	"strings"
)

func FormatDuration(seconds int64) string {
	var year, day, hour, min, cnt int64
	var res string
	if seconds == 0 {return "now"}
	if seconds / 31536000 > 0 {
		year = seconds/31536000
		seconds -= year * 31536000
		if year > 1 {
			res += fmt.Sprintf("%d years and ", year)
			cnt++
		} else {
			res += fmt.Sprintf("%d year and ", year)
			cnt++
		}
	}
	if seconds / 86400 > 0  {
		day = seconds/86400
		seconds -= day * 86400
		if day > 1 {
			res += fmt.Sprintf("%d days and ", day)
			cnt++
		} else {
			res += fmt.Sprintf("%d day and ", day)
			cnt++
		}
	}
	if seconds / 3600 > 0  {
		hour = seconds/3600
		seconds -= hour * 3600
		if hour > 1 {
			res += fmt.Sprintf("%d hours and ", hour)
			cnt++
		} else {
			res += fmt.Sprintf("%d hour and ", hour)
			cnt++
		}
	}
	if seconds / 60 > 0  {
		min = seconds/60
		seconds -= min * 60
		if min > 1 {
			res += fmt.Sprintf("%d minutes and ", min)
			cnt++
		} else {
			res += fmt.Sprintf("%d minute and ", min)
			cnt++
		}
	}
	if seconds > 1 {
		res += fmt.Sprintf("%d seconds and ", seconds)
		cnt++
	}
	if seconds == 1 {
		res += fmt.Sprintf("%d second and ", seconds)
		cnt++
	}
	res = res[:len(res)-5]
	if cnt == 5 {
		res = strings.Replace(res, " and ", ", ", 3)
	}
	if cnt == 4 {
		res = strings.Replace(res, " and ", ", ", 2)
	}
	if cnt == 3 {
		res = strings.Replace(res, " and ", ", ", 1)
	}
	return res
}

func main() {
	var a, b, c int64 = 100, 1000000, 1000000000
	var res string
	res = FormatDuration(a)
	fmt.Println(res)
	res = FormatDuration(b)
	fmt.Println(res)
	res = FormatDuration(c)
	fmt.Println(res)
}
