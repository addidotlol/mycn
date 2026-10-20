package mycn

import "github.com/egoist/mygo/ui"

type baseColor struct {
	foreground, primary, primaryForeground, secondary, mutedForeground, border, ring ui.Color
}

type darkBaseColor struct {
	background, foreground, card, primary, primaryForeground, secondary, mutedForeground, ring ui.Color
}

var (
	white            = ui.Oklch(1, 0, 0)
	destructiveLight = ui.Oklch(0.577, 0.245, 27.325)
	destructiveDark  = ui.Oklch(0.704, 0.191, 22.216)
)

func lightTheme(b baseColor) Theme {
	return Theme{
		Background:            white,
		Foreground:            b.foreground,
		Card:                  white,
		CardForeground:        b.foreground,
		Popover:               white,
		PopoverForeground:     b.foreground,
		Primary:               b.primary,
		PrimaryForeground:     b.primaryForeground,
		Secondary:             b.secondary,
		SecondaryForeground:   b.primary,
		Muted:                 b.secondary,
		MutedForeground:       b.mutedForeground,
		Accent:                b.secondary,
		AccentForeground:      b.primary,
		Destructive:           destructiveLight,
		DestructiveForeground: white,
		Border:                b.border,
		Input:                 b.border,
		Ring:                  b.ring,
		Radius:                DefaultRadius,
	}
}

func darkTheme(b darkBaseColor) Theme {
	return Theme{
		Dark:                  true,
		Background:            b.background,
		Foreground:            b.foreground,
		Card:                  b.card,
		CardForeground:        b.foreground,
		Popover:               b.card,
		PopoverForeground:     b.foreground,
		Primary:               b.primary,
		PrimaryForeground:     b.primaryForeground,
		Secondary:             b.secondary,
		SecondaryForeground:   b.foreground,
		Muted:                 b.secondary,
		MutedForeground:       b.mutedForeground,
		Accent:                b.secondary,
		AccentForeground:      b.foreground,
		Destructive:           destructiveDark,
		DestructiveForeground: white,
		Border:                white.Alpha(0.1),
		Input:                 white.Alpha(0.15),
		Ring:                  b.ring,
		Radius:                DefaultRadius,
	}
}

var Neutral = Preset{
	Light: lightTheme(baseColor{
		foreground:        ui.Oklch(0.145, 0, 0),
		primary:           ui.Oklch(0.205, 0, 0),
		primaryForeground: ui.Oklch(0.985, 0, 0),
		secondary:         ui.Oklch(0.97, 0, 0),
		mutedForeground:   ui.Oklch(0.556, 0, 0),
		border:            ui.Oklch(0.922, 0, 0),
		ring:              ui.Oklch(0.708, 0, 0),
	}),
	Dark: darkTheme(darkBaseColor{
		background:        ui.Oklch(0.145, 0, 0),
		foreground:        ui.Oklch(0.985, 0, 0),
		card:              ui.Oklch(0.205, 0, 0),
		primary:           ui.Oklch(0.922, 0, 0),
		primaryForeground: ui.Oklch(0.205, 0, 0),
		secondary:         ui.Oklch(0.269, 0, 0),
		mutedForeground:   ui.Oklch(0.708, 0, 0),
		ring:              ui.Oklch(0.556, 0, 0),
	}),
}

var Zinc = Preset{
	Light: lightTheme(baseColor{
		foreground:        ui.Oklch(0.141, 0.005, 285.823),
		primary:           ui.Oklch(0.21, 0.006, 285.885),
		primaryForeground: ui.Oklch(0.985, 0, 0),
		secondary:         ui.Oklch(0.967, 0.001, 286.375),
		mutedForeground:   ui.Oklch(0.552, 0.016, 285.938),
		border:            ui.Oklch(0.92, 0.004, 286.32),
		ring:              ui.Oklch(0.705, 0.015, 286.067),
	}),
	Dark: darkTheme(darkBaseColor{
		background:        ui.Oklch(0.141, 0.005, 285.823),
		foreground:        ui.Oklch(0.985, 0, 0),
		card:              ui.Oklch(0.21, 0.006, 285.885),
		primary:           ui.Oklch(0.92, 0.004, 286.32),
		primaryForeground: ui.Oklch(0.21, 0.006, 285.885),
		secondary:         ui.Oklch(0.274, 0.006, 286.033),
		mutedForeground:   ui.Oklch(0.705, 0.015, 286.067),
		ring:              ui.Oklch(0.552, 0.016, 285.938),
	}),
}

var Stone = Preset{
	Light: lightTheme(baseColor{
		foreground:        ui.Oklch(0.147, 0.004, 49.25),
		primary:           ui.Oklch(0.216, 0.006, 56.043),
		primaryForeground: ui.Oklch(0.985, 0.001, 106.423),
		secondary:         ui.Oklch(0.97, 0.001, 106.424),
		mutedForeground:   ui.Oklch(0.553, 0.013, 58.071),
		border:            ui.Oklch(0.923, 0.003, 48.717),
		ring:              ui.Oklch(0.709, 0.01, 56.259),
	}),
	Dark: darkTheme(darkBaseColor{
		background:        ui.Oklch(0.147, 0.004, 49.25),
		foreground:        ui.Oklch(0.985, 0.001, 106.423),
		card:              ui.Oklch(0.216, 0.006, 56.043),
		primary:           ui.Oklch(0.923, 0.003, 48.717),
		primaryForeground: ui.Oklch(0.216, 0.006, 56.043),
		secondary:         ui.Oklch(0.268, 0.007, 34.298),
		mutedForeground:   ui.Oklch(0.709, 0.01, 56.259),
		ring:              ui.Oklch(0.553, 0.013, 58.071),
	}),
}

var Gray = Preset{
	Light: lightTheme(baseColor{
		foreground:        ui.Oklch(0.13, 0.028, 261.692),
		primary:           ui.Oklch(0.21, 0.034, 264.665),
		primaryForeground: ui.Oklch(0.985, 0.002, 247.839),
		secondary:         ui.Oklch(0.967, 0.003, 264.542),
		mutedForeground:   ui.Oklch(0.551, 0.027, 264.364),
		border:            ui.Oklch(0.928, 0.006, 264.531),
		ring:              ui.Oklch(0.707, 0.022, 261.325),
	}),
	Dark: darkTheme(darkBaseColor{
		background:        ui.Oklch(0.13, 0.028, 261.692),
		foreground:        ui.Oklch(0.985, 0.002, 247.839),
		card:              ui.Oklch(0.21, 0.034, 264.665),
		primary:           ui.Oklch(0.928, 0.006, 264.531),
		primaryForeground: ui.Oklch(0.21, 0.034, 264.665),
		secondary:         ui.Oklch(0.278, 0.033, 256.848),
		mutedForeground:   ui.Oklch(0.707, 0.022, 261.325),
		ring:              ui.Oklch(0.551, 0.027, 264.364),
	}),
}

var Slate = Preset{
	Light: lightTheme(baseColor{
		foreground:        ui.Oklch(0.129, 0.042, 264.695),
		primary:           ui.Oklch(0.208, 0.042, 265.755),
		primaryForeground: ui.Oklch(0.984, 0.003, 247.858),
		secondary:         ui.Oklch(0.968, 0.007, 247.896),
		mutedForeground:   ui.Oklch(0.554, 0.046, 257.417),
		border:            ui.Oklch(0.929, 0.013, 255.508),
		ring:              ui.Oklch(0.704, 0.04, 256.788),
	}),
	Dark: darkTheme(darkBaseColor{
		background:        ui.Oklch(0.129, 0.042, 264.695),
		foreground:        ui.Oklch(0.984, 0.003, 247.858),
		card:              ui.Oklch(0.208, 0.042, 265.755),
		primary:           ui.Oklch(0.929, 0.013, 255.508),
		primaryForeground: ui.Oklch(0.208, 0.042, 265.755),
		secondary:         ui.Oklch(0.279, 0.041, 260.031),
		mutedForeground:   ui.Oklch(0.704, 0.04, 256.788),
		ring:              ui.Oklch(0.551, 0.027, 264.364),
	}),
}

var Presets = []struct {
	Name   string
	Preset *Preset
}{
	{"Neutral", &Neutral},
	{"Zinc", &Zinc},
	{"Stone", &Stone},
	{"Gray", &Gray},
	{"Slate", &Slate},
}
