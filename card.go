package mycn

import "github.com/egoist/mygo/ui"

func Card(c *ui.Context) ui.Element {
	t, u := Active(c), c.Theme()
	card := ui.Column(c).Gap(u.Space(6)).PaddingY(u.Space(6)).Radius(t.Radius+4).
		Border(1, t.Border).Background(t.Card).TextColor(t.CardForeground)
	return card.Shadow(0, 1, 3, 0, ui.RGBA(0, 0, 0, 0.1)).Shadow(0, 1, 2, -1, ui.RGBA(0, 0, 0, 0.1))
}

func CardHeader(c *ui.Context) ui.Element {
	u := c.Theme()
	return ui.Column(c).Gap(u.Space(1.5)).PaddingX(u.Space(6))
}

func CardTitle(c *ui.Context, title string) ui.Element {
	return ui.Text(c, title).FontWeight(600).LineHeight(1)
}

func CardDescription(c *ui.Context, description string) ui.Element {
	return ui.Text(c, description).TextColor(Active(c).MutedForeground)
}

func CardContent(c *ui.Context) ui.Element {
	return ui.Column(c).PaddingX(c.Theme().Space(6))
}

func CardFooter(c *ui.Context) ui.Element {
	u := c.Theme()
	return ui.Row(c).AlignItems(ui.Center).Gap(u.Space(2)).PaddingX(u.Space(6))
}
