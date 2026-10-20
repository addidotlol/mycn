package mycn

import "github.com/egoist/mygo/ui"

func Dialog(c *ui.Context, open *bool, content func(panel ui.Element)) ui.Element {
	t, u := Active(c), c.Theme()
	return ui.DialogBase(c, open, func(backdrop, panel ui.Element) {
		backdrop.Background(ui.RGBA(0, 0, 0, 0.5)).Padding(u.Space(4))
		panel.Width(u.Space(128)).MaxWidthPercent(100).Padding(u.Space(6)).Gap(u.Space(4)).
			Radius(t.Radius).Border(1, t.Border).Background(t.Background).TextColor(t.Foreground).
			Shadow(0, 10, 15, -3, ui.RGBA(0, 0, 0, 0.1)).Shadow(0, 4, 6, -4, ui.RGBA(0, 0, 0, 0.1))
		dialogClose(c, t, open)
		content(panel)
	})
}

func dialogClose(c *ui.Context, t *Theme, open *bool) {
	u := c.Theme()
	closeButton := ui.ButtonBase(c.Key("mycn-dialog-close")).Label("Close").Absolute().Top(u.Space(4)).Right(u.Space(4)).
		Radius(4).Opacity(0.7).TextColor(t.Foreground)
	if closeButton.Hovered() {
		closeButton.Opacity(1)
	}
	closeButton.Children(func() { icon(c, IconX, u.Space(4)) })
	if closeButton.Clicked() {
		*open = false
	}
}

func DialogHeader(c *ui.Context) ui.Element {
	return ui.Column(c).Gap(c.Theme().Space(2))
}

func DialogTitle(c *ui.Context, title string) ui.Element {
	return ui.Text(c, title).FontSize(c.Theme().Rem(18.0 / 14.0)).FontWeight(600).LineHeight(1).Role(ui.RoleHeading)
}

func DialogDescription(c *ui.Context, description string) ui.Element {
	return ui.Text(c, description).TextColor(Active(c).MutedForeground)
}

func DialogFooter(c *ui.Context) ui.Element {
	return ui.Row(c).Justify(ui.End).Gap(c.Theme().Space(2))
}
