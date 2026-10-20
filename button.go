package mycn

import "github.com/egoist/mygo/ui"

type disabledOption bool

func (disabledOption) mycnOption() {}

func Disabled(d bool) Option { return disabledOption(d) }

func disabledOf(opts []Option) bool {
	for _, o := range opts {
		if d, ok := o.(disabledOption); ok && bool(d) {
			return true
		}
	}
	return false
}

func controlRadius(t *Theme) float32 {
	return max(t.Radius-2, 0)
}

func shadowXS(e ui.Element) ui.Element {
	return e.Shadow(0, 1, 2, 0, ui.RGBA(0, 0, 0, 0.05))
}

type buttonColors struct {
	background, hover, foreground, border ui.Color
	shadow                                bool
}

func buttonPalette(t *Theme, v Variant) buttonColors {
	switch v {
	case Secondary:
		return buttonColors{
			background: t.Secondary,
			hover:      t.Secondary.Alpha(0.8).Over(t.Background),
			foreground: t.SecondaryForeground,
		}
	case Outline:
		bg, hover, border := t.Background, t.Accent, t.Border
		if t.Dark {
			bg, hover, border = t.Input.Alpha(0.3).Over(t.Background), t.Input.Alpha(0.5).Over(t.Background), t.Input
		}
		return buttonColors{background: bg, hover: hover, foreground: t.Foreground, border: border, shadow: true}
	case Ghost:
		hover := t.Accent
		if t.Dark {
			hover = t.Accent.Alpha(0.5).Over(t.Background)
		}
		return buttonColors{hover: hover, foreground: t.Foreground}
	case Destructive:
		bg := t.Destructive
		if t.Dark {
			bg = t.Destructive.Alpha(0.6).Over(t.Background)
		}
		return buttonColors{
			background: bg,
			hover:      bg.Alpha(0.9).Over(t.Background),
			foreground: t.DestructiveForeground,
			shadow:     true,
		}
	case Link:
		return buttonColors{foreground: t.Primary}
	}
	return buttonColors{
		background: t.Primary,
		hover:      t.Primary.Alpha(0.9).Over(t.Background),
		foreground: t.PrimaryForeground,
		shadow:     true,
	}
}

func sizeButton(b ui.Element, u *ui.Theme, s Size, hasLabel bool) {
	switch s {
	case SizeSm:
		b.Height(u.Space(8)).Gap(u.Space(1.5)).PaddingX(u.Space(3))
	case SizeLg:
		b.Height(u.Space(10)).Gap(u.Space(2)).PaddingX(u.Space(6))
	case SizeIcon:
		b.Size(u.Space(9), u.Space(9))
	default:
		b.Height(u.Space(9)).Gap(u.Space(2)).PaddingX(u.Space(4))
	}
	if !hasLabel && s != SizeIcon {
		b.PaddingX(u.Space(3))
	}
}

func Button(c *ui.Context, label string, opts ...Option) ui.Element {
	t, u := Active(c), c.Theme()
	v := variantOf(opts, Default, Secondary, Outline, Ghost, Destructive, Link)
	colors := buttonPalette(t, v)

	b := ui.ButtonBase(c).Radius(controlRadius(t)).FontWeight(500).FontSize(u.Rem(1)).TextColor(colors.foreground)
	sizeButton(b, u, sizeOf(opts), label != "")

	disabled := disabledOf(opts) || b.IsDisabled()
	if disabled {
		b.Disabled(true).Opacity(0.5)
	}

	active := !disabled && b.Hovered()
	switch {
	case active && colors.hover.A > 0:
		b.Background(colors.hover)
	case colors.background.A > 0:
		b.Background(colors.background)
	}
	if active && (v == Outline || v == Ghost) {
		b.TextColor(t.AccentForeground)
	}
	if colors.border.A > 0 {
		b.Border(1, colors.border)
	}
	if colors.shadow {
		shadowXS(b)
	}
	if v == Link {
		b.Cursor(ui.CursorPointer)
	}

	if label != "" {
		b.Children(func() {
			text := ui.Text(c, label).SingleLine()
			if v == Link && active {
				text.Underline()
			}
		})
	}
	return b
}
