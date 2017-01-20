package types

import (
	"fmt"
	"time"
)

// FBPage is a struct that will hold all the information we're populating
type FBPage struct {
	PageName    string
	ID          string
	LastUpdates UpdateTimes
}

// UpdateTimes is a struct that will hold timestamps for the last times data was updated from
// the facebook graph api.
type UpdateTimes struct {
	GetPageInfo  time.Time
	GetAlbumInfo time.Time
}

// NewFBPage creates a new FBPage struct with the timestamps set to unix time 0
func NewFBPage(name, i string) *FBPage {
	return &FBPage{
		PageName: name,
		ID:       i,
		LastUpdates: UpdateTimes{
			GetPageInfo:  time.Unix(0, 0),
			GetAlbumInfo: time.Unix(0, 0),
		},
	}
}

func (t *UpdateTimes) String() string {
	return fmt.Sprintf("{Last Updates:: Page Info: %s Album Info: %s}", t.GetPageInfo.String(), t.GetAlbumInfo.String())
}

func (f *FBPage) String() string {
	return fmt.Sprintf("[PageName: %s ID: %s %s]", f.PageName, f.ID, f.LastUpdates.String())
}
