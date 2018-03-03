package bakedlightmap

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"

	"github.com/shadowapex/godot-go/godot/class/bakedlightmapdata"
	"github.com/shadowapex/godot-go/godot/class/visualinstance"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewBakedLightmapFromPointer(ptr gdnative.Pointer) *BakedLightmap {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := BakedLightmap{}
	obj.SetOwner(owner)

	return &obj

}

/*

 */
type BakedLightmap struct {
	visualinstance.VisualInstance
}

func (o *BakedLightmap) BaseClass() string {
	return "BakedLightmap"
}

/*

	Args: [{Null true from_node Object} {False true create_visual_debug bool}], Returns: enum.BakedLightmap::BakeError
*/

/*

	Args: [], Returns: void
*/

func (o *BakedLightmap) DebugBake() {
	log.Println("Calling BakedLightmap.DebugBake()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "debug_bake")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: float
*/

func (o *BakedLightmap) GetBakeCellSize() gdnative.Float {
	log.Println("Calling BakedLightmap.GetBakeCellSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "get_bake_cell_size")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.BakedLightmap::BakeMode
*/

/*
        Undocumented
	Args: [], Returns: enum.BakedLightmap::BakeQuality
*/

/*
        Undocumented
	Args: [], Returns: float
*/

func (o *BakedLightmap) GetCaptureCellSize() gdnative.Float {
	log.Println("Calling BakedLightmap.GetCaptureCellSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "get_capture_cell_size")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/

func (o *BakedLightmap) GetEnergy() gdnative.Float {
	log.Println("Calling BakedLightmap.GetEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "get_energy")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector3
*/

func (o *BakedLightmap) GetExtents() gdnative.Vector3 {
	log.Println("Calling BakedLightmap.GetExtents()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "get_extents")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/

func (o *BakedLightmap) GetImagePath() gdnative.String {
	log.Println("Calling BakedLightmap.GetImagePath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "get_image_path")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: BakedLightmapData
*/

func (o *BakedLightmap) GetLightData() bakedlightmapdata.BakedLightmapData {
	log.Println("Calling BakedLightmap.GetLightData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "get_light_data")

	// Call the parent method.
	// BakedLightmapData
	retPtr := bakedlightmapdata.NewEmptyBakedLightmapData()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := bakedlightmapdata.NewBakedLightmapDataFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/

func (o *BakedLightmap) GetPropagation() gdnative.Float {
	log.Println("Calling BakedLightmap.GetPropagation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "get_propagation")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/

func (o *BakedLightmap) IsHdr() gdnative.Bool {
	log.Println("Calling BakedLightmap.IsHdr()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "is_hdr")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false bake_cell_size float}], Returns: void
*/

func (o *BakedLightmap) SetBakeCellSize(bakeCellSize gdnative.Float) {
	log.Println("Calling BakedLightmap.SetBakeCellSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(bakeCellSize)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "set_bake_cell_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bake_mode int}], Returns: void
*/

func (o *BakedLightmap) SetBakeMode(bakeMode gdnative.Int) {
	log.Println("Calling BakedLightmap.SetBakeMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bakeMode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "set_bake_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bake_quality int}], Returns: void
*/

func (o *BakedLightmap) SetBakeQuality(bakeQuality gdnative.Int) {
	log.Println("Calling BakedLightmap.SetBakeQuality()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bakeQuality)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "set_bake_quality")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false capture_cell_size float}], Returns: void
*/

func (o *BakedLightmap) SetCaptureCellSize(captureCellSize gdnative.Float) {
	log.Println("Calling BakedLightmap.SetCaptureCellSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(captureCellSize)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "set_capture_cell_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false energy float}], Returns: void
*/

func (o *BakedLightmap) SetEnergy(energy gdnative.Float) {
	log.Println("Calling BakedLightmap.SetEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(energy)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "set_energy")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false extents Vector3}], Returns: void
*/

func (o *BakedLightmap) SetExtents(extents gdnative.Vector3) {
	log.Println("Calling BakedLightmap.SetExtents()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(extents)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "set_extents")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false hdr bool}], Returns: void
*/

func (o *BakedLightmap) SetHdr(hdr gdnative.Bool) {
	log.Println("Calling BakedLightmap.SetHdr()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(hdr)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "set_hdr")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false image_path String}], Returns: void
*/

func (o *BakedLightmap) SetImagePath(imagePath gdnative.String) {
	log.Println("Calling BakedLightmap.SetImagePath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(imagePath)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "set_image_path")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false data BakedLightmapData}], Returns: void
*/

func (o *BakedLightmap) SetLightData(data bakedlightmapdata.BakedLightmapData) {
	log.Println("Calling BakedLightmap.SetLightData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(data.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "set_light_data")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false propagation float}], Returns: void
*/

func (o *BakedLightmap) SetPropagation(propagation gdnative.Float) {
	log.Println("Calling BakedLightmap.SetPropagation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(propagation)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmap", "set_propagation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
