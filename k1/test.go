package k1

import "time"

func NowUnixTS() string {
	return "qweqw"
}

func NowUnixTSMill() int64{
	return time.Now().UnixNano() / 100000
}

func NowUnixTSNano() int64{
	return time.Now().UnixNano()
}
