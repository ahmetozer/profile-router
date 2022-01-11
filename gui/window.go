package gui

import (
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/ahmetozer/profile-router/config"
)

var (
	list = &widget.List{
		List: layout.List{
			Axis: layout.Vertical,
		},
	}
)

type (
	D = layout.Dimensions
	C = layout.Context
)

// Window
// Create app window
func Window(w *app.Window, Config *config.Settings, ConfigError *error) error {
	margins := layout.Inset{
		Top:    unit.Dp(15),
		Bottom: unit.Dp(15),
		Right:  unit.Dp(15),
		Left:   unit.Dp(15),
	}
	var ops op.Ops

	startButtons := make([]widget.Clickable, len(Config.Profile))
	// th meterial style theme
	th := material.NewTheme(gofont.Collection())
	editor := new(widget.Editor)
	widgets := []layout.Widget{}
	var hasError error
	if ConfigError == nil {
		_, hasError = os.Stat(Config.ChromePath) //errors.Is(err, os.ErrNotExist)
	} else {
		hasError = *ConfigError
	}
	if hasError != nil {

		widgets = []layout.Widget{
			material.H4(th, "Error").Layout,

			func(gtx C) D {
				gtx.Constraints.Max.Y = gtx.Px(unit.Dp(200))
				return material.Editor(th, editor, hasError.Error()+"\n"+Config.ChromePath).Layout(gtx)
			},
		}
	} else {
		widgets = []layout.Widget{
			func(gtx C) D {
				return layout.Flex{
					Axis:    layout.Vertical,
					Spacing: layout.SpaceStart,
				}.Layout(gtx,
					append(ProfileButtons(Config, &margins, th, &startButtons)[:], layout.Rigid(
						layout.Spacer{Height: unit.Dp(25)}.Layout,
					))...,
				)
			},
		}
	}

	for e := range w.Events() {

		switch e := e.(type) {

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			material.List(th, list).Layout(gtx, len(widgets), func(gtx C, i int) D {
				return layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
			})
			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}
