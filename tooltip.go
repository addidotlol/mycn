package mycn

import "github.com/egoist/mygo/ui"

func Tooltip(c *ui.Context, anchor ui.Element, text string) ui.Element {
	t, u := Active(c), c.Theme()
	anchor.Description(text)
	return ui.TooltipBase(c, anchor, func(tip ui.Element) {
		tip.Margin(u.Space(1), 0).Padding(u.Space(1.5), u.Space(3)).Radius(controlRadius(t)).
			Background(t.Primary).TextColor(t.PrimaryForeground).FontSize(u.Rem(smallTextRem))
		tip.Children(func() { ui.Text(c, text).SingleLine() })
	})
}
