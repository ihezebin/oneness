package util

import "time"

// DefaultTZ 不使用LoadLocation()，它依赖于 IANA Time Zone Database这个数据库，一般linux系统都自带，如果no-linux系统没有带，调用LoadLocation就会报错。
var DefaultTZ *time.Location = time.FixedZone("Asia/Shanghai", 8*3600)
