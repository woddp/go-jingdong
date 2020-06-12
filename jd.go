package  go_jingdong

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"go-jingdong/base"
	"go-jingdong/config"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

type JD struct {
	config *config.Config
	//系统参数
	sysParams map[string]string
}

func NewDj(config *config.Config) *JD {
	jd := new(JD)
	jd.config = config
	return jd
}
func (t *JD) buildParams(system base.System) map[string][]string {
	t.sysParams=make(map[string]string)
	//t.sysParams["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	t.sysParams["timestamp"] = "2020-06-12 18:52:50"
	t.sysParams["v"] = t.config.V
	t.sysParams["sign_method"] = "md5"
	t.sysParams["format"] = "json"
	t.sysParams["method"] = system.Method
	t.sysParams["param_json"] = system.Param_json
	//t.sysParams["access_token"]=system.Access_token
	t.sysParams["app_key"] = t.config.AppKey
	//把所有请求参数按照参数名称的ASCII码表顺序进行排序
	keys := make([]string, 0)
	for s := range t.sysParams {
		keys = append(keys, s)
	}
	sort.Strings(keys)
	//把所有参数名和参数值进行拼接(拼接过程中param_json里面的参数顺序应该与请求时传入的顺序保持一致)
	signStr_ := ""
	for i := range keys {
		signStr_ += keys[i] + (t.sysParams[keys[i]])
	}
	// 把appSecret的值拼接在字符串的两端
	signStr := t.config.Secretkey + signStr_ + t.config.Secretkey
	//使用MD5进行加密（各个语言的处理方式略有不同），并转化成大写
	signStr = strings.ToUpper(md5V(signStr))
	t.sysParams["sign"] = signStr
	values := make(map[string][]string)
	value := make([]string, 0)
	for s := range t.sysParams {
		values[s] = append(value, t.sysParams[s])
	}
	return values
}
func (t *JD) Get(system base.System) (maps map[string]interface{}, err error) {
	maps = make(map[string]interface{})
	resp, err := http.PostForm(t.config.Url,
		t.buildParams(system))
	if err != nil {
		// handle error
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	_ = json.Unmarshal(body, &maps)
	return
}
func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
