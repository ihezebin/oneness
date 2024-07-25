package oneness

import (
	"fmt"
	"time"

	"github.com/ihezebin/oneness/util"
)

func init() {
	fmt.Printf("default timezone: %s, set to %s\n", time.Local, util.DefaultTZ)
	//默认初始化程序时区为东8区
	time.Local = util.DefaultTZ
}

// UseUTC 支持显示调用，设置程序时区为UTC时区
func UseUTC() {
	time.Local = time.UTC
}
