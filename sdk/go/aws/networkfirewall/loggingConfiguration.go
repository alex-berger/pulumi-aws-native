// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkfirewall

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource type definition for AWS::NetworkFirewall::LoggingConfiguration
type LoggingConfiguration struct {
	pulumi.CustomResourceState

	FirewallArn          pulumi.StringOutput            `pulumi:"firewallArn"`
	FirewallName         pulumi.StringPtrOutput         `pulumi:"firewallName"`
	LoggingConfiguration LoggingConfigurationTypeOutput `pulumi:"loggingConfiguration"`
}

// NewLoggingConfiguration registers a new resource with the given unique name, arguments, and options.
func NewLoggingConfiguration(ctx *pulumi.Context,
	name string, args *LoggingConfigurationArgs, opts ...pulumi.ResourceOption) (*LoggingConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FirewallArn == nil {
		return nil, errors.New("invalid value for required argument 'FirewallArn'")
	}
	if args.LoggingConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'LoggingConfiguration'")
	}
	var resource LoggingConfiguration
	err := ctx.RegisterResource("aws-native:networkfirewall:LoggingConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoggingConfiguration gets an existing LoggingConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoggingConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoggingConfigurationState, opts ...pulumi.ResourceOption) (*LoggingConfiguration, error) {
	var resource LoggingConfiguration
	err := ctx.ReadResource("aws-native:networkfirewall:LoggingConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoggingConfiguration resources.
type loggingConfigurationState struct {
}

type LoggingConfigurationState struct {
}

func (LoggingConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingConfigurationState)(nil)).Elem()
}

type loggingConfigurationArgs struct {
	FirewallArn          string                   `pulumi:"firewallArn"`
	FirewallName         *string                  `pulumi:"firewallName"`
	LoggingConfiguration LoggingConfigurationType `pulumi:"loggingConfiguration"`
}

// The set of arguments for constructing a LoggingConfiguration resource.
type LoggingConfigurationArgs struct {
	FirewallArn          pulumi.StringInput
	FirewallName         pulumi.StringPtrInput
	LoggingConfiguration LoggingConfigurationTypeInput
}

func (LoggingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingConfigurationArgs)(nil)).Elem()
}

type LoggingConfigurationInput interface {
	pulumi.Input

	ToLoggingConfigurationOutput() LoggingConfigurationOutput
	ToLoggingConfigurationOutputWithContext(ctx context.Context) LoggingConfigurationOutput
}

func (*LoggingConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingConfiguration)(nil)).Elem()
}

func (i *LoggingConfiguration) ToLoggingConfigurationOutput() LoggingConfigurationOutput {
	return i.ToLoggingConfigurationOutputWithContext(context.Background())
}

func (i *LoggingConfiguration) ToLoggingConfigurationOutputWithContext(ctx context.Context) LoggingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingConfigurationOutput)
}

type LoggingConfigurationOutput struct{ *pulumi.OutputState }

func (LoggingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingConfiguration)(nil)).Elem()
}

func (o LoggingConfigurationOutput) ToLoggingConfigurationOutput() LoggingConfigurationOutput {
	return o
}

func (o LoggingConfigurationOutput) ToLoggingConfigurationOutputWithContext(ctx context.Context) LoggingConfigurationOutput {
	return o
}

func (o LoggingConfigurationOutput) FirewallArn() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingConfiguration) pulumi.StringOutput { return v.FirewallArn }).(pulumi.StringOutput)
}

func (o LoggingConfigurationOutput) FirewallName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggingConfiguration) pulumi.StringPtrOutput { return v.FirewallName }).(pulumi.StringPtrOutput)
}

func (o LoggingConfigurationOutput) LoggingConfiguration() LoggingConfigurationTypeOutput {
	return o.ApplyT(func(v *LoggingConfiguration) LoggingConfigurationTypeOutput { return v.LoggingConfiguration }).(LoggingConfigurationTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoggingConfigurationInput)(nil)).Elem(), &LoggingConfiguration{})
	pulumi.RegisterOutputType(LoggingConfigurationOutput{})
}
