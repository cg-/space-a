package facebook

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cg-/space-a/common"
)

func Get(url string) map[string]interface{} {
	resp, err := http.Get(url)
	if err != nil {
		common.Debug(err.Error())
	}
	defer resp.Body.Close()
	myBytes, _ := ioutil.ReadAll(resp.Body)
	var f interface{}
	json.Unmarshal(myBytes, &f)
	return f.(map[string]interface{})
}
