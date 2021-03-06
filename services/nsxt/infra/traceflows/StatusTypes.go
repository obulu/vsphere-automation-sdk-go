/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Status.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package traceflows

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func statusGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traceflow_id"] = bindings.NewStringType()
	fields["traceflow_status_request"] = bindings.NewReferenceType(model.TraceflowStatusRequestBindingType)
	fieldNameMap["traceflow_id"] = "TraceflowId"
	fieldNameMap["traceflow_status_request"] = "TraceflowStatusRequest"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func statusGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TraceflowStatusResponseBindingType)
}

func statusGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["traceflow_id"] = bindings.NewStringType()
	fields["traceflow_status_request"] = bindings.NewReferenceType(model.TraceflowStatusRequestBindingType)
	fieldNameMap["traceflow_id"] = "TraceflowId"
	fieldNameMap["traceflow_status_request"] = "TraceflowStatusRequest"
	paramsTypeMap["traceflow_id"] = bindings.NewStringType()
	paramsTypeMap["traceflow_status_request"] = bindings.NewReferenceType(model.TraceflowStatusRequestBindingType)
	paramsTypeMap["traceflowId"] = bindings.NewStringType()
	pathParams["traceflow_id"] = "traceflowId"
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
		"traceflow_status_request",
		"GET",
		"/policy/api/v1/infra/traceflows/{traceflowId}/status",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}


