// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::SageMaker::DeviceFleet
func LookupDeviceFleet(ctx *pulumi.Context, args *LookupDeviceFleetArgs, opts ...pulumi.InvokeOption) (*LookupDeviceFleetResult, error) {
	var rv LookupDeviceFleetResult
	err := ctx.Invoke("aws-native:sagemaker:getDeviceFleet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeviceFleetArgs struct {
	// The name of the edge device fleet
	DeviceFleetName string `pulumi:"deviceFleetName"`
}

type LookupDeviceFleetResult struct {
	// Description for the edge device fleet
	Description *string `pulumi:"description"`
	// S3 bucket and an ecryption key id (if available) to store outputs for the fleet
	OutputConfig *DeviceFleetEdgeOutputConfig `pulumi:"outputConfig"`
	// Role associated with the device fleet
	RoleArn *string `pulumi:"roleArn"`
	// Associate tags with the resource
	Tags []DeviceFleetTag `pulumi:"tags"`
}

func LookupDeviceFleetOutput(ctx *pulumi.Context, args LookupDeviceFleetOutputArgs, opts ...pulumi.InvokeOption) LookupDeviceFleetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeviceFleetResult, error) {
			args := v.(LookupDeviceFleetArgs)
			r, err := LookupDeviceFleet(ctx, &args, opts...)
			var s LookupDeviceFleetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeviceFleetResultOutput)
}

type LookupDeviceFleetOutputArgs struct {
	// The name of the edge device fleet
	DeviceFleetName pulumi.StringInput `pulumi:"deviceFleetName"`
}

func (LookupDeviceFleetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceFleetArgs)(nil)).Elem()
}

type LookupDeviceFleetResultOutput struct{ *pulumi.OutputState }

func (LookupDeviceFleetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceFleetResult)(nil)).Elem()
}

func (o LookupDeviceFleetResultOutput) ToLookupDeviceFleetResultOutput() LookupDeviceFleetResultOutput {
	return o
}

func (o LookupDeviceFleetResultOutput) ToLookupDeviceFleetResultOutputWithContext(ctx context.Context) LookupDeviceFleetResultOutput {
	return o
}

// Description for the edge device fleet
func (o LookupDeviceFleetResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeviceFleetResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// S3 bucket and an ecryption key id (if available) to store outputs for the fleet
func (o LookupDeviceFleetResultOutput) OutputConfig() DeviceFleetEdgeOutputConfigPtrOutput {
	return o.ApplyT(func(v LookupDeviceFleetResult) *DeviceFleetEdgeOutputConfig { return v.OutputConfig }).(DeviceFleetEdgeOutputConfigPtrOutput)
}

// Role associated with the device fleet
func (o LookupDeviceFleetResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeviceFleetResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// Associate tags with the resource
func (o LookupDeviceFleetResultOutput) Tags() DeviceFleetTagArrayOutput {
	return o.ApplyT(func(v LookupDeviceFleetResult) []DeviceFleetTag { return v.Tags }).(DeviceFleetTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeviceFleetResultOutput{})
}
