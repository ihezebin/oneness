package util

import (
	"fmt"
	"testing"
	"time"
)

func TestTimezone(t *testing.T) {
	now := time.Now()
	beforeTime := now.Format(time.DateTime)
	//默认初始化程序时区为东8区
	time.Local = DefaultTZ
	afterTime := now.In(time.Local).Format(time.DateTime)
	fmt.Printf("[Oneness] Now default TimeZone: %s, Set to: %s\n", beforeTime, afterTime)

}
