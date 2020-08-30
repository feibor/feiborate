package feiborate

import "time"

// DateGetNowYYYYMMDDhhmmss 日期转换，time
func DateGetNowYYYYMMDDhhmmss(t time.Time) string {
	// GO的奇葩日期格式化，2006-01-02 15:04:05这个每组数字都是有独特的含义 类似于yyyy-mm-dd xxxxx
	return t.Format("2006-01-02 15:04:05")
}

// String2Timestamp 字符串转时间戳
//
// date 待转换时间字符串 如：2019/09/17 10:16:56
//
// format 时间字符串格式化类型 如：2006/01/02 15:04:05
//
// zone 时区 如：time.Local / time.UTC
func String2Timestamp(date, format string, zone *time.Location) (int64, error) {
	var (
		theTime time.Time
		err     error
	)
	if theTime, err = time.ParseInLocation(format, date, zone); nil != err {
		return 0, err
	}
	return theTime.Unix(), nil
}
