package mycn

import "github.com/egoist/mygo/ui"

func Tabs(c *ui.Context, selected *int, labels ...string) ui.TabsParts {
	t, u := Active(c), c.Theme()
	tabs := ui.TabsBase(c, selected, len(labels))
	tabs.List.Height(u.Space(9)).Shrink(0).Padding(3).Radius(t.Radius).Background(t.Muted).AlignItems(ui.Center).Children(func() {
		for i, label := range labels {
			tab := tabs.Tab(i).Grow(1).FillHeight().Gap(u.Space(1.5)).Padding(u.Space(1), u.Space(2)).
				Radius(controlRadius(t)).Border(1, ui.Transparent).FontWeight(500).Justify(ui.Center).AlignItems(ui.Center)
			tabColor := t.Foreground
			if t.Dark {
				tabColor = t.MutedForeground
			}
			if i == *selected {
				tabColor = t.Foreground
				if t.Dark {
					tab.Background(t.Input.Alpha(0.3).Over(t.Muted)).BorderColor(t.Input)
				} else {
					tab.Background(t.Background)
				}
				tab.Shadow(0, 1, 3, 0, ui.RGBA(0, 0, 0, 0.1)).Shadow(0, 1, 2, -1, ui.RGBA(0, 0, 0, 0.1))
			}
			tab.TextColor(tabColor).Children(func() { ui.Text(c, label).NoWrap() })
		}
	})
	return tabs
}
