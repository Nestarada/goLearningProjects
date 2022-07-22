package main

import (
	"fmt"
	"strings"
	"time"
)
func main() {
	input := "12 мин. 13 сек."
	minsecs := strings.Split(input, " ")
	formatedInput := minsecs[0] + "m" + minsecs[2] + "s"
	durTime, _ := time.ParseDuration(formatedInput)
	const now = 1589570165
	curTime := time.Unix(now, 0).UTC().Add(durTime)
	fmt.Println(curTime.Format(time.UnixDate))
}
