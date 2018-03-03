package editor

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot/class/resource"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewEditorSettingsFromPointer(ptr gdnative.Pointer) *EditorSettings {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorSettings{}
	obj.SetOwner(owner)

	return &obj

}

/*
Object that holds the project-independent editor settings. These settings are generally visible in the Editor Settings menu. Accessing the settings is done by using the regular [Object] API, such as: [codeblock] settings.set(prop,value) settings.get(prop) list_of_settings = settings.get_property_list() [/codeblock]
*/
type EditorSettings struct {
	resource.Resource
}

func (o *EditorSettings) BaseClass() string {
	return "EditorSettings"
}

/*
        Add a custom property info to a property. The dictionary must contain: name:[String](the name of the property) and type:[int](see TYPE_* in [@GlobalScope]), and optionally hint:[int](see PROPERTY_HINT_* in [@GlobalScope]), hint_string:[String]. Example: [codeblock] editor_settings.set("category/property_name", 0) var property_info = { "name": "category/property_name", "type": TYPE_INT, "hint": PROPERTY_HINT_ENUM, "hint_string": "one,two,three" } editor_settings.add_property_info(property_info) [/codeblock]
	Args: [{ false info Dictionary}], Returns: void
*/

func (o *EditorSettings) AddPropertyInfo(info gdnative.Dictionary) {
	log.Println("Calling EditorSettings.AddPropertyInfo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromDictionary(info)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSettings", "add_property_info")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Erase a given setting (pass full property path).
	Args: [{ false property String}], Returns: void
*/

func (o *EditorSettings) Erase(property gdnative.String) {
	log.Println("Calling EditorSettings.Erase()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(property)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSettings", "erase")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Get the list of favorite directories for this project.
	Args: [], Returns: PoolStringArray
*/

func (o *EditorSettings) GetFavoriteDirs() gdnative.PoolStringArray {
	log.Println("Calling EditorSettings.GetFavoriteDirs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSettings", "get_favorite_dirs")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Get the specific project settings path. Projects all have a unique sub-directory inside the settings path where project specific settings are saved.
	Args: [], Returns: String
*/

func (o *EditorSettings) GetProjectSettingsDir() gdnative.String {
	log.Println("Calling EditorSettings.GetProjectSettingsDir()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSettings", "get_project_settings_dir")

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
        Get the list of recently visited folders in the file dialog for this project.
	Args: [], Returns: PoolStringArray
*/

func (o *EditorSettings) GetRecentDirs() gdnative.PoolStringArray {
	log.Println("Calling EditorSettings.GetRecentDirs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSettings", "get_recent_dirs")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*

	Args: [{ false name String}], Returns: Variant
*/

func (o *EditorSettings) GetSetting(name gdnative.String) gdnative.Variant {
	log.Println("Calling EditorSettings.GetSetting()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSettings", "get_setting")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Get the global settings path for the engine. Inside this path you can find some standard paths such as: settings/tmp - used for temporary storage of files settings/templates - where export templates are located
	Args: [], Returns: String
*/

func (o *EditorSettings) GetSettingsDir() gdnative.String {
	log.Println("Calling EditorSettings.GetSettingsDir()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSettings", "get_settings_dir")

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

	Args: [{ false name String}], Returns: bool
*/

func (o *EditorSettings) HasSetting(name gdnative.String) gdnative.Bool {
	log.Println("Calling EditorSettings.HasSetting()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSettings", "has_setting")

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

	Args: [{ false name String}], Returns: bool
*/

func (o *EditorSettings) PropertyCanRevert(name gdnative.String) gdnative.Bool {
	log.Println("Calling EditorSettings.PropertyCanRevert()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSettings", "property_can_revert")

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

	Args: [{ false name String}], Returns: Variant
*/

func (o *EditorSettings) PropertyGetRevert(name gdnative.String) gdnative.Variant {
	log.Println("Calling EditorSettings.PropertyGetRevert()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSettings", "property_get_revert")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Set the list of favorite directories for this project.
	Args: [{ false dirs PoolStringArray}], Returns: void
*/

func (o *EditorSettings) SetFavoriteDirs(dirs gdnative.PoolStringArray) {
	log.Println("Calling EditorSettings.SetFavoriteDirs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolStringArray(dirs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSettings", "set_favorite_dirs")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*

	Args: [{ false name String} { false value Variant} { false update_current bool}], Returns: void
*/

func (o *EditorSettings) SetInitialValue(name gdnative.String, value gdnative.Variant, updateCurrent gdnative.Bool) {
	log.Println("Calling EditorSettings.SetInitialValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromVariant(value)
	ptrArguments[2] = gdnative.NewPointerFromBool(updateCurrent)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSettings", "set_initial_value")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Set the list of recently visited folders in the file dialog for this project.
	Args: [{ false dirs PoolStringArray}], Returns: void
*/

func (o *EditorSettings) SetRecentDirs(dirs gdnative.PoolStringArray) {
	log.Println("Calling EditorSettings.SetRecentDirs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolStringArray(dirs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSettings", "set_recent_dirs")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*

	Args: [{ false name String} { false value Variant}], Returns: void
*/

func (o *EditorSettings) SetSetting(name gdnative.String, value gdnative.Variant) {
	log.Println("Calling EditorSettings.SetSetting()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromVariant(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSettings", "set_setting")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
