package feiborate

import (
  "bytes"
  "encoding/json"
  "fmt"
  "hash/crc32"
  "math/rand"
  "strings"
)

// StringIsEmpty 判断字符串是否为空，是则返回true，否则返回false
func StringIsEmpty(str string) bool {
  return len(str) == 0
}

// StringIsNotEmpty 和 IsEmpty 的语义相反
func StringIsNotEmpty(str string) bool {
  return !StringIsEmpty(str)
}

// StringConvert 下划线转换，首字母小写变大写，
// 下划线去掉并将下划线后的首字母大写
func StringConvert(oriString string) string {
  cb := []byte(oriString)
  em := make([]byte, 0, 10)
  b := false
  for i, by := range cb {
    // 首字母如果是小写，则转换成大写
    if i == 0 && (97 <= by && by <= 122) {
      by = by - 32
    } else if by == 95 {
      // 下一个单词要变成大写
      b = true
      continue
    }
    if b {
      if 97 <= by && by <= 122 {
        by = by - 32
      }
      b = false
    }
    em = append(em, by)
  }
  return string(em)
}

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// StringRandSeq 创建指定长度的随机字符串
func StringRandSeq(n int) string {
  b := make([]rune, n)
  for i := range b {
    b[i] = letterRunes[rand.Intn(len(letterRunes))]
  }
  return string(b)
}

// StringRandSeq16 创建长度为16的随机字符串
func StringRandSeq16() string {
  return StringRandSeq(16)
}

// StringTrim 去除字符串中的空格和换行符
func StringTrim(str string) string {
  // 去除空格
  str = strings.Replace(str, " ", "", -1)
  return StringTrimN(str)
}

// StringTrimN 去除字符串中的换行符
func StringTrimN(str string) string {
  // 去除换行符
  str = strings.Replace(str, "\n", "", -1)
  return str
}

// ToString 将对象格式化成字符串
func ToString(i interface{}) string {
  b, err := json.Marshal(i)
  if err != nil {
    return fmt.Sprintf("%+v", i)
  }
  var out bytes.Buffer
  err = json.Indent(&out, b, "", "    ")
  if err != nil {
    return fmt.Sprintf("%+v", i)
  }
  return out.String()
}

// StringSingleValue 将字符串内所有连续value替换为单个value
func StringSingleValue(res string, value string) string {
  doubleValue := StringBuild(value, value)
  for skip := false; !skip; {
    resNew := strings.Replace(res, doubleValue, value, -1)
    if res == resNew {
      skip = true
    }
    res = resNew
  }
  return res
}

// StringSingleSpace 将字符串内所有连续空格替换为单个空格
func StringSingleSpace(res string) string {
  return StringSingleValue(res, " ")
}

// StringPrefixSupplementZero 当字符串长度不满足时，将字符串前几位补充0
//
// str 字符串内容
//
// offset 字符串期望的长度
func StringPrefixSupplementZero(str string, offset int) string {
  backZero := offset - len(str)
  if backZero <= 0 {
    return str
  }
  for i := 0; i < backZero; i++ {
    str = strings.Join([]string{"0", str}, "")
  }
  return str
}

// SubString 截取字符串
func SubString(res string, start, end int) string {
  if start > end || start < 0 || end > len(res) {
    return ""
  }
  return res[start:end]
}

// StringBuild 拼接字符串
func StringBuild(arrString ...string) string {
  return strings.Join(arrString, "")
}

// StringBuildSep 拼接字符串
func StringBuildSep(sep string, arrString ...string) string {
  return strings.Join(arrString, sep)
}

// HashCode 获取字符串的hashcode值
func HashCode(s string) int {
  v := int(crc32.ChecksumIEEE([]byte (s)))
  if v >= 0 {
    return v
  }
  if -v >= 0 {
    return -v
  }
  return 0
}
