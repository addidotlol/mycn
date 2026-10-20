package mycn

import "github.com/egoist/mygo/ui"

type SelectParts[T comparable] struct {
	base     ui.SelectParts[T]
	c        *ui.Context
	selected *T
}

func Select[T comparable](c *ui.Context, selected *T) SelectParts[T] {
	return SelectParts[T]{base: ui.SelectBase(c, selected), c: c, selected: selected}
}

func (s SelectParts[T]) Trigger(label string) ui.Element {
	t, u := Active(s.c), s.c.Theme()
	var zero T
	trigger := s.base.Trigger.Height(u.Space(9)).Gap(u.Space(2)).PaddingX(u.Space(3)).Justify(ui.SpaceBetween).AlignItems(ui.Center)
	styleField(trigger, t, nil).Cursor(ui.CursorPointer)
	textColor := t.Foreground
	if *s.selected == zero {
		textColor = t.MutedForeground
	}
	trigger.Children(func() {
		ui.Text(s.c, label).SingleLine().TextColor(textColor)
		icon(s.c, IconChevronDown, u.Space(4)).Opacity(0.5)
	})
	return trigger
}

func (s SelectParts[T]) Content(items func()) ui.Element {
	t, u := Active(s.c), s.c.Theme()
	return s.base.Popup(func(panel ui.Element) {
		stylePanel(panel, t).Margin(u.Space(1), 0).MinWidth(u.Space(32)).Padding(u.Space(1))
		items()
	})
}

func (s SelectParts[T]) Item(value T, label string) ui.Element {
	t, u := Active(s.c), s.c.Theme()
	item := s.base.Item(value).Gap(u.Space(2)).Padding(u.Space(1.5), u.Space(2)).Radius(max(t.Radius-4, 0)).AlignItems(ui.Center)
	if item.Highlighted() {
		item.Background(t.Accent).TextColor(t.AccentForeground)
	}
	item.Children(func() {
		ui.Text(s.c, label).SingleLine().Grow(1)
		check := icon(s.c, IconCheck, u.Space(4))
		if *s.selected != value {
			check.Invisible()
		}
	})
	return item
}

func (s SelectParts[T]) Open() bool { return s.base.Open() }
