package mycn

import "github.com/egoist/mygo/ui"

func Checkbox(c *ui.Context, checked *bool, label string, opts ...Option) ui.Element {
	t, u := Active(c), c.Theme()
	row := ui.CheckboxBase(c, checked).Gap(u.Space(2)).AlignItems(ui.Center).FocusRing(false)
	disabled := disabledOf(opts) || row.IsDisabled()
	if disabled {
		row.Disabled(true).Opacity(0.5)
	}
	on := *checked
	row.Children(func() {
		box := ui.Box(c).Size(u.Space(4), u.Space(4)).Radius(4).Shrink(0).Center()
		switch {
		case on:
			box.Background(t.Primary).Border(1, t.Primary).TextColor(t.PrimaryForeground)
		case t.Dark:
			box.Background(t.Input.Alpha(0.3).Over(t.Background)).Border(1, t.Input)
		default:
			box.Border(1, t.Input)
		}
		if row.FocusVisible() {
			box.BorderColor(t.Ring).Shadow(0, 0, 0, 3, t.Ring.Alpha(0.5))
		}
		shadowXS(box)
		if on {
			box.Children(func() { icon(c, IconCheck, u.Space(3.5)) })
		}
		if label != "" {
			ui.Text(c, label).FontWeight(500).LineHeight(1)
		}
	})
	return row
}
