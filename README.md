# mycn

shadcn/ui-style components for [MyGo](https://mygo.egoist.dev) Native UI.

## Install

```sh
go get github.com/addidotlol/mycn
```

Needs Go 1.27.1 or newer.

## Usage

Call `mycn.Use` (or `mycn.UsePreset`) at the top of your view, every frame, and then use the components the way you'd use `ui`:

```go
func (a *app) view(c *ui.Context) {
	mycn.UsePreset(c, &mycn.Zinc)
	mycn.Toaster(c)

	if mycn.Button(c.Key("save"), "Save").Clicked() {
		c.Toast("Saved")
	}
	mycn.Button(c.Key("delete"), "Delete", mycn.Destructive, mycn.SizeSm)
	mycn.Input(c.Key("email"), &a.email).Placeholder("Email")
	mycn.Checkbox(c.Key("terms"), &a.terms, "Accept terms")
}
```

## Components

| Component | API |
|---|---|
| Button | `Button(c, label, opts...)` with `Default`, `Secondary`, `Outline`, `Ghost`, `Destructive`, `Link` and `SizeSm`, `SizeDefault`, `SizeLg`, `SizeIcon`; `Disabled(true)` |
| Badge | `Badge(c, label, opts...)` with `Default`, `Secondary`, `Destructive`, `Outline` |
| Card | `Card`, `CardHeader`, `CardTitle`, `CardDescription`, `CardContent`, `CardFooter` |
| Separator | `Separator(c)`, horizontal in a column and vertical in a row |
| Input, Textarea | `Input(c, &value, opts...)`, `Textarea(c, &value, opts...)` |
| Checkbox, Switch | `Checkbox(c, &on, label, opts...)`, `Switch(c, &on, label, opts...)` |
| Tabs | `Tabs(c, &selected, labels...)` returns `ui.TabsParts`; build the content from `selected` |
| Select | `sel := Select(c, &value)`, then `sel.Trigger(label)`, `sel.Content(func() { sel.Item(v, label) })` |
| Tooltip | `Tooltip(c, anchor, text)` |
| Popover | `Popover(c, anchor, &open, func(panel ui.Element) { ... })` |
| Dialog | `Dialog(c, &open, func(panel ui.Element) { ... })` with `DialogHeader`, `DialogTitle`, `DialogDescription`, `DialogFooter` |
| Toast | `Toaster(c)` once in the view, then `c.Toast`, `c.ToastAction` or `c.AddToast(ui.Toast{...})`; `Type: mycn.ToastError` colors the title |

## Theming

`mycn.Theme` holds shadcn's full token set, named after its CSS variables (`Background`, `Foreground`, `Card`, `Popover`, `Primary`, `Secondary`, `Muted`, `Accent`, `Destructive`, each with a `...Foreground`, plus `Border`, `Input`, `Ring`, `Radius`), so you can paste OKLCH values from a shadcn theme editor with `ui.Oklch(l, c, h)`:

```go
light := mycn.Neutral.Light
light.Primary = ui.Oklch(0.488, 0.243, 264.376)
dark := mycn.Neutral.Dark
dark.Primary = ui.Oklch(0.488, 0.243, 264.376)
mycn.Use(c, &light, &dark)
```

Presets: `Neutral`, `Zinc`, `Stone`, `Gray` and `Slate`, each with `Light` and `Dark`, using shadcn's base colors and a radius of 10. `Spacing` and `FontSize` scale every component, since sizes come from `t.Space` and `t.Rem`.

## Development

```sh
go build ./...
go vet ./... && staticcheck ./... && gofmt -l .
go run ./examples/gallery
```
