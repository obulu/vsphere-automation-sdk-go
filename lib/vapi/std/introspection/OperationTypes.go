/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Operation.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package introspection

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)



// The OperationDataDefinition structure describes a vAPI data type.
type OperationDataDefinition struct {
    // Data type of the value.
	Type_ OperationDataDefinitionDataType
    // Contains the element definition for generic data types like List and Optional.
	ElementDefinition *OperationDataDefinition
    // Fully qualified name of the structure.
	Name *string
    // Fields of the structure type. The key of the map is the canonical name of the field and the value is the OperationDataDefinition for the field. The order of the structure fields defined in IDL is not maintained by the Operation service.
	Fields map[string]OperationDataDefinition
}

// The OperationDataDefinitionDataType enumeration provides values representing the data types supported by the vAPI infrastructure.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type OperationDataDefinitionDataType string

const (
    // Indicates the value is a binary type.
	OperationDataDefinitionDataType_BINARY OperationDataDefinitionDataType = "BINARY"
    // Indicates the value is a boolean type. The possible values are True and False equivalent of the language used to invoke this operation.
	OperationDataDefinitionDataType_BOOLEAN OperationDataDefinitionDataType = "BOOLEAN"
    // Indicates the value is a double type. It is a 64 bit floating point number.
	OperationDataDefinitionDataType_DOUBLE OperationDataDefinitionDataType = "DOUBLE"
    // Indicates the value is a dynamic structure. This means, any data of type OperationDataDefinitionDataType#DataType_STRUCTURE can be used.
	OperationDataDefinitionDataType_DYNAMIC_STRUCTURE OperationDataDefinitionDataType = "DYNAMIC_STRUCTURE"
    // Indicates the value is a specific error type.
	OperationDataDefinitionDataType_ERROR OperationDataDefinitionDataType = "ERROR"
    // Indicates the value is arbitrary error type. This means, any data of type OperationDataDefinitionDataType#DataType_ERROR can be used.
	OperationDataDefinitionDataType_ANY_ERROR OperationDataDefinitionDataType = "ANY_ERROR"
    // Indicates the value is a list data type. Any value of this type can have zero or more elements in the list.
	OperationDataDefinitionDataType_LIST OperationDataDefinitionDataType = "LIST"
    // Indicates the value is a long data type. It is a 64 bit signed integer number.
	OperationDataDefinitionDataType_LONG OperationDataDefinitionDataType = "LONG"
    // Indicates the value is an opaque type. This means, data of any OperationDataDefinitionDataType can be used.
	OperationDataDefinitionDataType_OPAQUE OperationDataDefinitionDataType = "OPAQUE"
    // Indicates the value is an optional data type. Any value of this type can be null.
	OperationDataDefinitionDataType_OPTIONAL OperationDataDefinitionDataType = "OPTIONAL"
    // Indicates the value is a secret data type. This is used for sensitive information. The server will not log any data of this type and if possible wipe the data from the memory after usage.
	OperationDataDefinitionDataType_SECRET OperationDataDefinitionDataType = "SECRET"
    // Indicates the value is a string data type. This is a unicode string.
	OperationDataDefinitionDataType_STRING OperationDataDefinitionDataType = "STRING"
    // Indicates the value is a structure data type. A structure has string identifier and a set of fields with corresponding values.
	OperationDataDefinitionDataType_STRUCTURE OperationDataDefinitionDataType = "STRUCTURE"
    // Indicates the value is a structure reference. This is used to break circular dependencies in the type references. This just has a string identifier of the structure. Clients have to maintain a list of structures already visited and use that to resolve this reference.
	OperationDataDefinitionDataType_STRUCTURE_REF OperationDataDefinitionDataType = "STRUCTURE_REF"
    // Indicates the value is a void data type.
	OperationDataDefinitionDataType_VOID OperationDataDefinitionDataType = "VOID"
)

func (d OperationDataDefinitionDataType) OperationDataDefinitionDataType() bool {
	switch d {
	case OperationDataDefinitionDataType_BINARY:
		return true
	case OperationDataDefinitionDataType_BOOLEAN:
		return true
	case OperationDataDefinitionDataType_DOUBLE:
		return true
	case OperationDataDefinitionDataType_DYNAMIC_STRUCTURE:
		return true
	case OperationDataDefinitionDataType_ERROR:
		return true
	case OperationDataDefinitionDataType_ANY_ERROR:
		return true
	case OperationDataDefinitionDataType_LIST:
		return true
	case OperationDataDefinitionDataType_LONG:
		return true
	case OperationDataDefinitionDataType_OPAQUE:
		return true
	case OperationDataDefinitionDataType_OPTIONAL:
		return true
	case OperationDataDefinitionDataType_SECRET:
		return true
	case OperationDataDefinitionDataType_STRING:
		return true
	case OperationDataDefinitionDataType_STRUCTURE:
		return true
	case OperationDataDefinitionDataType_STRUCTURE_REF:
		return true
	case OperationDataDefinitionDataType_VOID:
		return true
	default:
		return false
	}
}


// Information about a vAPI operation.
type OperationInfo struct {
    // OperationDataDefinition describing the operation input. 
    //
    //  The OperationDataDefinition#type of this field will be OperationDataDefinitionDataType#DataType_STRUCTURE. The keys of OperationDataDefinition#fields are the names of the operation parameters, and the values of OperationDataDefinition#fields describe the type of the operation parameters.
	InputDefinition OperationDataDefinition
    // OperationDataDefinition describing the operation output.
	OutputDefinition OperationDataDefinition
    // List of OperationDataDefinition describing the errors that the operation might report. 
    //
    //  The OperationDataDefinition#type of every element in this list will be OperationDataDefinitionDataType#DataType_ERROR.
	ErrorDefinitions []OperationDataDefinition
}



func operationListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fieldNameMap["service_id"] = "ServiceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func operationListOutputType() bindings.BindingType {
	return bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vapi.operation"}, ""), reflect.TypeOf(map[string]bool{}))
}

func operationListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["service_id"] = bindings.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fieldNameMap["service_id"] = "ServiceId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404})
}

func operationGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fields["operation_id"] = bindings.NewIdType([]string{"com.vmware.vapi.operation"}, "")
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["operation_id"] = "OperationId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func operationGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(OperationInfoBindingType)
}

func operationGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["service_id"] = bindings.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fields["operation_id"] = bindings.NewIdType([]string{"com.vmware.vapi.operation"}, "")
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["operation_id"] = "OperationId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404})
}


func OperationDataDefinitionBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vapi.std.introspection.operation.data_definition.data_type", reflect.TypeOf(OperationDataDefinitionDataType(OperationDataDefinitionDataType_BINARY)))
	fieldNameMap["type"] = "Type_"
	fields["element_definition"] = bindings.NewOptionalType(bindings.NewReferenceType(OperationDataDefinitionBindingType))
	fieldNameMap["element_definition"] = "ElementDefinition"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["fields"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(OperationDataDefinitionBindingType),reflect.TypeOf(map[string]OperationDataDefinition{})))
	fieldNameMap["fields"] = "Fields"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"OPTIONAL": []bindings.FieldData{
				bindings.NewFieldData("element_definition", true),
			},
			"LIST": []bindings.FieldData{
				bindings.NewFieldData("element_definition", true),
			},
			"STRUCTURE": []bindings.FieldData{
				bindings.NewFieldData("name", true),
				bindings.NewFieldData("fields", true),
			},
			"STRUCTURE_REF": []bindings.FieldData{
				bindings.NewFieldData("name", true),
			},
			"ERROR": []bindings.FieldData{
				bindings.NewFieldData("name", true),
				bindings.NewFieldData("fields", true),
			},
			"BINARY": []bindings.FieldData{},
			"BOOLEAN": []bindings.FieldData{},
			"DOUBLE": []bindings.FieldData{},
			"DYNAMIC_STRUCTURE": []bindings.FieldData{},
			"ANY_ERROR": []bindings.FieldData{},
			"LONG": []bindings.FieldData{},
			"OPAQUE": []bindings.FieldData{},
			"SECRET": []bindings.FieldData{},
			"STRING": []bindings.FieldData{},
			"VOID": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vapi.std.introspection.operation.data_definition", fields, reflect.TypeOf(OperationDataDefinition{}), fieldNameMap, validators)
}

func OperationInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["input_definition"] = bindings.NewReferenceType(OperationDataDefinitionBindingType)
	fieldNameMap["input_definition"] = "InputDefinition"
	fields["output_definition"] = bindings.NewReferenceType(OperationDataDefinitionBindingType)
	fieldNameMap["output_definition"] = "OutputDefinition"
	fields["error_definitions"] = bindings.NewListType(bindings.NewReferenceType(OperationDataDefinitionBindingType), reflect.TypeOf([]OperationDataDefinition{}))
	fieldNameMap["error_definitions"] = "ErrorDefinitions"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.std.introspection.operation.info", fields, reflect.TypeOf(OperationInfo{}), fieldNameMap, validators)
}

