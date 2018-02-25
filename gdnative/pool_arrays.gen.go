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
#include <gdnative/pool_arrays.h>
*/
import "C"

type PoolArrayReadAccess struct {
	base *C.godot_pool_array_read_access
}

func (t *PoolArrayReadAccess) getBase() *C.godot_pool_array_read_access {
	return t.base
}

type PoolByteArrayReadAccess PoolArrayReadAccess

type PoolIntArrayReadAccess PoolArrayReadAccess

type PoolRealArrayReadAccess PoolArrayReadAccess

type PoolStringArrayReadAccess PoolArrayReadAccess

type PoolVector2ArrayReadAccess PoolArrayReadAccess

type PoolVector3ArrayReadAccess PoolArrayReadAccess

type PoolColorArrayReadAccess PoolArrayReadAccess

type PoolArrayWriteAccess struct {
	base *C.godot_pool_array_write_access
}

func (t *PoolArrayWriteAccess) getBase() *C.godot_pool_array_write_access {
	return t.base
}

type PoolByteArrayWriteAccess PoolArrayWriteAccess

type PoolIntArrayWriteAccess PoolArrayWriteAccess

type PoolRealArrayWriteAccess PoolArrayWriteAccess

type PoolStringArrayWriteAccess PoolArrayWriteAccess

type PoolVector2ArrayWriteAccess PoolArrayWriteAccess

type PoolVector3ArrayWriteAccess PoolArrayWriteAccess

type PoolColorArrayWriteAccess PoolArrayWriteAccess

type PoolByteArray struct {
	base *C.godot_pool_byte_array
}

func (t *PoolByteArray) getBase() *C.godot_pool_byte_array {
	return t.base
}

type PoolIntArray struct {
	base *C.godot_pool_int_array
}

func (t *PoolIntArray) getBase() *C.godot_pool_int_array {
	return t.base
}

type PoolRealArray struct {
	base *C.godot_pool_real_array
}

func (t *PoolRealArray) getBase() *C.godot_pool_real_array {
	return t.base
}

type PoolStringArray struct {
	base *C.godot_pool_string_array
}

func (t *PoolStringArray) getBase() *C.godot_pool_string_array {
	return t.base
}

type PoolVector2Array struct {
	base *C.godot_pool_vector2_array
}

func (t *PoolVector2Array) getBase() *C.godot_pool_vector2_array {
	return t.base
}

type PoolVector3Array struct {
	base *C.godot_pool_vector3_array
}

func (t *PoolVector3Array) getBase() *C.godot_pool_vector3_array {
	return t.base
}

type PoolColorArray struct {
	base *C.godot_pool_color_array
}

func (t *PoolColorArray) getBase() *C.godot_pool_color_array {
	return t.base
}
