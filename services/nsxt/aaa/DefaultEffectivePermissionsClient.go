
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: EffectivePermissions
 * Functions that implement the generated EffectivePermissionsClient interface
 */


package aaa

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

type DefaultEffectivePermissionsClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultEffectivePermissionsClient(connector client.Connector) *DefaultEffectivePermissionsClient {
	interfaceName := "com.vmware.nsx_policy.aaa.effective_permissions"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	eIface := DefaultEffectivePermissionsClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	eIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	eIface.methodNameToDefMap["get"] = eIface.getMethodDefinition()
	return &eIface
}

func (eIface *DefaultEffectivePermissionsClient) Get(featureNameParam string, objectPathParam string) (model.PathPermissionGroup, error) {
	typeConverter := eIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(eIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(effectivePermissionsGetInputType(), typeConverter)
	sv.AddStructField("FeatureName", featureNameParam)
	sv.AddStructField("ObjectPath", objectPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PathPermissionGroup
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := effectivePermissionsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.Invoke(eIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.PathPermissionGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), effectivePermissionsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PathPermissionGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}


func (eIface *DefaultEffectivePermissionsClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := eIface.connector.GetApiProvider().Invoke(eIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (eIface *DefaultEffectivePermissionsClient) getMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(eIface.interfaceName)
	typeConverter := eIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(effectivePermissionsGetInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(effectivePermissionsGetOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultEffectivePermissionsClient.get method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultEffectivePermissionsClient.get method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	eIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultEffectivePermissionsClient.get method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	eIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultEffectivePermissionsClient.get method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	eIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultEffectivePermissionsClient.get method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	eIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultEffectivePermissionsClient.get method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	eIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultEffectivePermissionsClient.get method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}