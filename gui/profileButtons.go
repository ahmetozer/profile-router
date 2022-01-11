package gui

import (
	"fmt"
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/ahmetozer/profile-router/config"
	linkrouter "github.com/ahmetozer/profile-router/linkRouter"
)

// ProfileButtons
// Generate Profile buttons
func ProfileButtons(Config *config.Settings, margins *layout.Inset, th *material.Theme, b1 *[]widget.Clickable) []layout.FlexChild {
	plen := len(Config.Profile)
	buttons := make([]layout.FlexChild, plen)

	// Create buttons for each line
	for p := range Config.Profile {
		displayName := Config.Profile[p].DisplayName
		profile := Config.Profile[p].Profile
		if displayName == "" {
			displayName = profile
		}
		var buttonLayout layout.FlexChild
		if displayName == "" {
			buttonLayout = layout.Rigid(
				func(gtx C) D {
					return margins.Layout(gtx,
						func(gtx C) D {
							p += 1
							k := (p % plen)
							btn := material.Button(th, &(*b1)[k], fmt.Sprintf("No Profile %v", k))
							btn.Background = color.NRGBA{A: 0xff, R: 0x90, G: 0xA4, B: 0xAE}
							gtx = gtx.Disabled()
							return btn.Layout(gtx)
						},
					)
				},
			)
		} else {
			buttonLayout = layout.Rigid(
				func(gtx C) D {

					return margins.Layout(gtx,
						func(gtx C) D {
							p += 1
							k := (p % plen)
							btn := material.Button(th, &(*b1)[k], displayName)
							if (*b1)[k].Clicked() {
								linkrouter.ProfileExecuter(Config, profile)
							}
							return btn.Layout(gtx)
						},
					)
				},
			)
		}
		buttons[p] = buttonLayout
	}
	return buttons

}
