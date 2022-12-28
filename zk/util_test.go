package zk

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	zkCli "github.com/pengkai242/go-zookeeper/zk"
	"io/ioutil"
	"testing"
	"time"
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

	c, _, err := Connect([]string{"14.29.105.178"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	authData := fmt.Sprintf("%s:%s", "admin", "admin")
	err = c.AddAuth("sasl", []byte(authData))
	children, _, _, err := c.ChildrenW("/zookeeper")
	if err != nil {
		panic(err)
	}
	fmt.Println(children)

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
