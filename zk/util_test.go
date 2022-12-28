package zk

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

//
//func TestFormatServers(t *testing.T) {
//	t.Parallel()
//	servers := []string{"127.0.0.1:2181", "127.0.0.42", "127.0.42.1:8811"}
//	r := []string{"127.0.0.1:2181", "127.0.0.42:2181", "127.0.42.1:8811"}
//	for i, s := range FormatServers(servers) {
//		if s != r[i] {
//			t.Errorf("%v should equal %v", s, r[i])
//		}
//	}
//}
//
//func TestValidatePath(t *testing.T) {
//	tt := []struct {
//		path  string
//		seq   bool
//		valid bool
//	}{
//		{"/this is / a valid/path", false, true},
//		{"/", false, true},
//		{"", false, false},
//		{"not/valid", false, false},
//		{"/ends/with/slash/", false, false},
//		{"/sequential/", true, true},
//		{"/test\u0000", false, false},
//		{"/double//slash", false, false},
//		{"/single/./period", false, false},
//		{"/double/../period", false, false},
//		{"/double/..ok/period", false, true},
//		{"/double/alsook../period", false, true},
//		{"/double/period/at/end/..", false, false},
//		{"/name/with.period", false, true},
//		{"/test\u0001", false, false},
//		{"/test\u001f", false, false},
//		{"/test\u0020", false, true}, // first allowable
//		{"/test\u007e", false, true}, // last valid ascii
//		{"/test\u007f", false, false},
//		{"/test\u009f", false, false},
//		{"/test\uf8ff", false, false},
//		{"/test\uffef", false, true},
//		{"/test\ufff0", false, false},
//	}
//
//	for _, tc := range tt {
//		err := validatePath(tc.path, tc.seq)
//		if (err != nil) == tc.valid {
//			t.Errorf("failed to validate path %q", tc.path)
//		}
//	}
//}

func TestText(t *testing.T) {
	str := "Server {\n\norg.apache.zookeeper.server.auth.DigestLoginModule required\n\nuser_beacon_datallink=\"(%@&BEacon39\";\n\n};"
	t.Log(str)
	var i int64
	i = 12343566
	str1 := strconv.FormatInt(i, 10)

	t.Log(str1)
	cosUrl := "http://tse-migrate-cd-1306669067.cos.ap-chengdu.myqcloud.com"
	startIndex := strings.LastIndex(cosUrl, "//")
	if startIndex == -1 {
		startIndex = 0
	} else {
		startIndex += 2
	}
	endIndex := strings.LastIndex(cosUrl, ".cos")
	t.Log(cosUrl[startIndex:endIndex])

	t.Log(MD5File("/Users/pengkan/Downloads/VSCode-darwin-universal.zip"))

	s := "12345/ins-1231/dd.txt"
	endIndex = len(s)
	startIndex = strings.LastIndex(s, "/") + 1
	t.Log(s[startIndex:endIndex])

}
func MD5Bytes(s []byte) string {
	ret := md5.Sum(s)
	return hex.EncodeToString(ret[:])
}

//计算字符串MD5值
func MD5(s string) string {
	return MD5Bytes([]byte(s))
}

//计算文件MD5值
func MD5File(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return MD5Bytes(data), nil
}
