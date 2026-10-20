package mycn

import "github.com/egoist/mygo/ui"

func Textarea(c *ui.Context, value *string, opts ...Option) ui.Element {
	t, u := Active(c), c.Theme()
	area := ui.TextAreaBase(c, value).MinHeight(u.Space(16)).Padding(u.Space(2), u.Space(3)).FillWidth()
	return styleField(area, t, opts)
}
