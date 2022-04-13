// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Greengrass::LoggerDefinitionVersion
//
// Deprecated: LoggerDefinitionVersion is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type LoggerDefinitionVersion struct {
	pulumi.CustomResourceState

	LoggerDefinitionId pulumi.StringOutput                      `pulumi:"loggerDefinitionId"`
	Loggers            LoggerDefinitionVersionLoggerArrayOutput `pulumi:"loggers"`
}

// NewLoggerDefinitionVersion registers a new resource with the given unique name, arguments, and options.
func NewLoggerDefinitionVersion(ctx *pulumi.Context,
	name string, args *LoggerDefinitionVersionArgs, opts ...pulumi.ResourceOption) (*LoggerDefinitionVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoggerDefinitionId == nil {
		return nil, errors.New("invalid value for required argument 'LoggerDefinitionId'")
	}
	if args.Loggers == nil {
		return nil, errors.New("invalid value for required argument 'Loggers'")
	}
	var resource LoggerDefinitionVersion
	err := ctx.RegisterResource("aws-native:greengrass:LoggerDefinitionVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoggerDefinitionVersion gets an existing LoggerDefinitionVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoggerDefinitionVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoggerDefinitionVersionState, opts ...pulumi.ResourceOption) (*LoggerDefinitionVersion, error) {
	var resource LoggerDefinitionVersion
	err := ctx.ReadResource("aws-native:greengrass:LoggerDefinitionVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoggerDefinitionVersion resources.
type loggerDefinitionVersionState struct {
}

type LoggerDefinitionVersionState struct {
}

func (LoggerDefinitionVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*loggerDefinitionVersionState)(nil)).Elem()
}

type loggerDefinitionVersionArgs struct {
	LoggerDefinitionId string                          `pulumi:"loggerDefinitionId"`
	Loggers            []LoggerDefinitionVersionLogger `pulumi:"loggers"`
}

// The set of arguments for constructing a LoggerDefinitionVersion resource.
type LoggerDefinitionVersionArgs struct {
	LoggerDefinitionId pulumi.StringInput
	Loggers            LoggerDefinitionVersionLoggerArrayInput
}

func (LoggerDefinitionVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loggerDefinitionVersionArgs)(nil)).Elem()
}

type LoggerDefinitionVersionInput interface {
	pulumi.Input

	ToLoggerDefinitionVersionOutput() LoggerDefinitionVersionOutput
	ToLoggerDefinitionVersionOutputWithContext(ctx context.Context) LoggerDefinitionVersionOutput
}

func (*LoggerDefinitionVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggerDefinitionVersion)(nil)).Elem()
}

func (i *LoggerDefinitionVersion) ToLoggerDefinitionVersionOutput() LoggerDefinitionVersionOutput {
	return i.ToLoggerDefinitionVersionOutputWithContext(context.Background())
}

func (i *LoggerDefinitionVersion) ToLoggerDefinitionVersionOutputWithContext(ctx context.Context) LoggerDefinitionVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggerDefinitionVersionOutput)
}

type LoggerDefinitionVersionOutput struct{ *pulumi.OutputState }

func (LoggerDefinitionVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggerDefinitionVersion)(nil)).Elem()
}

func (o LoggerDefinitionVersionOutput) ToLoggerDefinitionVersionOutput() LoggerDefinitionVersionOutput {
	return o
}

func (o LoggerDefinitionVersionOutput) ToLoggerDefinitionVersionOutputWithContext(ctx context.Context) LoggerDefinitionVersionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoggerDefinitionVersionInput)(nil)).Elem(), &LoggerDefinitionVersion{})
	pulumi.RegisterOutputType(LoggerDefinitionVersionOutput{})
}
