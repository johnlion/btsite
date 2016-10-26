package util
import(
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"regexp"
	"strings"
	"path"
	"os"
)

const(
	// Spider中启用Keyin的初始值
	USE_KEYIN = "\r\t\n"
)

var (
	re = regexp.MustCompile(">[ \t\n\v\f\r]+<")
)


// JsonpToJson modify jsonp string to json string
// Example: forbar({a:"1",b:2}) to {"a":"1","b":2}
func JsonpToJson( json string ) string {
	start := strings.Index(json, "{")
	end := strings.LastIndex(json, "}")
	start1 := strings.Index(json, "[")
	if start1 >0 && start > start1{
		start = start1
		end = strings.LastIndex( json, "]" )
	}
	if end >start && end !=-1 && start != -1{
		json = json[ start: end+1 ]
	}
	json = strings.Replace(json, "\\'", "", -1)
	regDetail, _ := regexp.Compile("([^\\s\\:\\{\\,\\d\"]+|[a-z][a-z\\d]*)\\s*\\:")
	return regDetail.ReplaceAllString(json, "\"$1\":")
}

// 创建目录
func Mkdir( Path string ){
	p, _ := path.Split(Path)
	if p == ""{
		return ""
	}
	d, err := os.Stat(p)
	if err != nil || !d.IsDir(){
		if err = os.MkdirAll( p, 0777 ); err != nil{

		}
	}
}