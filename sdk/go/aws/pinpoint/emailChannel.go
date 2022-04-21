// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Pinpoint::EmailChannel
//
// Deprecated: EmailChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type EmailChannel struct {
	pulumi.CustomResourceState

	ApplicationId    pulumi.StringOutput    `pulumi:"applicationId"`
	ConfigurationSet pulumi.StringPtrOutput `pulumi:"configurationSet"`
	Enabled          pulumi.BoolPtrOutput   `pulumi:"enabled"`
	FromAddress      pulumi.StringOutput    `pulumi:"fromAddress"`
	Identity         pulumi.StringOutput    `pulumi:"identity"`
	RoleArn          pulumi.StringPtrOutput `pulumi:"roleArn"`
}

// NewEmailChannel registers a new resource with the given unique name, arguments, and options.
func NewEmailChannel(ctx *pulumi.Context,
	name string, args *EmailChannelArgs, opts ...pulumi.ResourceOption) (*EmailChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.FromAddress == nil {
		return nil, errors.New("invalid value for required argument 'FromAddress'")
	}
	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
	}
	var resource EmailChannel
	err := ctx.RegisterResource("aws-native:pinpoint:EmailChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailChannel gets an existing EmailChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailChannelState, opts ...pulumi.ResourceOption) (*EmailChannel, error) {
	var resource EmailChannel
	err := ctx.ReadResource("aws-native:pinpoint:EmailChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailChannel resources.
type emailChannelState struct {
}

type EmailChannelState struct {
}

func (EmailChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailChannelState)(nil)).Elem()
}

type emailChannelArgs struct {
	ApplicationId    string  `pulumi:"applicationId"`
	ConfigurationSet *string `pulumi:"configurationSet"`
	Enabled          *bool   `pulumi:"enabled"`
	FromAddress      string  `pulumi:"fromAddress"`
	Identity         string  `pulumi:"identity"`
	RoleArn          *string `pulumi:"roleArn"`
}

// The set of arguments for constructing a EmailChannel resource.
type EmailChannelArgs struct {
	ApplicationId    pulumi.StringInput
	ConfigurationSet pulumi.StringPtrInput
	Enabled          pulumi.BoolPtrInput
	FromAddress      pulumi.StringInput
	Identity         pulumi.StringInput
	RoleArn          pulumi.StringPtrInput
}

func (EmailChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailChannelArgs)(nil)).Elem()
}

type EmailChannelInput interface {
	pulumi.Input

	ToEmailChannelOutput() EmailChannelOutput
	ToEmailChannelOutputWithContext(ctx context.Context) EmailChannelOutput
}

func (*EmailChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailChannel)(nil)).Elem()
}

func (i *EmailChannel) ToEmailChannelOutput() EmailChannelOutput {
	return i.ToEmailChannelOutputWithContext(context.Background())
}

func (i *EmailChannel) ToEmailChannelOutputWithContext(ctx context.Context) EmailChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailChannelOutput)
}

type EmailChannelOutput struct{ *pulumi.OutputState }

func (EmailChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailChannel)(nil)).Elem()
}

func (o EmailChannelOutput) ToEmailChannelOutput() EmailChannelOutput {
	return o
}

func (o EmailChannelOutput) ToEmailChannelOutputWithContext(ctx context.Context) EmailChannelOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EmailChannelInput)(nil)).Elem(), &EmailChannel{})
	pulumi.RegisterOutputType(EmailChannelOutput{})
}
