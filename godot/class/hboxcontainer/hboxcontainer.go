package hboxcontainer

import (
	"github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot/class/boxcontainer"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewHBoxContainerFromPointer(ptr gdnative.Pointer) *HBoxContainer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := HBoxContainer{}
	obj.SetOwner(owner)

	return &obj

}

/*
Horizontal box container. See [BoxContainer].
*/
type HBoxContainer struct {
	boxcontainer.BoxContainer
}

func (o *HBoxContainer) BaseClass() string {
	return "HBoxContainer"
}
