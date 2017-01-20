package main

import (
	"flag"

	"io/ioutil"

	"strings"

	"github.com/cg-/space-a/common"
	"github.com/cg-/space-a/facebook"
	"github.com/cg-/space-a/types"
)

var (
	fbCon   *facebook.FBConnector
	fBPages map[string](*types.FBPage) = make(map[string](*types.FBPage))
)

func init() {
	in, err := ioutil.ReadFile(".conf")
	if err != nil {
		common.Debug("trouble reading api settings conf")
	}
	tokens := strings.Split(string(in), "\n")
	app_id := strings.TrimSpace(tokens[0])
	app_secret := strings.TrimSpace(tokens[1])
	accessKey := app_id + "|" + app_secret
	fbCon = facebook.NewFBConnector("https://graph.gacebook.com", accessKey)
	flag.Parse()
}

func updateAllPageData() {
	for k, v := range facebook.PAGE_NAMES {
		common.Debug("Updating info on " + k)
		page, err := fbCon.GetPageInfo(v)
		if err != nil {
			common.Debug("Error encountered in updating page data: " + err.Error())
		}
		fBPages[k] = page
		err = fbCon.GetAlbumInfo(fBPages[k])
		if err != nil {
			common.Debug("Error encountered in updating album data: " + err.Error())
		}
	}
}

func main() {
	common.Debug("Starting Up")
	updateAllPageData()
}
