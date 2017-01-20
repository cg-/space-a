package facebook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cg-/space-a/common"
	"github.com/cg-/space-a/types"
)

// FBConnector is an interface to facebook's graph api.
type FBConnector struct {
	prefixURL string
	accessKey string
}

// NewFBConnector will return a new FBConnector object set up to use the prefix url and key provided.
func NewFBConnector(prefix, key string) *FBConnector {
	return &FBConnector{
		prefixURL: prefix,
		accessKey: key,
	}
}

// GetPageInfo returns a new FBPage object populated with the ID and Page Name grabbed from the facebook
// graph api.
func (f *FBConnector) GetPageInfo(pageName string) (*types.FBPage, error) {
	r, err := get(fmt.Sprintf("%s/%s?access_token=%s", f.prefixURL, pageName, f.accessKey))
	if err != nil {
		common.Debug(err.Error())
		return nil, fmt.Errorf("Error in GetPageInfo: %s", err.Error())
	}
	return &types.FBPage{
		PageName: r["name"].(string),
		ID:       r["id"].(string),
	}, nil
}

// GetAlbumInfo takes a FBPage object (probably created with GetPageInfo) and updates it's album information
// from the facebook graph api.
func (f *FBConnector) GetAlbumInfo(page *types.FBPage) error {
	r, err := get(fmt.Sprintf("%s/%s?access_token=%s", f.prefixURL, page.ID, f.accessKey))
	if err != nil {
		common.Debug(err.Error())
		return fmt.Errorf("Error in GetAlbumInfo: %s", err.Error())
	}
	fmt.Println(r)
	return nil
}

// get does the dirty work of making the http request and unmarshalling the json into something we can use.
func get(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		common.Debug(err.Error())
		return nil, fmt.Errorf("Error in get (http request): %s", err.Error())
	}
	defer resp.Body.Close()
	myBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		common.Debug(err.Error())
		return nil, fmt.Errorf("Error in get (readall portion): %s", err.Error())
	}
	var f interface{}
	json.Unmarshal(myBytes, &f)
	return f.(map[string]interface{}), nil
}
