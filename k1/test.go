package k1

import (
      "time"
      "github.com/lihaifeng111/gitskills/k2"
)

func NowUnixTS() (int64,string,string) {
	return time.Now().Unix(),"666666",k2.GetUserMessage()
}

func NowUnixTSMill() int64{
	return time.Now().UnixNano() / 100000
}

func NowUnixTSNano() int64{
	return time.Now().UnixNano()
}
