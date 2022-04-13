// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Lambda::EventInvokeConfig
//
// Deprecated: EventInvokeConfig is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type EventInvokeConfig struct {
	pulumi.CustomResourceState

	DestinationConfig        EventInvokeConfigDestinationConfigPtrOutput `pulumi:"destinationConfig"`
	FunctionName             pulumi.StringOutput                         `pulumi:"functionName"`
	MaximumEventAgeInSeconds pulumi.IntPtrOutput                         `pulumi:"maximumEventAgeInSeconds"`
	MaximumRetryAttempts     pulumi.IntPtrOutput                         `pulumi:"maximumRetryAttempts"`
	Qualifier                pulumi.StringOutput                         `pulumi:"qualifier"`
}

// NewEventInvokeConfig registers a new resource with the given unique name, arguments, and options.
func NewEventInvokeConfig(ctx *pulumi.Context,
	name string, args *EventInvokeConfigArgs, opts ...pulumi.ResourceOption) (*EventInvokeConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FunctionName == nil {
		return nil, errors.New("invalid value for required argument 'FunctionName'")
	}
	if args.Qualifier == nil {
		return nil, errors.New("invalid value for required argument 'Qualifier'")
	}
	var resource EventInvokeConfig
	err := ctx.RegisterResource("aws-native:lambda:EventInvokeConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventInvokeConfig gets an existing EventInvokeConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventInvokeConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventInvokeConfigState, opts ...pulumi.ResourceOption) (*EventInvokeConfig, error) {
	var resource EventInvokeConfig
	err := ctx.ReadResource("aws-native:lambda:EventInvokeConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventInvokeConfig resources.
type eventInvokeConfigState struct {
}

type EventInvokeConfigState struct {
}

func (EventInvokeConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventInvokeConfigState)(nil)).Elem()
}

type eventInvokeConfigArgs struct {
	DestinationConfig        *EventInvokeConfigDestinationConfig `pulumi:"destinationConfig"`
	FunctionName             string                              `pulumi:"functionName"`
	MaximumEventAgeInSeconds *int                                `pulumi:"maximumEventAgeInSeconds"`
	MaximumRetryAttempts     *int                                `pulumi:"maximumRetryAttempts"`
	Qualifier                string                              `pulumi:"qualifier"`
}

// The set of arguments for constructing a EventInvokeConfig resource.
type EventInvokeConfigArgs struct {
	DestinationConfig        EventInvokeConfigDestinationConfigPtrInput
	FunctionName             pulumi.StringInput
	MaximumEventAgeInSeconds pulumi.IntPtrInput
	MaximumRetryAttempts     pulumi.IntPtrInput
	Qualifier                pulumi.StringInput
}

func (EventInvokeConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventInvokeConfigArgs)(nil)).Elem()
}

type EventInvokeConfigInput interface {
	pulumi.Input

	ToEventInvokeConfigOutput() EventInvokeConfigOutput
	ToEventInvokeConfigOutputWithContext(ctx context.Context) EventInvokeConfigOutput
}

func (*EventInvokeConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**EventInvokeConfig)(nil)).Elem()
}

func (i *EventInvokeConfig) ToEventInvokeConfigOutput() EventInvokeConfigOutput {
	return i.ToEventInvokeConfigOutputWithContext(context.Background())
}

func (i *EventInvokeConfig) ToEventInvokeConfigOutputWithContext(ctx context.Context) EventInvokeConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventInvokeConfigOutput)
}

type EventInvokeConfigOutput struct{ *pulumi.OutputState }

func (EventInvokeConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventInvokeConfig)(nil)).Elem()
}

func (o EventInvokeConfigOutput) ToEventInvokeConfigOutput() EventInvokeConfigOutput {
	return o
}

func (o EventInvokeConfigOutput) ToEventInvokeConfigOutputWithContext(ctx context.Context) EventInvokeConfigOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventInvokeConfigInput)(nil)).Elem(), &EventInvokeConfig{})
	pulumi.RegisterOutputType(EventInvokeConfigOutput{})
}
