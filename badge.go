package mycn

import "github.com/egoist/mygo/ui"

const smallTextRem = 12.0 / 14.0

func Badge(c *ui.Context, label string, opts ...Option) ui.Element {
	t, u := Active(c), c.Theme()
	bg, fg, border := t.Primary, t.PrimaryForeground, ui.Transparent
	switch variantOf(opts, Secondary, Destructive, Outline) {
	case Secondary:
		bg, fg = t.Secondary, t.SecondaryForeground
	case Destructive:
		bg, fg = t.Destructive, t.DestructiveForeground
		if t.Dark {
			bg = t.Destructive.Alpha(0.6).Over(t.Background)
		}
	case Outline:
		bg, fg, border = ui.Transparent, t.Foreground, t.Border
	}
	b := ui.Row(c).AlignItems(ui.Center).Gap(u.Space(1)).Padding(u.Space(0.5), u.Space(2)).
		Radius(controlRadius(t)).Border(1, border).Background(bg).TextColor(fg).
		FontSize(u.Rem(smallTextRem)).FontWeight(500).Shrink(0)
	if label != "" {
		b.Children(func() { ui.Text(c, label).SingleLine() })
	}
	return b
}
