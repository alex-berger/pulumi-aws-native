// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotwireless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Device Profile's resource schema demonstrating some basic constructs and validation rules.
type DeviceProfile struct {
	pulumi.CustomResourceState

	// Service profile Arn. Returned after successful create.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// LoRaWANDeviceProfile supports all LoRa specific attributes for service profile for CreateDeviceProfile operation
	LoRaWAN DeviceProfileLoRaWANDeviceProfilePtrOutput `pulumi:"loRaWAN"`
	// Name of service profile
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// A list of key-value pairs that contain metadata for the device profile.
	Tags DeviceProfileTagArrayOutput `pulumi:"tags"`
}

// NewDeviceProfile registers a new resource with the given unique name, arguments, and options.
func NewDeviceProfile(ctx *pulumi.Context,
	name string, args *DeviceProfileArgs, opts ...pulumi.ResourceOption) (*DeviceProfile, error) {
	if args == nil {
		args = &DeviceProfileArgs{}
	}

	var resource DeviceProfile
	err := ctx.RegisterResource("aws-native:iotwireless:DeviceProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeviceProfile gets an existing DeviceProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceProfileState, opts ...pulumi.ResourceOption) (*DeviceProfile, error) {
	var resource DeviceProfile
	err := ctx.ReadResource("aws-native:iotwireless:DeviceProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeviceProfile resources.
type deviceProfileState struct {
}

type DeviceProfileState struct {
}

func (DeviceProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceProfileState)(nil)).Elem()
}

type deviceProfileArgs struct {
	// LoRaWANDeviceProfile supports all LoRa specific attributes for service profile for CreateDeviceProfile operation
	LoRaWAN *DeviceProfileLoRaWANDeviceProfile `pulumi:"loRaWAN"`
	// Name of service profile
	Name *string `pulumi:"name"`
	// A list of key-value pairs that contain metadata for the device profile.
	Tags []DeviceProfileTag `pulumi:"tags"`
}

// The set of arguments for constructing a DeviceProfile resource.
type DeviceProfileArgs struct {
	// LoRaWANDeviceProfile supports all LoRa specific attributes for service profile for CreateDeviceProfile operation
	LoRaWAN DeviceProfileLoRaWANDeviceProfilePtrInput
	// Name of service profile
	Name pulumi.StringPtrInput
	// A list of key-value pairs that contain metadata for the device profile.
	Tags DeviceProfileTagArrayInput
}

func (DeviceProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceProfileArgs)(nil)).Elem()
}

type DeviceProfileInput interface {
	pulumi.Input

	ToDeviceProfileOutput() DeviceProfileOutput
	ToDeviceProfileOutputWithContext(ctx context.Context) DeviceProfileOutput
}

func (*DeviceProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceProfile)(nil)).Elem()
}

func (i *DeviceProfile) ToDeviceProfileOutput() DeviceProfileOutput {
	return i.ToDeviceProfileOutputWithContext(context.Background())
}

func (i *DeviceProfile) ToDeviceProfileOutputWithContext(ctx context.Context) DeviceProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceProfileOutput)
}

type DeviceProfileOutput struct{ *pulumi.OutputState }

func (DeviceProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceProfile)(nil)).Elem()
}

func (o DeviceProfileOutput) ToDeviceProfileOutput() DeviceProfileOutput {
	return o
}

func (o DeviceProfileOutput) ToDeviceProfileOutputWithContext(ctx context.Context) DeviceProfileOutput {
	return o
}

// Service profile Arn. Returned after successful create.
func (o DeviceProfileOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DeviceProfile) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// LoRaWANDeviceProfile supports all LoRa specific attributes for service profile for CreateDeviceProfile operation
func (o DeviceProfileOutput) LoRaWAN() DeviceProfileLoRaWANDeviceProfilePtrOutput {
	return o.ApplyT(func(v *DeviceProfile) DeviceProfileLoRaWANDeviceProfilePtrOutput { return v.LoRaWAN }).(DeviceProfileLoRaWANDeviceProfilePtrOutput)
}

// Name of service profile
func (o DeviceProfileOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeviceProfile) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// A list of key-value pairs that contain metadata for the device profile.
func (o DeviceProfileOutput) Tags() DeviceProfileTagArrayOutput {
	return o.ApplyT(func(v *DeviceProfile) DeviceProfileTagArrayOutput { return v.Tags }).(DeviceProfileTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceProfileInput)(nil)).Elem(), &DeviceProfile{})
	pulumi.RegisterOutputType(DeviceProfileOutput{})
}
