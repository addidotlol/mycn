package mycn

import "github.com/egoist/mygo/ui"

func styleField(e ui.Element, t *Theme, opts []Option) ui.Element {
	bg := ui.Transparent
	if t.Dark {
		bg = t.Input.Alpha(0.3).Over(t.Background)
	}
	border := t.Input
	if e.Focused() {
		border = t.Ring
	}
	e.Radius(controlRadius(t)).Border(1, border).Background(bg).Cursor(ui.CursorText)
	if disabledOf(opts) || e.IsDisabled() {
		e.Disabled(true).Opacity(0.5)
	}
	return shadowXS(e)
}

func Input(c *ui.Context, value *string, opts ...Option) ui.Element {
	t, u := Active(c), c.Theme()
	in := ui.TextInputBase(c, value).MinHeight(u.Space(9)).Padding(u.Space(1.5), u.Space(3)).FillWidth().MinWidth(0)
	return styleField(in, t, opts)
}
