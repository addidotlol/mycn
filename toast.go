package mycn

import "github.com/egoist/mygo/ui"

const ToastError = "error"

func Toaster(c *ui.Context) ui.Element {
	t, u := Active(c), c.Theme()
	return ui.ToastViewportBase(c, func(viewport ui.Element, toasts []ui.Toast) {
		viewport.Padding(u.Space(6)).AlignItems(ui.End).Gap(u.Space(3.5))
		for _, toast := range toasts {
			toastCard(c, t, toast)
		}
	})
}

func toastCard(c *ui.Context, t *Theme, toast ui.Toast) {
	u := c.Theme()
	parts := ui.ToastBase(c.Key(toast.ID), toast)
	root := stylePanel(parts.Root, t).Row().AlignItems(ui.Center).Gap(u.Space(1.5)).
		Width(u.Space(89)).MaxWidthPercent(100).Padding(u.Space(4)).Radius(t.Radius)
	root.Transition(ui.ElementTransition{Enter: &ui.Motion{Y: 16, Opacity: 0}, Exit: &ui.Motion{Opacity: 0}})
	titleColor := t.PopoverForeground
	if toast.Type == ToastError {
		titleColor = t.Destructive
	}
	root.Children(func() {
		ui.Column(c).Grow(1).Gap(u.Space(0.5)).Children(func() {
			ui.Text(c, toast.Title).FontWeight(500).TextColor(titleColor)
			if toast.Description != "" {
				ui.Text(c, toast.Description).TextColor(t.MutedForeground)
			}
		})
		if toast.Action != "" {
			action := parts.ActionButton().Height(u.Space(6)).PaddingX(u.Space(2)).Radius(4).
				Background(t.Primary).TextColor(t.PrimaryForeground).FontSize(u.Rem(smallTextRem)).FontWeight(500)
			action.Children(func() { ui.Text(c, toast.Action).SingleLine() })
		}
		closeButton := parts.CloseButton().Label("Close").Radius(4).Padding(u.Space(1)).TextColor(t.MutedForeground)
		if closeButton.Hovered() {
			closeButton.TextColor(t.Foreground)
		}
		closeButton.Children(func() { icon(c, IconX, u.Space(3.5)) })
	})
}
