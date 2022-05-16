// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package qldb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::QLDB::Stream.
type Stream struct {
	pulumi.CustomResourceState

	Arn                  pulumi.StringOutput              `pulumi:"arn"`
	ExclusiveEndTime     pulumi.StringPtrOutput           `pulumi:"exclusiveEndTime"`
	InclusiveStartTime   pulumi.StringOutput              `pulumi:"inclusiveStartTime"`
	KinesisConfiguration StreamKinesisConfigurationOutput `pulumi:"kinesisConfiguration"`
	LedgerName           pulumi.StringOutput              `pulumi:"ledgerName"`
	RoleArn              pulumi.StringOutput              `pulumi:"roleArn"`
	StreamName           pulumi.StringOutput              `pulumi:"streamName"`
	// An array of key-value pairs to apply to this resource.
	Tags StreamTagArrayOutput `pulumi:"tags"`
}

// NewStream registers a new resource with the given unique name, arguments, and options.
func NewStream(ctx *pulumi.Context,
	name string, args *StreamArgs, opts ...pulumi.ResourceOption) (*Stream, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InclusiveStartTime == nil {
		return nil, errors.New("invalid value for required argument 'InclusiveStartTime'")
	}
	if args.KinesisConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'KinesisConfiguration'")
	}
	if args.LedgerName == nil {
		return nil, errors.New("invalid value for required argument 'LedgerName'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource Stream
	err := ctx.RegisterResource("aws-native:qldb:Stream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStream gets an existing Stream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamState, opts ...pulumi.ResourceOption) (*Stream, error) {
	var resource Stream
	err := ctx.ReadResource("aws-native:qldb:Stream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stream resources.
type streamState struct {
}

type StreamState struct {
}

func (StreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamState)(nil)).Elem()
}

type streamArgs struct {
	ExclusiveEndTime     *string                    `pulumi:"exclusiveEndTime"`
	InclusiveStartTime   string                     `pulumi:"inclusiveStartTime"`
	KinesisConfiguration StreamKinesisConfiguration `pulumi:"kinesisConfiguration"`
	LedgerName           string                     `pulumi:"ledgerName"`
	RoleArn              string                     `pulumi:"roleArn"`
	StreamName           *string                    `pulumi:"streamName"`
	// An array of key-value pairs to apply to this resource.
	Tags []StreamTag `pulumi:"tags"`
}

// The set of arguments for constructing a Stream resource.
type StreamArgs struct {
	ExclusiveEndTime     pulumi.StringPtrInput
	InclusiveStartTime   pulumi.StringInput
	KinesisConfiguration StreamKinesisConfigurationInput
	LedgerName           pulumi.StringInput
	RoleArn              pulumi.StringInput
	StreamName           pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags StreamTagArrayInput
}

func (StreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamArgs)(nil)).Elem()
}

type StreamInput interface {
	pulumi.Input

	ToStreamOutput() StreamOutput
	ToStreamOutputWithContext(ctx context.Context) StreamOutput
}

func (*Stream) ElementType() reflect.Type {
	return reflect.TypeOf((**Stream)(nil)).Elem()
}

func (i *Stream) ToStreamOutput() StreamOutput {
	return i.ToStreamOutputWithContext(context.Background())
}

func (i *Stream) ToStreamOutputWithContext(ctx context.Context) StreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamOutput)
}

type StreamOutput struct{ *pulumi.OutputState }

func (StreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Stream)(nil)).Elem()
}

func (o StreamOutput) ToStreamOutput() StreamOutput {
	return o
}

func (o StreamOutput) ToStreamOutputWithContext(ctx context.Context) StreamOutput {
	return o
}

func (o StreamOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o StreamOutput) ExclusiveEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringPtrOutput { return v.ExclusiveEndTime }).(pulumi.StringPtrOutput)
}

func (o StreamOutput) InclusiveStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringOutput { return v.InclusiveStartTime }).(pulumi.StringOutput)
}

func (o StreamOutput) KinesisConfiguration() StreamKinesisConfigurationOutput {
	return o.ApplyT(func(v *Stream) StreamKinesisConfigurationOutput { return v.KinesisConfiguration }).(StreamKinesisConfigurationOutput)
}

func (o StreamOutput) LedgerName() pulumi.StringOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringOutput { return v.LedgerName }).(pulumi.StringOutput)
}

func (o StreamOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

func (o StreamOutput) StreamName() pulumi.StringOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringOutput { return v.StreamName }).(pulumi.StringOutput)
}

// An array of key-value pairs to apply to this resource.
func (o StreamOutput) Tags() StreamTagArrayOutput {
	return o.ApplyT(func(v *Stream) StreamTagArrayOutput { return v.Tags }).(StreamTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StreamInput)(nil)).Elem(), &Stream{})
	pulumi.RegisterOutputType(StreamOutput{})
}
