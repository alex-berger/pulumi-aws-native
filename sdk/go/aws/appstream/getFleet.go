// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appstream

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppStream::Fleet
func LookupFleet(ctx *pulumi.Context, args *LookupFleetArgs, opts ...pulumi.InvokeOption) (*LookupFleetResult, error) {
	var rv LookupFleetResult
	err := ctx.Invoke("aws-native:appstream:getFleet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFleetArgs struct {
	Id string `pulumi:"id"`
}

type LookupFleetResult struct {
	ComputeCapacity                *FleetComputeCapacity `pulumi:"computeCapacity"`
	Description                    *string               `pulumi:"description"`
	DisconnectTimeoutInSeconds     *int                  `pulumi:"disconnectTimeoutInSeconds"`
	DisplayName                    *string               `pulumi:"displayName"`
	DomainJoinInfo                 *FleetDomainJoinInfo  `pulumi:"domainJoinInfo"`
	EnableDefaultInternetAccess    *bool                 `pulumi:"enableDefaultInternetAccess"`
	IamRoleArn                     *string               `pulumi:"iamRoleArn"`
	Id                             *string               `pulumi:"id"`
	IdleDisconnectTimeoutInSeconds *int                  `pulumi:"idleDisconnectTimeoutInSeconds"`
	ImageArn                       *string               `pulumi:"imageArn"`
	ImageName                      *string               `pulumi:"imageName"`
	InstanceType                   *string               `pulumi:"instanceType"`
	MaxConcurrentSessions          *int                  `pulumi:"maxConcurrentSessions"`
	MaxUserDurationInSeconds       *int                  `pulumi:"maxUserDurationInSeconds"`
	Platform                       *string               `pulumi:"platform"`
	SessionScriptS3Location        *FleetS3Location      `pulumi:"sessionScriptS3Location"`
	StreamView                     *string               `pulumi:"streamView"`
	Tags                           []FleetTag            `pulumi:"tags"`
	UsbDeviceFilterStrings         []string              `pulumi:"usbDeviceFilterStrings"`
	VpcConfig                      *FleetVpcConfig       `pulumi:"vpcConfig"`
}

func LookupFleetOutput(ctx *pulumi.Context, args LookupFleetOutputArgs, opts ...pulumi.InvokeOption) LookupFleetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFleetResult, error) {
			args := v.(LookupFleetArgs)
			r, err := LookupFleet(ctx, &args, opts...)
			var s LookupFleetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFleetResultOutput)
}

type LookupFleetOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupFleetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetArgs)(nil)).Elem()
}

type LookupFleetResultOutput struct{ *pulumi.OutputState }

func (LookupFleetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetResult)(nil)).Elem()
}

func (o LookupFleetResultOutput) ToLookupFleetResultOutput() LookupFleetResultOutput {
	return o
}

func (o LookupFleetResultOutput) ToLookupFleetResultOutputWithContext(ctx context.Context) LookupFleetResultOutput {
	return o
}

func (o LookupFleetResultOutput) ComputeCapacity() FleetComputeCapacityPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *FleetComputeCapacity { return v.ComputeCapacity }).(FleetComputeCapacityPtrOutput)
}

func (o LookupFleetResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFleetResultOutput) DisconnectTimeoutInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *int { return v.DisconnectTimeoutInSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupFleetResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupFleetResultOutput) DomainJoinInfo() FleetDomainJoinInfoPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *FleetDomainJoinInfo { return v.DomainJoinInfo }).(FleetDomainJoinInfoPtrOutput)
}

func (o LookupFleetResultOutput) EnableDefaultInternetAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *bool { return v.EnableDefaultInternetAccess }).(pulumi.BoolPtrOutput)
}

func (o LookupFleetResultOutput) IamRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *string { return v.IamRoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupFleetResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupFleetResultOutput) IdleDisconnectTimeoutInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *int { return v.IdleDisconnectTimeoutInSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupFleetResultOutput) ImageArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *string { return v.ImageArn }).(pulumi.StringPtrOutput)
}

func (o LookupFleetResultOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *string { return v.ImageName }).(pulumi.StringPtrOutput)
}

func (o LookupFleetResultOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o LookupFleetResultOutput) MaxConcurrentSessions() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *int { return v.MaxConcurrentSessions }).(pulumi.IntPtrOutput)
}

func (o LookupFleetResultOutput) MaxUserDurationInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *int { return v.MaxUserDurationInSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupFleetResultOutput) Platform() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *string { return v.Platform }).(pulumi.StringPtrOutput)
}

func (o LookupFleetResultOutput) SessionScriptS3Location() FleetS3LocationPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *FleetS3Location { return v.SessionScriptS3Location }).(FleetS3LocationPtrOutput)
}

func (o LookupFleetResultOutput) StreamView() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *string { return v.StreamView }).(pulumi.StringPtrOutput)
}

func (o LookupFleetResultOutput) Tags() FleetTagArrayOutput {
	return o.ApplyT(func(v LookupFleetResult) []FleetTag { return v.Tags }).(FleetTagArrayOutput)
}

func (o LookupFleetResultOutput) UsbDeviceFilterStrings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFleetResult) []string { return v.UsbDeviceFilterStrings }).(pulumi.StringArrayOutput)
}

func (o LookupFleetResultOutput) VpcConfig() FleetVpcConfigPtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *FleetVpcConfig { return v.VpcConfig }).(FleetVpcConfigPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFleetResultOutput{})
}
