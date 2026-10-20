package main

import (
	"github.com/addidotlol/mycn"
	"github.com/egoist/mygo/ui"
)

const (
	appearanceSystem = iota
	appearanceLight
	appearanceDark
)

type gallery struct {
	preset     string
	appearance int

	email, bio      string
	terms, airplane bool
	notifications   bool
	tab             int
	fruit           string

	popoverOpen, dialogOpen bool
	name                    string
}

func newGallery() *gallery {
	return &gallery{preset: "Neutral", terms: true, notifications: true, name: "Pedro Duarte"}
}

func (g *gallery) theme() *mycn.Preset {
	for _, p := range mycn.Presets {
		if p.Name == g.preset {
			return p.Preset
		}
	}
	return &mycn.Neutral
}

func (g *gallery) applyTheme(c *ui.Context) {
	p := g.theme()
	switch g.appearance {
	case appearanceLight:
		mycn.Use(c, &p.Light, &p.Light)
	case appearanceDark:
		mycn.Use(c, &p.Dark, &p.Dark)
	default:
		mycn.Use(c, &p.Light, &p.Dark)
	}
}

func (g *gallery) view(c *ui.Context) {
	g.applyTheme(c)
	mycn.Toaster(c)
	ui.Scroll(c).Fill().Children(func() {
		ui.Column(c).FillWidth().Padding(32).Gap(28).Children(func() {
			g.header(c)
			section(c, "Button", g.buttons)
			section(c, "Badge", g.badges)
			section(c, "Input and Textarea", g.fields)
			section(c, "Checkbox and Switch", g.toggles)
			section(c, "Tabs", g.tabs)
			section(c, "Select", g.selects)
			section(c, "Overlays", g.overlays)
			section(c, "Card", g.card)
		})
	})
	g.dialog(c)
}

func section(c *ui.Context, title string, body func(c *ui.Context)) {
	ui.Column(c).Gap(12).Children(func() {
		ui.Text(c, title).FontWeight(600).FontSize(c.Theme().Rem(1.15))
		mycn.Separator(c)
		ui.Column(c).Gap(12).AlignItems(ui.Start).Children(func() { body(c) })
	})
}

func (g *gallery) header(c *ui.Context) {
	ui.Row(c).Gap(12).AlignItems(ui.Center).Children(func() {
		ui.Column(c).Grow(1).Gap(4).Children(func() {
			ui.Text(c, "mycn").FontWeight(700).FontSize(c.Theme().Rem(1.6))
			ui.Text(c, "shadcn/ui for MyGo native UI").TextColor(mycn.Active(c).MutedForeground)
		})
		sel := mycn.Select(c.Key("preset"), &g.preset)
		sel.Trigger(g.preset).Label("Base color").Width(120)
		sel.Content(func() {
			for _, p := range mycn.Presets {
				sel.Item(p.Name, p.Name)
			}
		})
		mycn.Tabs(c.Key("appearance"), &g.appearance, "System", "Light", "Dark")
	})
}

var variants = []struct {
	name    string
	variant mycn.Variant
}{
	{"Default", mycn.Default},
	{"Secondary", mycn.Secondary},
	{"Outline", mycn.Outline},
	{"Ghost", mycn.Ghost},
	{"Destructive", mycn.Destructive},
	{"Link", mycn.Link},
}

func (g *gallery) buttons(c *ui.Context) {
	ui.Row(c).Gap(8).Wrap().Children(func() {
		for _, v := range variants {
			if mycn.Button(c.Key(v.name), v.name, v.variant).Clicked() {
				c.Toast(v.name + " clicked")
			}
		}
	})
	ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
		mycn.Button(c.Key("sm"), "Small", mycn.SizeSm)
		mycn.Button(c.Key("md"), "Default")
		mycn.Button(c.Key("lg"), "Large", mycn.SizeLg)
		mycn.Button(c.Key("icon"), "", mycn.SizeIcon, mycn.Outline).Label("Close").Children(func() {
			ui.Icon(c, mycn.IconX).Size(16, 16)
		})
		mycn.Button(c.Key("disabled"), "Disabled", mycn.Disabled(true))
	})
}

func (g *gallery) badges(c *ui.Context) {
	ui.Row(c).Gap(8).Children(func() {
		mycn.Badge(c, "Badge")
		mycn.Badge(c, "Secondary", mycn.Secondary)
		mycn.Badge(c, "Destructive", mycn.Destructive)
		mycn.Badge(c, "Outline", mycn.Outline)
	})
}

func (g *gallery) fields(c *ui.Context) {
	ui.Column(c).Gap(8).Width(360).Children(func() {
		mycn.Input(c.Key("email"), &g.email).Placeholder("Email").Label("Email")
		mycn.Textarea(c.Key("bio"), &g.bio).Placeholder("Tell us a little bit about yourself").Label("Bio")
		mycn.Input(c.Key("disabled-input"), &g.name, mycn.Disabled(true))
	})
}

func (g *gallery) toggles(c *ui.Context) {
	mycn.Checkbox(c.Key("terms"), &g.terms, "Accept terms and conditions")
	mycn.Checkbox(c.Key("notifications"), &g.notifications, "Enable notifications")
	mycn.Switch(c.Key("airplane"), &g.airplane, "Airplane mode")
}

func (g *gallery) tabs(c *ui.Context) {
	tabs := []string{"Account", "Password"}
	mycn.Tabs(c.Key("tabs"), &g.tab, tabs...).List.Width(360)
	mycn.Card(c).Width(360).Children(func() {
		mycn.CardHeader(c).Children(func() {
			mycn.CardTitle(c, tabs[g.tab])
			if g.tab == 0 {
				mycn.CardDescription(c, "Make changes to your account here.")
			} else {
				mycn.CardDescription(c, "Change your password here.")
			}
		})
	})
}

func (g *gallery) selects(c *ui.Context) {
	label := g.fruit
	if label == "" {
		label = "Select a fruit"
	}
	sel := mycn.Select(c.Key("fruit"), &g.fruit)
	sel.Trigger(label).Label("Fruit").Width(180)
	sel.Content(func() {
		for _, f := range []string{"Apple", "Banana", "Blueberry", "Grapes", "Pineapple"} {
			sel.Item(f, f)
		}
	})
}

func (g *gallery) overlays(c *ui.Context) {
	ui.Row(c).Gap(8).Children(func() {
		hover := mycn.Button(c.Key("tooltip"), "Hover", mycn.Outline)
		mycn.Tooltip(c, hover, "Add to library")

		open := mycn.Button(c.Key("popover"), "Open popover", mycn.Outline)
		if open.Clicked() {
			g.popoverOpen = !g.popoverOpen
		}
		mycn.Popover(c.Key("popover-panel"), open, &g.popoverOpen, func(ui.Element) {
			ui.Text(c, "Dimensions").FontWeight(500)
			ui.Text(c, "Set the dimensions for the layer.").TextColor(mycn.Active(c).MutedForeground)
		})

		if mycn.Button(c.Key("dialog"), "Edit profile", mycn.Outline).Clicked() {
			g.dialogOpen = true
		}
		if mycn.Button(c.Key("toast"), "Show toast", mycn.Outline).Clicked() {
			c.AddToast(ui.Toast{Title: "Event has been created", Description: "Sunday, December 03, 2023 at 9:00 AM", Action: "Undo"})
		}
	})
}

func (g *gallery) dialog(c *ui.Context) {
	mycn.Dialog(c.Key("profile"), &g.dialogOpen, func(ui.Element) {
		mycn.DialogHeader(c).Children(func() {
			mycn.DialogTitle(c, "Edit profile")
			mycn.DialogDescription(c, "Make changes to your profile here. Click save when you're done.")
		})
		mycn.Input(c.Key("name"), &g.name).Label("Name")
		mycn.DialogFooter(c).Children(func() {
			if mycn.Button(c.Key("cancel"), "Cancel", mycn.Outline).Clicked() {
				g.dialogOpen = false
			}
			if mycn.Button(c.Key("save"), "Save changes").Clicked() {
				g.dialogOpen = false
				c.Toast("Profile saved")
			}
		})
	})
}

func (g *gallery) card(c *ui.Context) {
	mycn.Card(c).Width(360).Children(func() {
		mycn.CardHeader(c).Children(func() {
			mycn.CardTitle(c, "Create project")
			mycn.CardDescription(c, "Deploy your new project in one click.")
		})
		mycn.CardContent(c).Children(func() {
			mycn.Input(c.Key("project"), &g.email).Placeholder("Name of your project")
		})
		mycn.CardFooter(c).Justify(ui.SpaceBetween).Children(func() {
			mycn.Button(c.Key("card-cancel"), "Cancel", mycn.Outline)
			mycn.Button(c.Key("card-deploy"), "Deploy")
		})
	})
}
