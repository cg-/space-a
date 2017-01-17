package types

import (
	"fmt"
)

type FBPage struct {
	PageName string
	ID       string
}

func (f *FBPage) String() string {
	return fmt.Sprintf("[PageName: %s ID: %s]", f.PageName, f.ID)
}
