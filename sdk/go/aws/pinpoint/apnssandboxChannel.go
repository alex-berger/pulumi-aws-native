// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Pinpoint::APNSSandboxChannel
//
// Deprecated: APNSSandboxChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type APNSSandboxChannel struct {
	pulumi.CustomResourceState

	ApplicationId               pulumi.StringOutput    `pulumi:"applicationId"`
	BundleId                    pulumi.StringPtrOutput `pulumi:"bundleId"`
	Certificate                 pulumi.StringPtrOutput `pulumi:"certificate"`
	DefaultAuthenticationMethod pulumi.StringPtrOutput `pulumi:"defaultAuthenticationMethod"`
	Enabled                     pulumi.BoolPtrOutput   `pulumi:"enabled"`
	PrivateKey                  pulumi.StringPtrOutput `pulumi:"privateKey"`
	TeamId                      pulumi.StringPtrOutput `pulumi:"teamId"`
	TokenKey                    pulumi.StringPtrOutput `pulumi:"tokenKey"`
	TokenKeyId                  pulumi.StringPtrOutput `pulumi:"tokenKeyId"`
}

// NewAPNSSandboxChannel registers a new resource with the given unique name, arguments, and options.
func NewAPNSSandboxChannel(ctx *pulumi.Context,
	name string, args *APNSSandboxChannelArgs, opts ...pulumi.ResourceOption) (*APNSSandboxChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	var resource APNSSandboxChannel
	err := ctx.RegisterResource("aws-native:pinpoint:APNSSandboxChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAPNSSandboxChannel gets an existing APNSSandboxChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAPNSSandboxChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *APNSSandboxChannelState, opts ...pulumi.ResourceOption) (*APNSSandboxChannel, error) {
	var resource APNSSandboxChannel
	err := ctx.ReadResource("aws-native:pinpoint:APNSSandboxChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering APNSSandboxChannel resources.
type apnssandboxChannelState struct {
}

type APNSSandboxChannelState struct {
}

func (APNSSandboxChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*apnssandboxChannelState)(nil)).Elem()
}

type apnssandboxChannelArgs struct {
	ApplicationId               string  `pulumi:"applicationId"`
	BundleId                    *string `pulumi:"bundleId"`
	Certificate                 *string `pulumi:"certificate"`
	DefaultAuthenticationMethod *string `pulumi:"defaultAuthenticationMethod"`
	Enabled                     *bool   `pulumi:"enabled"`
	PrivateKey                  *string `pulumi:"privateKey"`
	TeamId                      *string `pulumi:"teamId"`
	TokenKey                    *string `pulumi:"tokenKey"`
	TokenKeyId                  *string `pulumi:"tokenKeyId"`
}

// The set of arguments for constructing a APNSSandboxChannel resource.
type APNSSandboxChannelArgs struct {
	ApplicationId               pulumi.StringInput
	BundleId                    pulumi.StringPtrInput
	Certificate                 pulumi.StringPtrInput
	DefaultAuthenticationMethod pulumi.StringPtrInput
	Enabled                     pulumi.BoolPtrInput
	PrivateKey                  pulumi.StringPtrInput
	TeamId                      pulumi.StringPtrInput
	TokenKey                    pulumi.StringPtrInput
	TokenKeyId                  pulumi.StringPtrInput
}

func (APNSSandboxChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apnssandboxChannelArgs)(nil)).Elem()
}

type APNSSandboxChannelInput interface {
	pulumi.Input

	ToAPNSSandboxChannelOutput() APNSSandboxChannelOutput
	ToAPNSSandboxChannelOutputWithContext(ctx context.Context) APNSSandboxChannelOutput
}

func (*APNSSandboxChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**APNSSandboxChannel)(nil)).Elem()
}

func (i *APNSSandboxChannel) ToAPNSSandboxChannelOutput() APNSSandboxChannelOutput {
	return i.ToAPNSSandboxChannelOutputWithContext(context.Background())
}

func (i *APNSSandboxChannel) ToAPNSSandboxChannelOutputWithContext(ctx context.Context) APNSSandboxChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APNSSandboxChannelOutput)
}

type APNSSandboxChannelOutput struct{ *pulumi.OutputState }

func (APNSSandboxChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**APNSSandboxChannel)(nil)).Elem()
}

func (o APNSSandboxChannelOutput) ToAPNSSandboxChannelOutput() APNSSandboxChannelOutput {
	return o
}

func (o APNSSandboxChannelOutput) ToAPNSSandboxChannelOutputWithContext(ctx context.Context) APNSSandboxChannelOutput {
	return o
}

func (o APNSSandboxChannelOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *APNSSandboxChannel) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

func (o APNSSandboxChannelOutput) BundleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APNSSandboxChannel) pulumi.StringPtrOutput { return v.BundleId }).(pulumi.StringPtrOutput)
}

func (o APNSSandboxChannelOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APNSSandboxChannel) pulumi.StringPtrOutput { return v.Certificate }).(pulumi.StringPtrOutput)
}

func (o APNSSandboxChannelOutput) DefaultAuthenticationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APNSSandboxChannel) pulumi.StringPtrOutput { return v.DefaultAuthenticationMethod }).(pulumi.StringPtrOutput)
}

func (o APNSSandboxChannelOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *APNSSandboxChannel) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o APNSSandboxChannelOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APNSSandboxChannel) pulumi.StringPtrOutput { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

func (o APNSSandboxChannelOutput) TeamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APNSSandboxChannel) pulumi.StringPtrOutput { return v.TeamId }).(pulumi.StringPtrOutput)
}

func (o APNSSandboxChannelOutput) TokenKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APNSSandboxChannel) pulumi.StringPtrOutput { return v.TokenKey }).(pulumi.StringPtrOutput)
}

func (o APNSSandboxChannelOutput) TokenKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APNSSandboxChannel) pulumi.StringPtrOutput { return v.TokenKeyId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*APNSSandboxChannelInput)(nil)).Elem(), &APNSSandboxChannel{})
	pulumi.RegisterOutputType(APNSSandboxChannelOutput{})
}
