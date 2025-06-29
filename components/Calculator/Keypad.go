package calculator

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

func KeyboardLayout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,

		// First Child
		layout.Flexed(3, func(gtx layout.Context) layout.Dimensions {
			return material.Body1(th, "numbers").Layout(gtx)
		}),

		//Second Child containing inner childs
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return material.Body1(th, "*").Layout(gtx)
				}),
				layout.Flexed(2, func(gtx layout.Context) layout.Dimensions {
					return material.Body1(th, "+").Layout(gtx)
				}),
				layout.Flexed(2, func(gtx layout.Context) layout.Dimensions {
					return material.Body1(th, "=").Layout(gtx)
				}),
			)
		}),
	)
}
