package mycn

import "github.com/egoist/mygo/ui"

func shadowMD(e ui.Element) ui.Element {
	return e.Shadow(0, 4, 6, -1, ui.RGBA(0, 0, 0, 0.1)).Shadow(0, 2, 4, -2, ui.RGBA(0, 0, 0, 0.1))
}

func stylePanel(e ui.Element, t *Theme) ui.Element {
	e.Radius(controlRadius(t)).Border(1, t.Border).Background(t.Popover).TextColor(t.PopoverForeground)
	return shadowMD(e)
}

func Popover(c *ui.Context, anchor ui.Element, open *bool, content func(panel ui.Element)) ui.Element {
	t, u := Active(c), c.Theme()
	return ui.PopoverBase(c, anchor, open, func(panel ui.Element) {
		stylePanel(panel, t).Margin(u.Space(1), 0).Width(u.Space(72)).Padding(u.Space(4)).Gap(u.Space(2))
		content(panel)
	})
}
