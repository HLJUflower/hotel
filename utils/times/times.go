package times

//时间
import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

// 时间 String => Time
func ToTime(str string) (time.Time, error) {
	p := strings.TrimSpace(str)
	if p == "" {
		return time.Time{}, errors.New("%s不能为空")
	}
	t, err := time.ParseInLocation(TimeFormat, str, time.Local)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// 时间 Time => String
func ToStr(t ...time.Time) string {
	if len(t) == 0 {
		return time.Now().Format(TimeFormat)
	} else {
		return t[0].Format(TimeFormat)
	}
}

// 返回当天剩余时间秒数
func GetRemainSecondsOneDay() time.Duration {
	todayLast := time.Now().Format("2006-01-02") + " 23:59:59"
	todayLastTime, _ := time.ParseInLocation(TimeFormat, todayLast, time.Local)
	remainSecond := time.Duration(todayLastTime.Unix()-time.Now().Local().Unix()) * time.Second
	return remainSecond
}

type JsonTime struct {
	time.Time
}

// 时间反序列化
func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = JsonTime{
		now,
	}
	return
}

// 时间序列化
func (t JsonTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(TimeFormat))
	return []byte(formatted), nil
}

// 当前时间
func (t JsonTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 判断是否为时间
func (t *JsonTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JsonTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
