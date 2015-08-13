package dmm

import (
	"fmt"
)

type Actress struct {
	ID   int
	Name string
	Image string
}

func (a *Actress) VideoURL() string {
	return fmt.Sprintf("http://www.dmm.co.jp/digital/videoa/-/list/=/article=actress/id=%d/", a.ID)
}

func (a *Actress) ImageURL() string {
	return fmt.Sprintf("http://pics.dmm.co.jp/mono/actjpgs/%s", a.Image)
}
