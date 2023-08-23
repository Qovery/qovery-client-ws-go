/*
websocket-gateway

Describe the weboscket endpoints

API version: 0.1.0
Contact: erebe@erebe.eu
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery-ws

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// LogsApiService LogsApi service
type LogsApiService service

type ApiHandleInfraLogsRequestRequest struct {
	ctx context.Context
	ApiService *LogsApiService
	organization string
	cluster string
	project string
	environment string
	service string
	infraComponentType string
}

func (r ApiHandleInfraLogsRequestRequest) Execute() (*ServiceInfraLogResponseDto, *http.Response, error) {
	return r.ApiService.HandleInfraLogsRequestExecute(r)
}

/*
HandleInfraLogsRequest Method for HandleInfraLogsRequest

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param cluster
 @param project
 @param environment
 @param service
 @param infraComponentType
 @return ApiHandleInfraLogsRequestRequest
*/
func (a *LogsApiService) HandleInfraLogsRequest(ctx context.Context, organization string, cluster string, project string, environment string, service string, infraComponentType string) ApiHandleInfraLogsRequestRequest {
	return ApiHandleInfraLogsRequestRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		cluster: cluster,
		project: project,
		environment: environment,
		service: service,
		infraComponentType: infraComponentType,
	}
}

// Execute executes the request
//  @return ServiceInfraLogResponseDto
func (a *LogsApiService) HandleInfraLogsRequestExecute(r ApiHandleInfraLogsRequestRequest) (*ServiceInfraLogResponseDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceInfraLogResponseDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsApiService.HandleInfraLogsRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/infra/logs"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterToString(r.organization, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cluster"+"}", url.PathEscape(parameterToString(r.cluster, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterToString(r.project, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment"+"}", url.PathEscape(parameterToString(r.environment, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"service"+"}", url.PathEscape(parameterToString(r.service, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"infra_component_type"+"}", url.PathEscape(parameterToString(r.infraComponentType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiHandleServiceLogsRequestRequest struct {
	ctx context.Context
	ApiService *LogsApiService
	organization string
	cluster string
	project string
	environment string
	service string
}

func (r ApiHandleServiceLogsRequestRequest) Execute() (*ServiceLogResponseDto, *http.Response, error) {
	return r.ApiService.HandleServiceLogsRequestExecute(r)
}

/*
HandleServiceLogsRequest Method for HandleServiceLogsRequest

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organization
 @param cluster
 @param project
 @param environment
 @param service
 @return ApiHandleServiceLogsRequestRequest
*/
func (a *LogsApiService) HandleServiceLogsRequest(ctx context.Context, organization string, cluster string, project string, environment string, service string) ApiHandleServiceLogsRequestRequest {
	return ApiHandleServiceLogsRequestRequest{
		ApiService: a,
		ctx: ctx,
		organization: organization,
		cluster: cluster,
		project: project,
		environment: environment,
		service: service,
	}
}

// Execute executes the request
//  @return ServiceLogResponseDto
func (a *LogsApiService) HandleServiceLogsRequestExecute(r ApiHandleServiceLogsRequestRequest) (*ServiceLogResponseDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceLogResponseDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsApiService.HandleServiceLogsRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/logs"
	localVarPath = strings.Replace(localVarPath, "{"+"organization"+"}", url.PathEscape(parameterToString(r.organization, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cluster"+"}", url.PathEscape(parameterToString(r.cluster, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", url.PathEscape(parameterToString(r.project, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment"+"}", url.PathEscape(parameterToString(r.environment, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"service"+"}", url.PathEscape(parameterToString(r.service, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
