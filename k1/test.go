package k1

import "time"

func NowUnixTS() (int64,string) {
	return time.Now().Unix(),"lll"
}

func NowUnixTSMill() int64{
	return time.Now().UnixNano() / 100000
}

func NowUnixTSNano() int64{
	return time.Now().UnixNano()
}
