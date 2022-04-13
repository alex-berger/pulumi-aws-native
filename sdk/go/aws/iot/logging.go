// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Logging Options enable you to configure your IoT V2 logging role and default logging level so that you can monitor progress events logs as it passes from your devices through Iot core service.
type Logging struct {
	pulumi.CustomResourceState

	// Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The log level to use. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.
	DefaultLogLevel LoggingDefaultLogLevelOutput `pulumi:"defaultLogLevel"`
	// The ARN of the role that allows IoT to write to Cloudwatch logs.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
}

// NewLogging registers a new resource with the given unique name, arguments, and options.
func NewLogging(ctx *pulumi.Context,
	name string, args *LoggingArgs, opts ...pulumi.ResourceOption) (*Logging, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	if args.DefaultLogLevel == nil {
		return nil, errors.New("invalid value for required argument 'DefaultLogLevel'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource Logging
	err := ctx.RegisterResource("aws-native:iot:Logging", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogging gets an existing Logging resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogging(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoggingState, opts ...pulumi.ResourceOption) (*Logging, error) {
	var resource Logging
	err := ctx.ReadResource("aws-native:iot:Logging", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Logging resources.
type loggingState struct {
}

type LoggingState struct {
}

func (LoggingState) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingState)(nil)).Elem()
}

type loggingArgs struct {
	// Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).
	AccountId string `pulumi:"accountId"`
	// The log level to use. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.
	DefaultLogLevel LoggingDefaultLogLevel `pulumi:"defaultLogLevel"`
	// The ARN of the role that allows IoT to write to Cloudwatch logs.
	RoleArn string `pulumi:"roleArn"`
}

// The set of arguments for constructing a Logging resource.
type LoggingArgs struct {
	// Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).
	AccountId pulumi.StringInput
	// The log level to use. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.
	DefaultLogLevel LoggingDefaultLogLevelInput
	// The ARN of the role that allows IoT to write to Cloudwatch logs.
	RoleArn pulumi.StringInput
}

func (LoggingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingArgs)(nil)).Elem()
}

type LoggingInput interface {
	pulumi.Input

	ToLoggingOutput() LoggingOutput
	ToLoggingOutputWithContext(ctx context.Context) LoggingOutput
}

func (*Logging) ElementType() reflect.Type {
	return reflect.TypeOf((**Logging)(nil)).Elem()
}

func (i *Logging) ToLoggingOutput() LoggingOutput {
	return i.ToLoggingOutputWithContext(context.Background())
}

func (i *Logging) ToLoggingOutputWithContext(ctx context.Context) LoggingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingOutput)
}

type LoggingOutput struct{ *pulumi.OutputState }

func (LoggingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Logging)(nil)).Elem()
}

func (o LoggingOutput) ToLoggingOutput() LoggingOutput {
	return o
}

func (o LoggingOutput) ToLoggingOutputWithContext(ctx context.Context) LoggingOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoggingInput)(nil)).Elem(), &Logging{})
	pulumi.RegisterOutputType(LoggingOutput{})
}
