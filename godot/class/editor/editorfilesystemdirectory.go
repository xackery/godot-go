package editor

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot/class/object"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewEditorFileSystemDirectoryFromPointer(ptr gdnative.Pointer) *EditorFileSystemDirectory {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorFileSystemDirectory{}
	obj.SetOwner(owner)

	return &obj

}

/*
A more generalized, low-level variation of the directory concept.
*/
type EditorFileSystemDirectory struct {
	object.Object
}

func (o *EditorFileSystemDirectory) BaseClass() string {
	return "EditorFileSystemDirectory"
}

/*
        Returns the index of the directory with name [code]name[/code] or [code]-1[/code] if not found.
	Args: [{ false name String}], Returns: int
*/

func (o *EditorFileSystemDirectory) FindDirIndex(name gdnative.String) gdnative.Int {
	log.Println("Calling EditorFileSystemDirectory.FindDirIndex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileSystemDirectory", "find_dir_index")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns the index of the file with name [code]name[/code] or [code]-1[/code] if not found.
	Args: [{ false name String}], Returns: int
*/

func (o *EditorFileSystemDirectory) FindFileIndex(name gdnative.String) gdnative.Int {
	log.Println("Calling EditorFileSystemDirectory.FindFileIndex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileSystemDirectory", "find_file_index")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns the name of the file at index [code]idx[/code].
	Args: [{ false idx int}], Returns: String
*/

func (o *EditorFileSystemDirectory) GetFile(idx gdnative.Int) gdnative.String {
	log.Println("Calling EditorFileSystemDirectory.GetFile()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileSystemDirectory", "get_file")

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
        Returns the number of files in this directory.
	Args: [], Returns: int
*/

func (o *EditorFileSystemDirectory) GetFileCount() gdnative.Int {
	log.Println("Calling EditorFileSystemDirectory.GetFileCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileSystemDirectory", "get_file_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns [code]true[/code] if the file at index [code]idx[/code] imported properly.
	Args: [{ false idx int}], Returns: bool
*/

func (o *EditorFileSystemDirectory) GetFileImportIsValid(idx gdnative.Int) gdnative.Bool {
	log.Println("Calling EditorFileSystemDirectory.GetFileImportIsValid()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileSystemDirectory", "get_file_import_is_valid")

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
        Returns the path to the file at index [code]idx[/code].
	Args: [{ false idx int}], Returns: String
*/

func (o *EditorFileSystemDirectory) GetFilePath(idx gdnative.Int) gdnative.String {
	log.Println("Calling EditorFileSystemDirectory.GetFilePath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileSystemDirectory", "get_file_path")

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
        Returns the file extension of the file at index [code]idx[/code].
	Args: [{ false idx int}], Returns: String
*/

func (o *EditorFileSystemDirectory) GetFileType(idx gdnative.Int) gdnative.String {
	log.Println("Calling EditorFileSystemDirectory.GetFileType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileSystemDirectory", "get_file_type")

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
        Returns the name of this directory.
	Args: [], Returns: String
*/

func (o *EditorFileSystemDirectory) GetName() gdnative.String {
	log.Println("Calling EditorFileSystemDirectory.GetName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileSystemDirectory", "get_name")

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
        Returns the parent directory for this directory or null if called on a directory at [code]res://[/code] or [code]user://[/code].
	Args: [], Returns: EditorFileSystemDirectory
*/

func (o *EditorFileSystemDirectory) GetParent() EditorFileSystemDirectory {
	log.Println("Calling EditorFileSystemDirectory.GetParent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileSystemDirectory", "get_parent")

	// Call the parent method.
	// EditorFileSystemDirectory
	retPtr := NewEmptyEditorFileSystemDirectory()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewEditorFileSystemDirectoryFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns the path to this directory.
	Args: [], Returns: String
*/

func (o *EditorFileSystemDirectory) GetPath() gdnative.String {
	log.Println("Calling EditorFileSystemDirectory.GetPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileSystemDirectory", "get_path")

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
        Returns the subdirectory at index [code]idx[/code].
	Args: [{ false idx int}], Returns: EditorFileSystemDirectory
*/

func (o *EditorFileSystemDirectory) GetSubdir(idx gdnative.Int) EditorFileSystemDirectory {
	log.Println("Calling EditorFileSystemDirectory.GetSubdir()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileSystemDirectory", "get_subdir")

	// Call the parent method.
	// EditorFileSystemDirectory
	retPtr := NewEmptyEditorFileSystemDirectory()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewEditorFileSystemDirectoryFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns the number of subdirectories in this directory.
	Args: [], Returns: int
*/

func (o *EditorFileSystemDirectory) GetSubdirCount() gdnative.Int {
	log.Println("Calling EditorFileSystemDirectory.GetSubdirCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileSystemDirectory", "get_subdir_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}
