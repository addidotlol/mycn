package mycn

import (
	"sync"

	"github.com/egoist/mygo/ui"
)

const DefaultRadius = 10

type Theme struct {
	Dark bool

	Background            ui.Color
	Foreground            ui.Color
	Card                  ui.Color
	CardForeground        ui.Color
	Popover               ui.Color
	PopoverForeground     ui.Color
	Primary               ui.Color
	PrimaryForeground     ui.Color
	Secondary             ui.Color
	SecondaryForeground   ui.Color
	Muted                 ui.Color
	MutedForeground       ui.Color
	Accent                ui.Color
	AccentForeground      ui.Color
	Destructive           ui.Color
	DestructiveForeground ui.Color
	Border                ui.Color
	Input                 ui.Color
	Ring                  ui.Color

	Radius   float32
	Font     string
	FontSize float32
	Spacing  float32
}

type Preset struct {
	Light Theme
	Dark  Theme
}

func (t *Theme) ToUI() ui.Theme {
	base := ui.LightTheme()
	if t.Dark {
		base = ui.DarkTheme()
	}
	u := *base
	u.Dark = t.Dark
	u.Background = t.Background
	u.Text = t.Foreground
	u.TextMuted = t.MutedForeground
	u.Surface = t.Secondary
	u.SurfaceHover = t.Accent
	u.SurfacePressed = t.Accent.Mix(t.Foreground, 0.08)
	u.Border = t.Border
	u.Accent = t.Primary
	u.AccentText = t.PrimaryForeground
	u.AccentHover = t.Primary.Alpha(0.9).Over(t.Background)
	u.AccentPressed = t.Primary.Alpha(0.8).Over(t.Background)
	u.Danger = t.Destructive
	u.Focus = t.Ring.Alpha(0.5)
	u.Selection = t.Primary.Alpha(0.25)
	u.Inverse = ui.Transparent
	u.InverseText = ui.Transparent
	u.Radius = t.Radius
	u.Spacing = t.Spacing
	u.Font = t.Font
	if t.FontSize > 0 {
		u.FontSize = t.FontSize
	}
	return u
}

type binding struct {
	source Theme
	ui     *ui.Theme
}

var registry = struct {
	sync.Mutex
	bySource map[*Theme]*binding
	byUI     map[*ui.Theme]*Theme
}{
	bySource: map[*Theme]*binding{},
	byUI:     map[*ui.Theme]*Theme{},
}

func Use(c *ui.Context, light, dark *Theme) {
	chosen := light
	if c.Theme().Dark {
		chosen = dark
	}
	resolved := *chosen
	if resolved.FontSize <= 0 {
		resolved.FontSize = c.Theme().FontSize
	}
	c.SetTheme(bind(chosen, resolved))
}

func UsePreset(c *ui.Context, p *Preset) {
	Use(c, &p.Light, &p.Dark)
}

func bind(key *Theme, resolved Theme) *ui.Theme {
	registry.Lock()
	defer registry.Unlock()
	b, ok := registry.bySource[key]
	if !ok {
		b = &binding{ui: new(ui.Theme)}
		registry.bySource[key] = b
	}
	if !ok || b.source != resolved {
		b.source = resolved
		*b.ui = resolved.ToUI()
		registry.byUI[b.ui] = &b.source
	}
	return b.ui
}

func Active(c *ui.Context) *Theme {
	u := c.Theme()
	registry.Lock()
	t, ok := registry.byUI[u]
	registry.Unlock()
	if ok {
		return t
	}
	return fromUI(u)
}

func fromUI(u *ui.Theme) *Theme {
	return &Theme{
		Dark:                  u.Dark,
		Background:            u.Background,
		Foreground:            u.Text,
		Card:                  u.Background,
		CardForeground:        u.Text,
		Popover:               u.Background,
		PopoverForeground:     u.Text,
		Primary:               u.Accent,
		PrimaryForeground:     u.AccentText,
		Secondary:             u.Surface,
		SecondaryForeground:   u.Text,
		Muted:                 u.Surface,
		MutedForeground:       u.TextMuted,
		Accent:                u.SurfaceHover,
		AccentForeground:      u.Text,
		Destructive:           u.Danger,
		DestructiveForeground: ui.Hex("#ffffff"),
		Border:                u.Border,
		Input:                 u.Border,
		Ring:                  u.Focus,
		Radius:                u.Radius,
		Font:                  u.Font,
		FontSize:              u.FontSize,
		Spacing:               u.Spacing,
	}
}
