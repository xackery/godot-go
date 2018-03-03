package hscrollbar

import (
	"github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot/class/scrollbar"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewHScrollBarFromPointer(ptr gdnative.Pointer) *HScrollBar {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := HScrollBar{}
	obj.SetOwner(owner)

	return &obj

}

/*
Horizontal scroll bar. See [ScrollBar]. This one goes from left (min) to right (max).
*/
type HScrollBar struct {
	scrollbar.ScrollBar
}

func (o *HScrollBar) BaseClass() string {
	return "HScrollBar"
}
