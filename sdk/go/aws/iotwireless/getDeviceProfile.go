// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iotwireless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Device Profile's resource schema demonstrating some basic constructs and validation rules.
func LookupDeviceProfile(ctx *pulumi.Context, args *LookupDeviceProfileArgs, opts ...pulumi.InvokeOption) (*LookupDeviceProfileResult, error) {
	var rv LookupDeviceProfileResult
	err := ctx.Invoke("aws-native:iotwireless:getDeviceProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeviceProfileArgs struct {
	// Service profile Id. Returned after successful create.
	Id string `pulumi:"id"`
}

type LookupDeviceProfileResult struct {
	// Service profile Arn. Returned after successful create.
	Arn *string `pulumi:"arn"`
	// Service profile Id. Returned after successful create.
	Id *string `pulumi:"id"`
	// LoRaWANDeviceProfile supports all LoRa specific attributes for service profile for CreateDeviceProfile operation
	LoRaWAN *DeviceProfileLoRaWANDeviceProfile `pulumi:"loRaWAN"`
	// Name of service profile
	Name *string `pulumi:"name"`
	// A list of key-value pairs that contain metadata for the device profile.
	Tags []DeviceProfileTag `pulumi:"tags"`
}

func LookupDeviceProfileOutput(ctx *pulumi.Context, args LookupDeviceProfileOutputArgs, opts ...pulumi.InvokeOption) LookupDeviceProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeviceProfileResult, error) {
			args := v.(LookupDeviceProfileArgs)
			r, err := LookupDeviceProfile(ctx, &args, opts...)
			var s LookupDeviceProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeviceProfileResultOutput)
}

type LookupDeviceProfileOutputArgs struct {
	// Service profile Id. Returned after successful create.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupDeviceProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceProfileArgs)(nil)).Elem()
}

type LookupDeviceProfileResultOutput struct{ *pulumi.OutputState }

func (LookupDeviceProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceProfileResult)(nil)).Elem()
}

func (o LookupDeviceProfileResultOutput) ToLookupDeviceProfileResultOutput() LookupDeviceProfileResultOutput {
	return o
}

func (o LookupDeviceProfileResultOutput) ToLookupDeviceProfileResultOutputWithContext(ctx context.Context) LookupDeviceProfileResultOutput {
	return o
}

// Service profile Arn. Returned after successful create.
func (o LookupDeviceProfileResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeviceProfileResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Service profile Id. Returned after successful create.
func (o LookupDeviceProfileResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeviceProfileResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// LoRaWANDeviceProfile supports all LoRa specific attributes for service profile for CreateDeviceProfile operation
func (o LookupDeviceProfileResultOutput) LoRaWAN() DeviceProfileLoRaWANDeviceProfilePtrOutput {
	return o.ApplyT(func(v LookupDeviceProfileResult) *DeviceProfileLoRaWANDeviceProfile { return v.LoRaWAN }).(DeviceProfileLoRaWANDeviceProfilePtrOutput)
}

// Name of service profile
func (o LookupDeviceProfileResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeviceProfileResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A list of key-value pairs that contain metadata for the device profile.
func (o LookupDeviceProfileResultOutput) Tags() DeviceProfileTagArrayOutput {
	return o.ApplyT(func(v LookupDeviceProfileResult) []DeviceProfileTag { return v.Tags }).(DeviceProfileTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeviceProfileResultOutput{})
}
