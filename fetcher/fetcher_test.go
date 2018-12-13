package fetcher

import (
	"testing"
	"net/http"
	"fmt"
	"io/ioutil"
)
//http://album.zhenai.com/u/1445144756
func TestFetch(t *testing.T) {
	resp, err := http.Get("https://zhidao.baidu.com/question/115916435.html")
	fmt.Println(err)
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	htmlStr := string(bytes)
	fmt.Println(htmlStr)
}
