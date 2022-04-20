// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iotwireless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create and manage wireless gateways, including LoRa gateways.
func LookupWirelessDevice(ctx *pulumi.Context, args *LookupWirelessDeviceArgs, opts ...pulumi.InvokeOption) (*LookupWirelessDeviceResult, error) {
	var rv LookupWirelessDeviceResult
	err := ctx.Invoke("aws-native:iotwireless:getWirelessDevice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWirelessDeviceArgs struct {
	// Wireless device Id. Returned after successful create.
	Id string `pulumi:"id"`
}

type LookupWirelessDeviceResult struct {
	// Wireless device arn. Returned after successful create.
	Arn *string `pulumi:"arn"`
	// Wireless device description
	Description *string `pulumi:"description"`
	// Wireless device destination name
	DestinationName *string `pulumi:"destinationName"`
	// Wireless device Id. Returned after successful create.
	Id *string `pulumi:"id"`
	// The date and time when the most recent uplink was received.
	LastUplinkReceivedAt *string `pulumi:"lastUplinkReceivedAt"`
	// The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Device.
	LoRaWAN *WirelessDeviceLoRaWANDevice `pulumi:"loRaWAN"`
	// Wireless device name
	Name *string `pulumi:"name"`
	// A list of key-value pairs that contain metadata for the device. Currently not supported, will not create if tags are passed.
	Tags []WirelessDeviceTag `pulumi:"tags"`
	// Thing arn. Passed into update to associate Thing with Wireless device.
	ThingArn *string `pulumi:"thingArn"`
	// Thing Arn. If there is a Thing created, this can be returned with a Get call.
	ThingName *string `pulumi:"thingName"`
	// Wireless device type, currently only Sidewalk and LoRa
	Type *WirelessDeviceType `pulumi:"type"`
}

func LookupWirelessDeviceOutput(ctx *pulumi.Context, args LookupWirelessDeviceOutputArgs, opts ...pulumi.InvokeOption) LookupWirelessDeviceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWirelessDeviceResult, error) {
			args := v.(LookupWirelessDeviceArgs)
			r, err := LookupWirelessDevice(ctx, &args, opts...)
			var s LookupWirelessDeviceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWirelessDeviceResultOutput)
}

type LookupWirelessDeviceOutputArgs struct {
	// Wireless device Id. Returned after successful create.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupWirelessDeviceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWirelessDeviceArgs)(nil)).Elem()
}

type LookupWirelessDeviceResultOutput struct{ *pulumi.OutputState }

func (LookupWirelessDeviceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWirelessDeviceResult)(nil)).Elem()
}

func (o LookupWirelessDeviceResultOutput) ToLookupWirelessDeviceResultOutput() LookupWirelessDeviceResultOutput {
	return o
}

func (o LookupWirelessDeviceResultOutput) ToLookupWirelessDeviceResultOutputWithContext(ctx context.Context) LookupWirelessDeviceResultOutput {
	return o
}

// Wireless device arn. Returned after successful create.
func (o LookupWirelessDeviceResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWirelessDeviceResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Wireless device description
func (o LookupWirelessDeviceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWirelessDeviceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Wireless device destination name
func (o LookupWirelessDeviceResultOutput) DestinationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWirelessDeviceResult) *string { return v.DestinationName }).(pulumi.StringPtrOutput)
}

// Wireless device Id. Returned after successful create.
func (o LookupWirelessDeviceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWirelessDeviceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The date and time when the most recent uplink was received.
func (o LookupWirelessDeviceResultOutput) LastUplinkReceivedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWirelessDeviceResult) *string { return v.LastUplinkReceivedAt }).(pulumi.StringPtrOutput)
}

// The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Device.
func (o LookupWirelessDeviceResultOutput) LoRaWAN() WirelessDeviceLoRaWANDevicePtrOutput {
	return o.ApplyT(func(v LookupWirelessDeviceResult) *WirelessDeviceLoRaWANDevice { return v.LoRaWAN }).(WirelessDeviceLoRaWANDevicePtrOutput)
}

// Wireless device name
func (o LookupWirelessDeviceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWirelessDeviceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A list of key-value pairs that contain metadata for the device. Currently not supported, will not create if tags are passed.
func (o LookupWirelessDeviceResultOutput) Tags() WirelessDeviceTagArrayOutput {
	return o.ApplyT(func(v LookupWirelessDeviceResult) []WirelessDeviceTag { return v.Tags }).(WirelessDeviceTagArrayOutput)
}

// Thing arn. Passed into update to associate Thing with Wireless device.
func (o LookupWirelessDeviceResultOutput) ThingArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWirelessDeviceResult) *string { return v.ThingArn }).(pulumi.StringPtrOutput)
}

// Thing Arn. If there is a Thing created, this can be returned with a Get call.
func (o LookupWirelessDeviceResultOutput) ThingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWirelessDeviceResult) *string { return v.ThingName }).(pulumi.StringPtrOutput)
}

// Wireless device type, currently only Sidewalk and LoRa
func (o LookupWirelessDeviceResultOutput) Type() WirelessDeviceTypePtrOutput {
	return o.ApplyT(func(v LookupWirelessDeviceResult) *WirelessDeviceType { return v.Type }).(WirelessDeviceTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWirelessDeviceResultOutput{})
}
