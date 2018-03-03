package javascript

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

func NewjavaScriptFromPointer(ptr gdnative.Pointer) *javaScript {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := javaScript{}
	obj.SetOwner(owner)

	return &obj

}

func newSingletonJavaScript() *javaScript {
	obj := &javaScript{}
	ptr := C.godot_global_get_singleton(C.CString("JavaScript"))
	obj.owner = (*C.godot_object)(ptr)
	return obj
}

/*
   The JavaScript singleton is implemented only in HTML5 export. It's used to access the browser's JavaScript context. This allows interaction with embedding pages or calling third-party JavaScript APIs.
*/
var JavaScript = newSingletonJavaScript()

/*
The JavaScript singleton is implemented only in HTML5 export. It's used to access the browser's JavaScript context. This allows interaction with embedding pages or calling third-party JavaScript APIs.
*/
type javaScript struct {
	object.Object
}

func (o *javaScript) BaseClass() string {
	return "JavaScript"
}

/*
        Execute the string [code]code[/code] as JavaScript code within the browser window. This is a call to the actual global JavaScript function [code]eval()[/code]. If [code]use_global_execution_context[/code] is [code]true[/code], the code will be evaluated in the global execution context. Otherwise, it is evaluated in the execution context of a function within the engine's runtime environment.
	Args: [{ false code String} {False true use_global_execution_context bool}], Returns: Variant
*/

func (o *javaScript) Eval(code gdnative.String, useGlobalExecutionContext gdnative.Bool) gdnative.Variant {
	log.Println("Calling JavaScript.Eval()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(code)
	ptrArguments[1] = gdnative.NewPointerFromBool(useGlobalExecutionContext)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("JavaScript", "eval")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}
