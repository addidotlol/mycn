package mycn

import (
	"time"

	"github.com/egoist/mygo/ui"
)

const switchDuration = 150 * time.Millisecond

func Switch(c *ui.Context, on *bool, label string, opts ...Option) ui.Element {
	t, u := Active(c), c.Theme()
	row := ui.SwitchBase(c, on).Gap(u.Space(2)).AlignItems(ui.Center).FocusRing(false)
	disabled := disabledOf(opts) || row.IsDisabled()
	if disabled {
		row.Disabled(true).Opacity(0.5)
	}

	width, height, thumb := u.Space(8), u.Space(4.6), u.Space(4)
	travel := width - thumb - 2
	target := float32(0)
	if *on {
		target = 1
	}
	progress := row.Animate("thumb", target, switchDuration)

	trackOff, thumbOff, thumbOn := t.Input, t.Background, t.Background
	if t.Dark {
		trackOff, thumbOff, thumbOn = t.Input.Alpha(0.8).Over(t.Background), t.Foreground, t.PrimaryForeground
	}

	row.Children(func() {
		track := ui.Row(c).Size(width, height).Radius(999).Shrink(0).AlignItems(ui.Center).
			Background(trackOff.Mix(t.Primary, progress)).Padding(0, 1)
		if row.FocusVisible() {
			track.Shadow(0, 0, 0, 3, t.Ring.Alpha(0.5))
		}
		shadowXS(track).Children(func() {
			ui.Box(c).Size(thumb, thumb).Radius(999).Background(thumbOff.Mix(thumbOn, progress)).Translate(travel*progress, 0)
		})
		if label != "" {
			ui.Text(c, label).FontWeight(500).LineHeight(1)
		}
	})
	return row
}
