package until

import (
	"time"
)

var cstZone = time.FixedZone("CST", 8*3600)

func MemberTokenAddValidPeriod() string {

	//fmt.Println(time.Now().In(cstZone).Add(time.Minute * 30).Format("2006-01-02 15:04:05"))
	add := time.Now().In(cstZone).Add(time.Minute * 30)
	//fmt.Println(add.String())
	return add.Format("20060102150405")
}

/**
首页加载的 最热数据时差
*/
func DataTimeDifferenceIndexHome() time.Time {
	t_30 := time.Now().AddDate(0, 0, -30)
	return t_30
}

func DataTimeDifference(day int) time.Time {
	t_30 := time.Now().AddDate(0, 0, day)
	return t_30
}
