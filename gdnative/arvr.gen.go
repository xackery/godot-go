package gdnative

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "types.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
#include <arvr/godot_arvr.h>
*/
import "C"

type ArvrInterfaceGdnative struct {
	base *C.godot_arvr_interface_gdnative
}

func (t *ArvrInterfaceGdnative) getBase() *C.godot_arvr_interface_gdnative {
	return t.base
}
