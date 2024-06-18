// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddServiceSharedAccountsRequest struct {
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-63b8a060e9d54cxxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// This parameter is required.
	SharedAccounts []*AddServiceSharedAccountsRequestSharedAccounts `json:"SharedAccounts,omitempty" xml:"SharedAccounts,omitempty" type:"Repeated"`
	// example:
	//
	// SharedAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddServiceSharedAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddServiceSharedAccountsRequest) GoString() string {
	return s.String()
}

func (s *AddServiceSharedAccountsRequest) SetClientToken(v string) *AddServiceSharedAccountsRequest {
	s.ClientToken = &v
	return s
}

func (s *AddServiceSharedAccountsRequest) SetRegionId(v string) *AddServiceSharedAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *AddServiceSharedAccountsRequest) SetServiceId(v string) *AddServiceSharedAccountsRequest {
	s.ServiceId = &v
	return s
}

func (s *AddServiceSharedAccountsRequest) SetSharedAccounts(v []*AddServiceSharedAccountsRequestSharedAccounts) *AddServiceSharedAccountsRequest {
	s.SharedAccounts = v
	return s
}

func (s *AddServiceSharedAccountsRequest) SetType(v string) *AddServiceSharedAccountsRequest {
	s.Type = &v
	return s
}

type AddServiceSharedAccountsRequestSharedAccounts struct {
	// This parameter is required.
	//
	// example:
	//
	// Accessible
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	UserAliUid *string `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
}

func (s AddServiceSharedAccountsRequestSharedAccounts) String() string {
	return tea.Prettify(s)
}

func (s AddServiceSharedAccountsRequestSharedAccounts) GoString() string {
	return s.String()
}

func (s *AddServiceSharedAccountsRequestSharedAccounts) SetPermission(v string) *AddServiceSharedAccountsRequestSharedAccounts {
	s.Permission = &v
	return s
}

func (s *AddServiceSharedAccountsRequestSharedAccounts) SetUserAliUid(v string) *AddServiceSharedAccountsRequestSharedAccounts {
	s.UserAliUid = &v
	return s
}

type AddServiceSharedAccountsResponseBody struct {
	// example:
	//
	// E2815213-EA4F-5759-8EA1-56DD051BB3FD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddServiceSharedAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddServiceSharedAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *AddServiceSharedAccountsResponseBody) SetRequestId(v string) *AddServiceSharedAccountsResponseBody {
	s.RequestId = &v
	return s
}

type AddServiceSharedAccountsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddServiceSharedAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddServiceSharedAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddServiceSharedAccountsResponse) GoString() string {
	return s.String()
}

func (s *AddServiceSharedAccountsResponse) SetHeaders(v map[string]*string) *AddServiceSharedAccountsResponse {
	s.Headers = v
	return s
}

func (s *AddServiceSharedAccountsResponse) SetStatusCode(v int32) *AddServiceSharedAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddServiceSharedAccountsResponse) SetBody(v *AddServiceSharedAccountsResponseBody) *AddServiceSharedAccountsResponse {
	s.Body = v
	return s
}

type ContinueDeployServiceInstanceRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// {"NodeCount": 3, "SystemDiskSize": 40, "InstancePassword": "******"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// si-0e6fca6a51a54420****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s ContinueDeployServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceRequest) SetClientToken(v string) *ContinueDeployServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetDryRun(v bool) *ContinueDeployServiceInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetParameters(v string) *ContinueDeployServiceInstanceRequest {
	s.Parameters = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetRegionId(v string) *ContinueDeployServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetServiceInstanceId(v string) *ContinueDeployServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type ContinueDeployServiceInstanceResponseBody struct {
	DryRunResult *ContinueDeployServiceInstanceResponseBodyDryRunResult `json:"DryRunResult,omitempty" xml:"DryRunResult,omitempty" type:"Struct"`
	// example:
	//
	// 82DF27ED-E538-5AC0-A11C-39334A873189
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s ContinueDeployServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceResponseBody) SetDryRunResult(v *ContinueDeployServiceInstanceResponseBodyDryRunResult) *ContinueDeployServiceInstanceResponseBody {
	s.DryRunResult = v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBody) SetRequestId(v string) *ContinueDeployServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBody) SetServiceInstanceId(v string) *ContinueDeployServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

type ContinueDeployServiceInstanceResponseBodyDryRunResult struct {
	ParametersAllowedToBeModified              []*string `json:"ParametersAllowedToBeModified,omitempty" xml:"ParametersAllowedToBeModified,omitempty" type:"Repeated"`
	ParametersConditionallyAllowedToBeModified []*string `json:"ParametersConditionallyAllowedToBeModified,omitempty" xml:"ParametersConditionallyAllowedToBeModified,omitempty" type:"Repeated"`
	ParametersNotAllowedToBeModified           []*string `json:"ParametersNotAllowedToBeModified,omitempty" xml:"ParametersNotAllowedToBeModified,omitempty" type:"Repeated"`
}

func (s ContinueDeployServiceInstanceResponseBodyDryRunResult) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceResponseBodyDryRunResult) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) SetParametersAllowedToBeModified(v []*string) *ContinueDeployServiceInstanceResponseBodyDryRunResult {
	s.ParametersAllowedToBeModified = v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) SetParametersConditionallyAllowedToBeModified(v []*string) *ContinueDeployServiceInstanceResponseBodyDryRunResult {
	s.ParametersConditionallyAllowedToBeModified = v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) SetParametersNotAllowedToBeModified(v []*string) *ContinueDeployServiceInstanceResponseBodyDryRunResult {
	s.ParametersNotAllowedToBeModified = v
	return s
}

type ContinueDeployServiceInstanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContinueDeployServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContinueDeployServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceResponse) SetHeaders(v map[string]*string) *ContinueDeployServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *ContinueDeployServiceInstanceResponse) SetStatusCode(v int32) *ContinueDeployServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinueDeployServiceInstanceResponse) SetBody(v *ContinueDeployServiceInstanceResponseBody) *ContinueDeployServiceInstanceResponse {
	s.Body = v
	return s
}

type CreateArtifactRequest struct {
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId       *string                                `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	ArtifactProperty *CreateArtifactRequestArtifactProperty `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name             *string                     `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceGroupId  *string                     `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SupportRegionIds []*string                   `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	Tag              []*CreateArtifactRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequest) SetArtifactId(v string) *CreateArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *CreateArtifactRequest) SetArtifactProperty(v *CreateArtifactRequestArtifactProperty) *CreateArtifactRequest {
	s.ArtifactProperty = v
	return s
}

func (s *CreateArtifactRequest) SetArtifactType(v string) *CreateArtifactRequest {
	s.ArtifactType = &v
	return s
}

func (s *CreateArtifactRequest) SetDescription(v string) *CreateArtifactRequest {
	s.Description = &v
	return s
}

func (s *CreateArtifactRequest) SetName(v string) *CreateArtifactRequest {
	s.Name = &v
	return s
}

func (s *CreateArtifactRequest) SetResourceGroupId(v string) *CreateArtifactRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateArtifactRequest) SetSupportRegionIds(v []*string) *CreateArtifactRequest {
	s.SupportRegionIds = v
	return s
}

func (s *CreateArtifactRequest) SetTag(v []*CreateArtifactRequestTag) *CreateArtifactRequest {
	s.Tag = v
	return s
}

func (s *CreateArtifactRequest) SetVersionName(v string) *CreateArtifactRequest {
	s.VersionName = &v
	return s
}

type CreateArtifactRequestArtifactProperty struct {
	// example:
	//
	// cmjj00xxxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// V1.0
	CommodityVersion *string `json:"CommodityVersion,omitempty" xml:"CommodityVersion,omitempty"`
	// example:
	//
	// {\\"WorkDir\\":\\"/root\\",\\"CommandType\\":\\"RunShellScript\\",\\"Platform\\":\\"Linux\\",\\"Script\\":\\"echo hello\\"}
	FileScriptMetadata *string `json:"FileScriptMetadata,omitempty" xml:"FileScriptMetadata,omitempty"`
	// example:
	//
	// m-0xij191j9cuev6xxxxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RepoId   *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// example:
	//
	// {"ScriptMetadata":"{\\"CommandType\\":\\"RunShellScript\\",\\"Platform\\":\\"Linux\\",\\"Script\\":\\"ls\\"}"}
	ScriptMetadata *string `json:"ScriptMetadata,omitempty" xml:"ScriptMetadata,omitempty"`
	Tag            *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// https://service-info-private.oss-cn-hangzhou.aliyuncs.com/1309208528xxxxxx/template/2e1ce8fc-xxxx-481c-9e8e-789ba9db487d.json
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateArtifactRequestArtifactProperty) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactRequestArtifactProperty) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequestArtifactProperty) SetCommodityCode(v string) *CreateArtifactRequestArtifactProperty {
	s.CommodityCode = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetCommodityVersion(v string) *CreateArtifactRequestArtifactProperty {
	s.CommodityVersion = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetFileScriptMetadata(v string) *CreateArtifactRequestArtifactProperty {
	s.FileScriptMetadata = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetImageId(v string) *CreateArtifactRequestArtifactProperty {
	s.ImageId = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetRegionId(v string) *CreateArtifactRequestArtifactProperty {
	s.RegionId = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetRepoId(v string) *CreateArtifactRequestArtifactProperty {
	s.RepoId = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetRepoName(v string) *CreateArtifactRequestArtifactProperty {
	s.RepoName = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetScriptMetadata(v string) *CreateArtifactRequestArtifactProperty {
	s.ScriptMetadata = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetTag(v string) *CreateArtifactRequestArtifactProperty {
	s.Tag = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetUrl(v string) *CreateArtifactRequestArtifactProperty {
	s.Url = &v
	return s
}

type CreateArtifactRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateArtifactRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactRequestTag) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequestTag) SetKey(v string) *CreateArtifactRequestTag {
	s.Key = &v
	return s
}

func (s *CreateArtifactRequestTag) SetValue(v string) *CreateArtifactRequestTag {
	s.Value = &v
	return s
}

type CreateArtifactShrinkRequest struct {
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId             *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	ArtifactPropertyShrink *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name             *string                           `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceGroupId  *string                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SupportRegionIds []*string                         `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	Tag              []*CreateArtifactShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateArtifactShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactShrinkRequest) SetArtifactId(v string) *CreateArtifactShrinkRequest {
	s.ArtifactId = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetArtifactPropertyShrink(v string) *CreateArtifactShrinkRequest {
	s.ArtifactPropertyShrink = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetArtifactType(v string) *CreateArtifactShrinkRequest {
	s.ArtifactType = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetDescription(v string) *CreateArtifactShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetName(v string) *CreateArtifactShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetResourceGroupId(v string) *CreateArtifactShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetSupportRegionIds(v []*string) *CreateArtifactShrinkRequest {
	s.SupportRegionIds = v
	return s
}

func (s *CreateArtifactShrinkRequest) SetTag(v []*CreateArtifactShrinkRequestTag) *CreateArtifactShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateArtifactShrinkRequest) SetVersionName(v string) *CreateArtifactShrinkRequest {
	s.VersionName = &v
	return s
}

type CreateArtifactShrinkRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateArtifactShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateArtifactShrinkRequestTag) SetKey(v string) *CreateArtifactShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateArtifactShrinkRequestTag) SetValue(v string) *CreateArtifactShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateArtifactResponseBody struct {
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// example:
	//
	// {\\"CommodityCode\\":\\"cmjj0005xxxx\\",\\"CommodityVersion\\":\\"V2022xxxx\\"}
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// example:
	//
	// 2
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2022-11-11T12:00:00Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1
	MaxVersion *int64  `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// B1A0198B-F316-1B72-B8DD-28B6F6D6XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// [
	//
	// 			"cn-beijing",
	//
	// 			"cn-hangzhou",
	//
	// 			"cn-shanghai"
	//
	// 		]
	SupportRegionIds *string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty"`
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateArtifactResponseBody) SetArtifactId(v string) *CreateArtifactResponseBody {
	s.ArtifactId = &v
	return s
}

func (s *CreateArtifactResponseBody) SetArtifactProperty(v string) *CreateArtifactResponseBody {
	s.ArtifactProperty = &v
	return s
}

func (s *CreateArtifactResponseBody) SetArtifactType(v string) *CreateArtifactResponseBody {
	s.ArtifactType = &v
	return s
}

func (s *CreateArtifactResponseBody) SetArtifactVersion(v string) *CreateArtifactResponseBody {
	s.ArtifactVersion = &v
	return s
}

func (s *CreateArtifactResponseBody) SetDescription(v string) *CreateArtifactResponseBody {
	s.Description = &v
	return s
}

func (s *CreateArtifactResponseBody) SetGmtModified(v string) *CreateArtifactResponseBody {
	s.GmtModified = &v
	return s
}

func (s *CreateArtifactResponseBody) SetMaxVersion(v int64) *CreateArtifactResponseBody {
	s.MaxVersion = &v
	return s
}

func (s *CreateArtifactResponseBody) SetName(v string) *CreateArtifactResponseBody {
	s.Name = &v
	return s
}

func (s *CreateArtifactResponseBody) SetRequestId(v string) *CreateArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateArtifactResponseBody) SetStatus(v string) *CreateArtifactResponseBody {
	s.Status = &v
	return s
}

func (s *CreateArtifactResponseBody) SetSupportRegionIds(v string) *CreateArtifactResponseBody {
	s.SupportRegionIds = &v
	return s
}

func (s *CreateArtifactResponseBody) SetVersionName(v string) *CreateArtifactResponseBody {
	s.VersionName = &v
	return s
}

type CreateArtifactResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactResponse) GoString() string {
	return s.String()
}

func (s *CreateArtifactResponse) SetHeaders(v map[string]*string) *CreateArtifactResponse {
	s.Headers = v
	return s
}

func (s *CreateArtifactResponse) SetStatusCode(v int32) *CreateArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateArtifactResponse) SetBody(v *CreateArtifactResponseBody) *CreateArtifactResponse {
	s.Body = v
	return s
}

type CreateServiceRequest struct {
	AlarmMetadata *string `json:"AlarmMetadata,omitempty" xml:"AlarmMetadata,omitempty"`
	// example:
	//
	// Manual
	ApprovalType    *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	BuildParameters *string `json:"BuildParameters,omitempty" xml:"BuildParameters,omitempty"`
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// example:
	//
	// 0
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// false
	IsSupportOperated *bool   `json:"IsSupportOperated,omitempty" xml:"IsSupportOperated,omitempty"`
	LicenseMetadata   *string `json:"LicenseMetadata,omitempty" xml:"LicenseMetadata,omitempty"`
	LogMetadata       *string `json:"LogMetadata,omitempty" xml:"LogMetadata,omitempty"`
	// example:
	//
	// {\\"PrometheusConfigMap\\":{\\"New_Vpc_Ack_And_Jumpserver\\":{}}}
	OperationMetadata *string `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty"`
	// example:
	//
	// policyName1, policyName2
	PolicyNames *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Resellable      *bool   `json:"Resellable,omitempty" xml:"Resellable,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId   *string                            `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceInfo []*CreateServiceRequestServiceInfo `json:"ServiceInfo,omitempty" xml:"ServiceInfo,omitempty" type:"Repeated"`
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// example:
	//
	// Public
	ShareType            *string                    `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	SourceServiceId      *string                    `json:"SourceServiceId,omitempty" xml:"SourceServiceId,omitempty"`
	SourceServiceVersion *string                    `json:"SourceServiceVersion,omitempty" xml:"SourceServiceVersion,omitempty"`
	Tag                  []*CreateServiceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// example:
	//
	// SingleTenant
	TenantType *string `json:"TenantType,omitempty" xml:"TenantType,omitempty"`
	// example:
	//
	// 7
	TrialDuration *int64 `json:"TrialDuration,omitempty" xml:"TrialDuration,omitempty"`
	// example:
	//
	// {\\"Description\\":\\"xxx\\",\\"SupportRollback\\":true,\\"SupportUpgradeFromVersions\\":[],\\"UpgradeComponents\\":[\\"Configuration\\"]}
	UpgradeMetadata *string `json:"UpgradeMetadata,omitempty" xml:"UpgradeMetadata,omitempty"`
	// example:
	//
	// Draft
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) SetAlarmMetadata(v string) *CreateServiceRequest {
	s.AlarmMetadata = &v
	return s
}

func (s *CreateServiceRequest) SetApprovalType(v string) *CreateServiceRequest {
	s.ApprovalType = &v
	return s
}

func (s *CreateServiceRequest) SetBuildParameters(v string) *CreateServiceRequest {
	s.BuildParameters = &v
	return s
}

func (s *CreateServiceRequest) SetClientToken(v string) *CreateServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceRequest) SetDeployMetadata(v string) *CreateServiceRequest {
	s.DeployMetadata = &v
	return s
}

func (s *CreateServiceRequest) SetDeployType(v string) *CreateServiceRequest {
	s.DeployType = &v
	return s
}

func (s *CreateServiceRequest) SetDuration(v int64) *CreateServiceRequest {
	s.Duration = &v
	return s
}

func (s *CreateServiceRequest) SetIsSupportOperated(v bool) *CreateServiceRequest {
	s.IsSupportOperated = &v
	return s
}

func (s *CreateServiceRequest) SetLicenseMetadata(v string) *CreateServiceRequest {
	s.LicenseMetadata = &v
	return s
}

func (s *CreateServiceRequest) SetLogMetadata(v string) *CreateServiceRequest {
	s.LogMetadata = &v
	return s
}

func (s *CreateServiceRequest) SetOperationMetadata(v string) *CreateServiceRequest {
	s.OperationMetadata = &v
	return s
}

func (s *CreateServiceRequest) SetPolicyNames(v string) *CreateServiceRequest {
	s.PolicyNames = &v
	return s
}

func (s *CreateServiceRequest) SetRegionId(v string) *CreateServiceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceRequest) SetResellable(v bool) *CreateServiceRequest {
	s.Resellable = &v
	return s
}

func (s *CreateServiceRequest) SetResourceGroupId(v string) *CreateServiceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServiceRequest) SetServiceId(v string) *CreateServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceRequest) SetServiceInfo(v []*CreateServiceRequestServiceInfo) *CreateServiceRequest {
	s.ServiceInfo = v
	return s
}

func (s *CreateServiceRequest) SetServiceType(v string) *CreateServiceRequest {
	s.ServiceType = &v
	return s
}

func (s *CreateServiceRequest) SetShareType(v string) *CreateServiceRequest {
	s.ShareType = &v
	return s
}

func (s *CreateServiceRequest) SetSourceServiceId(v string) *CreateServiceRequest {
	s.SourceServiceId = &v
	return s
}

func (s *CreateServiceRequest) SetSourceServiceVersion(v string) *CreateServiceRequest {
	s.SourceServiceVersion = &v
	return s
}

func (s *CreateServiceRequest) SetTag(v []*CreateServiceRequestTag) *CreateServiceRequest {
	s.Tag = v
	return s
}

func (s *CreateServiceRequest) SetTenantType(v string) *CreateServiceRequest {
	s.TenantType = &v
	return s
}

func (s *CreateServiceRequest) SetTrialDuration(v int64) *CreateServiceRequest {
	s.TrialDuration = &v
	return s
}

func (s *CreateServiceRequest) SetUpgradeMetadata(v string) *CreateServiceRequest {
	s.UpgradeMetadata = &v
	return s
}

func (s *CreateServiceRequest) SetVersionName(v string) *CreateServiceRequest {
	s.VersionName = &v
	return s
}

type CreateServiceRequestServiceInfo struct {
	Agreements []*CreateServiceRequestServiceInfoAgreements `json:"Agreements,omitempty" xml:"Agreements,omitempty" type:"Repeated"`
	// example:
	//
	// http://img.tidb.oss.url
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh-CN
	Locale             *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	LongDescriptionUrl *string `json:"LongDescriptionUrl,omitempty" xml:"LongDescriptionUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TiDB Database
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s CreateServiceRequestServiceInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequestServiceInfo) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestServiceInfo) SetAgreements(v []*CreateServiceRequestServiceInfoAgreements) *CreateServiceRequestServiceInfo {
	s.Agreements = v
	return s
}

func (s *CreateServiceRequestServiceInfo) SetImage(v string) *CreateServiceRequestServiceInfo {
	s.Image = &v
	return s
}

func (s *CreateServiceRequestServiceInfo) SetLocale(v string) *CreateServiceRequestServiceInfo {
	s.Locale = &v
	return s
}

func (s *CreateServiceRequestServiceInfo) SetLongDescriptionUrl(v string) *CreateServiceRequestServiceInfo {
	s.LongDescriptionUrl = &v
	return s
}

func (s *CreateServiceRequestServiceInfo) SetName(v string) *CreateServiceRequestServiceInfo {
	s.Name = &v
	return s
}

func (s *CreateServiceRequestServiceInfo) SetShortDescription(v string) *CreateServiceRequestServiceInfo {
	s.ShortDescription = &v
	return s
}

type CreateServiceRequestServiceInfoAgreements struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateServiceRequestServiceInfoAgreements) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequestServiceInfoAgreements) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestServiceInfoAgreements) SetName(v string) *CreateServiceRequestServiceInfoAgreements {
	s.Name = &v
	return s
}

func (s *CreateServiceRequestServiceInfoAgreements) SetUrl(v string) *CreateServiceRequestServiceInfoAgreements {
	s.Url = &v
	return s
}

type CreateServiceRequestTag struct {
	// example:
	//
	// Usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// Web
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestTag) SetKey(v string) *CreateServiceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceRequestTag) SetValue(v string) *CreateServiceRequestTag {
	s.Value = &v
	return s
}

type CreateServiceResponseBody struct {
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// draft
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBody) SetRequestId(v string) *CreateServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceResponseBody) SetServiceId(v string) *CreateServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceResponseBody) SetStatus(v string) *CreateServiceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateServiceResponseBody) SetVersion(v string) *CreateServiceResponseBody {
	s.Version = &v
	return s
}

type CreateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceResponse) SetHeaders(v map[string]*string) *CreateServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceResponse) SetStatusCode(v int32) *CreateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceResponse) SetBody(v *CreateServiceResponseBody) *CreateServiceResponse {
	s.Body = v
	return s
}

type CreateServiceInstanceRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {"NodeCount": 3, "SystemDiskSize": 40, "InstancePassword": "******"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aekzuqyxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// 1
	ServiceVersion    *string                            `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	SpecificationName *string                            `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	Tag               []*CreateServiceInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TemplateName      *string                            `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 1563457855xxxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequest) SetClientToken(v string) *CreateServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetDryRun(v bool) *CreateServiceInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetEndTime(v string) *CreateServiceInstanceRequest {
	s.EndTime = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetName(v string) *CreateServiceInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetParameters(v map[string]interface{}) *CreateServiceInstanceRequest {
	s.Parameters = v
	return s
}

func (s *CreateServiceInstanceRequest) SetRegionId(v string) *CreateServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetResourceGroupId(v string) *CreateServiceInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetServiceId(v string) *CreateServiceInstanceRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetServiceVersion(v string) *CreateServiceInstanceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetSpecificationName(v string) *CreateServiceInstanceRequest {
	s.SpecificationName = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetTag(v []*CreateServiceInstanceRequestTag) *CreateServiceInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateServiceInstanceRequest) SetTemplateName(v string) *CreateServiceInstanceRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetUserId(v string) *CreateServiceInstanceRequest {
	s.UserId = &v
	return s
}

type CreateServiceInstanceRequestTag struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceInstanceRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequestTag) SetKey(v string) *CreateServiceInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceInstanceRequestTag) SetValue(v string) *CreateServiceInstanceRequestTag {
	s.Value = &v
	return s
}

type CreateServiceInstanceShrinkRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {"NodeCount": 3, "SystemDiskSize": 40, "InstancePassword": "******"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aekzuqyxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// 1
	ServiceVersion    *string                                  `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	SpecificationName *string                                  `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	Tag               []*CreateServiceInstanceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TemplateName      *string                                  `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 1563457855xxxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateServiceInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequest) SetClientToken(v string) *CreateServiceInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetDryRun(v bool) *CreateServiceInstanceShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetEndTime(v string) *CreateServiceInstanceShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetName(v string) *CreateServiceInstanceShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetParametersShrink(v string) *CreateServiceInstanceShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetRegionId(v string) *CreateServiceInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetResourceGroupId(v string) *CreateServiceInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetServiceId(v string) *CreateServiceInstanceShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetServiceVersion(v string) *CreateServiceInstanceShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetSpecificationName(v string) *CreateServiceInstanceShrinkRequest {
	s.SpecificationName = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetTag(v []*CreateServiceInstanceShrinkRequestTag) *CreateServiceInstanceShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetTemplateName(v string) *CreateServiceInstanceShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetUserId(v string) *CreateServiceInstanceShrinkRequest {
	s.UserId = &v
	return s
}

type CreateServiceInstanceShrinkRequestTag struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceInstanceShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequestTag) SetKey(v string) *CreateServiceInstanceShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestTag) SetValue(v string) *CreateServiceInstanceShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateServiceInstanceResponseBody struct {
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceResponseBody) SetRequestId(v string) *CreateServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetServiceInstanceId(v string) *CreateServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetStatus(v string) *CreateServiceInstanceResponseBody {
	s.Status = &v
	return s
}

type CreateServiceInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceResponse) SetHeaders(v map[string]*string) *CreateServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceInstanceResponse) SetStatusCode(v int32) *CreateServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceInstanceResponse) SetBody(v *CreateServiceInstanceResponseBody) *CreateServiceInstanceResponse {
	s.Body = v
	return s
}

type DeleteArtifactRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// example:
	//
	// 2
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
}

func (s DeleteArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteArtifactRequest) GoString() string {
	return s.String()
}

func (s *DeleteArtifactRequest) SetArtifactId(v string) *DeleteArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *DeleteArtifactRequest) SetArtifactVersion(v string) *DeleteArtifactRequest {
	s.ArtifactVersion = &v
	return s
}

type DeleteArtifactResponseBody struct {
	// example:
	//
	// B1A0198B-F316-1B72-B8DD-28B6F6D6XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteArtifactResponseBody) SetRequestId(v string) *DeleteArtifactResponseBody {
	s.RequestId = &v
	return s
}

type DeleteArtifactResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteArtifactResponse) GoString() string {
	return s.String()
}

func (s *DeleteArtifactResponse) SetHeaders(v map[string]*string) *DeleteArtifactResponse {
	s.Headers = v
	return s
}

func (s *DeleteArtifactResponse) SetStatusCode(v int32) *DeleteArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteArtifactResponse) SetBody(v *DeleteArtifactResponseBody) *DeleteArtifactResponse {
	s.Body = v
	return s
}

type DeleteServiceRequest struct {
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s DeleteServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceRequest) SetClientToken(v string) *DeleteServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServiceRequest) SetRegionId(v string) *DeleteServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServiceRequest) SetServiceId(v string) *DeleteServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *DeleteServiceRequest) SetServiceVersion(v string) *DeleteServiceRequest {
	s.ServiceVersion = &v
	return s
}

type DeleteServiceResponseBody struct {
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponseBody) SetRequestId(v string) *DeleteServiceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponse) SetHeaders(v map[string]*string) *DeleteServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceResponse) SetStatusCode(v int32) *DeleteServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceResponse) SetBody(v *DeleteServiceResponseBody) *DeleteServiceResponse {
	s.Body = v
	return s
}

type DeleteServiceInstancesRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ServiceInstanceId []*string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty" type:"Repeated"`
}

func (s DeleteServiceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesRequest) SetClientToken(v string) *DeleteServiceInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServiceInstancesRequest) SetRegionId(v string) *DeleteServiceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServiceInstancesRequest) SetServiceInstanceId(v []*string) *DeleteServiceInstancesRequest {
	s.ServiceInstanceId = v
	return s
}

type DeleteServiceInstancesResponseBody struct {
	// example:
	//
	// DB140E67-D75F-5585-946E-41D8DC8F4E00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesResponseBody) SetRequestId(v string) *DeleteServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesResponse) SetHeaders(v map[string]*string) *DeleteServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceInstancesResponse) SetStatusCode(v int32) *DeleteServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceInstancesResponse) SetBody(v *DeleteServiceInstancesResponseBody) *DeleteServiceInstancesResponse {
	s.Body = v
	return s
}

type DeployServiceInstanceRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// si-0e6fca6a51a54420****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s DeployServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeployServiceInstanceRequest) SetClientToken(v string) *DeployServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeployServiceInstanceRequest) SetRegionId(v string) *DeployServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeployServiceInstanceRequest) SetServiceInstanceId(v string) *DeployServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type DeployServiceInstanceResponseBody struct {
	// example:
	//
	// B8A6AEA6-0D8F-589A-A7FF-B44FD842456E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeployServiceInstanceResponseBody) SetRequestId(v string) *DeployServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeployServiceInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeployServiceInstanceResponse) SetHeaders(v map[string]*string) *DeployServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeployServiceInstanceResponse) SetStatusCode(v int32) *DeployServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployServiceInstanceResponse) SetBody(v *DeployServiceInstanceResponseBody) *DeployServiceInstanceResponse {
	s.Body = v
	return s
}

type GetArtifactRequest struct {
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// example:
	//
	// gpu-test
	ArtifactName *string `json:"ArtifactName,omitempty" xml:"ArtifactName,omitempty"`
	// example:
	//
	// 1
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
}

func (s GetArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactRequest) SetArtifactId(v string) *GetArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *GetArtifactRequest) SetArtifactName(v string) *GetArtifactRequest {
	s.ArtifactName = &v
	return s
}

func (s *GetArtifactRequest) SetArtifactVersion(v string) *GetArtifactRequest {
	s.ArtifactVersion = &v
	return s
}

type GetArtifactResponseBody struct {
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// example:
	//
	// {\\"CommodityCode\\":\\"cmjj0005xxxx\\",\\"CommodityVersion\\":\\"V2022xxxx\\"}
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// example:
	//
	// 2
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2022-10-20T02:19:55Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 2
	MaxVersion *int64  `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// B1A0198B-F316-1B72-B8DD-28B6F6D6XXXX
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ["cn-hangzhou","cn-beijing"]
	SupportRegionIds *string                        `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty"`
	Tags             []*GetArtifactResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactResponseBody) SetArtifactId(v string) *GetArtifactResponseBody {
	s.ArtifactId = &v
	return s
}

func (s *GetArtifactResponseBody) SetArtifactProperty(v string) *GetArtifactResponseBody {
	s.ArtifactProperty = &v
	return s
}

func (s *GetArtifactResponseBody) SetArtifactType(v string) *GetArtifactResponseBody {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactResponseBody) SetArtifactVersion(v string) *GetArtifactResponseBody {
	s.ArtifactVersion = &v
	return s
}

func (s *GetArtifactResponseBody) SetDescription(v string) *GetArtifactResponseBody {
	s.Description = &v
	return s
}

func (s *GetArtifactResponseBody) SetGmtModified(v string) *GetArtifactResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetArtifactResponseBody) SetMaxVersion(v int64) *GetArtifactResponseBody {
	s.MaxVersion = &v
	return s
}

func (s *GetArtifactResponseBody) SetName(v string) *GetArtifactResponseBody {
	s.Name = &v
	return s
}

func (s *GetArtifactResponseBody) SetProgress(v string) *GetArtifactResponseBody {
	s.Progress = &v
	return s
}

func (s *GetArtifactResponseBody) SetRequestId(v string) *GetArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArtifactResponseBody) SetResourceGroupId(v string) *GetArtifactResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetArtifactResponseBody) SetStatus(v string) *GetArtifactResponseBody {
	s.Status = &v
	return s
}

func (s *GetArtifactResponseBody) SetSupportRegionIds(v string) *GetArtifactResponseBody {
	s.SupportRegionIds = &v
	return s
}

func (s *GetArtifactResponseBody) SetTags(v []*GetArtifactResponseBodyTags) *GetArtifactResponseBody {
	s.Tags = v
	return s
}

func (s *GetArtifactResponseBody) SetVersionName(v string) *GetArtifactResponseBody {
	s.VersionName = &v
	return s
}

type GetArtifactResponseBodyTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetArtifactResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetArtifactResponseBodyTags) SetKey(v string) *GetArtifactResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetArtifactResponseBodyTags) SetValue(v string) *GetArtifactResponseBodyTags {
	s.Value = &v
	return s
}

type GetArtifactResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactResponse) SetHeaders(v map[string]*string) *GetArtifactResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactResponse) SetStatusCode(v int32) *GetArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactResponse) SetBody(v *GetArtifactResponseBody) *GetArtifactResponse {
	s.Body = v
	return s
}

type GetArtifactRepositoryCredentialsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// File
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// example:
	//
	// cn-hangzhou
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
}

func (s GetArtifactRepositoryCredentialsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsRequest) SetArtifactType(v string) *GetArtifactRepositoryCredentialsRequest {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsRequest) SetDeployRegionId(v string) *GetArtifactRepositoryCredentialsRequest {
	s.DeployRegionId = &v
	return s
}

type GetArtifactRepositoryCredentialsResponseBody struct {
	AvailableResources []*GetArtifactRepositoryCredentialsResponseBodyAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Repeated"`
	Credentials        *GetArtifactRepositoryCredentialsResponseBodyCredentials          `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Struct"`
	// example:
	//
	// 1526549792000
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// example:
	//
	// 94E89857-B994-44B6-9C4F-DBD200E9XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetArtifactRepositoryCredentialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsResponseBody) SetAvailableResources(v []*GetArtifactRepositoryCredentialsResponseBodyAvailableResources) *GetArtifactRepositoryCredentialsResponseBody {
	s.AvailableResources = v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBody) SetCredentials(v *GetArtifactRepositoryCredentialsResponseBodyCredentials) *GetArtifactRepositoryCredentialsResponseBody {
	s.Credentials = v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBody) SetExpireDate(v string) *GetArtifactRepositoryCredentialsResponseBody {
	s.ExpireDate = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBody) SetRequestId(v string) *GetArtifactRepositoryCredentialsResponseBody {
	s.RequestId = &v
	return s
}

type GetArtifactRepositoryCredentialsResponseBodyAvailableResources struct {
	// example:
	//
	// "/xxx/"
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// oss-cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// computenest-artifacts-draft-cn-hangzhou
	RepositoryName *string `json:"RepositoryName,omitempty" xml:"RepositoryName,omitempty"`
}

func (s GetArtifactRepositoryCredentialsResponseBodyAvailableResources) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsResponseBodyAvailableResources) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsResponseBodyAvailableResources) SetPath(v string) *GetArtifactRepositoryCredentialsResponseBodyAvailableResources {
	s.Path = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyAvailableResources) SetRegionId(v string) *GetArtifactRepositoryCredentialsResponseBodyAvailableResources {
	s.RegionId = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyAvailableResources) SetRepositoryName(v string) *GetArtifactRepositoryCredentialsResponseBodyAvailableResources {
	s.RepositoryName = &v
	return s
}

type GetArtifactRepositoryCredentialsResponseBodyCredentials struct {
	// example:
	//
	// STS.xxx
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// example:
	//
	// xxx
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// example:
	//
	// eyJ0aW1lIjoiMTUyNjU0OTc5:0705733****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// xxx
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// xxx
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetArtifactRepositoryCredentialsResponseBodyCredentials) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsResponseBodyCredentials) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetAccessKeyId(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.AccessKeyId = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetAccessKeySecret(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetPassword(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.Password = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetSecurityToken(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.SecurityToken = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetUsername(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.Username = &v
	return s
}

type GetArtifactRepositoryCredentialsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactRepositoryCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactRepositoryCredentialsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsResponse) SetHeaders(v map[string]*string) *GetArtifactRepositoryCredentialsResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponse) SetStatusCode(v int32) *GetArtifactRepositoryCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponse) SetBody(v *GetArtifactRepositoryCredentialsResponseBody) *GetArtifactRepositoryCredentialsResponse {
	s.Body = v
	return s
}

type GetServiceRequest struct {
	FilterAliUid *bool `json:"FilterAliUid,omitempty" xml:"FilterAliUid,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-4ee86df83fd948******
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// 1
	ServiceVersion    *string   `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	SharedAccountType *string   `json:"SharedAccountType,omitempty" xml:"SharedAccountType,omitempty"`
	ShowDetail        []*string `json:"ShowDetail,omitempty" xml:"ShowDetail,omitempty" type:"Repeated"`
}

func (s GetServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceRequest) SetFilterAliUid(v bool) *GetServiceRequest {
	s.FilterAliUid = &v
	return s
}

func (s *GetServiceRequest) SetRegionId(v string) *GetServiceRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceRequest) SetServiceId(v string) *GetServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceRequest) SetServiceVersion(v string) *GetServiceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceRequest) SetSharedAccountType(v string) *GetServiceRequest {
	s.SharedAccountType = &v
	return s
}

func (s *GetServiceRequest) SetShowDetail(v []*string) *GetServiceRequest {
	s.ShowDetail = v
	return s
}

type GetServiceResponseBody struct {
	AlarmMetadata *string `json:"AlarmMetadata,omitempty" xml:"AlarmMetadata,omitempty"`
	// example:
	//
	// Manual
	ApprovalType *string                          `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	BuildInfo    *string                          `json:"BuildInfo,omitempty" xml:"BuildInfo,omitempty"`
	Categories   *string                          `json:"Categories,omitempty" xml:"Categories,omitempty"`
	Commodity    *GetServiceResponseBodyCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// example:
	//
	// cmjj00xxxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime                  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CrossRegionConnectionStatus *string `json:"CrossRegionConnectionStatus,omitempty" xml:"CrossRegionConnectionStatus,omitempty"`
	// example:
	//
	// 1
	DefaultLicenseDays *int64  `json:"DefaultLicenseDays,omitempty" xml:"DefaultLicenseDays,omitempty"`
	DeployMetadata     *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// example:
	//
	// 259200
	Duration     *int64             `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EntitySource map[string]*string `json:"EntitySource,omitempty" xml:"EntitySource,omitempty"`
	// example:
	//
	// false
	IsSupportOperated *bool   `json:"IsSupportOperated,omitempty" xml:"IsSupportOperated,omitempty"`
	LicenseMetadata   *string `json:"LicenseMetadata,omitempty" xml:"LicenseMetadata,omitempty"`
	LogMetadata       *string `json:"LogMetadata,omitempty" xml:"LogMetadata,omitempty"`
	// example:
	//
	// {\\"PrometheusConfigMap\\":{\\"New_Vpc_Ack_And_Jumpserver\\":{}}}
	OperationMetadata *string `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty"`
	// example:
	//
	// None
	PayFromType *string `json:"PayFromType,omitempty" xml:"PayFromType,omitempty"`
	// example:
	//
	// Permanent
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// Deployable
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// example:
	//
	// policyName1, policyName2
	PolicyNames *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
	// example:
	//
	// 90
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 2021-05-21T00:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// example:
	//
	// sr-04056c2ab4b94bxxxxxx
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// example:
	//
	// B1A0198B-F316-1B72-B8DD-28B6F6D6XXXX
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resellable *bool   `json:"Resellable,omitempty" xml:"Resellable,omitempty"`
	// example:
	//
	// rg-aekzuqyxxxxxx
	ResourceGroupId         *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServiceAuditDocumentUrl *string `json:"ServiceAuditDocumentUrl,omitempty" xml:"ServiceAuditDocumentUrl,omitempty"`
	ServiceDiscoverable     *string `json:"ServiceDiscoverable,omitempty" xml:"ServiceDiscoverable,omitempty"`
	// example:
	//
	// http://example1.com
	ServiceDocUrl *string `json:"ServiceDocUrl,omitempty" xml:"ServiceDocUrl,omitempty"`
	// example:
	//
	// service-70a3b15bb62643xxxxxx
	ServiceId    *string                               `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceInfos []*GetServiceResponseBodyServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// http://example2.com
	ServiceProductUrl *string `json:"ServiceProductUrl,omitempty" xml:"ServiceProductUrl,omitempty"`
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// example:
	//
	// Public
	ShareType            *string                          `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	ShareTypeStatus      *string                          `json:"ShareTypeStatus,omitempty" xml:"ShareTypeStatus,omitempty"`
	SourceServiceId      *string                          `json:"SourceServiceId,omitempty" xml:"SourceServiceId,omitempty"`
	SourceServiceVersion *string                          `json:"SourceServiceVersion,omitempty" xml:"SourceServiceVersion,omitempty"`
	SourceSupplierName   *string                          `json:"SourceSupplierName,omitempty" xml:"SourceSupplierName,omitempty"`
	Statistic            *GetServiceResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Struct"`
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// deploy successfully
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// example:
	//
	// http://example.com
	SupplierUrl *string                       `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	Tags        []*GetServiceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// SingleTenant
	TenantType *string `json:"TenantType,omitempty" xml:"TenantType,omitempty"`
	// example:
	//
	// SERVICE_TEST_SUCCEED
	TestStatus *string `json:"TestStatus,omitempty" xml:"TestStatus,omitempty"`
	// example:
	//
	// 7
	TrialDuration *int64 `json:"TrialDuration,omitempty" xml:"TrialDuration,omitempty"`
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
	// example:
	//
	// 2021-05-22T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// {\\"Description\\":\\"xxx\\",\\"SupportRollback\\":true,\\"SupportUpgradeFromVersions\\":[],\\"UpgradeComponents\\":[\\"Configuration\\"]}
	UpgradeMetadata *string `json:"UpgradeMetadata,omitempty" xml:"UpgradeMetadata,omitempty"`
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// v1
	VersionName              *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	VirtualInternetService   *string `json:"VirtualInternetService,omitempty" xml:"VirtualInternetService,omitempty"`
	VirtualInternetServiceId *string `json:"VirtualInternetServiceId,omitempty" xml:"VirtualInternetServiceId,omitempty"`
}

func (s GetServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBody) SetAlarmMetadata(v string) *GetServiceResponseBody {
	s.AlarmMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetApprovalType(v string) *GetServiceResponseBody {
	s.ApprovalType = &v
	return s
}

func (s *GetServiceResponseBody) SetBuildInfo(v string) *GetServiceResponseBody {
	s.BuildInfo = &v
	return s
}

func (s *GetServiceResponseBody) SetCategories(v string) *GetServiceResponseBody {
	s.Categories = &v
	return s
}

func (s *GetServiceResponseBody) SetCommodity(v *GetServiceResponseBodyCommodity) *GetServiceResponseBody {
	s.Commodity = v
	return s
}

func (s *GetServiceResponseBody) SetCommodityCode(v string) *GetServiceResponseBody {
	s.CommodityCode = &v
	return s
}

func (s *GetServiceResponseBody) SetCreateTime(v string) *GetServiceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetServiceResponseBody) SetCrossRegionConnectionStatus(v string) *GetServiceResponseBody {
	s.CrossRegionConnectionStatus = &v
	return s
}

func (s *GetServiceResponseBody) SetDefaultLicenseDays(v int64) *GetServiceResponseBody {
	s.DefaultLicenseDays = &v
	return s
}

func (s *GetServiceResponseBody) SetDeployMetadata(v string) *GetServiceResponseBody {
	s.DeployMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetDeployType(v string) *GetServiceResponseBody {
	s.DeployType = &v
	return s
}

func (s *GetServiceResponseBody) SetDuration(v int64) *GetServiceResponseBody {
	s.Duration = &v
	return s
}

func (s *GetServiceResponseBody) SetEntitySource(v map[string]*string) *GetServiceResponseBody {
	s.EntitySource = v
	return s
}

func (s *GetServiceResponseBody) SetIsSupportOperated(v bool) *GetServiceResponseBody {
	s.IsSupportOperated = &v
	return s
}

func (s *GetServiceResponseBody) SetLicenseMetadata(v string) *GetServiceResponseBody {
	s.LicenseMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetLogMetadata(v string) *GetServiceResponseBody {
	s.LogMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetOperationMetadata(v string) *GetServiceResponseBody {
	s.OperationMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetPayFromType(v string) *GetServiceResponseBody {
	s.PayFromType = &v
	return s
}

func (s *GetServiceResponseBody) SetPayType(v string) *GetServiceResponseBody {
	s.PayType = &v
	return s
}

func (s *GetServiceResponseBody) SetPermission(v string) *GetServiceResponseBody {
	s.Permission = &v
	return s
}

func (s *GetServiceResponseBody) SetPolicyNames(v string) *GetServiceResponseBody {
	s.PolicyNames = &v
	return s
}

func (s *GetServiceResponseBody) SetProgress(v int64) *GetServiceResponseBody {
	s.Progress = &v
	return s
}

func (s *GetServiceResponseBody) SetPublishTime(v string) *GetServiceResponseBody {
	s.PublishTime = &v
	return s
}

func (s *GetServiceResponseBody) SetRegistrationId(v string) *GetServiceResponseBody {
	s.RegistrationId = &v
	return s
}

func (s *GetServiceResponseBody) SetRequestId(v string) *GetServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceResponseBody) SetResellable(v bool) *GetServiceResponseBody {
	s.Resellable = &v
	return s
}

func (s *GetServiceResponseBody) SetResourceGroupId(v string) *GetServiceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceAuditDocumentUrl(v string) *GetServiceResponseBody {
	s.ServiceAuditDocumentUrl = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceDiscoverable(v string) *GetServiceResponseBody {
	s.ServiceDiscoverable = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceDocUrl(v string) *GetServiceResponseBody {
	s.ServiceDocUrl = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceId(v string) *GetServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceInfos(v []*GetServiceResponseBodyServiceInfos) *GetServiceResponseBody {
	s.ServiceInfos = v
	return s
}

func (s *GetServiceResponseBody) SetServiceProductUrl(v string) *GetServiceResponseBody {
	s.ServiceProductUrl = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceType(v string) *GetServiceResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetServiceResponseBody) SetShareType(v string) *GetServiceResponseBody {
	s.ShareType = &v
	return s
}

func (s *GetServiceResponseBody) SetShareTypeStatus(v string) *GetServiceResponseBody {
	s.ShareTypeStatus = &v
	return s
}

func (s *GetServiceResponseBody) SetSourceServiceId(v string) *GetServiceResponseBody {
	s.SourceServiceId = &v
	return s
}

func (s *GetServiceResponseBody) SetSourceServiceVersion(v string) *GetServiceResponseBody {
	s.SourceServiceVersion = &v
	return s
}

func (s *GetServiceResponseBody) SetSourceSupplierName(v string) *GetServiceResponseBody {
	s.SourceSupplierName = &v
	return s
}

func (s *GetServiceResponseBody) SetStatistic(v *GetServiceResponseBodyStatistic) *GetServiceResponseBody {
	s.Statistic = v
	return s
}

func (s *GetServiceResponseBody) SetStatus(v string) *GetServiceResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceResponseBody) SetStatusDetail(v string) *GetServiceResponseBody {
	s.StatusDetail = &v
	return s
}

func (s *GetServiceResponseBody) SetSupplierName(v string) *GetServiceResponseBody {
	s.SupplierName = &v
	return s
}

func (s *GetServiceResponseBody) SetSupplierUrl(v string) *GetServiceResponseBody {
	s.SupplierUrl = &v
	return s
}

func (s *GetServiceResponseBody) SetTags(v []*GetServiceResponseBodyTags) *GetServiceResponseBody {
	s.Tags = v
	return s
}

func (s *GetServiceResponseBody) SetTenantType(v string) *GetServiceResponseBody {
	s.TenantType = &v
	return s
}

func (s *GetServiceResponseBody) SetTestStatus(v string) *GetServiceResponseBody {
	s.TestStatus = &v
	return s
}

func (s *GetServiceResponseBody) SetTrialDuration(v int64) *GetServiceResponseBody {
	s.TrialDuration = &v
	return s
}

func (s *GetServiceResponseBody) SetTrialType(v string) *GetServiceResponseBody {
	s.TrialType = &v
	return s
}

func (s *GetServiceResponseBody) SetUpdateTime(v string) *GetServiceResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetServiceResponseBody) SetUpgradeMetadata(v string) *GetServiceResponseBody {
	s.UpgradeMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetVersion(v string) *GetServiceResponseBody {
	s.Version = &v
	return s
}

func (s *GetServiceResponseBody) SetVersionName(v string) *GetServiceResponseBody {
	s.VersionName = &v
	return s
}

func (s *GetServiceResponseBody) SetVirtualInternetService(v string) *GetServiceResponseBody {
	s.VirtualInternetService = &v
	return s
}

func (s *GetServiceResponseBody) SetVirtualInternetServiceId(v string) *GetServiceResponseBody {
	s.VirtualInternetServiceId = &v
	return s
}

type GetServiceResponseBodyCommodity struct {
	ChargeType          *string                                             `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CommodityCode       *string                                             `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Components          []*string                                           `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	CssMetadata         *GetServiceResponseBodyCommodityCssMetadata         `json:"CssMetadata,omitempty" xml:"CssMetadata,omitempty" type:"Struct"`
	MarketplaceMetadata *GetServiceResponseBodyCommodityMarketplaceMetadata `json:"MarketplaceMetadata,omitempty" xml:"MarketplaceMetadata,omitempty" type:"Struct"`
	MeteringEntities    []*GetServiceResponseBodyCommodityMeteringEntities  `json:"MeteringEntities,omitempty" xml:"MeteringEntities,omitempty" type:"Repeated"`
	SaasBoostMetadata   *string                                             `json:"SaasBoostMetadata,omitempty" xml:"SaasBoostMetadata,omitempty"`
	Specifications      []*GetServiceResponseBodyCommoditySpecifications    `json:"Specifications,omitempty" xml:"Specifications,omitempty" type:"Repeated"`
	Type                *string                                             `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceResponseBodyCommodity) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodity) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodity) SetChargeType(v string) *GetServiceResponseBodyCommodity {
	s.ChargeType = &v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetCommodityCode(v string) *GetServiceResponseBodyCommodity {
	s.CommodityCode = &v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetComponents(v []*string) *GetServiceResponseBodyCommodity {
	s.Components = v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetCssMetadata(v *GetServiceResponseBodyCommodityCssMetadata) *GetServiceResponseBodyCommodity {
	s.CssMetadata = v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetMarketplaceMetadata(v *GetServiceResponseBodyCommodityMarketplaceMetadata) *GetServiceResponseBodyCommodity {
	s.MarketplaceMetadata = v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetMeteringEntities(v []*GetServiceResponseBodyCommodityMeteringEntities) *GetServiceResponseBodyCommodity {
	s.MeteringEntities = v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetSaasBoostMetadata(v string) *GetServiceResponseBodyCommodity {
	s.SaasBoostMetadata = &v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetSpecifications(v []*GetServiceResponseBodyCommoditySpecifications) *GetServiceResponseBodyCommodity {
	s.Specifications = v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetType(v string) *GetServiceResponseBodyCommodity {
	s.Type = &v
	return s
}

type GetServiceResponseBodyCommodityCssMetadata struct {
	ComponentsMappings       []*GetServiceResponseBodyCommodityCssMetadataComponentsMappings       `json:"ComponentsMappings,omitempty" xml:"ComponentsMappings,omitempty" type:"Repeated"`
	MeteringEntityExtraInfos []*GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos `json:"MeteringEntityExtraInfos,omitempty" xml:"MeteringEntityExtraInfos,omitempty" type:"Repeated"`
	MeteringEntityMappings   []*GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings   `json:"MeteringEntityMappings,omitempty" xml:"MeteringEntityMappings,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyCommodityCssMetadata) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityCssMetadata) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityCssMetadata) SetComponentsMappings(v []*GetServiceResponseBodyCommodityCssMetadataComponentsMappings) *GetServiceResponseBodyCommodityCssMetadata {
	s.ComponentsMappings = v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadata) SetMeteringEntityExtraInfos(v []*GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) *GetServiceResponseBodyCommodityCssMetadata {
	s.MeteringEntityExtraInfos = v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadata) SetMeteringEntityMappings(v []*GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) *GetServiceResponseBodyCommodityCssMetadata {
	s.MeteringEntityMappings = v
	return s
}

type GetServiceResponseBodyCommodityCssMetadataComponentsMappings struct {
	Mappings     map[string]*string `json:"Mappings,omitempty" xml:"Mappings,omitempty"`
	TemplateName *string            `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceResponseBodyCommodityCssMetadataComponentsMappings) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityCssMetadataComponentsMappings) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityCssMetadataComponentsMappings) SetMappings(v map[string]*string) *GetServiceResponseBodyCommodityCssMetadataComponentsMappings {
	s.Mappings = v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadataComponentsMappings) SetTemplateName(v string) *GetServiceResponseBodyCommodityCssMetadataComponentsMappings {
	s.TemplateName = &v
	return s
}

type GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos struct {
	EntityId   *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Promql     *string `json:"Promql,omitempty" xml:"Promql,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) SetEntityId(v string) *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos {
	s.EntityId = &v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) SetMetricName(v string) *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos {
	s.MetricName = &v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) SetPromql(v string) *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos {
	s.Promql = &v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) SetType(v string) *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos {
	s.Type = &v
	return s
}

type GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings struct {
	EntityIds         *string `json:"EntityIds,omitempty" xml:"EntityIds,omitempty"`
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	TemplateName      *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) SetEntityIds(v string) *GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings {
	s.EntityIds = &v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) SetSpecificationName(v string) *GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) SetTemplateName(v string) *GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings {
	s.TemplateName = &v
	return s
}

type GetServiceResponseBodyCommodityMarketplaceMetadata struct {
	MeteringEntityExtraInfos []*GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos `json:"MeteringEntityExtraInfos,omitempty" xml:"MeteringEntityExtraInfos,omitempty" type:"Repeated"`
	MeteringEntityMappings   []*GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings   `json:"MeteringEntityMappings,omitempty" xml:"MeteringEntityMappings,omitempty" type:"Repeated"`
	SpecificationMappings    []*GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings    `json:"SpecificationMappings,omitempty" xml:"SpecificationMappings,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadata) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadata) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadata) SetMeteringEntityExtraInfos(v []*GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) *GetServiceResponseBodyCommodityMarketplaceMetadata {
	s.MeteringEntityExtraInfos = v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadata) SetMeteringEntityMappings(v []*GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) *GetServiceResponseBodyCommodityMarketplaceMetadata {
	s.MeteringEntityMappings = v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadata) SetSpecificationMappings(v []*GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) *GetServiceResponseBodyCommodityMarketplaceMetadata {
	s.SpecificationMappings = v
	return s
}

type GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos struct {
	EntityId   *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Promql     *string `json:"Promql,omitempty" xml:"Promql,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) SetEntityId(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos {
	s.EntityId = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) SetMetricName(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos {
	s.MetricName = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) SetPromql(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos {
	s.Promql = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) SetType(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos {
	s.Type = &v
	return s
}

type GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings struct {
	EntityIds         *string `json:"EntityIds,omitempty" xml:"EntityIds,omitempty"`
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	TemplateName      *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) SetEntityIds(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings {
	s.EntityIds = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) SetSpecificationName(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) SetTemplateName(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings {
	s.TemplateName = &v
	return s
}

type GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings struct {
	SpecificationCode *string `json:"SpecificationCode,omitempty" xml:"SpecificationCode,omitempty"`
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	TemplateName      *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TrialType         *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) SetSpecificationCode(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings {
	s.SpecificationCode = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) SetSpecificationName(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) SetTemplateName(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings {
	s.TemplateName = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) SetTrialType(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings {
	s.TrialType = &v
	return s
}

type GetServiceResponseBodyCommodityMeteringEntities struct {
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetServiceResponseBodyCommodityMeteringEntities) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMeteringEntities) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMeteringEntities) SetEntityId(v string) *GetServiceResponseBodyCommodityMeteringEntities {
	s.EntityId = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMeteringEntities) SetName(v string) *GetServiceResponseBodyCommodityMeteringEntities {
	s.Name = &v
	return s
}

type GetServiceResponseBodyCommoditySpecifications struct {
	Code  *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Times []*string `json:"Times,omitempty" xml:"Times,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyCommoditySpecifications) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommoditySpecifications) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommoditySpecifications) SetCode(v string) *GetServiceResponseBodyCommoditySpecifications {
	s.Code = &v
	return s
}

func (s *GetServiceResponseBodyCommoditySpecifications) SetName(v string) *GetServiceResponseBodyCommoditySpecifications {
	s.Name = &v
	return s
}

func (s *GetServiceResponseBodyCommoditySpecifications) SetTimes(v []*string) *GetServiceResponseBodyCommoditySpecifications {
	s.Times = v
	return s
}

type GetServiceResponseBodyServiceInfos struct {
	Agreements []*GetServiceResponseBodyServiceInfosAgreements `json:"Agreements,omitempty" xml:"Agreements,omitempty" type:"Repeated"`
	// example:
	//
	// https://example.com/service-image/c1c4a559-cc60-4af1-b976-98f356602462.png
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// example:
	//
	// zh-CN
	Locale             *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	LongDescriptionUrl *string `json:"LongDescriptionUrl,omitempty" xml:"LongDescriptionUrl,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription   *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s GetServiceResponseBodyServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyServiceInfos) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyServiceInfos) SetAgreements(v []*GetServiceResponseBodyServiceInfosAgreements) *GetServiceResponseBodyServiceInfos {
	s.Agreements = v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetImage(v string) *GetServiceResponseBodyServiceInfos {
	s.Image = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetLocale(v string) *GetServiceResponseBodyServiceInfos {
	s.Locale = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetLongDescriptionUrl(v string) *GetServiceResponseBodyServiceInfos {
	s.LongDescriptionUrl = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetName(v string) *GetServiceResponseBodyServiceInfos {
	s.Name = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetShortDescription(v string) *GetServiceResponseBodyServiceInfos {
	s.ShortDescription = &v
	return s
}

type GetServiceResponseBodyServiceInfosAgreements struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetServiceResponseBodyServiceInfosAgreements) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyServiceInfosAgreements) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyServiceInfosAgreements) SetName(v string) *GetServiceResponseBodyServiceInfosAgreements {
	s.Name = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfosAgreements) SetUrl(v string) *GetServiceResponseBodyServiceInfosAgreements {
	s.Url = &v
	return s
}

type GetServiceResponseBodyStatistic struct {
	AccumulativeInstanceCount    *int32   `json:"AccumulativeInstanceCount,omitempty" xml:"AccumulativeInstanceCount,omitempty"`
	AccumulativePocAmount        *float64 `json:"AccumulativePocAmount,omitempty" xml:"AccumulativePocAmount,omitempty"`
	AccumulativeUserCount        *int32   `json:"AccumulativeUserCount,omitempty" xml:"AccumulativeUserCount,omitempty"`
	AveragePocAmount             *float64 `json:"AveragePocAmount,omitempty" xml:"AveragePocAmount,omitempty"`
	AveragePocDuration           *float64 `json:"AveragePocDuration,omitempty" xml:"AveragePocDuration,omitempty"`
	AveragePocUnitAmount         *float64 `json:"AveragePocUnitAmount,omitempty" xml:"AveragePocUnitAmount,omitempty"`
	DeployedServiceInstanceCount *int32   `json:"DeployedServiceInstanceCount,omitempty" xml:"DeployedServiceInstanceCount,omitempty"`
	DeployedUserCount            *int32   `json:"DeployedUserCount,omitempty" xml:"DeployedUserCount,omitempty"`
	SubmittedUsageCount          *int32   `json:"SubmittedUsageCount,omitempty" xml:"SubmittedUsageCount,omitempty"`
}

func (s GetServiceResponseBodyStatistic) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyStatistic) SetAccumulativeInstanceCount(v int32) *GetServiceResponseBodyStatistic {
	s.AccumulativeInstanceCount = &v
	return s
}

func (s *GetServiceResponseBodyStatistic) SetAccumulativePocAmount(v float64) *GetServiceResponseBodyStatistic {
	s.AccumulativePocAmount = &v
	return s
}

func (s *GetServiceResponseBodyStatistic) SetAccumulativeUserCount(v int32) *GetServiceResponseBodyStatistic {
	s.AccumulativeUserCount = &v
	return s
}

func (s *GetServiceResponseBodyStatistic) SetAveragePocAmount(v float64) *GetServiceResponseBodyStatistic {
	s.AveragePocAmount = &v
	return s
}

func (s *GetServiceResponseBodyStatistic) SetAveragePocDuration(v float64) *GetServiceResponseBodyStatistic {
	s.AveragePocDuration = &v
	return s
}

func (s *GetServiceResponseBodyStatistic) SetAveragePocUnitAmount(v float64) *GetServiceResponseBodyStatistic {
	s.AveragePocUnitAmount = &v
	return s
}

func (s *GetServiceResponseBodyStatistic) SetDeployedServiceInstanceCount(v int32) *GetServiceResponseBodyStatistic {
	s.DeployedServiceInstanceCount = &v
	return s
}

func (s *GetServiceResponseBodyStatistic) SetDeployedUserCount(v int32) *GetServiceResponseBodyStatistic {
	s.DeployedUserCount = &v
	return s
}

func (s *GetServiceResponseBodyStatistic) SetSubmittedUsageCount(v int32) *GetServiceResponseBodyStatistic {
	s.SubmittedUsageCount = &v
	return s
}

type GetServiceResponseBodyTags struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetServiceResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyTags) SetKey(v string) *GetServiceResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetServiceResponseBodyTags) SetValue(v string) *GetServiceResponseBodyTags {
	s.Value = &v
	return s
}

type GetServiceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceResponse) SetHeaders(v map[string]*string) *GetServiceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceResponse) SetStatusCode(v int32) *GetServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceResponse) SetBody(v *GetServiceResponseBody) *GetServiceResponse {
	s.Body = v
	return s
}

type GetServiceEstimateCostRequest struct {
	// example:
	//
	// mRdxWuW2ts
	ClientToken *string                                 `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Commodity   *GetServiceEstimateCostRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// example:
	//
	// {\\"PayType\\":\\"PostPaid\\",\\"InstancePassword\\":\\"xxxxxxxxxx\\",\\"EcsInstanceType\\":\\"ecs.g6.large\\",\\"VSwitchId\\":\\"vsw-0jlueyydpuekoxxxxxxxx\\",\\"VpcId\\":\\"vpc-0jlps6mjbgvpqxxxxxxxx\\",\\"ZoneId\\":\\"cn-wulanchabu-a\\",\\"Enable\\":false,\\"RegionId\\":\\"cn-wulanchabu\\"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-16fbd358d75e49xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// example:
	//
	// draft
	ServiceVersion    *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// example:
	//
	// Custom_Image_Ecs
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceEstimateCostRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceEstimateCostRequest) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostRequest) SetClientToken(v string) *GetServiceEstimateCostRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetCommodity(v *GetServiceEstimateCostRequestCommodity) *GetServiceEstimateCostRequest {
	s.Commodity = v
	return s
}

func (s *GetServiceEstimateCostRequest) SetParameters(v map[string]interface{}) *GetServiceEstimateCostRequest {
	s.Parameters = v
	return s
}

func (s *GetServiceEstimateCostRequest) SetRegionId(v string) *GetServiceEstimateCostRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetServiceId(v string) *GetServiceEstimateCostRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetServiceInstanceId(v string) *GetServiceEstimateCostRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetServiceVersion(v string) *GetServiceEstimateCostRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetSpecificationName(v string) *GetServiceEstimateCostRequest {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetTemplateName(v string) *GetServiceEstimateCostRequest {
	s.TemplateName = &v
	return s
}

type GetServiceEstimateCostRequestCommodity struct {
	PayPeriod     *int32  `json:"PayPeriod,omitempty" xml:"PayPeriod,omitempty"`
	PayPeriodUnit *string `json:"PayPeriodUnit,omitempty" xml:"PayPeriodUnit,omitempty"`
}

func (s GetServiceEstimateCostRequestCommodity) String() string {
	return tea.Prettify(s)
}

func (s GetServiceEstimateCostRequestCommodity) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostRequestCommodity) SetPayPeriod(v int32) *GetServiceEstimateCostRequestCommodity {
	s.PayPeriod = &v
	return s
}

func (s *GetServiceEstimateCostRequestCommodity) SetPayPeriodUnit(v string) *GetServiceEstimateCostRequestCommodity {
	s.PayPeriodUnit = &v
	return s
}

type GetServiceEstimateCostShrinkRequest struct {
	// example:
	//
	// mRdxWuW2ts
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CommodityShrink *string `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	// example:
	//
	// {\\"PayType\\":\\"PostPaid\\",\\"InstancePassword\\":\\"xxxxxxxxxx\\",\\"EcsInstanceType\\":\\"ecs.g6.large\\",\\"VSwitchId\\":\\"vsw-0jlueyydpuekoxxxxxxxx\\",\\"VpcId\\":\\"vpc-0jlps6mjbgvpqxxxxxxxx\\",\\"ZoneId\\":\\"cn-wulanchabu-a\\",\\"Enable\\":false,\\"RegionId\\":\\"cn-wulanchabu\\"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-16fbd358d75e49xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// example:
	//
	// draft
	ServiceVersion    *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// example:
	//
	// Custom_Image_Ecs
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceEstimateCostShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceEstimateCostShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostShrinkRequest) SetClientToken(v string) *GetServiceEstimateCostShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetCommodityShrink(v string) *GetServiceEstimateCostShrinkRequest {
	s.CommodityShrink = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetParametersShrink(v string) *GetServiceEstimateCostShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetRegionId(v string) *GetServiceEstimateCostShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetServiceId(v string) *GetServiceEstimateCostShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetServiceInstanceId(v string) *GetServiceEstimateCostShrinkRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetServiceVersion(v string) *GetServiceEstimateCostShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetSpecificationName(v string) *GetServiceEstimateCostShrinkRequest {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetTemplateName(v string) *GetServiceEstimateCostShrinkRequest {
	s.TemplateName = &v
	return s
}

type GetServiceEstimateCostResponseBody struct {
	Commodity map[string]interface{} `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	// example:
	//
	// E73F09DC-6C13-5CB1-A10F-7A4E125ABD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {
	//
	//       "ECSInstances":{
	//
	//         "Type":"ALIYUN::ECS::InstanceGroup",
	//
	//         "Success":true,
	//
	//         "Result":{
	//
	//           "Order":{
	//
	//             "Currency":"CNY",
	//
	//             "RuleIds":[
	//
	//               1752723
	//
	//             ],
	//
	//             "DetailInfos":{
	//
	//               "ResourcePriceModel":[
	//
	//                 {
	//
	//                   "Resource":"bandwidth",
	//
	//                   "TradeAmount":0.0,
	//
	//                   "SubRuleIds":[],
	//
	//                   "OriginalAmount":0.0,
	//
	//                   "DiscountAmount":0.0
	//
	//                 },
	//
	//                 {
	//
	//                   "Resource":"image",
	//
	//                   "TradeAmount":0.0,
	//
	//                   "SubRuleIds":[],
	//
	//                   "OriginalAmount":0.0,
	//
	//                   "DiscountAmount":0.0
	//
	//                 },
	//
	//                 {
	//
	//                   "Resource":"instanceType",
	//
	//                   "TradeAmount":0.006966,
	//
	//                   "SubRuleIds":[],
	//
	//                   "OriginalAmount":0.45,
	//
	//                   "DiscountAmount":0.443034
	//
	//                 },
	//
	//                 {
	//
	//                   "Resource":"systemDisk",
	//
	//                   "TradeAmount":0.000867,
	//
	//                   "SubRuleIds":[],
	//
	//                   "OriginalAmount":0.056,
	//
	//                   "DiscountAmount":0.055133
	//
	//                 },
	//
	//                 {
	//
	//                   "Resource":"dataDisk",
	//
	//                   "TradeAmount":0.002167,
	//
	//                   "SubRuleIds":[],
	//
	//                   "OriginalAmount":0.14,
	//
	//                   "DiscountAmount":0.137833
	//
	//                 }
	//
	//               ]
	//
	//             }
	Resources map[string]interface{} `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s GetServiceEstimateCostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceEstimateCostResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostResponseBody) SetCommodity(v map[string]interface{}) *GetServiceEstimateCostResponseBody {
	s.Commodity = v
	return s
}

func (s *GetServiceEstimateCostResponseBody) SetRequestId(v string) *GetServiceEstimateCostResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceEstimateCostResponseBody) SetResources(v map[string]interface{}) *GetServiceEstimateCostResponseBody {
	s.Resources = v
	return s
}

type GetServiceEstimateCostResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceEstimateCostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceEstimateCostResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceEstimateCostResponse) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostResponse) SetHeaders(v map[string]*string) *GetServiceEstimateCostResponse {
	s.Headers = v
	return s
}

func (s *GetServiceEstimateCostResponse) SetStatusCode(v int32) *GetServiceEstimateCostResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceEstimateCostResponse) SetBody(v *GetServiceEstimateCostResponseBody) *GetServiceEstimateCostResponse {
	s.Body = v
	return s
}

type GetServiceInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s GetServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceRequest) SetRegionId(v string) *GetServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceInstanceRequest) SetServiceInstanceId(v string) *GetServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type GetServiceInstanceResponseBody struct {
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	EnableInstanceOps *bool `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// example:
	//
	// 2022-01-01T12:00:00
	EndTime             *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GrafanaDashBoardUrl *string `json:"GrafanaDashBoardUrl,omitempty" xml:"GrafanaDashBoardUrl,omitempty"`
	// example:
	//
	// true
	IsOperated *bool `json:"IsOperated,omitempty" xml:"IsOperated,omitempty"`
	// example:
	//
	// {"renewType":"MONTHLY"}
	LicenseMetadata *string `json:"LicenseMetadata,omitempty" xml:"LicenseMetadata,omitempty"`
	// example:
	//
	// TestName
	Name          *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	NetworkConfig *GetServiceInstanceResponseBodyNetworkConfig `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty" type:"Struct"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	OperatedServiceInstanceId *string `json:"OperatedServiceInstanceId,omitempty" xml:"OperatedServiceInstanceId,omitempty"`
	// example:
	//
	// 2022-01-28T06:48:56Z
	OperationEndTime *string `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	// example:
	//
	// 2021-12-29T06:48:56Z
	OperationStartTime *string `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	// example:
	//
	// {
	//
	//       "InstanceIds": [
	//
	//             "i-hp38ofxl0dsyfi7z****"
	//
	//       ]
	//
	// }
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// example:
	//
	// {
	//
	//       "param": "value"
	//
	// }
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// Subscription
	PayType                 *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PredefinedParameterName *string `json:"PredefinedParameterName,omitempty" xml:"PredefinedParameterName,omitempty"`
	// example:
	//
	// 90
	Progress          *int64  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RdAccountLoginUrl *string `json:"RdAccountLoginUrl,omitempty" xml:"RdAccountLoginUrl,omitempty"`
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// [
	//
	//       {
	//
	//             "StackId": "stack-xxx"
	//
	//       }
	//
	// ]
	Resources *string                                `json:"Resources,omitempty" xml:"Resources,omitempty"`
	Service   *GetServiceInstanceResponseBodyService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// example:
	//
	// User
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// Deployed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// deploy successfully
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// example:
	//
	// 1964460391538545
	SupplierUid  *int64                                `json:"SupplierUid,omitempty" xml:"SupplierUid,omitempty"`
	Tags         []*GetServiceInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateName *string                               `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 2021-05-20T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1234567
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBody) SetBizStatus(v string) *GetServiceInstanceResponseBody {
	s.BizStatus = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetCreateTime(v string) *GetServiceInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEnableInstanceOps(v bool) *GetServiceInstanceResponseBody {
	s.EnableInstanceOps = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEnableUserPrometheus(v bool) *GetServiceInstanceResponseBody {
	s.EnableUserPrometheus = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEndTime(v string) *GetServiceInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetGrafanaDashBoardUrl(v string) *GetServiceInstanceResponseBody {
	s.GrafanaDashBoardUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetIsOperated(v bool) *GetServiceInstanceResponseBody {
	s.IsOperated = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetLicenseMetadata(v string) *GetServiceInstanceResponseBody {
	s.LicenseMetadata = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetName(v string) *GetServiceInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetNetworkConfig(v *GetServiceInstanceResponseBodyNetworkConfig) *GetServiceInstanceResponseBody {
	s.NetworkConfig = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperatedServiceInstanceId(v string) *GetServiceInstanceResponseBody {
	s.OperatedServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperationEndTime(v string) *GetServiceInstanceResponseBody {
	s.OperationEndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperationStartTime(v string) *GetServiceInstanceResponseBody {
	s.OperationStartTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOutputs(v string) *GetServiceInstanceResponseBody {
	s.Outputs = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetParameters(v string) *GetServiceInstanceResponseBody {
	s.Parameters = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetPayType(v string) *GetServiceInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetPredefinedParameterName(v string) *GetServiceInstanceResponseBody {
	s.PredefinedParameterName = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetProgress(v int64) *GetServiceInstanceResponseBody {
	s.Progress = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetRdAccountLoginUrl(v string) *GetServiceInstanceResponseBody {
	s.RdAccountLoginUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetRequestId(v string) *GetServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetResourceGroupId(v string) *GetServiceInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetResources(v string) *GetServiceInstanceResponseBody {
	s.Resources = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetService(v *GetServiceInstanceResponseBodyService) *GetServiceInstanceResponseBody {
	s.Service = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetServiceInstanceId(v string) *GetServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetServiceType(v string) *GetServiceInstanceResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetSource(v string) *GetServiceInstanceResponseBody {
	s.Source = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetStatus(v string) *GetServiceInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetStatusDetail(v string) *GetServiceInstanceResponseBody {
	s.StatusDetail = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetSupplierUid(v int64) *GetServiceInstanceResponseBody {
	s.SupplierUid = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetTags(v []*GetServiceInstanceResponseBodyTags) *GetServiceInstanceResponseBody {
	s.Tags = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetTemplateName(v string) *GetServiceInstanceResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetUpdateTime(v string) *GetServiceInstanceResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetUserId(v int64) *GetServiceInstanceResponseBody {
	s.UserId = &v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfig struct {
	// example:
	//
	// ep-m5ei37240541816b****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// example:
	//
	// epsrv-5ei07324541816bxxxx
	EndpointServiceId            *string                                                                    `json:"EndpointServiceId,omitempty" xml:"EndpointServiceId,omitempty"`
	PrivateVpcConnections        []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections        `json:"PrivateVpcConnections,omitempty" xml:"PrivateVpcConnections,omitempty" type:"Repeated"`
	ReversePrivateVpcConnections []*GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections `json:"ReversePrivateVpcConnections,omitempty" xml:"ReversePrivateVpcConnections,omitempty" type:"Repeated"`
}

func (s GetServiceInstanceResponseBodyNetworkConfig) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfig) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfig {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetEndpointServiceId(v string) *GetServiceInstanceResponseBodyNetworkConfig {
	s.EndpointServiceId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetPrivateVpcConnections(v []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) *GetServiceInstanceResponseBodyNetworkConfig {
	s.PrivateVpcConnections = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetReversePrivateVpcConnections(v []*GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) *GetServiceInstanceResponseBodyNetworkConfig {
	s.ReversePrivateVpcConnections = v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections struct {
	ConnectionConfigs []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs `json:"ConnectionConfigs,omitempty" xml:"ConnectionConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// ep-m5ei37240541816b****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// example:
	//
	// epsrv-5ei07324541816bxxxx
	EndpointServiceId *string `json:"EndpointServiceId,omitempty" xml:"EndpointServiceId,omitempty"`
	// example:
	//
	// test.computenest.aliyuncs.com
	PrivateZoneName *string `json:"PrivateZoneName,omitempty" xml:"PrivateZoneName,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetConnectionConfigs(v []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.ConnectionConfigs = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetEndpointServiceId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.EndpointServiceId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetPrivateZoneName(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.PrivateZoneName = &v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs struct {
	ConnectBandwidth *int32    `json:"ConnectBandwidth,omitempty" xml:"ConnectBandwidth,omitempty"`
	DomainName       *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndpointIps      []*string `json:"EndpointIps,omitempty" xml:"EndpointIps,omitempty" type:"Repeated"`
	// example:
	//
	// Ready
	IngressEndpointStatus *string `json:"IngressEndpointStatus,omitempty" xml:"IngressEndpointStatus,omitempty"`
	// example:
	//
	// Ready
	NetworkServiceStatus *string   `json:"NetworkServiceStatus,omitempty" xml:"NetworkServiceStatus,omitempty"`
	SecurityGroups       []*string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	VSwitches            []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-bp1vwnn14rqpyiczj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetConnectBandwidth(v int32) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.ConnectBandwidth = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetDomainName(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.DomainName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetEndpointIps(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.EndpointIps = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetIngressEndpointStatus(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.IngressEndpointStatus = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetNetworkServiceStatus(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.NetworkServiceStatus = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetSecurityGroups(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.SecurityGroups = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetVSwitches(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.VSwitches = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetVpcId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.VpcId = &v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections struct {
	// example:
	//
	// ep-m5ei42370541816b****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// example:
	//
	// epsrv-5ei07324541816bxxxx
	EndpointServiceId *string `json:"EndpointServiceId,omitempty" xml:"EndpointServiceId,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) SetEndpointServiceId(v string) *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections {
	s.EndpointServiceId = &v
	return s
}

type GetServiceInstanceResponseBodyService struct {
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// example:
	//
	// 2021-05-21T00:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// example:
	//
	// http://example.com
	ServiceDocUrl *string `json:"ServiceDocUrl,omitempty" xml:"ServiceDocUrl,omitempty"`
	// example:
	//
	// service-9c8a3522528b4fe8****
	ServiceId    *string                                              `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceInfos []*GetServiceInstanceResponseBodyServiceServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// https://service-info-private.oss-cn-hangzhou.aliyuncs.com
	ServiceProductUrl *string `json:"ServiceProductUrl,omitempty" xml:"ServiceProductUrl,omitempty"`
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// example:
	//
	// Online
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// example:
	//
	// http://example.com
	SupplierUrl               *string   `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	UpgradableServiceVersions []*string `json:"UpgradableServiceVersions,omitempty" xml:"UpgradableServiceVersions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetServiceInstanceResponseBodyService) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyService) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyService) SetDeployMetadata(v string) *GetServiceInstanceResponseBodyService {
	s.DeployMetadata = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetDeployType(v string) *GetServiceInstanceResponseBodyService {
	s.DeployType = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetPublishTime(v string) *GetServiceInstanceResponseBodyService {
	s.PublishTime = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceDocUrl(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceDocUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceId(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceInfos(v []*GetServiceInstanceResponseBodyServiceServiceInfos) *GetServiceInstanceResponseBodyService {
	s.ServiceInfos = v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceProductUrl(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceProductUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceType(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceType = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetStatus(v string) *GetServiceInstanceResponseBodyService {
	s.Status = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetSupplierName(v string) *GetServiceInstanceResponseBodyService {
	s.SupplierName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetSupplierUrl(v string) *GetServiceInstanceResponseBodyService {
	s.SupplierUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetUpgradableServiceVersions(v []*string) *GetServiceInstanceResponseBodyService {
	s.UpgradableServiceVersions = v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetVersion(v string) *GetServiceInstanceResponseBodyService {
	s.Version = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetVersionName(v string) *GetServiceInstanceResponseBodyService {
	s.VersionName = &v
	return s
}

type GetServiceInstanceResponseBodyServiceServiceInfos struct {
	// example:
	//
	// https://example.com/service-image/c1c4a559-cc60-4af1-b976-98f356602462.png
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// example:
	//
	// zh-CN
	Locale           *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s GetServiceInstanceResponseBodyServiceServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyServiceServiceInfos) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetImage(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Image = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetLocale(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Locale = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetName(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Name = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetShortDescription(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.ShortDescription = &v
	return s
}

type GetServiceInstanceResponseBodyTags struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetServiceInstanceResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyTags) SetKey(v string) *GetServiceInstanceResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetServiceInstanceResponseBodyTags) SetValue(v string) *GetServiceInstanceResponseBodyTags {
	s.Value = &v
	return s
}

type GetServiceInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponse) SetHeaders(v map[string]*string) *GetServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceInstanceResponse) SetStatusCode(v int32) *GetServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceInstanceResponse) SetBody(v *GetServiceInstanceResponseBody) *GetServiceInstanceResponse {
	s.Body = v
	return s
}

type GetServiceTemplateParameterConstraintsRequest struct {
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	// example:
	//
	// true
	EnablePrivateVpcConnection *bool                                                      `json:"EnablePrivateVpcConnection,omitempty" xml:"EnablePrivateVpcConnection,omitempty"`
	Parameters                 []*GetServiceTemplateParameterConstraintsRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-1c11f365190c44xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// example:
	//
	// draft
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// This parameter is required.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsRequest) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetClientToken(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetDeployRegionId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.DeployRegionId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetEnablePrivateVpcConnection(v bool) *GetServiceTemplateParameterConstraintsRequest {
	s.EnablePrivateVpcConnection = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetParameters(v []*GetServiceTemplateParameterConstraintsRequestParameters) *GetServiceTemplateParameterConstraintsRequest {
	s.Parameters = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetRegionId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceInstanceId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceVersion(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetTemplateName(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.TemplateName = &v
	return s
}

type GetServiceTemplateParameterConstraintsRequestParameters struct {
	// example:
	//
	// PayType
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// example:
	//
	// PostPaid
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsRequestParameters) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) SetParameterKey(v string) *GetServiceTemplateParameterConstraintsRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) SetParameterValue(v string) *GetServiceTemplateParameterConstraintsRequestParameters {
	s.ParameterValue = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponseBody struct {
	FamilyConstraints    []*string                                                                 `json:"FamilyConstraints,omitempty" xml:"FamilyConstraints,omitempty" type:"Repeated"`
	ParameterConstraints []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints `json:"ParameterConstraints,omitempty" xml:"ParameterConstraints,omitempty" type:"Repeated"`
	// example:
	//
	// C81C0732-DEBC-559C-B563-7EB2BEB21088
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) SetFamilyConstraints(v []*string) *GetServiceTemplateParameterConstraintsResponseBody {
	s.FamilyConstraints = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) SetParameterConstraints(v []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) *GetServiceTemplateParameterConstraintsResponseBody {
	s.ParameterConstraints = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) SetRequestId(v string) *GetServiceTemplateParameterConstraintsResponseBody {
	s.RequestId = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints struct {
	AllowedValues             []*string `json:"AllowedValues,omitempty" xml:"AllowedValues,omitempty" type:"Repeated"`
	AssociationParameterNames []*string `json:"AssociationParameterNames,omitempty" xml:"AssociationParameterNames,omitempty" type:"Repeated"`
	// example:
	//
	// NoLimit
	Behavior *string `json:"Behavior,omitempty" xml:"Behavior,omitempty"`
	// example:
	//
	// none
	BehaviorReason      *string                                                                                      `json:"BehaviorReason,omitempty" xml:"BehaviorReason,omitempty"`
	OriginalConstraints []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints `json:"OriginalConstraints,omitempty" xml:"OriginalConstraints,omitempty" type:"Repeated"`
	// example:
	//
	// PayType
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetAllowedValues(v []*string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.AllowedValues = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetAssociationParameterNames(v []*string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.AssociationParameterNames = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetBehavior(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.Behavior = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetBehaviorReason(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.BehaviorReason = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetOriginalConstraints(v []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.OriginalConstraints = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetParameterKey(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.ParameterKey = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetType(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.Type = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints struct {
	AllowedValues []*string `json:"AllowedValues,omitempty" xml:"AllowedValues,omitempty" type:"Repeated"`
	// example:
	//
	// lnch_Source
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// example:
	//
	// i-8vb0smn1lf6g77md****
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// serviceinstance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetAllowedValues(v []*string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.AllowedValues = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetPropertyName(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.PropertyName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetResourceName(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.ResourceName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetResourceType(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.ResourceType = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceTemplateParameterConstraintsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponse) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponse) SetHeaders(v map[string]*string) *GetServiceTemplateParameterConstraintsResponse {
	s.Headers = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponse) SetStatusCode(v int32) *GetServiceTemplateParameterConstraintsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponse) SetBody(v *GetServiceTemplateParameterConstraintsResponseBody) *GetServiceTemplateParameterConstraintsResponse {
	s.Body = v
	return s
}

type GetUploadCredentialsRequest struct {
	// This parameter is required.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetUploadCredentialsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUploadCredentialsRequest) GoString() string {
	return s.String()
}

func (s *GetUploadCredentialsRequest) SetFileName(v string) *GetUploadCredentialsRequest {
	s.FileName = &v
	return s
}

type GetUploadCredentialsResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetUploadCredentialsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FCC3321E-D518-1BC4-861E-588E9D4DAFB7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUploadCredentialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUploadCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadCredentialsResponseBody) SetCode(v string) *GetUploadCredentialsResponseBody {
	s.Code = &v
	return s
}

func (s *GetUploadCredentialsResponseBody) SetData(v *GetUploadCredentialsResponseBodyData) *GetUploadCredentialsResponseBody {
	s.Data = v
	return s
}

func (s *GetUploadCredentialsResponseBody) SetHttpStatusCode(v int32) *GetUploadCredentialsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUploadCredentialsResponseBody) SetMessage(v string) *GetUploadCredentialsResponseBody {
	s.Message = &v
	return s
}

func (s *GetUploadCredentialsResponseBody) SetRequestId(v string) *GetUploadCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUploadCredentialsResponseBody) SetSuccess(v bool) *GetUploadCredentialsResponseBody {
	s.Success = &v
	return s
}

type GetUploadCredentialsResponseBodyData struct {
	// example:
	//
	// STS.NUCe19W1FKaHAYAhe********
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// example:
	//
	// 8LQGp59mY23pcXeTdcvSA1cUQZBeD92sFrXi********
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// example:
	//
	// service-info-private
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// 2023-05-18T12:27:59Z
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// example:
	//
	// 221514575922756034/cn-hangzhou/d57c62fbd508xxxxxxxx.json
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// CAISzQN1q6Ft5B2yfSjIr5b2LouNuu5n/KOjQ3/wjGUHYdlagYGdmzz2IH1Le3NrBO8esfgymGFU6v8dlo1dYLQeHhadQI5cs80HtFqLSNaE65LswPlZ2M2ISETPJzfV9pCK
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetUploadCredentialsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUploadCredentialsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUploadCredentialsResponseBodyData) SetAccessKeyId(v string) *GetUploadCredentialsResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetAccessKeySecret(v string) *GetUploadCredentialsResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetBucketName(v string) *GetUploadCredentialsResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetExpireDate(v string) *GetUploadCredentialsResponseBodyData {
	s.ExpireDate = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetKey(v string) *GetUploadCredentialsResponseBodyData {
	s.Key = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetRegionId(v string) *GetUploadCredentialsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetSecurityToken(v string) *GetUploadCredentialsResponseBodyData {
	s.SecurityToken = &v
	return s
}

type GetUploadCredentialsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadCredentialsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUploadCredentialsResponse) GoString() string {
	return s.String()
}

func (s *GetUploadCredentialsResponse) SetHeaders(v map[string]*string) *GetUploadCredentialsResponse {
	s.Headers = v
	return s
}

func (s *GetUploadCredentialsResponse) SetStatusCode(v int32) *GetUploadCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadCredentialsResponse) SetBody(v *GetUploadCredentialsResponseBody) *GetUploadCredentialsResponse {
	s.Body = v
	return s
}

type ListAcrImageRepositoriesRequest struct {
	// example:
	//
	// AcrImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// wordpress
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
}

func (s ListAcrImageRepositoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAcrImageRepositoriesRequest) GoString() string {
	return s.String()
}

func (s *ListAcrImageRepositoriesRequest) SetArtifactType(v string) *ListAcrImageRepositoriesRequest {
	s.ArtifactType = &v
	return s
}

func (s *ListAcrImageRepositoriesRequest) SetMaxResults(v int32) *ListAcrImageRepositoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAcrImageRepositoriesRequest) SetNextToken(v string) *ListAcrImageRepositoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAcrImageRepositoriesRequest) SetRepoName(v string) *ListAcrImageRepositoriesRequest {
	s.RepoName = &v
	return s
}

type ListAcrImageRepositoriesResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken    *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Repositories []*ListAcrImageRepositoriesResponseBodyRepositories `json:"Repositories,omitempty" xml:"Repositories,omitempty" type:"Repeated"`
	// example:
	//
	// C4A145D8-6F6C-532A-9001-9730CDA27578
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAcrImageRepositoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAcrImageRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAcrImageRepositoriesResponseBody) SetMaxResults(v int32) *ListAcrImageRepositoriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBody) SetNextToken(v string) *ListAcrImageRepositoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBody) SetRepositories(v []*ListAcrImageRepositoriesResponseBodyRepositories) *ListAcrImageRepositoriesResponseBody {
	s.Repositories = v
	return s
}

func (s *ListAcrImageRepositoriesResponseBody) SetRequestId(v string) *ListAcrImageRepositoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBody) SetTotalCount(v int32) *ListAcrImageRepositoriesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAcrImageRepositoriesResponseBodyRepositories struct {
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2021-05-20T00:00:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// crr-3gqhkza0wbxxxxxx
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// example:
	//
	// wordpress
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
}

func (s ListAcrImageRepositoriesResponseBodyRepositories) String() string {
	return tea.Prettify(s)
}

func (s ListAcrImageRepositoriesResponseBodyRepositories) GoString() string {
	return s.String()
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) SetCreateTime(v string) *ListAcrImageRepositoriesResponseBodyRepositories {
	s.CreateTime = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) SetModifiedTime(v string) *ListAcrImageRepositoriesResponseBodyRepositories {
	s.ModifiedTime = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) SetRepoId(v string) *ListAcrImageRepositoriesResponseBodyRepositories {
	s.RepoId = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) SetRepoName(v string) *ListAcrImageRepositoriesResponseBodyRepositories {
	s.RepoName = &v
	return s
}

type ListAcrImageRepositoriesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAcrImageRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAcrImageRepositoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAcrImageRepositoriesResponse) GoString() string {
	return s.String()
}

func (s *ListAcrImageRepositoriesResponse) SetHeaders(v map[string]*string) *ListAcrImageRepositoriesResponse {
	s.Headers = v
	return s
}

func (s *ListAcrImageRepositoriesResponse) SetStatusCode(v int32) *ListAcrImageRepositoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAcrImageRepositoriesResponse) SetBody(v *ListAcrImageRepositoriesResponseBody) *ListAcrImageRepositoriesResponse {
	s.Body = v
	return s
}

type ListAcrImageTagsRequest struct {
	// example:
	//
	// AcrImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAW8kZY+u1sYOaYf5JmgmDQQ=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// crr-3gqhkza0wbxxxxxx
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s ListAcrImageTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAcrImageTagsRequest) GoString() string {
	return s.String()
}

func (s *ListAcrImageTagsRequest) SetArtifactType(v string) *ListAcrImageTagsRequest {
	s.ArtifactType = &v
	return s
}

func (s *ListAcrImageTagsRequest) SetMaxResults(v int32) *ListAcrImageTagsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAcrImageTagsRequest) SetNextToken(v string) *ListAcrImageTagsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAcrImageTagsRequest) SetRepoId(v string) *ListAcrImageTagsRequest {
	s.RepoId = &v
	return s
}

type ListAcrImageTagsResponseBody struct {
	Images []*ListAcrImageTagsResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// ey14..
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// FEF343B9-1A15-5789-BE88-7B36190F5BF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAcrImageTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAcrImageTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAcrImageTagsResponseBody) SetImages(v []*ListAcrImageTagsResponseBodyImages) *ListAcrImageTagsResponseBody {
	s.Images = v
	return s
}

func (s *ListAcrImageTagsResponseBody) SetMaxResults(v int32) *ListAcrImageTagsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAcrImageTagsResponseBody) SetNextToken(v string) *ListAcrImageTagsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAcrImageTagsResponseBody) SetRequestId(v string) *ListAcrImageTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAcrImageTagsResponseBody) SetTotalCount(v int32) *ListAcrImageTagsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAcrImageTagsResponseBodyImages struct {
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 188394616
	ImageSize *string `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// example:
	//
	// 2021-05-20T00:00:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 5.7.2
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListAcrImageTagsResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListAcrImageTagsResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListAcrImageTagsResponseBodyImages) SetCreateTime(v string) *ListAcrImageTagsResponseBodyImages {
	s.CreateTime = &v
	return s
}

func (s *ListAcrImageTagsResponseBodyImages) SetImageSize(v string) *ListAcrImageTagsResponseBodyImages {
	s.ImageSize = &v
	return s
}

func (s *ListAcrImageTagsResponseBodyImages) SetModifiedTime(v string) *ListAcrImageTagsResponseBodyImages {
	s.ModifiedTime = &v
	return s
}

func (s *ListAcrImageTagsResponseBodyImages) SetTag(v string) *ListAcrImageTagsResponseBodyImages {
	s.Tag = &v
	return s
}

type ListAcrImageTagsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAcrImageTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAcrImageTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAcrImageTagsResponse) GoString() string {
	return s.String()
}

func (s *ListAcrImageTagsResponse) SetHeaders(v map[string]*string) *ListAcrImageTagsResponse {
	s.Headers = v
	return s
}

func (s *ListAcrImageTagsResponse) SetStatusCode(v int32) *ListAcrImageTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAcrImageTagsResponse) SetBody(v *ListAcrImageTagsResponseBody) *ListAcrImageTagsResponse {
	s.Body = v
	return s
}

type ListArtifactVersionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbfVwapgMwCN1wYzPVzLbItEdB0uWSY7AGnM3qCgm/YnjuEfwSnMwiMkcUoI0hRQzE=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListArtifactVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsRequest) SetArtifactId(v string) *ListArtifactVersionsRequest {
	s.ArtifactId = &v
	return s
}

func (s *ListArtifactVersionsRequest) SetMaxResults(v int32) *ListArtifactVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactVersionsRequest) SetNextToken(v string) *ListArtifactVersionsRequest {
	s.NextToken = &v
	return s
}

type ListArtifactVersionsResponseBody struct {
	Artifacts []*ListArtifactVersionsResponseBodyArtifacts `json:"Artifacts,omitempty" xml:"Artifacts,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbfVwapgMwCN1wYzPVzLbItEdB0uWSY7AGnM3qCgm/YnjuEfwSnMwiMkcUoI0hRQzE=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 46577928-3162-15A6-9084-69820EB9xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListArtifactVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsResponseBody) SetArtifacts(v []*ListArtifactVersionsResponseBodyArtifacts) *ListArtifactVersionsResponseBody {
	s.Artifacts = v
	return s
}

func (s *ListArtifactVersionsResponseBody) SetMaxResults(v int32) *ListArtifactVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactVersionsResponseBody) SetNextToken(v string) *ListArtifactVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListArtifactVersionsResponseBody) SetRequestId(v string) *ListArtifactVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactVersionsResponseBody) SetTotalCount(v int32) *ListArtifactVersionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListArtifactVersionsResponseBodyArtifacts struct {
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// example:
	//
	// {\\"CommodityCode\\":\\"cmjj0005xxxx\\",\\"CommodityVersion\\":\\"V2022xxxx\\"}
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// example:
	//
	// 2
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	// example:
	//
	// 2022-10-20T02:19:53Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2022-10-20T02:19:55Z
	GmtModified   *string            `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ImageDelivery map[string]*string `json:"ImageDelivery,omitempty" xml:"ImageDelivery,omitempty"`
	// example:
	//
	// 100
	Progress   *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ResultFile *string `json:"ResultFile,omitempty" xml:"ResultFile,omitempty"`
	// example:
	//
	// Normal
	SecurityAuditResult *string `json:"SecurityAuditResult,omitempty" xml:"SecurityAuditResult,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// [
	//
	// 					"cn-beijing",
	//
	// 					"cn-hangzhou",
	//
	// 					"cn-shanghai"
	//
	// 				]
	SupportRegionIds *string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty"`
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListArtifactVersionsResponseBodyArtifacts) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactVersionsResponseBodyArtifacts) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactId(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactId = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactProperty(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactProperty = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactType(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactType = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactVersion(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactVersion = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetGmtCreate(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.GmtCreate = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetGmtModified(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.GmtModified = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetImageDelivery(v map[string]*string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ImageDelivery = v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetProgress(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.Progress = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetResultFile(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ResultFile = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetSecurityAuditResult(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.SecurityAuditResult = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetStatus(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.Status = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetSupportRegionIds(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.SupportRegionIds = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetVersionName(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.VersionName = &v
	return s
}

type ListArtifactVersionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsResponse) SetHeaders(v map[string]*string) *ListArtifactVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactVersionsResponse) SetStatusCode(v int32) *ListArtifactVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactVersionsResponse) SetBody(v *ListArtifactVersionsResponseBody) *ListArtifactVersionsResponse {
	s.Body = v
	return s
}

type ListArtifactsRequest struct {
	Filter []*ListArtifactsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbfVwapgMwCN1wYzPVzLbItEdB0uWSY7AGnM3qCgm/YnjuEfwSnMwiMkcUoI0hRQzE=
	NextToken       *string                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceGroupId *string                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*ListArtifactsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListArtifactsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactsRequest) SetFilter(v []*ListArtifactsRequestFilter) *ListArtifactsRequest {
	s.Filter = v
	return s
}

func (s *ListArtifactsRequest) SetMaxResults(v int32) *ListArtifactsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactsRequest) SetNextToken(v string) *ListArtifactsRequest {
	s.NextToken = &v
	return s
}

func (s *ListArtifactsRequest) SetResourceGroupId(v string) *ListArtifactsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListArtifactsRequest) SetTag(v []*ListArtifactsRequestTag) *ListArtifactsRequest {
	s.Tag = v
	return s
}

type ListArtifactsRequestFilter struct {
	// example:
	//
	// ArtifactType
	Name   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListArtifactsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListArtifactsRequestFilter) SetName(v string) *ListArtifactsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListArtifactsRequestFilter) SetValues(v []*string) *ListArtifactsRequestFilter {
	s.Values = v
	return s
}

type ListArtifactsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListArtifactsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsRequestTag) GoString() string {
	return s.String()
}

func (s *ListArtifactsRequestTag) SetKey(v string) *ListArtifactsRequestTag {
	s.Key = &v
	return s
}

func (s *ListArtifactsRequestTag) SetValue(v string) *ListArtifactsRequestTag {
	s.Value = &v
	return s
}

type ListArtifactsResponseBody struct {
	Artifacts []*ListArtifactsResponseBodyArtifacts `json:"Artifacts,omitempty" xml:"Artifacts,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbfVwapgMwCN1wYzPVzLbItEdB0uWSY7AGnM3qCgm/YnjuEfwSnMwiMkcUoI0hRQzE=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 46577928-3162-15A6-9084-69820EB9xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListArtifactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponseBody) SetArtifacts(v []*ListArtifactsResponseBodyArtifacts) *ListArtifactsResponseBody {
	s.Artifacts = v
	return s
}

func (s *ListArtifactsResponseBody) SetMaxResults(v int32) *ListArtifactsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactsResponseBody) SetNextToken(v string) *ListArtifactsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListArtifactsResponseBody) SetRequestId(v string) *ListArtifactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactsResponseBody) SetTotalCount(v int32) *ListArtifactsResponseBody {
	s.TotalCount = &v
	return s
}

type ListArtifactsResponseBodyArtifacts struct {
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2022-10-20T02:19:55Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 2
	MaxVersion      *string `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// Created
	Status *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*ListArtifactsResponseBodyArtifactsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListArtifactsResponseBodyArtifacts) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsResponseBodyArtifacts) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponseBodyArtifacts) SetArtifactId(v string) *ListArtifactsResponseBodyArtifacts {
	s.ArtifactId = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetArtifactType(v string) *ListArtifactsResponseBodyArtifacts {
	s.ArtifactType = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetDescription(v string) *ListArtifactsResponseBodyArtifacts {
	s.Description = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetGmtModified(v string) *ListArtifactsResponseBodyArtifacts {
	s.GmtModified = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetMaxVersion(v string) *ListArtifactsResponseBodyArtifacts {
	s.MaxVersion = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetName(v string) *ListArtifactsResponseBodyArtifacts {
	s.Name = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetResourceGroupId(v string) *ListArtifactsResponseBodyArtifacts {
	s.ResourceGroupId = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetStatus(v string) *ListArtifactsResponseBodyArtifacts {
	s.Status = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetTags(v []*ListArtifactsResponseBodyArtifactsTags) *ListArtifactsResponseBodyArtifacts {
	s.Tags = v
	return s
}

type ListArtifactsResponseBodyArtifactsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListArtifactsResponseBodyArtifactsTags) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsResponseBodyArtifactsTags) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponseBodyArtifactsTags) SetKey(v string) *ListArtifactsResponseBodyArtifactsTags {
	s.Key = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifactsTags) SetValue(v string) *ListArtifactsResponseBodyArtifactsTags {
	s.Value = &v
	return s
}

type ListArtifactsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponse) SetHeaders(v map[string]*string) *ListArtifactsResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactsResponse) SetStatusCode(v int32) *ListArtifactsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactsResponse) SetBody(v *ListArtifactsResponseBody) *ListArtifactsResponse {
	s.Body = v
	return s
}

type ListServiceInstancesRequest struct {
	Filter []*ListServiceInstancesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// false
	ShowDeleted *bool                             `json:"ShowDeleted,omitempty" xml:"ShowDeleted,omitempty"`
	Tag         []*ListServiceInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequest) SetFilter(v []*ListServiceInstancesRequestFilter) *ListServiceInstancesRequest {
	s.Filter = v
	return s
}

func (s *ListServiceInstancesRequest) SetMaxResults(v int32) *ListServiceInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstancesRequest) SetNextToken(v string) *ListServiceInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstancesRequest) SetRegionId(v string) *ListServiceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstancesRequest) SetResourceGroupId(v string) *ListServiceInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServiceInstancesRequest) SetShowDeleted(v bool) *ListServiceInstancesRequest {
	s.ShowDeleted = &v
	return s
}

func (s *ListServiceInstancesRequest) SetTag(v []*ListServiceInstancesRequestTag) *ListServiceInstancesRequest {
	s.Tag = v
	return s
}

type ListServiceInstancesRequestFilter struct {
	// example:
	//
	// ServiceInstanceId
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestFilter) SetName(v string) *ListServiceInstancesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesRequestFilter) SetValue(v []*string) *ListServiceInstancesRequestFilter {
	s.Value = v
	return s
}

type ListServiceInstancesRequestTag struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestTag) SetKey(v string) *ListServiceInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServiceInstancesRequestTag) SetValue(v string) *ListServiceInstancesRequestTag {
	s.Value = &v
	return s
}

type ListServiceInstancesResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// E50287CB-AABF-4877-92C0-289B339A1546
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceInstances []*ListServiceInstancesResponseBodyServiceInstances `json:"ServiceInstances,omitempty" xml:"ServiceInstances,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBody) SetMaxResults(v int32) *ListServiceInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetNextToken(v string) *ListServiceInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetRequestId(v string) *ListServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetServiceInstances(v []*ListServiceInstancesResponseBodyServiceInstances) *ListServiceInstancesResponseBody {
	s.ServiceInstances = v
	return s
}

func (s *ListServiceInstancesResponseBody) SetTotalCount(v int32) *ListServiceInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstances struct {
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	EnableInstanceOps *bool `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	// example:
	//
	// 2022-01-01T12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// true
	IsOperated *bool `json:"IsOperated,omitempty" xml:"IsOperated,omitempty"`
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	OperatedServiceInstanceId *string `json:"OperatedServiceInstanceId,omitempty" xml:"OperatedServiceInstanceId,omitempty"`
	// example:
	//
	// 2021-12-29T06:48:56Z
	OperationEndTime *string `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	// example:
	//
	// 2021-12-29T06:48:56Z
	OperationStartTime *string `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	// example:
	//
	// {"param":"value"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// Subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// 90
	Progress        *int64                                                   `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ResourceGroupId *string                                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Service         *ListServiceInstancesResponseBodyServiceInstancesService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// example:
	//
	// Supplier
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// Deployed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// deploy successfully
	StatusDetail *string                                                 `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	Tags         []*ListServiceInstancesResponseBodyServiceInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateName *string                                                 `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 2021-05-20T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1234567
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstances) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstances) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetBizStatus(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.BizStatus = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetCreateTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.CreateTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetEnableInstanceOps(v bool) *ListServiceInstancesResponseBodyServiceInstances {
	s.EnableInstanceOps = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetEndTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.EndTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetIsOperated(v bool) *ListServiceInstancesResponseBodyServiceInstances {
	s.IsOperated = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetName(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperatedServiceInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperatedServiceInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperationEndTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperationEndTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperationStartTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperationStartTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetParameters(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Parameters = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetPayType(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.PayType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetProgress(v int64) *ListServiceInstancesResponseBodyServiceInstances {
	s.Progress = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetResourceGroupId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetService(v *ListServiceInstancesResponseBodyServiceInstancesService) *ListServiceInstancesResponseBodyServiceInstances {
	s.Service = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetServiceInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetServiceType(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ServiceType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetSource(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Source = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetStatus(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Status = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetStatusDetail(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.StatusDetail = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetTags(v []*ListServiceInstancesResponseBodyServiceInstancesTags) *ListServiceInstancesResponseBodyServiceInstances {
	s.Tags = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetTemplateName(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.TemplateName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetUpdateTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.UpdateTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetUserId(v int64) *ListServiceInstancesResponseBodyServiceInstances {
	s.UserId = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesService struct {
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// example:
	//
	// false
	EnablePrivateVpcConnection *bool `json:"EnablePrivateVpcConnection,omitempty" xml:"EnablePrivateVpcConnection,omitempty"`
	// example:
	//
	// 2021-05-21T00:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId    *string                                                                `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceInfos []*ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// private
	ServiceType        *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	SourceSupplierName *string `json:"SourceSupplierName,omitempty" xml:"SourceSupplierName,omitempty"`
	// example:
	//
	// Online
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// example:
	//
	// http://example.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// example:
	//
	// 1.0
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesService) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesService) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetDeployMetadata(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.DeployMetadata = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetDeployType(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.DeployType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetEnablePrivateVpcConnection(v bool) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.EnablePrivateVpcConnection = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetPublishTime(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.PublishTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceId(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceInfos(v []*ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceInfos = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceType(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSourceSupplierName(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SourceSupplierName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetStatus(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Status = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSupplierName(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SupplierName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSupplierUrl(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SupplierUrl = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetVersion(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Version = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetVersionName(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.VersionName = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos struct {
	// example:
	//
	// https://example.com/service-image/c1c4a559-cc60-4af1-b976-98f356602462.png
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// example:
	//
	// zh-CN
	Locale           *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetImage(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Image = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetLocale(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Locale = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetName(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetShortDescription(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.ShortDescription = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesTags struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesTags) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesTags) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesTags) SetKey(v string) *ListServiceInstancesResponseBodyServiceInstancesTags {
	s.Key = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesTags) SetValue(v string) *ListServiceInstancesResponseBodyServiceInstancesTags {
	s.Value = &v
	return s
}

type ListServiceInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponse) SetHeaders(v map[string]*string) *ListServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstancesResponse) SetStatusCode(v int32) *ListServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstancesResponse) SetBody(v *ListServiceInstancesResponseBody) *ListServiceInstancesResponse {
	s.Body = v
	return s
}

type ListServiceUsagesRequest struct {
	Filter []*ListServiceUsagesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SupplierRole *string `json:"SupplierRole,omitempty" xml:"SupplierRole,omitempty"`
}

func (s ListServiceUsagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesRequest) SetFilter(v []*ListServiceUsagesRequestFilter) *ListServiceUsagesRequest {
	s.Filter = v
	return s
}

func (s *ListServiceUsagesRequest) SetMaxResults(v int32) *ListServiceUsagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceUsagesRequest) SetNextToken(v string) *ListServiceUsagesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceUsagesRequest) SetSupplierRole(v string) *ListServiceUsagesRequest {
	s.SupplierRole = &v
	return s
}

type ListServiceUsagesRequestFilter struct {
	// example:
	//
	// ServiceId
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceUsagesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesRequestFilter) SetName(v string) *ListServiceUsagesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceUsagesRequestFilter) SetValue(v []*string) *ListServiceUsagesRequestFilter {
	s.Value = v
	return s
}

type ListServiceUsagesResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAYChudnQUoBH+mGWFpb6oP0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 18AD0960-A9FE-1AC8-ADF8-22131Fxxxxxx
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceUsages []*ListServiceUsagesResponseBodyServiceUsages `json:"ServiceUsages,omitempty" xml:"ServiceUsages,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceUsagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesResponseBody) SetMaxResults(v int32) *ListServiceUsagesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceUsagesResponseBody) SetNextToken(v string) *ListServiceUsagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceUsagesResponseBody) SetRequestId(v string) *ListServiceUsagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceUsagesResponseBody) SetServiceUsages(v []*ListServiceUsagesResponseBodyServiceUsages) *ListServiceUsagesResponseBody {
	s.ServiceUsages = v
	return s
}

func (s *ListServiceUsagesResponseBody) SetTotalCount(v int32) *ListServiceUsagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceUsagesResponseBodyServiceUsages struct {
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// example:
	//
	// 2022-05-25T02:02:02Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// service-c9f36ec6d19b4exxxxxx
	ServiceId   *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// example:
	//
	// Submitted
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// example:
	//
	// 2022-05-25T02:02:02Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 127383705958xxxx
	UserAliUid      *int64             `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
	UserInformation map[string]*string `json:"UserInformation,omitempty" xml:"UserInformation,omitempty"`
}

func (s ListServiceUsagesResponseBodyServiceUsages) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesResponseBodyServiceUsages) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetComments(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.Comments = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetCreateTime(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.CreateTime = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetServiceId(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.ServiceId = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetServiceName(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.ServiceName = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetStatus(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.Status = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetSupplierName(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.SupplierName = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetUpdateTime(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.UpdateTime = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetUserAliUid(v int64) *ListServiceUsagesResponseBodyServiceUsages {
	s.UserAliUid = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetUserInformation(v map[string]*string) *ListServiceUsagesResponseBodyServiceUsages {
	s.UserInformation = v
	return s
}

type ListServiceUsagesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceUsagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceUsagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesResponse) SetHeaders(v map[string]*string) *ListServiceUsagesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceUsagesResponse) SetStatusCode(v int32) *ListServiceUsagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceUsagesResponse) SetBody(v *ListServiceUsagesResponseBody) *ListServiceUsagesResponse {
	s.Body = v
	return s
}

type ListServicesRequest struct {
	// example:
	//
	// false
	AllVersions *bool                        `json:"AllVersions,omitempty" xml:"AllVersions,omitempty"`
	Filter      []*ListServicesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*ListServicesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) SetAllVersions(v bool) *ListServicesRequest {
	s.AllVersions = &v
	return s
}

func (s *ListServicesRequest) SetFilter(v []*ListServicesRequestFilter) *ListServicesRequest {
	s.Filter = v
	return s
}

func (s *ListServicesRequest) SetMaxResults(v int32) *ListServicesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServicesRequest) SetNextToken(v string) *ListServicesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServicesRequest) SetRegionId(v string) *ListServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServicesRequest) SetResourceGroupId(v string) *ListServicesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServicesRequest) SetTag(v []*ListServicesRequestTag) *ListServicesRequest {
	s.Tag = v
	return s
}

type ListServicesRequestFilter struct {
	// example:
	//
	// Status
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServicesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServicesRequestFilter) SetName(v string) *ListServicesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServicesRequestFilter) SetValue(v []*string) *ListServicesRequestFilter {
	s.Value = v
	return s
}

type ListServicesRequestTag struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServicesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServicesRequestTag) SetKey(v string) *ListServicesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServicesRequestTag) SetValue(v string) *ListServicesRequestTag {
	s.Value = &v
	return s
}

type ListServicesResponseBody struct {
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Services  []*ListServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) SetMaxResults(v int32) *ListServicesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServicesResponseBody) SetNextToken(v string) *ListServicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServicesResponseBody) SetServices(v []*ListServicesResponseBodyServices) *ListServicesResponseBody {
	s.Services = v
	return s
}

func (s *ListServicesResponseBody) SetTotalCount(v int32) *ListServicesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServicesResponseBodyServices struct {
	// example:
	//
	// AutoPass
	ApprovalType *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// example:
	//
	// artifact-21ca53ac16a643xxxxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// example:
	//
	// draft
	ArtifactVersion *string                                    `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	BuildInfo       *string                                    `json:"BuildInfo,omitempty" xml:"BuildInfo,omitempty"`
	Categories      *string                                    `json:"Categories,omitempty" xml:"Categories,omitempty"`
	Commodity       *ListServicesResponseBodyServicesCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// example:
	//
	// cmjj00xxxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	DefaultVersion *bool `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// example:
	//
	// ros
	DeployType                       *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	HasBeta                          *bool   `json:"HasBeta,omitempty" xml:"HasBeta,omitempty"`
	HasDraft                         *bool   `json:"HasDraft,omitempty" xml:"HasDraft,omitempty"`
	LatestResellSourceServiceVersion *string `json:"LatestResellSourceServiceVersion,omitempty" xml:"LatestResellSourceServiceVersion,omitempty"`
	// example:
	//
	// 2021-05-21T00:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// example:
	//
	// ServiceDeployment
	RelationType        *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	ResellApplyStatus   *string `json:"ResellApplyStatus,omitempty" xml:"ResellApplyStatus,omitempty"`
	ResellServiceId     *string `json:"ResellServiceId,omitempty" xml:"ResellServiceId,omitempty"`
	ResourceGroupId     *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServiceDiscoverable *string `json:"ServiceDiscoverable,omitempty" xml:"ServiceDiscoverable,omitempty"`
	// example:
	//
	// service-70a3b15bb62643xxxxxx
	ServiceId    *string                                         `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceInfos []*ListServicesResponseBodyServicesServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// example:
	//
	// Public
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// example:
	//
	// centos_7
	SourceImage          *string `json:"SourceImage,omitempty" xml:"SourceImage,omitempty"`
	SourceServiceId      *string `json:"SourceServiceId,omitempty" xml:"SourceServiceId,omitempty"`
	SourceServiceVersion *string `json:"SourceServiceVersion,omitempty" xml:"SourceServiceVersion,omitempty"`
	SourceSupplierName   *string `json:"SourceSupplierName,omitempty" xml:"SourceSupplierName,omitempty"`
	// example:
	//
	// Online
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// example:
	//
	// http://example.com
	SupplierUrl *string                                 `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	Tags        []*ListServicesResponseBodyServicesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// SingleTenant
	TenantType *string `json:"TenantType,omitempty" xml:"TenantType,omitempty"`
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
	// example:
	//
	// 2021-05-21T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// v2.0.0
	VersionName            *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	VirtualInternetService *string `json:"VirtualInternetService,omitempty" xml:"VirtualInternetService,omitempty"`
}

func (s ListServicesResponseBodyServices) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServices) SetApprovalType(v string) *ListServicesResponseBodyServices {
	s.ApprovalType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetArtifactId(v string) *ListServicesResponseBodyServices {
	s.ArtifactId = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetArtifactVersion(v string) *ListServicesResponseBodyServices {
	s.ArtifactVersion = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetBuildInfo(v string) *ListServicesResponseBodyServices {
	s.BuildInfo = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetCategories(v string) *ListServicesResponseBodyServices {
	s.Categories = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetCommodity(v *ListServicesResponseBodyServicesCommodity) *ListServicesResponseBodyServices {
	s.Commodity = v
	return s
}

func (s *ListServicesResponseBodyServices) SetCommodityCode(v string) *ListServicesResponseBodyServices {
	s.CommodityCode = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetCreateTime(v string) *ListServicesResponseBodyServices {
	s.CreateTime = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetDefaultVersion(v bool) *ListServicesResponseBodyServices {
	s.DefaultVersion = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetDeployType(v string) *ListServicesResponseBodyServices {
	s.DeployType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetHasBeta(v bool) *ListServicesResponseBodyServices {
	s.HasBeta = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetHasDraft(v bool) *ListServicesResponseBodyServices {
	s.HasDraft = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetLatestResellSourceServiceVersion(v string) *ListServicesResponseBodyServices {
	s.LatestResellSourceServiceVersion = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetPublishTime(v string) *ListServicesResponseBodyServices {
	s.PublishTime = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetRelationType(v string) *ListServicesResponseBodyServices {
	s.RelationType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetResellApplyStatus(v string) *ListServicesResponseBodyServices {
	s.ResellApplyStatus = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetResellServiceId(v string) *ListServicesResponseBodyServices {
	s.ResellServiceId = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetResourceGroupId(v string) *ListServicesResponseBodyServices {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceDiscoverable(v string) *ListServicesResponseBodyServices {
	s.ServiceDiscoverable = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceId(v string) *ListServicesResponseBodyServices {
	s.ServiceId = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceInfos(v []*ListServicesResponseBodyServicesServiceInfos) *ListServicesResponseBodyServices {
	s.ServiceInfos = v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceType(v string) *ListServicesResponseBodyServices {
	s.ServiceType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetShareType(v string) *ListServicesResponseBodyServices {
	s.ShareType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSourceImage(v string) *ListServicesResponseBodyServices {
	s.SourceImage = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSourceServiceId(v string) *ListServicesResponseBodyServices {
	s.SourceServiceId = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSourceServiceVersion(v string) *ListServicesResponseBodyServices {
	s.SourceServiceVersion = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSourceSupplierName(v string) *ListServicesResponseBodyServices {
	s.SourceSupplierName = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetStatus(v string) *ListServicesResponseBodyServices {
	s.Status = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSupplierName(v string) *ListServicesResponseBodyServices {
	s.SupplierName = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSupplierUrl(v string) *ListServicesResponseBodyServices {
	s.SupplierUrl = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetTags(v []*ListServicesResponseBodyServicesTags) *ListServicesResponseBodyServices {
	s.Tags = v
	return s
}

func (s *ListServicesResponseBodyServices) SetTenantType(v string) *ListServicesResponseBodyServices {
	s.TenantType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetTrialType(v string) *ListServicesResponseBodyServices {
	s.TrialType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetUpdateTime(v string) *ListServicesResponseBodyServices {
	s.UpdateTime = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetVersion(v string) *ListServicesResponseBodyServices {
	s.Version = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetVersionName(v string) *ListServicesResponseBodyServices {
	s.VersionName = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetVirtualInternetService(v string) *ListServicesResponseBodyServices {
	s.VirtualInternetService = &v
	return s
}

type ListServicesResponseBodyServicesCommodity struct {
	CommodityCode     *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	SaasBoostMetadata *string `json:"SaasBoostMetadata,omitempty" xml:"SaasBoostMetadata,omitempty"`
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListServicesResponseBodyServicesCommodity) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServicesCommodity) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesCommodity) SetCommodityCode(v string) *ListServicesResponseBodyServicesCommodity {
	s.CommodityCode = &v
	return s
}

func (s *ListServicesResponseBodyServicesCommodity) SetSaasBoostMetadata(v string) *ListServicesResponseBodyServicesCommodity {
	s.SaasBoostMetadata = &v
	return s
}

func (s *ListServicesResponseBodyServicesCommodity) SetType(v string) *ListServicesResponseBodyServicesCommodity {
	s.Type = &v
	return s
}

type ListServicesResponseBodyServicesServiceInfos struct {
	// example:
	//
	// http://img.tidb.oss.url
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// example:
	//
	// zh-CN
	Locale           *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s ListServicesResponseBodyServicesServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServicesServiceInfos) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesServiceInfos) SetImage(v string) *ListServicesResponseBodyServicesServiceInfos {
	s.Image = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceInfos) SetLocale(v string) *ListServicesResponseBodyServicesServiceInfos {
	s.Locale = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceInfos) SetName(v string) *ListServicesResponseBodyServicesServiceInfos {
	s.Name = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceInfos) SetShortDescription(v string) *ListServicesResponseBodyServicesServiceInfos {
	s.ShortDescription = &v
	return s
}

type ListServicesResponseBodyServicesTags struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServicesResponseBodyServicesTags) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServicesTags) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesTags) SetKey(v string) *ListServicesResponseBodyServicesTags {
	s.Key = &v
	return s
}

func (s *ListServicesResponseBodyServicesTags) SetValue(v string) *ListServicesResponseBodyServicesTags {
	s.Value = &v
	return s
}

type ListServicesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) SetHeaders(v map[string]*string) *ListServicesResponse {
	s.Headers = v
	return s
}

func (s *ListServicesResponse) SetStatusCode(v int32) *ListServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServicesResponse) SetBody(v *ListServicesResponseBody) *ListServicesResponse {
	s.Body = v
	return s
}

type ModifyServiceInstanceResourcesRequest struct {
	// example:
	//
	// {
	//
	//   "RegionId": "cn-hangzhou",
	//
	//   "Type": "ResourceIds",
	//
	//   "ResourceIds": {
	//
	//     "ALIYUN::ECS::INSTANCE": ["i-xxx", "i-yyy"],
	//
	//     "ALIYUN::RDS::INSTANCE": ["rm-xxx", "rm-yyy"],
	//
	//     "ALIYUN::VPC::VPC": ["vpc-xxx", "vpc-yyy"],
	//
	//     "ALIYUN::SLB::INSTANCE": ["lb-xxx", "lb-yyy"]
	//
	//   }
	//
	// }
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// si-d8a0cc2a1ee04dce****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// example:
	//
	// Import
	ServiceInstanceResourcesAction *string `json:"ServiceInstanceResourcesAction,omitempty" xml:"ServiceInstanceResourcesAction,omitempty"`
}

func (s ModifyServiceInstanceResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyServiceInstanceResourcesRequest) GoString() string {
	return s.String()
}

func (s *ModifyServiceInstanceResourcesRequest) SetResources(v string) *ModifyServiceInstanceResourcesRequest {
	s.Resources = &v
	return s
}

func (s *ModifyServiceInstanceResourcesRequest) SetServiceInstanceId(v string) *ModifyServiceInstanceResourcesRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *ModifyServiceInstanceResourcesRequest) SetServiceInstanceResourcesAction(v string) *ModifyServiceInstanceResourcesRequest {
	s.ServiceInstanceResourcesAction = &v
	return s
}

type ModifyServiceInstanceResourcesResponseBody struct {
	// example:
	//
	// 46577928-3162-15A6-9084-69820EB9xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyServiceInstanceResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyServiceInstanceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyServiceInstanceResourcesResponseBody) SetRequestId(v string) *ModifyServiceInstanceResourcesResponseBody {
	s.RequestId = &v
	return s
}

type ModifyServiceInstanceResourcesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyServiceInstanceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyServiceInstanceResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyServiceInstanceResourcesResponse) GoString() string {
	return s.String()
}

func (s *ModifyServiceInstanceResourcesResponse) SetHeaders(v map[string]*string) *ModifyServiceInstanceResourcesResponse {
	s.Headers = v
	return s
}

func (s *ModifyServiceInstanceResourcesResponse) SetStatusCode(v int32) *ModifyServiceInstanceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyServiceInstanceResourcesResponse) SetBody(v *ModifyServiceInstanceResourcesResponseBody) *ModifyServiceInstanceResourcesResponse {
	s.Body = v
	return s
}

type PushMeteringDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{\\"StartTime\\":1681264800,\\"EndTime\\":1681268400,\\"Entities\\":[{\\"Key\\":\\"Unit\\",\\"Value\\":\\"0\\"}]}]
	Metering *string `json:"Metering,omitempty" xml:"Metering,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// si-0e6fca6a51a54420****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s PushMeteringDataRequest) String() string {
	return tea.Prettify(s)
}

func (s PushMeteringDataRequest) GoString() string {
	return s.String()
}

func (s *PushMeteringDataRequest) SetMetering(v string) *PushMeteringDataRequest {
	s.Metering = &v
	return s
}

func (s *PushMeteringDataRequest) SetServiceInstanceId(v string) *PushMeteringDataRequest {
	s.ServiceInstanceId = &v
	return s
}

type PushMeteringDataResponseBody struct {
	// example:
	//
	// 94E89857-B994-44B6-9C4F-DBD200E9XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushMeteringDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushMeteringDataResponseBody) GoString() string {
	return s.String()
}

func (s *PushMeteringDataResponseBody) SetRequestId(v string) *PushMeteringDataResponseBody {
	s.RequestId = &v
	return s
}

type PushMeteringDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushMeteringDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushMeteringDataResponse) String() string {
	return tea.Prettify(s)
}

func (s PushMeteringDataResponse) GoString() string {
	return s.String()
}

func (s *PushMeteringDataResponse) SetHeaders(v map[string]*string) *PushMeteringDataResponse {
	s.Headers = v
	return s
}

func (s *PushMeteringDataResponse) SetStatusCode(v int32) *PushMeteringDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PushMeteringDataResponse) SetBody(v *PushMeteringDataResponseBody) *PushMeteringDataResponse {
	s.Body = v
	return s
}

type RegisterServiceRequest struct {
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-f7024a22ea5149xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s RegisterServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterServiceRequest) GoString() string {
	return s.String()
}

func (s *RegisterServiceRequest) SetClientToken(v string) *RegisterServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *RegisterServiceRequest) SetRegionId(v string) *RegisterServiceRequest {
	s.RegionId = &v
	return s
}

func (s *RegisterServiceRequest) SetServiceId(v string) *RegisterServiceRequest {
	s.ServiceId = &v
	return s
}

type RegisterServiceResponseBody struct {
	// example:
	//
	// sr-72dd5071e90c40xxxxxx
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// example:
	//
	// A361BA9E-2713-52C8-AFFC-C26E5180456E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterServiceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterServiceResponseBody) SetRegistrationId(v string) *RegisterServiceResponseBody {
	s.RegistrationId = &v
	return s
}

func (s *RegisterServiceResponseBody) SetRequestId(v string) *RegisterServiceResponseBody {
	s.RequestId = &v
	return s
}

type RegisterServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterServiceResponse) GoString() string {
	return s.String()
}

func (s *RegisterServiceResponse) SetHeaders(v map[string]*string) *RegisterServiceResponse {
	s.Headers = v
	return s
}

func (s *RegisterServiceResponse) SetStatusCode(v int32) *RegisterServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterServiceResponse) SetBody(v *RegisterServiceResponseBody) *RegisterServiceResponse {
	s.Body = v
	return s
}

type ReleaseArtifactRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// artifact-9feded91880e4c78xxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
}

func (s ReleaseArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseArtifactRequest) GoString() string {
	return s.String()
}

func (s *ReleaseArtifactRequest) SetArtifactId(v string) *ReleaseArtifactRequest {
	s.ArtifactId = &v
	return s
}

type ReleaseArtifactResponseBody struct {
	// example:
	//
	// artifact-9feded91880e4c78xxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// example:
	//
	// "{\\"Url\\":\\"https://computenest-artifacts-draft-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/130920852836xxxx/cn-hangzhou/service-8072a04e5a134382xxxx/165095355xxxx/changes.txt\\",\\"ConfigurationMetadata\\":\\"{\\\\\\"WorkDir\\\\\\":\\\\\\"/root\\\\\\",\\\\\\"Platform\\\\\\":\\\\\\"Linux\\\\\\",\\\\\\"CommandType\\\\\\":\\\\\\"RunShellScript\\\\\\",\\\\\\"UpgradeScript\\\\\\":\\\\\\"cd /root\\\\\\\\ncp changes.txt cpchanges.txt\\\\\\\\nmv changes.txt mvchangge.txt\\\\\\"}\\"}"
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// example:
	//
	// File
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// example:
	//
	// draft
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1650954178000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 3818BA7D-3F50-1A44-9FF3-04A52A59XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Created
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ReleaseArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseArtifactResponseBody) SetArtifactId(v string) *ReleaseArtifactResponseBody {
	s.ArtifactId = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetArtifactProperty(v string) *ReleaseArtifactResponseBody {
	s.ArtifactProperty = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetArtifactType(v string) *ReleaseArtifactResponseBody {
	s.ArtifactType = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetArtifactVersion(v string) *ReleaseArtifactResponseBody {
	s.ArtifactVersion = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetDescription(v string) *ReleaseArtifactResponseBody {
	s.Description = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetGmtModified(v string) *ReleaseArtifactResponseBody {
	s.GmtModified = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetRequestId(v string) *ReleaseArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetStatus(v string) *ReleaseArtifactResponseBody {
	s.Status = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetVersionName(v string) *ReleaseArtifactResponseBody {
	s.VersionName = &v
	return s
}

type ReleaseArtifactResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseArtifactResponse) GoString() string {
	return s.String()
}

func (s *ReleaseArtifactResponse) SetHeaders(v map[string]*string) *ReleaseArtifactResponse {
	s.Headers = v
	return s
}

func (s *ReleaseArtifactResponse) SetStatusCode(v int32) *ReleaseArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseArtifactResponse) SetBody(v *ReleaseArtifactResponseBody) *ReleaseArtifactResponse {
	s.Body = v
	return s
}

type RestartServiceInstanceRequest struct {
	// example:
	//
	// 10CM943JP0EN9****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s RestartServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartServiceInstanceRequest) SetClientToken(v string) *RestartServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartServiceInstanceRequest) SetRegionId(v string) *RestartServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RestartServiceInstanceRequest) SetServiceInstanceId(v string) *RestartServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type RestartServiceInstanceResponseBody struct {
	// example:
	//
	// B1A0198B-F316-1B72-B8DD-28B6F6D6XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartServiceInstanceResponseBody) SetRequestId(v string) *RestartServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestartServiceInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartServiceInstanceResponse) SetHeaders(v map[string]*string) *RestartServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartServiceInstanceResponse) SetStatusCode(v int32) *RestartServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartServiceInstanceResponse) SetBody(v *RestartServiceInstanceResponseBody) *RestartServiceInstanceResponse {
	s.Body = v
	return s
}

type StartServiceInstanceRequest struct {
	// example:
	//
	// 10CM943JP0EN9****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s StartServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartServiceInstanceRequest) SetClientToken(v string) *StartServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StartServiceInstanceRequest) SetRegionId(v string) *StartServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartServiceInstanceRequest) SetServiceInstanceId(v string) *StartServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type StartServiceInstanceResponseBody struct {
	// example:
	//
	// 2E91D771-0183-52CE-86CB-882D99B2CB77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartServiceInstanceResponseBody) SetRequestId(v string) *StartServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartServiceInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartServiceInstanceResponse) SetHeaders(v map[string]*string) *StartServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartServiceInstanceResponse) SetStatusCode(v int32) *StartServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartServiceInstanceResponse) SetBody(v *StartServiceInstanceResponseBody) *StartServiceInstanceResponse {
	s.Body = v
	return s
}

type StopServiceInstanceRequest struct {
	// example:
	//
	// 10CM943JP0EN9****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// si-c39ed4779cec449f****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s StopServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopServiceInstanceRequest) SetClientToken(v string) *StopServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StopServiceInstanceRequest) SetRegionId(v string) *StopServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StopServiceInstanceRequest) SetServiceInstanceId(v string) *StopServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type StopServiceInstanceResponseBody struct {
	// example:
	//
	// 49A369EF-A302-5006-B0CE-94CED47C38CF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopServiceInstanceResponseBody) SetRequestId(v string) *StopServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StopServiceInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopServiceInstanceResponse) SetHeaders(v map[string]*string) *StopServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopServiceInstanceResponse) SetStatusCode(v int32) *StopServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopServiceInstanceResponse) SetBody(v *StopServiceInstanceResponseBody) *StopServiceInstanceResponse {
	s.Body = v
	return s
}

type UpdateArtifactRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// This parameter is required.
	ArtifactProperty *UpdateArtifactRequestArtifactProperty `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty" type:"Struct"`
	Description      *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	SupportRegionIds []*string                              `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactRequest) GoString() string {
	return s.String()
}

func (s *UpdateArtifactRequest) SetArtifactId(v string) *UpdateArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *UpdateArtifactRequest) SetArtifactProperty(v *UpdateArtifactRequestArtifactProperty) *UpdateArtifactRequest {
	s.ArtifactProperty = v
	return s
}

func (s *UpdateArtifactRequest) SetDescription(v string) *UpdateArtifactRequest {
	s.Description = &v
	return s
}

func (s *UpdateArtifactRequest) SetSupportRegionIds(v []*string) *UpdateArtifactRequest {
	s.SupportRegionIds = v
	return s
}

func (s *UpdateArtifactRequest) SetVersionName(v string) *UpdateArtifactRequest {
	s.VersionName = &v
	return s
}

type UpdateArtifactRequestArtifactProperty struct {
	// example:
	//
	// cmjj00xxxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// V1.0
	CommodityVersion *string `json:"CommodityVersion,omitempty" xml:"CommodityVersion,omitempty"`
	// example:
	//
	// {\\"WorkDir\\":\\"/root\\",\\"CommandType\\":\\"RunShellScript\\",\\"Platform\\":\\"Linux\\",\\"Script\\":\\"echo hello\\"}
	FileScriptMetadata *string `json:"FileScriptMetadata,omitempty" xml:"FileScriptMetadata,omitempty"`
	// example:
	//
	// m-0xij191j9cuev6ucxxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// {"ScriptMetadata":"{\\"CommandType\\":\\"RunShellScript\\",\\"Platform\\":\\"Linux\\",\\"Script\\":\\"ls\\"}"}
	ScriptMetadata *string `json:"ScriptMetadata,omitempty" xml:"ScriptMetadata,omitempty"`
	// example:
	//
	// https://service-info-private.oss-cn-hangzhou.aliyuncs.com/1309208528xxxxxx/template/2e1ce8fc-xxxx-481c-9e8e-789ba9db487d.json
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateArtifactRequestArtifactProperty) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactRequestArtifactProperty) GoString() string {
	return s.String()
}

func (s *UpdateArtifactRequestArtifactProperty) SetCommodityCode(v string) *UpdateArtifactRequestArtifactProperty {
	s.CommodityCode = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetCommodityVersion(v string) *UpdateArtifactRequestArtifactProperty {
	s.CommodityVersion = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetFileScriptMetadata(v string) *UpdateArtifactRequestArtifactProperty {
	s.FileScriptMetadata = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetImageId(v string) *UpdateArtifactRequestArtifactProperty {
	s.ImageId = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetRegionId(v string) *UpdateArtifactRequestArtifactProperty {
	s.RegionId = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetScriptMetadata(v string) *UpdateArtifactRequestArtifactProperty {
	s.ScriptMetadata = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetUrl(v string) *UpdateArtifactRequestArtifactProperty {
	s.Url = &v
	return s
}

type UpdateArtifactShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// This parameter is required.
	ArtifactPropertyShrink *string   `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	Description            *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	SupportRegionIds       []*string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateArtifactShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateArtifactShrinkRequest) SetArtifactId(v string) *UpdateArtifactShrinkRequest {
	s.ArtifactId = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetArtifactPropertyShrink(v string) *UpdateArtifactShrinkRequest {
	s.ArtifactPropertyShrink = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetDescription(v string) *UpdateArtifactShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetSupportRegionIds(v []*string) *UpdateArtifactShrinkRequest {
	s.SupportRegionIds = v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetVersionName(v string) *UpdateArtifactShrinkRequest {
	s.VersionName = &v
	return s
}

type UpdateArtifactResponseBody struct {
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// example:
	//
	// {\\"CommodityCode\\":\\"cmjj0005xxxx\\",\\"CommodityVersion\\":\\"V2022xxxx\\"}
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// example:
	//
	// 2
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2022-10-20T02:19:55Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// B1A0198B-F316-1B72-B8DD-28B6F6D6XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// [
	//
	// 			"cn-beijing",
	//
	// 			"cn-hangzhou",
	//
	// 			"cn-shanghai"
	//
	// 		]
	SupportRegionIds *string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty"`
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateArtifactResponseBody) SetArtifactId(v string) *UpdateArtifactResponseBody {
	s.ArtifactId = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetArtifactProperty(v string) *UpdateArtifactResponseBody {
	s.ArtifactProperty = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetArtifactType(v string) *UpdateArtifactResponseBody {
	s.ArtifactType = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetArtifactVersion(v string) *UpdateArtifactResponseBody {
	s.ArtifactVersion = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetDescription(v string) *UpdateArtifactResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetGmtModified(v string) *UpdateArtifactResponseBody {
	s.GmtModified = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetRequestId(v string) *UpdateArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetStatus(v string) *UpdateArtifactResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetSupportRegionIds(v string) *UpdateArtifactResponseBody {
	s.SupportRegionIds = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetVersionName(v string) *UpdateArtifactResponseBody {
	s.VersionName = &v
	return s
}

type UpdateArtifactResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactResponse) GoString() string {
	return s.String()
}

func (s *UpdateArtifactResponse) SetHeaders(v map[string]*string) *UpdateArtifactResponse {
	s.Headers = v
	return s
}

func (s *UpdateArtifactResponse) SetStatusCode(v int32) *UpdateArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateArtifactResponse) SetBody(v *UpdateArtifactResponseBody) *UpdateArtifactResponse {
	s.Body = v
	return s
}

type UpdateServiceRequest struct {
	// example:
	//
	// {\\"CmsTemplateId\\":1162921,\\"TemplateUrl\\":\\"https://service-info-private.oss-cn-hangzhou.aliyuncs.com/1760465342xxxxxx/template/c072ef50-6c03-4d9c-8f0e-d1c440xxxxxx.json\\"}
	AlarmMetadata *string `json:"AlarmMetadata,omitempty" xml:"AlarmMetadata,omitempty"`
	// example:
	//
	// 788E7CP0EN9D51P
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// {\\"EstimateTime\\":null,\\"SupplierDeployMetadata\\":{\\"DeployTimeout\\":7200},\\"EnableVnc\\":false}
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// example:
	//
	// 259200
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// false
	IsSupportOperated *bool   `json:"IsSupportOperated,omitempty" xml:"IsSupportOperated,omitempty"`
	LicenseMetadata   *string `json:"LicenseMetadata,omitempty" xml:"LicenseMetadata,omitempty"`
	LogMetadata       *string `json:"LogMetadata,omitempty" xml:"LogMetadata,omitempty"`
	// example:
	//
	// {\\"PrometheusConfigMap\\":{\\"Custom_Image_Ecs\\":{\\"EnablePrometheus\\":false}}}
	OperationMetadata *string `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty"`
	// example:
	//
	// policyName1, policyName2
	PolicyNames *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Resellable *bool   `json:"Resellable,omitempty" xml:"Resellable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-1dda29c3eca648xxxxxx
	ServiceId   *string                            `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceInfo []*UpdateServiceRequestServiceInfo `json:"ServiceInfo,omitempty" xml:"ServiceInfo,omitempty" type:"Repeated"`
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// example:
	//
	// SingleTenant
	TenantType *string `json:"TenantType,omitempty" xml:"TenantType,omitempty"`
	// example:
	//
	// 7
	TrialDuration *int32                            `json:"TrialDuration,omitempty" xml:"TrialDuration,omitempty"`
	UpdateOption  *UpdateServiceRequestUpdateOption `json:"UpdateOption,omitempty" xml:"UpdateOption,omitempty" type:"Struct"`
	// example:
	//
	// {\\"Description\\":\\"xxx\\",\\"SupportRollback\\":true,\\"SupportUpgradeFromVersions\\":[],\\"UpgradeComponents\\":[\\"Configuration\\"]}
	UpgradeMetadata *string `json:"UpgradeMetadata,omitempty" xml:"UpgradeMetadata,omitempty"`
	// example:
	//
	// Draft
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequest) SetAlarmMetadata(v string) *UpdateServiceRequest {
	s.AlarmMetadata = &v
	return s
}

func (s *UpdateServiceRequest) SetClientToken(v string) *UpdateServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceRequest) SetDeployMetadata(v string) *UpdateServiceRequest {
	s.DeployMetadata = &v
	return s
}

func (s *UpdateServiceRequest) SetDeployType(v string) *UpdateServiceRequest {
	s.DeployType = &v
	return s
}

func (s *UpdateServiceRequest) SetDuration(v int64) *UpdateServiceRequest {
	s.Duration = &v
	return s
}

func (s *UpdateServiceRequest) SetIsSupportOperated(v bool) *UpdateServiceRequest {
	s.IsSupportOperated = &v
	return s
}

func (s *UpdateServiceRequest) SetLicenseMetadata(v string) *UpdateServiceRequest {
	s.LicenseMetadata = &v
	return s
}

func (s *UpdateServiceRequest) SetLogMetadata(v string) *UpdateServiceRequest {
	s.LogMetadata = &v
	return s
}

func (s *UpdateServiceRequest) SetOperationMetadata(v string) *UpdateServiceRequest {
	s.OperationMetadata = &v
	return s
}

func (s *UpdateServiceRequest) SetPolicyNames(v string) *UpdateServiceRequest {
	s.PolicyNames = &v
	return s
}

func (s *UpdateServiceRequest) SetRegionId(v string) *UpdateServiceRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceRequest) SetResellable(v bool) *UpdateServiceRequest {
	s.Resellable = &v
	return s
}

func (s *UpdateServiceRequest) SetServiceId(v string) *UpdateServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateServiceRequest) SetServiceInfo(v []*UpdateServiceRequestServiceInfo) *UpdateServiceRequest {
	s.ServiceInfo = v
	return s
}

func (s *UpdateServiceRequest) SetServiceType(v string) *UpdateServiceRequest {
	s.ServiceType = &v
	return s
}

func (s *UpdateServiceRequest) SetServiceVersion(v string) *UpdateServiceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *UpdateServiceRequest) SetTenantType(v string) *UpdateServiceRequest {
	s.TenantType = &v
	return s
}

func (s *UpdateServiceRequest) SetTrialDuration(v int32) *UpdateServiceRequest {
	s.TrialDuration = &v
	return s
}

func (s *UpdateServiceRequest) SetUpdateOption(v *UpdateServiceRequestUpdateOption) *UpdateServiceRequest {
	s.UpdateOption = v
	return s
}

func (s *UpdateServiceRequest) SetUpgradeMetadata(v string) *UpdateServiceRequest {
	s.UpgradeMetadata = &v
	return s
}

func (s *UpdateServiceRequest) SetVersionName(v string) *UpdateServiceRequest {
	s.VersionName = &v
	return s
}

type UpdateServiceRequestServiceInfo struct {
	Agreements []*UpdateServiceRequestServiceInfoAgreements `json:"Agreements,omitempty" xml:"Agreements,omitempty" type:"Repeated"`
	// example:
	//
	// http://img.tidb.oss.url
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// example:
	//
	// zh-CN
	Locale             *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	LongDescriptionUrl *string `json:"LongDescriptionUrl,omitempty" xml:"LongDescriptionUrl,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription   *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s UpdateServiceRequestServiceInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequestServiceInfo) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestServiceInfo) SetAgreements(v []*UpdateServiceRequestServiceInfoAgreements) *UpdateServiceRequestServiceInfo {
	s.Agreements = v
	return s
}

func (s *UpdateServiceRequestServiceInfo) SetImage(v string) *UpdateServiceRequestServiceInfo {
	s.Image = &v
	return s
}

func (s *UpdateServiceRequestServiceInfo) SetLocale(v string) *UpdateServiceRequestServiceInfo {
	s.Locale = &v
	return s
}

func (s *UpdateServiceRequestServiceInfo) SetLongDescriptionUrl(v string) *UpdateServiceRequestServiceInfo {
	s.LongDescriptionUrl = &v
	return s
}

func (s *UpdateServiceRequestServiceInfo) SetName(v string) *UpdateServiceRequestServiceInfo {
	s.Name = &v
	return s
}

func (s *UpdateServiceRequestServiceInfo) SetShortDescription(v string) *UpdateServiceRequestServiceInfo {
	s.ShortDescription = &v
	return s
}

type UpdateServiceRequestServiceInfoAgreements struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateServiceRequestServiceInfoAgreements) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequestServiceInfoAgreements) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestServiceInfoAgreements) SetName(v string) *UpdateServiceRequestServiceInfoAgreements {
	s.Name = &v
	return s
}

func (s *UpdateServiceRequestServiceInfoAgreements) SetUrl(v string) *UpdateServiceRequestServiceInfoAgreements {
	s.Url = &v
	return s
}

type UpdateServiceRequestUpdateOption struct {
	UpdateFrom *string `json:"UpdateFrom,omitempty" xml:"UpdateFrom,omitempty"`
}

func (s UpdateServiceRequestUpdateOption) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequestUpdateOption) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestUpdateOption) SetUpdateFrom(v string) *UpdateServiceRequestUpdateOption {
	s.UpdateFrom = &v
	return s
}

type UpdateServiceShrinkRequest struct {
	// example:
	//
	// {\\"CmsTemplateId\\":1162921,\\"TemplateUrl\\":\\"https://service-info-private.oss-cn-hangzhou.aliyuncs.com/1760465342xxxxxx/template/c072ef50-6c03-4d9c-8f0e-d1c440xxxxxx.json\\"}
	AlarmMetadata *string `json:"AlarmMetadata,omitempty" xml:"AlarmMetadata,omitempty"`
	// example:
	//
	// 788E7CP0EN9D51P
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// {\\"EstimateTime\\":null,\\"SupplierDeployMetadata\\":{\\"DeployTimeout\\":7200},\\"EnableVnc\\":false}
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// example:
	//
	// 259200
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// false
	IsSupportOperated *bool   `json:"IsSupportOperated,omitempty" xml:"IsSupportOperated,omitempty"`
	LicenseMetadata   *string `json:"LicenseMetadata,omitempty" xml:"LicenseMetadata,omitempty"`
	LogMetadata       *string `json:"LogMetadata,omitempty" xml:"LogMetadata,omitempty"`
	// example:
	//
	// {\\"PrometheusConfigMap\\":{\\"Custom_Image_Ecs\\":{\\"EnablePrometheus\\":false}}}
	OperationMetadata *string `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty"`
	// example:
	//
	// policyName1, policyName2
	PolicyNames *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Resellable *bool   `json:"Resellable,omitempty" xml:"Resellable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-1dda29c3eca648xxxxxx
	ServiceId   *string                                  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceInfo []*UpdateServiceShrinkRequestServiceInfo `json:"ServiceInfo,omitempty" xml:"ServiceInfo,omitempty" type:"Repeated"`
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// example:
	//
	// SingleTenant
	TenantType *string `json:"TenantType,omitempty" xml:"TenantType,omitempty"`
	// example:
	//
	// 7
	TrialDuration      *int32  `json:"TrialDuration,omitempty" xml:"TrialDuration,omitempty"`
	UpdateOptionShrink *string `json:"UpdateOption,omitempty" xml:"UpdateOption,omitempty"`
	// example:
	//
	// {\\"Description\\":\\"xxx\\",\\"SupportRollback\\":true,\\"SupportUpgradeFromVersions\\":[],\\"UpgradeComponents\\":[\\"Configuration\\"]}
	UpgradeMetadata *string `json:"UpgradeMetadata,omitempty" xml:"UpgradeMetadata,omitempty"`
	// example:
	//
	// Draft
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateServiceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceShrinkRequest) SetAlarmMetadata(v string) *UpdateServiceShrinkRequest {
	s.AlarmMetadata = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetClientToken(v string) *UpdateServiceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetDeployMetadata(v string) *UpdateServiceShrinkRequest {
	s.DeployMetadata = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetDeployType(v string) *UpdateServiceShrinkRequest {
	s.DeployType = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetDuration(v int64) *UpdateServiceShrinkRequest {
	s.Duration = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetIsSupportOperated(v bool) *UpdateServiceShrinkRequest {
	s.IsSupportOperated = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetLicenseMetadata(v string) *UpdateServiceShrinkRequest {
	s.LicenseMetadata = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetLogMetadata(v string) *UpdateServiceShrinkRequest {
	s.LogMetadata = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetOperationMetadata(v string) *UpdateServiceShrinkRequest {
	s.OperationMetadata = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetPolicyNames(v string) *UpdateServiceShrinkRequest {
	s.PolicyNames = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetRegionId(v string) *UpdateServiceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetResellable(v bool) *UpdateServiceShrinkRequest {
	s.Resellable = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetServiceId(v string) *UpdateServiceShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetServiceInfo(v []*UpdateServiceShrinkRequestServiceInfo) *UpdateServiceShrinkRequest {
	s.ServiceInfo = v
	return s
}

func (s *UpdateServiceShrinkRequest) SetServiceType(v string) *UpdateServiceShrinkRequest {
	s.ServiceType = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetServiceVersion(v string) *UpdateServiceShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetTenantType(v string) *UpdateServiceShrinkRequest {
	s.TenantType = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetTrialDuration(v int32) *UpdateServiceShrinkRequest {
	s.TrialDuration = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetUpdateOptionShrink(v string) *UpdateServiceShrinkRequest {
	s.UpdateOptionShrink = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetUpgradeMetadata(v string) *UpdateServiceShrinkRequest {
	s.UpgradeMetadata = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetVersionName(v string) *UpdateServiceShrinkRequest {
	s.VersionName = &v
	return s
}

type UpdateServiceShrinkRequestServiceInfo struct {
	Agreements []*UpdateServiceShrinkRequestServiceInfoAgreements `json:"Agreements,omitempty" xml:"Agreements,omitempty" type:"Repeated"`
	// example:
	//
	// http://img.tidb.oss.url
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// example:
	//
	// zh-CN
	Locale             *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	LongDescriptionUrl *string `json:"LongDescriptionUrl,omitempty" xml:"LongDescriptionUrl,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription   *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s UpdateServiceShrinkRequestServiceInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceShrinkRequestServiceInfo) GoString() string {
	return s.String()
}

func (s *UpdateServiceShrinkRequestServiceInfo) SetAgreements(v []*UpdateServiceShrinkRequestServiceInfoAgreements) *UpdateServiceShrinkRequestServiceInfo {
	s.Agreements = v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfo) SetImage(v string) *UpdateServiceShrinkRequestServiceInfo {
	s.Image = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfo) SetLocale(v string) *UpdateServiceShrinkRequestServiceInfo {
	s.Locale = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfo) SetLongDescriptionUrl(v string) *UpdateServiceShrinkRequestServiceInfo {
	s.LongDescriptionUrl = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfo) SetName(v string) *UpdateServiceShrinkRequestServiceInfo {
	s.Name = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfo) SetShortDescription(v string) *UpdateServiceShrinkRequestServiceInfo {
	s.ShortDescription = &v
	return s
}

type UpdateServiceShrinkRequestServiceInfoAgreements struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateServiceShrinkRequestServiceInfoAgreements) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceShrinkRequestServiceInfoAgreements) GoString() string {
	return s.String()
}

func (s *UpdateServiceShrinkRequestServiceInfoAgreements) SetName(v string) *UpdateServiceShrinkRequestServiceInfoAgreements {
	s.Name = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfoAgreements) SetUrl(v string) *UpdateServiceShrinkRequestServiceInfoAgreements {
	s.Url = &v
	return s
}

type UpdateServiceResponseBody struct {
	// example:
	//
	// DF0F666F-FBBC-55C3-A368-C955DE7B4839
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBody) SetRequestId(v string) *UpdateServiceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponse) SetHeaders(v map[string]*string) *UpdateServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceResponse) SetStatusCode(v int32) *UpdateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceResponse) SetBody(v *UpdateServiceResponseBody) *UpdateServiceResponse {
	s.Body = v
	return s
}

type UpdateServiceInstanceAttributeRequest struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2023-12-25T02:28:40Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// si-3df88e962cdexxxxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s UpdateServiceInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributeRequest) SetEndTime(v string) *UpdateServiceInstanceAttributeRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateServiceInstanceAttributeRequest) SetRegionId(v string) *UpdateServiceInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceInstanceAttributeRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceAttributeRequest {
	s.ServiceInstanceId = &v
	return s
}

type UpdateServiceInstanceAttributeResponseBody struct {
	// example:
	//
	// 0CB2E0A9-B4DF-5C16-86AD-C511C483144B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributeResponseBody) SetRequestId(v string) *UpdateServiceInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceInstanceAttributeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributeResponse) SetHeaders(v map[string]*string) *UpdateServiceInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceInstanceAttributeResponse) SetStatusCode(v int32) *UpdateServiceInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceInstanceAttributeResponse) SetBody(v *UpdateServiceInstanceAttributeResponseBody) *UpdateServiceInstanceAttributeResponse {
	s.Body = v
	return s
}

type UpdateServiceInstanceSpecRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	EnableUserPrometheus *bool   `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	OperationName        *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// example:
	//
	// {\\"EcsInstanceParameter\\":\\"4vCPU 8GiB\\",\\"ZoneId\\":\\"cn-heyuan-a\\",\\"SystemDiskSize\\":50,\\"DataDiskSize\\":150,\\"InternetMaxBandwidthOut\\":2,\\"RegionId\\":\\"cn-heyuan\\"}
	Parameters               map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	PredefinedParametersName *string                `json:"PredefinedParametersName,omitempty" xml:"PredefinedParametersName,omitempty"`
	// example:
	//
	// si-0e6fca6a51a54420****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s UpdateServiceInstanceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecRequest) SetClientToken(v string) *UpdateServiceInstanceSpecRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetEnableUserPrometheus(v bool) *UpdateServiceInstanceSpecRequest {
	s.EnableUserPrometheus = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetOperationName(v string) *UpdateServiceInstanceSpecRequest {
	s.OperationName = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetParameters(v map[string]interface{}) *UpdateServiceInstanceSpecRequest {
	s.Parameters = v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetPredefinedParametersName(v string) *UpdateServiceInstanceSpecRequest {
	s.PredefinedParametersName = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceSpecRequest {
	s.ServiceInstanceId = &v
	return s
}

type UpdateServiceInstanceSpecShrinkRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	EnableUserPrometheus *bool   `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	OperationName        *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// example:
	//
	// {\\"EcsInstanceParameter\\":\\"4vCPU 8GiB\\",\\"ZoneId\\":\\"cn-heyuan-a\\",\\"SystemDiskSize\\":50,\\"DataDiskSize\\":150,\\"InternetMaxBandwidthOut\\":2,\\"RegionId\\":\\"cn-heyuan\\"}
	ParametersShrink         *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	PredefinedParametersName *string `json:"PredefinedParametersName,omitempty" xml:"PredefinedParametersName,omitempty"`
	// example:
	//
	// si-0e6fca6a51a54420****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s UpdateServiceInstanceSpecShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetClientToken(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetEnableUserPrometheus(v bool) *UpdateServiceInstanceSpecShrinkRequest {
	s.EnableUserPrometheus = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetOperationName(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.OperationName = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetParametersShrink(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetPredefinedParametersName(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.PredefinedParametersName = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ServiceInstanceId = &v
	return s
}

type UpdateServiceInstanceSpecResponseBody struct {
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// DF0F666F-FBBC-55C3-A368-C955DE7B4839
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceInstanceSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecResponseBody) SetOrderId(v string) *UpdateServiceInstanceSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpdateServiceInstanceSpecResponseBody) SetRequestId(v string) *UpdateServiceInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceInstanceSpecResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceInstanceSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecResponse) SetHeaders(v map[string]*string) *UpdateServiceInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceInstanceSpecResponse) SetStatusCode(v int32) *UpdateServiceInstanceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceInstanceSpecResponse) SetBody(v *UpdateServiceInstanceSpecResponseBody) *UpdateServiceInstanceSpecResponse {
	s.Body = v
	return s
}

type UpgradeServiceInstanceRequest struct {
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun     *string                `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// si-5cbae874da0e47xxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// example:
	//
	// 2
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s UpgradeServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeServiceInstanceRequest) SetClientToken(v string) *UpgradeServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetDryRun(v string) *UpgradeServiceInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetParameters(v map[string]interface{}) *UpgradeServiceInstanceRequest {
	s.Parameters = v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetRegionId(v string) *UpgradeServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetServiceInstanceId(v string) *UpgradeServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetServiceVersion(v string) *UpgradeServiceInstanceRequest {
	s.ServiceVersion = &v
	return s
}

type UpgradeServiceInstanceShrinkRequest struct {
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun           *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// si-5cbae874da0e47xxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// example:
	//
	// 2
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s UpgradeServiceInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeServiceInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpgradeServiceInstanceShrinkRequest) SetClientToken(v string) *UpgradeServiceInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetDryRun(v string) *UpgradeServiceInstanceShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetParametersShrink(v string) *UpgradeServiceInstanceShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetRegionId(v string) *UpgradeServiceInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetServiceInstanceId(v string) *UpgradeServiceInstanceShrinkRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetServiceVersion(v string) *UpgradeServiceInstanceShrinkRequest {
	s.ServiceVersion = &v
	return s
}

type UpgradeServiceInstanceResponseBody struct {
	// example:
	//
	// F224E002-AB0E-5FD9-A87E-54AEE56F6CAE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// si-5cbae874da0e47xxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// example:
	//
	// Created
	Status                    *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UpgradeRequiredParameters []*string `json:"UpgradeRequiredParameters,omitempty" xml:"UpgradeRequiredParameters,omitempty" type:"Repeated"`
}

func (s UpgradeServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeServiceInstanceResponseBody) SetRequestId(v string) *UpgradeServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeServiceInstanceResponseBody) SetServiceInstanceId(v string) *UpgradeServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpgradeServiceInstanceResponseBody) SetStatus(v string) *UpgradeServiceInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *UpgradeServiceInstanceResponseBody) SetUpgradeRequiredParameters(v []*string) *UpgradeServiceInstanceResponseBody {
	s.UpgradeRequiredParameters = v
	return s
}

type UpgradeServiceInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeServiceInstanceResponse) SetHeaders(v map[string]*string) *UpgradeServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeServiceInstanceResponse) SetStatusCode(v int32) *UpgradeServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeServiceInstanceResponse) SetBody(v *UpgradeServiceInstanceResponseBody) *UpgradeServiceInstanceResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("computenestsupplier"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddServiceSharedAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddServiceSharedAccountsResponse
func (client *Client) AddServiceSharedAccountsWithOptions(request *AddServiceSharedAccountsRequest, runtime *util.RuntimeOptions) (_result *AddServiceSharedAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.SharedAccounts)) {
		query["SharedAccounts"] = request.SharedAccounts
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddServiceSharedAccounts"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddServiceSharedAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddServiceSharedAccountsRequest
//
// @return AddServiceSharedAccountsResponse
func (client *Client) AddServiceSharedAccounts(request *AddServiceSharedAccountsRequest) (_result *AddServiceSharedAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddServiceSharedAccountsResponse{}
	_body, _err := client.AddServiceSharedAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ContinueDeployServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinueDeployServiceInstanceResponse
func (client *Client) ContinueDeployServiceInstanceWithOptions(request *ContinueDeployServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *ContinueDeployServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ContinueDeployServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ContinueDeployServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ContinueDeployServiceInstanceRequest
//
// @return ContinueDeployServiceInstanceResponse
func (client *Client) ContinueDeployServiceInstance(request *ContinueDeployServiceInstanceRequest) (_result *ContinueDeployServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContinueDeployServiceInstanceResponse{}
	_body, _err := client.ContinueDeployServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建部署物
//
// @param tmpReq - CreateArtifactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateArtifactResponse
func (client *Client) CreateArtifactWithOptions(tmpReq *CreateArtifactRequest, runtime *util.RuntimeOptions) (_result *CreateArtifactResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateArtifactShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ArtifactProperty)) {
		request.ArtifactPropertyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ArtifactProperty, tea.String("ArtifactProperty"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactPropertyShrink)) {
		query["ArtifactProperty"] = request.ArtifactPropertyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactType)) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SupportRegionIds)) {
		query["SupportRegionIds"] = request.SupportRegionIds
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VersionName)) {
		query["VersionName"] = request.VersionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateArtifact"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建部署物
//
// @param request - CreateArtifactRequest
//
// @return CreateArtifactResponse
func (client *Client) CreateArtifact(request *CreateArtifactRequest) (_result *CreateArtifactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateArtifactResponse{}
	_body, _err := client.CreateArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建新服务版本
//
// @param request - CreateServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceResponse
func (client *Client) CreateServiceWithOptions(request *CreateServiceRequest, runtime *util.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmMetadata)) {
		query["AlarmMetadata"] = request.AlarmMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.ApprovalType)) {
		query["ApprovalType"] = request.ApprovalType
	}

	if !tea.BoolValue(util.IsUnset(request.BuildParameters)) {
		query["BuildParameters"] = request.BuildParameters
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeployMetadata)) {
		query["DeployMetadata"] = request.DeployMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.DeployType)) {
		query["DeployType"] = request.DeployType
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.IsSupportOperated)) {
		query["IsSupportOperated"] = request.IsSupportOperated
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseMetadata)) {
		query["LicenseMetadata"] = request.LicenseMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.LogMetadata)) {
		query["LogMetadata"] = request.LogMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.OperationMetadata)) {
		query["OperationMetadata"] = request.OperationMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyNames)) {
		query["PolicyNames"] = request.PolicyNames
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resellable)) {
		query["Resellable"] = request.Resellable
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInfo)) {
		query["ServiceInfo"] = request.ServiceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.ShareType)) {
		query["ShareType"] = request.ShareType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceServiceId)) {
		query["SourceServiceId"] = request.SourceServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceServiceVersion)) {
		query["SourceServiceVersion"] = request.SourceServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TenantType)) {
		query["TenantType"] = request.TenantType
	}

	if !tea.BoolValue(util.IsUnset(request.TrialDuration)) {
		query["TrialDuration"] = request.TrialDuration
	}

	if !tea.BoolValue(util.IsUnset(request.UpgradeMetadata)) {
		query["UpgradeMetadata"] = request.UpgradeMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.VersionName)) {
		query["VersionName"] = request.VersionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateService"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建新服务版本
//
// @param request - CreateServiceRequest
//
// @return CreateServiceResponse
func (client *Client) CreateService(request *CreateServiceRequest) (_result *CreateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceResponse{}
	_body, _err := client.CreateServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 商家侧创建服务实例
//
// @param tmpReq - CreateServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceInstanceResponse
func (client *Client) CreateServiceInstanceWithOptions(tmpReq *CreateServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateServiceInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SpecificationName)) {
		query["SpecificationName"] = request.SpecificationName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 商家侧创建服务实例
//
// @param request - CreateServiceInstanceRequest
//
// @return CreateServiceInstanceResponse
func (client *Client) CreateServiceInstance(request *CreateServiceInstanceRequest) (_result *CreateServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceInstanceResponse{}
	_body, _err := client.CreateServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除部署物
//
// @param request - DeleteArtifactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteArtifactResponse
func (client *Client) DeleteArtifactWithOptions(request *DeleteArtifactRequest, runtime *util.RuntimeOptions) (_result *DeleteArtifactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactVersion)) {
		query["ArtifactVersion"] = request.ArtifactVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteArtifact"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除部署物
//
// @param request - DeleteArtifactRequest
//
// @return DeleteArtifactResponse
func (client *Client) DeleteArtifact(request *DeleteArtifactRequest) (_result *DeleteArtifactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteArtifactResponse{}
	_body, _err := client.DeleteArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceResponse
func (client *Client) DeleteServiceWithOptions(request *DeleteServiceRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteService"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteServiceRequest
//
// @return DeleteServiceResponse
func (client *Client) DeleteService(request *DeleteServiceRequest) (_result *DeleteServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServiceResponse{}
	_body, _err := client.DeleteServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteServiceInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceInstancesResponse
func (client *Client) DeleteServiceInstancesWithOptions(request *DeleteServiceInstancesRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceInstances"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteServiceInstancesRequest
//
// @return DeleteServiceInstancesResponse
func (client *Client) DeleteServiceInstances(request *DeleteServiceInstancesRequest) (_result *DeleteServiceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServiceInstancesResponse{}
	_body, _err := client.DeleteServiceInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeployServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployServiceInstanceResponse
func (client *Client) DeployServiceInstanceWithOptions(request *DeployServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *DeployServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeployServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeployServiceInstanceRequest
//
// @return DeployServiceInstanceResponse
func (client *Client) DeployServiceInstance(request *DeployServiceInstanceRequest) (_result *DeployServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployServiceInstanceResponse{}
	_body, _err := client.DeployServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取部署物信息
//
// @param request - GetArtifactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactResponse
func (client *Client) GetArtifactWithOptions(request *GetArtifactRequest, runtime *util.RuntimeOptions) (_result *GetArtifactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactName)) {
		query["ArtifactName"] = request.ArtifactName
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactVersion)) {
		query["ArtifactVersion"] = request.ArtifactVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetArtifact"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取部署物信息
//
// @param request - GetArtifactRequest
//
// @return GetArtifactResponse
func (client *Client) GetArtifact(request *GetArtifactRequest) (_result *GetArtifactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetArtifactResponse{}
	_body, _err := client.GetArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取部署物仓库访问凭证
//
// @param request - GetArtifactRepositoryCredentialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactRepositoryCredentialsResponse
func (client *Client) GetArtifactRepositoryCredentialsWithOptions(request *GetArtifactRepositoryCredentialsRequest, runtime *util.RuntimeOptions) (_result *GetArtifactRepositoryCredentialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactType)) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !tea.BoolValue(util.IsUnset(request.DeployRegionId)) {
		query["DeployRegionId"] = request.DeployRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetArtifactRepositoryCredentials"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetArtifactRepositoryCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取部署物仓库访问凭证
//
// @param request - GetArtifactRepositoryCredentialsRequest
//
// @return GetArtifactRepositoryCredentialsResponse
func (client *Client) GetArtifactRepositoryCredentials(request *GetArtifactRepositoryCredentialsRequest) (_result *GetArtifactRepositoryCredentialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetArtifactRepositoryCredentialsResponse{}
	_body, _err := client.GetArtifactRepositoryCredentialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询服务详情
//
// @param request - GetServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceResponse
func (client *Client) GetServiceWithOptions(request *GetServiceRequest, runtime *util.RuntimeOptions) (_result *GetServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterAliUid)) {
		query["FilterAliUid"] = request.FilterAliUid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SharedAccountType)) {
		query["SharedAccountType"] = request.SharedAccountType
	}

	if !tea.BoolValue(util.IsUnset(request.ShowDetail)) {
		query["ShowDetail"] = request.ShowDetail
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetService"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务详情
//
// @param request - GetServiceRequest
//
// @return GetServiceResponse
func (client *Client) GetService(request *GetServiceRequest) (_result *GetServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceResponse{}
	_body, _err := client.GetServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 计算巢服务部署询价
//
// @param tmpReq - GetServiceEstimateCostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceEstimateCostResponse
func (client *Client) GetServiceEstimateCostWithOptions(tmpReq *GetServiceEstimateCostRequest, runtime *util.RuntimeOptions) (_result *GetServiceEstimateCostResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetServiceEstimateCostShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Commodity)) {
		request.CommodityShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Commodity, tea.String("Commodity"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityShrink)) {
		query["Commodity"] = request.CommodityShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SpecificationName)) {
		query["SpecificationName"] = request.SpecificationName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceEstimateCost"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceEstimateCostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 计算巢服务部署询价
//
// @param request - GetServiceEstimateCostRequest
//
// @return GetServiceEstimateCostResponse
func (client *Client) GetServiceEstimateCost(request *GetServiceEstimateCostRequest) (_result *GetServiceEstimateCostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceEstimateCostResponse{}
	_body, _err := client.GetServiceEstimateCostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceInstanceResponse
func (client *Client) GetServiceInstanceWithOptions(request *GetServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *GetServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetServiceInstanceRequest
//
// @return GetServiceInstanceResponse
func (client *Client) GetServiceInstance(request *GetServiceInstanceRequest) (_result *GetServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceInstanceResponse{}
	_body, _err := client.GetServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取ROS模板参数限制
//
// @param request - GetServiceTemplateParameterConstraintsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceTemplateParameterConstraintsResponse
func (client *Client) GetServiceTemplateParameterConstraintsWithOptions(request *GetServiceTemplateParameterConstraintsRequest, runtime *util.RuntimeOptions) (_result *GetServiceTemplateParameterConstraintsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeployRegionId)) {
		query["DeployRegionId"] = request.DeployRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePrivateVpcConnection)) {
		query["EnablePrivateVpcConnection"] = request.EnablePrivateVpcConnection
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceTemplateParameterConstraints"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceTemplateParameterConstraintsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取ROS模板参数限制
//
// @param request - GetServiceTemplateParameterConstraintsRequest
//
// @return GetServiceTemplateParameterConstraintsResponse
func (client *Client) GetServiceTemplateParameterConstraints(request *GetServiceTemplateParameterConstraintsRequest) (_result *GetServiceTemplateParameterConstraintsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceTemplateParameterConstraintsResponse{}
	_body, _err := client.GetServiceTemplateParameterConstraintsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetUploadCredentialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUploadCredentialsResponse
func (client *Client) GetUploadCredentialsWithOptions(request *GetUploadCredentialsRequest, runtime *util.RuntimeOptions) (_result *GetUploadCredentialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUploadCredentials"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUploadCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetUploadCredentialsRequest
//
// @return GetUploadCredentialsResponse
func (client *Client) GetUploadCredentials(request *GetUploadCredentialsRequest) (_result *GetUploadCredentialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUploadCredentialsResponse{}
	_body, _err := client.GetUploadCredentialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 展示部署物
//
// @param request - ListAcrImageRepositoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAcrImageRepositoriesResponse
func (client *Client) ListAcrImageRepositoriesWithOptions(request *ListAcrImageRepositoriesRequest, runtime *util.RuntimeOptions) (_result *ListAcrImageRepositoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactType)) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAcrImageRepositories"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAcrImageRepositoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 展示部署物
//
// @param request - ListAcrImageRepositoriesRequest
//
// @return ListAcrImageRepositoriesResponse
func (client *Client) ListAcrImageRepositories(request *ListAcrImageRepositoriesRequest) (_result *ListAcrImageRepositoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAcrImageRepositoriesResponse{}
	_body, _err := client.ListAcrImageRepositoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 展示部署物
//
// @param request - ListAcrImageTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAcrImageTagsResponse
func (client *Client) ListAcrImageTagsWithOptions(request *ListAcrImageTagsRequest, runtime *util.RuntimeOptions) (_result *ListAcrImageTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactType)) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAcrImageTags"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAcrImageTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 展示部署物
//
// @param request - ListAcrImageTagsRequest
//
// @return ListAcrImageTagsResponse
func (client *Client) ListAcrImageTags(request *ListAcrImageTagsRequest) (_result *ListAcrImageTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAcrImageTagsResponse{}
	_body, _err := client.ListAcrImageTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 展示部署物版本
//
// @param request - ListArtifactVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListArtifactVersionsResponse
func (client *Client) ListArtifactVersionsWithOptions(request *ListArtifactVersionsRequest, runtime *util.RuntimeOptions) (_result *ListArtifactVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListArtifactVersions"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListArtifactVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 展示部署物版本
//
// @param request - ListArtifactVersionsRequest
//
// @return ListArtifactVersionsResponse
func (client *Client) ListArtifactVersions(request *ListArtifactVersionsRequest) (_result *ListArtifactVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListArtifactVersionsResponse{}
	_body, _err := client.ListArtifactVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 展示部署物
//
// @param request - ListArtifactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListArtifactsResponse
func (client *Client) ListArtifactsWithOptions(request *ListArtifactsRequest, runtime *util.RuntimeOptions) (_result *ListArtifactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListArtifacts"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListArtifactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 展示部署物
//
// @param request - ListArtifactsRequest
//
// @return ListArtifactsResponse
func (client *Client) ListArtifacts(request *ListArtifactsRequest) (_result *ListArtifactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListArtifactsResponse{}
	_body, _err := client.ListArtifactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListServiceInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstancesResponse
func (client *Client) ListServiceInstancesWithOptions(request *ListServiceInstancesRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowDeleted)) {
		query["ShowDeleted"] = request.ShowDeleted
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstances"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListServiceInstancesRequest
//
// @return ListServiceInstancesResponse
func (client *Client) ListServiceInstances(request *ListServiceInstancesRequest) (_result *ListServiceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.ListServiceInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 商家获取服务使用申请接口
//
// @param request - ListServiceUsagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceUsagesResponse
func (client *Client) ListServiceUsagesWithOptions(request *ListServiceUsagesRequest, runtime *util.RuntimeOptions) (_result *ListServiceUsagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SupplierRole)) {
		query["SupplierRole"] = request.SupplierRole
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceUsages"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceUsagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 商家获取服务使用申请接口
//
// @param request - ListServiceUsagesRequest
//
// @return ListServiceUsagesResponse
func (client *Client) ListServiceUsages(request *ListServiceUsagesRequest) (_result *ListServiceUsagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceUsagesResponse{}
	_body, _err := client.ListServiceUsagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询服务
//
// @param request - ListServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServicesResponse
func (client *Client) ListServicesWithOptions(request *ListServicesRequest, runtime *util.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllVersions)) {
		query["AllVersions"] = request.AllVersions
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServices"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务
//
// @param request - ListServicesRequest
//
// @return ListServicesResponse
func (client *Client) ListServices(request *ListServicesRequest) (_result *ListServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改服务实例资源
//
// @param request - ModifyServiceInstanceResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyServiceInstanceResourcesResponse
func (client *Client) ModifyServiceInstanceResourcesWithOptions(request *ModifyServiceInstanceResourcesRequest, runtime *util.RuntimeOptions) (_result *ModifyServiceInstanceResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		query["Resources"] = request.Resources
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceResourcesAction)) {
		query["ServiceInstanceResourcesAction"] = request.ServiceInstanceResourcesAction
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyServiceInstanceResources"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyServiceInstanceResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改服务实例资源
//
// @param request - ModifyServiceInstanceResourcesRequest
//
// @return ModifyServiceInstanceResourcesResponse
func (client *Client) ModifyServiceInstanceResources(request *ModifyServiceInstanceResourcesRequest) (_result *ModifyServiceInstanceResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyServiceInstanceResourcesResponse{}
	_body, _err := client.ModifyServiceInstanceResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - PushMeteringDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushMeteringDataResponse
func (client *Client) PushMeteringDataWithOptions(request *PushMeteringDataRequest, runtime *util.RuntimeOptions) (_result *PushMeteringDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Metering)) {
		query["Metering"] = request.Metering
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PushMeteringData"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushMeteringDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - PushMeteringDataRequest
//
// @return PushMeteringDataResponse
func (client *Client) PushMeteringData(request *PushMeteringDataRequest) (_result *PushMeteringDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushMeteringDataResponse{}
	_body, _err := client.PushMeteringDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RegisterServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterServiceResponse
func (client *Client) RegisterServiceWithOptions(request *RegisterServiceRequest, runtime *util.RuntimeOptions) (_result *RegisterServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterService"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - RegisterServiceRequest
//
// @return RegisterServiceResponse
func (client *Client) RegisterService(request *RegisterServiceRequest) (_result *RegisterServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterServiceResponse{}
	_body, _err := client.RegisterServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布部署物
//
// @param request - ReleaseArtifactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseArtifactResponse
func (client *Client) ReleaseArtifactWithOptions(request *ReleaseArtifactRequest, runtime *util.RuntimeOptions) (_result *ReleaseArtifactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseArtifact"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布部署物
//
// @param request - ReleaseArtifactRequest
//
// @return ReleaseArtifactResponse
func (client *Client) ReleaseArtifact(request *ReleaseArtifactRequest) (_result *ReleaseArtifactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseArtifactResponse{}
	_body, _err := client.ReleaseArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重启服务实例
//
// @param request - RestartServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartServiceInstanceResponse
func (client *Client) RestartServiceInstanceWithOptions(request *RestartServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重启服务实例
//
// @param request - RestartServiceInstanceRequest
//
// @return RestartServiceInstanceResponse
func (client *Client) RestartServiceInstance(request *RestartServiceInstanceRequest) (_result *RestartServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartServiceInstanceResponse{}
	_body, _err := client.RestartServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动服务实例
//
// @param request - StartServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartServiceInstanceResponse
func (client *Client) StartServiceInstanceWithOptions(request *StartServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *StartServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动服务实例
//
// @param request - StartServiceInstanceRequest
//
// @return StartServiceInstanceResponse
func (client *Client) StartServiceInstance(request *StartServiceInstanceRequest) (_result *StartServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartServiceInstanceResponse{}
	_body, _err := client.StartServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止服务实例
//
// @param request - StopServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopServiceInstanceResponse
func (client *Client) StopServiceInstanceWithOptions(request *StopServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *StopServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止服务实例
//
// @param request - StopServiceInstanceRequest
//
// @return StopServiceInstanceResponse
func (client *Client) StopServiceInstance(request *StopServiceInstanceRequest) (_result *StopServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopServiceInstanceResponse{}
	_body, _err := client.StopServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新部署物
//
// @param tmpReq - UpdateArtifactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateArtifactResponse
func (client *Client) UpdateArtifactWithOptions(tmpReq *UpdateArtifactRequest, runtime *util.RuntimeOptions) (_result *UpdateArtifactResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateArtifactShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ArtifactProperty)) {
		request.ArtifactPropertyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ArtifactProperty, tea.String("ArtifactProperty"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactPropertyShrink)) {
		query["ArtifactProperty"] = request.ArtifactPropertyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.SupportRegionIds)) {
		query["SupportRegionIds"] = request.SupportRegionIds
	}

	if !tea.BoolValue(util.IsUnset(request.VersionName)) {
		query["VersionName"] = request.VersionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateArtifact"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新部署物
//
// @param request - UpdateArtifactRequest
//
// @return UpdateArtifactResponse
func (client *Client) UpdateArtifact(request *UpdateArtifactRequest) (_result *UpdateArtifactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateArtifactResponse{}
	_body, _err := client.UpdateArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - UpdateServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceResponse
func (client *Client) UpdateServiceWithOptions(tmpReq *UpdateServiceRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateServiceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UpdateOption)) {
		request.UpdateOptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateOption, tea.String("UpdateOption"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmMetadata)) {
		query["AlarmMetadata"] = request.AlarmMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeployMetadata)) {
		query["DeployMetadata"] = request.DeployMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.DeployType)) {
		query["DeployType"] = request.DeployType
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.IsSupportOperated)) {
		query["IsSupportOperated"] = request.IsSupportOperated
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseMetadata)) {
		query["LicenseMetadata"] = request.LicenseMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.LogMetadata)) {
		query["LogMetadata"] = request.LogMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.OperationMetadata)) {
		query["OperationMetadata"] = request.OperationMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyNames)) {
		query["PolicyNames"] = request.PolicyNames
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resellable)) {
		query["Resellable"] = request.Resellable
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInfo)) {
		query["ServiceInfo"] = request.ServiceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TenantType)) {
		query["TenantType"] = request.TenantType
	}

	if !tea.BoolValue(util.IsUnset(request.TrialDuration)) {
		query["TrialDuration"] = request.TrialDuration
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateOptionShrink)) {
		query["UpdateOption"] = request.UpdateOptionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UpgradeMetadata)) {
		query["UpgradeMetadata"] = request.UpgradeMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.VersionName)) {
		query["VersionName"] = request.VersionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateService"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateServiceRequest
//
// @return UpdateServiceResponse
func (client *Client) UpdateService(request *UpdateServiceRequest) (_result *UpdateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServiceResponse{}
	_body, _err := client.UpdateServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新服务实例属性
//
// @param request - UpdateServiceInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceInstanceAttributeResponse
func (client *Client) UpdateServiceInstanceAttributeWithOptions(request *UpdateServiceInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceInstanceAttribute"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新服务实例属性
//
// @param request - UpdateServiceInstanceAttributeRequest
//
// @return UpdateServiceInstanceAttributeResponse
func (client *Client) UpdateServiceInstanceAttribute(request *UpdateServiceInstanceAttributeRequest) (_result *UpdateServiceInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServiceInstanceAttributeResponse{}
	_body, _err := client.UpdateServiceInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变配服务实例
//
// @param tmpReq - UpdateServiceInstanceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceInstanceSpecResponse
func (client *Client) UpdateServiceInstanceSpecWithOptions(tmpReq *UpdateServiceInstanceSpecRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceInstanceSpecResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateServiceInstanceSpecShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EnableUserPrometheus)) {
		query["EnableUserPrometheus"] = request.EnableUserPrometheus
	}

	if !tea.BoolValue(util.IsUnset(request.OperationName)) {
		query["OperationName"] = request.OperationName
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PredefinedParametersName)) {
		query["PredefinedParametersName"] = request.PredefinedParametersName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceInstanceSpec"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceInstanceSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变配服务实例
//
// @param request - UpdateServiceInstanceSpecRequest
//
// @return UpdateServiceInstanceSpecResponse
func (client *Client) UpdateServiceInstanceSpec(request *UpdateServiceInstanceSpecRequest) (_result *UpdateServiceInstanceSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServiceInstanceSpecResponse{}
	_body, _err := client.UpdateServiceInstanceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - UpgradeServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeServiceInstanceResponse
func (client *Client) UpgradeServiceInstanceWithOptions(tmpReq *UpgradeServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *UpgradeServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpgradeServiceInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpgradeServiceInstanceRequest
//
// @return UpgradeServiceInstanceResponse
func (client *Client) UpgradeServiceInstance(request *UpgradeServiceInstanceRequest) (_result *UpgradeServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeServiceInstanceResponse{}
	_body, _err := client.UpgradeServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
