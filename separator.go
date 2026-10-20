package mycn

import "github.com/egoist/mygo/ui"

func Separator(c *ui.Context) ui.Element {
	return ui.Divider(c).Background(Active(c).Border)
}
