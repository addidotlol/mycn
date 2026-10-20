package mycn

import "github.com/egoist/mygo/ui"

func lucide(paths string) *ui.SVG {
	return ui.MustParseSVG([]byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">` + paths + `</svg>`))
}

var (
	IconCheck       = lucide(`<path d="M20 6 9 17l-5-5"/>`)
	IconX           = lucide(`<path d="M18 6 6 18"/><path d="m6 6 12 12"/>`)
	IconChevronDown = lucide(`<path d="m6 9 6 6 6-6"/>`)
)

func icon(c *ui.Context, s *ui.SVG, size float32) ui.Element {
	return ui.Icon(c, s).Size(size, size).Shrink(0)
}
