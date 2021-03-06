package swag

import (
	"fmt"
	"go/ast"
)

type propertyName struct {
	SchemaType string
	ArrayType  string
}

func parseFieldSelectorExpr(astTypeSelectorExpr *ast.SelectorExpr) propertyName {
	// Support for time.Time as a structure field
	if "Time" == astTypeSelectorExpr.Sel.Name {
		return propertyName{SchemaType: "string", ArrayType: "string"}
	}

	// Support bson.ObjectId type
	if "ObjectId" == astTypeSelectorExpr.Sel.Name {
		return propertyName{SchemaType: "string", ArrayType: "string"}
	}

	panic("not supported 'astSelectorExpr' yet.")
}

// getPropertyName returns the string value for the given field if it exists, otherwise it panics.
// allowedValues: array, boolean, integer, null, number, object, string
func getPropertyName(field *ast.Field) propertyName {
	if astTypeSelectorExpr, ok := field.Type.(*ast.SelectorExpr); ok {
		return parseFieldSelectorExpr(astTypeSelectorExpr)
	}
	if astTypeIdent, ok := field.Type.(*ast.Ident); ok {
		name := astTypeIdent.Name
		schemeType := TransToValidSchemeType(name)
		return propertyName{SchemaType: schemeType, ArrayType: schemeType}
	}
	if ptr, ok := field.Type.(*ast.StarExpr); ok {
		if astTypeSelectorExpr, ok := ptr.X.(*ast.SelectorExpr); ok {
			return parseFieldSelectorExpr(astTypeSelectorExpr)
		}
		if astTypeIdent, ok := ptr.X.(*ast.Ident); ok {
			name := astTypeIdent.Name
			schemeType := TransToValidSchemeType(name)
			return propertyName{SchemaType: schemeType, ArrayType: schemeType}
		}
	}
	if _, ok := field.Type.(*ast.MapType); ok { // if map
		//TODO: support map
		return propertyName{SchemaType: "object", ArrayType: "object"}
	}
	if astTypeArray, ok := field.Type.(*ast.ArrayType); ok { // if array
		str := fmt.Sprintf("%s", astTypeArray.Elt)
		return propertyName{SchemaType: "array", ArrayType: str}
	}
	if _, ok := field.Type.(*ast.StructType); ok { // if struct
		return propertyName{SchemaType: "object", ArrayType: "object"}
	}
	if _, ok := field.Type.(*ast.InterfaceType); ok { // if interface{}
		return propertyName{SchemaType: "object", ArrayType: "object"}
	}
	panic("not supported" + fmt.Sprint(field.Type))
}
