package pckpacker

import (
	"github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot/class/reference"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewPCKPackerFromPointer(ptr gdnative.Pointer) *PCKPacker {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PCKPacker{}
	obj.SetOwner(owner)

	return &obj

}

/*

 */
type PCKPacker struct {
	reference.Reference
}

func (o *PCKPacker) BaseClass() string {
	return "PCKPacker"
}

/*

	Args: [{ false pck_path String} { false source_path String}], Returns: enum.Error
*/

/*

	Args: [{ false verbose bool}], Returns: enum.Error
*/

/*

	Args: [{ false pck_name String} { false alignment int}], Returns: enum.Error
*/
