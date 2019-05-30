package k1

import "time"

func NowUnixTS() int64 {
	return time.Now().Unix()
}

func NowUnixTSMill() int64{
	return time.Now().UnixNano() / 100000
}

func NowUnixTSNano() int64{
	return time.Now().UnixNano()
}
