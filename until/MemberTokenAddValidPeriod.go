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
