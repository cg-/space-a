package main

import (
	"flag"

	"fmt"

	"io/ioutil"

	"strings"

	"time"

	"github.com/cg-/space-a/common"
	"github.com/cg-/space-a/facebook"
	"github.com/cg-/space-a/types"
)

var (
	APP_SECRET string
	APP_ID     string
	FBPages    map[string](types.FBPage) = make(map[string](types.FBPage))
)

func init() {
	in, err := ioutil.ReadFile(".conf")
	if err != nil {
		common.Debug("trouble reading api settings conf")
	}

	tokens := strings.Split(string(in), "\n")
	APP_ID = strings.TrimSpace(tokens[0])
	APP_SECRET = strings.TrimSpace(tokens[1])
	flag.Parse()
}

func updateAllPageData() {
	for k, v := range facebook.PAGE_NAMES {
		common.Debug("Updating info on " + k)
		accessKey := APP_ID + "|" + APP_SECRET
		r := facebook.Get(fmt.Sprintf("https://graph.facebook.com/%s?access_token=%s", v, accessKey))
		FBPages[k] = types.FBPage{
			PageName: r["name"].(string),
			ID:       r["id"].(string),
		}
		time.Sleep(time.Millisecond * 100)
		r2 := facebook.Get(fmt.Sprintf("https://graph.facebook.com/v2.8/%s/albums?access_token=%s", FBPages[k].ID, APP_ID+"|"+APP_SECRET))
		i := ((r2["data"]).([]interface{}))
		for k := range i {
			fmt.Println(i[k].(map[string]interface{}))
		}
	}
}

func main() {
	common.Debug("Starting Up")
	common.Debug("App ID set to " + APP_ID)
	common.Debug("App secret set to " + APP_SECRET)
	updateAllPageData()
}
