/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: DetailedStatus.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package sessions

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)

// Possible value for ``source`` of method DetailedStatus#get.
const DetailedStatus_GET_SOURCE_REALTIME = "realtime"
// Possible value for ``source`` of method DetailedStatus#get.
const DetailedStatus_GET_SOURCE_CACHED = "cached"




func detailedStatusGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["service_id"] = bindings.NewStringType()
	fields["session_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["session_id"] = "SessionId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func detailedStatusGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AggregateL2VPNSessionStatusBindingType)
}

func detailedStatusGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["service_id"] = bindings.NewStringType()
	fields["session_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["session_id"] = "SessionId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["locale_service_id"] = bindings.NewStringType()
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["session_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["localeServiceId"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["sessionId"] = bindings.NewStringType()
	pathParams["tier1_id"] = "tier1Id"
	pathParams["session_id"] = "sessionId"
	pathParams["locale_service_id"] = "localeServiceId"
	pathParams["service_id"] = "serviceId"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
	queryParams["source"] = "source"
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
		"GET",
		"/policy/api/v1/infra/tier-1s/{tier1Id}/locale-services/{localeServiceId}/l2vpn-services/{serviceId}/sessions/{sessionId}/detailed-status",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}


