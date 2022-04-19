// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Pinpoint::VoiceChannel
//
// Deprecated: VoiceChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type VoiceChannel struct {
	pulumi.CustomResourceState

	ApplicationId pulumi.StringOutput  `pulumi:"applicationId"`
	Enabled       pulumi.BoolPtrOutput `pulumi:"enabled"`
}

// NewVoiceChannel registers a new resource with the given unique name, arguments, and options.
func NewVoiceChannel(ctx *pulumi.Context,
	name string, args *VoiceChannelArgs, opts ...pulumi.ResourceOption) (*VoiceChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	var resource VoiceChannel
	err := ctx.RegisterResource("aws-native:pinpoint:VoiceChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVoiceChannel gets an existing VoiceChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVoiceChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VoiceChannelState, opts ...pulumi.ResourceOption) (*VoiceChannel, error) {
	var resource VoiceChannel
	err := ctx.ReadResource("aws-native:pinpoint:VoiceChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VoiceChannel resources.
type voiceChannelState struct {
}

type VoiceChannelState struct {
}

func (VoiceChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*voiceChannelState)(nil)).Elem()
}

type voiceChannelArgs struct {
	ApplicationId string `pulumi:"applicationId"`
	Enabled       *bool  `pulumi:"enabled"`
}

// The set of arguments for constructing a VoiceChannel resource.
type VoiceChannelArgs struct {
	ApplicationId pulumi.StringInput
	Enabled       pulumi.BoolPtrInput
}

func (VoiceChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*voiceChannelArgs)(nil)).Elem()
}

type VoiceChannelInput interface {
	pulumi.Input

	ToVoiceChannelOutput() VoiceChannelOutput
	ToVoiceChannelOutputWithContext(ctx context.Context) VoiceChannelOutput
}

func (*VoiceChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**VoiceChannel)(nil)).Elem()
}

func (i *VoiceChannel) ToVoiceChannelOutput() VoiceChannelOutput {
	return i.ToVoiceChannelOutputWithContext(context.Background())
}

func (i *VoiceChannel) ToVoiceChannelOutputWithContext(ctx context.Context) VoiceChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceChannelOutput)
}

type VoiceChannelOutput struct{ *pulumi.OutputState }

func (VoiceChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VoiceChannel)(nil)).Elem()
}

func (o VoiceChannelOutput) ToVoiceChannelOutput() VoiceChannelOutput {
	return o
}

func (o VoiceChannelOutput) ToVoiceChannelOutputWithContext(ctx context.Context) VoiceChannelOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VoiceChannelInput)(nil)).Elem(), &VoiceChannel{})
	pulumi.RegisterOutputType(VoiceChannelOutput{})
}
