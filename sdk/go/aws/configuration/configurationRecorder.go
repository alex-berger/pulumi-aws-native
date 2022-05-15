// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Config::ConfigurationRecorder
//
// Deprecated: ConfigurationRecorder is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ConfigurationRecorder struct {
	pulumi.CustomResourceState

	Name           pulumi.StringPtrOutput                       `pulumi:"name"`
	RecordingGroup ConfigurationRecorderRecordingGroupPtrOutput `pulumi:"recordingGroup"`
	RoleARN        pulumi.StringOutput                          `pulumi:"roleARN"`
}

// NewConfigurationRecorder registers a new resource with the given unique name, arguments, and options.
func NewConfigurationRecorder(ctx *pulumi.Context,
	name string, args *ConfigurationRecorderArgs, opts ...pulumi.ResourceOption) (*ConfigurationRecorder, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleARN == nil {
		return nil, errors.New("invalid value for required argument 'RoleARN'")
	}
	var resource ConfigurationRecorder
	err := ctx.RegisterResource("aws-native:configuration:ConfigurationRecorder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationRecorder gets an existing ConfigurationRecorder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationRecorder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationRecorderState, opts ...pulumi.ResourceOption) (*ConfigurationRecorder, error) {
	var resource ConfigurationRecorder
	err := ctx.ReadResource("aws-native:configuration:ConfigurationRecorder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationRecorder resources.
type configurationRecorderState struct {
}

type ConfigurationRecorderState struct {
}

func (ConfigurationRecorderState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationRecorderState)(nil)).Elem()
}

type configurationRecorderArgs struct {
	Name           *string                              `pulumi:"name"`
	RecordingGroup *ConfigurationRecorderRecordingGroup `pulumi:"recordingGroup"`
	RoleARN        string                               `pulumi:"roleARN"`
}

// The set of arguments for constructing a ConfigurationRecorder resource.
type ConfigurationRecorderArgs struct {
	Name           pulumi.StringPtrInput
	RecordingGroup ConfigurationRecorderRecordingGroupPtrInput
	RoleARN        pulumi.StringInput
}

func (ConfigurationRecorderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationRecorderArgs)(nil)).Elem()
}

type ConfigurationRecorderInput interface {
	pulumi.Input

	ToConfigurationRecorderOutput() ConfigurationRecorderOutput
	ToConfigurationRecorderOutputWithContext(ctx context.Context) ConfigurationRecorderOutput
}

func (*ConfigurationRecorder) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationRecorder)(nil)).Elem()
}

func (i *ConfigurationRecorder) ToConfigurationRecorderOutput() ConfigurationRecorderOutput {
	return i.ToConfigurationRecorderOutputWithContext(context.Background())
}

func (i *ConfigurationRecorder) ToConfigurationRecorderOutputWithContext(ctx context.Context) ConfigurationRecorderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationRecorderOutput)
}

type ConfigurationRecorderOutput struct{ *pulumi.OutputState }

func (ConfigurationRecorderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationRecorder)(nil)).Elem()
}

func (o ConfigurationRecorderOutput) ToConfigurationRecorderOutput() ConfigurationRecorderOutput {
	return o
}

func (o ConfigurationRecorderOutput) ToConfigurationRecorderOutputWithContext(ctx context.Context) ConfigurationRecorderOutput {
	return o
}

func (o ConfigurationRecorderOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationRecorder) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConfigurationRecorderOutput) RecordingGroup() ConfigurationRecorderRecordingGroupPtrOutput {
	return o.ApplyT(func(v *ConfigurationRecorder) ConfigurationRecorderRecordingGroupPtrOutput { return v.RecordingGroup }).(ConfigurationRecorderRecordingGroupPtrOutput)
}

func (o ConfigurationRecorderOutput) RoleARN() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationRecorder) pulumi.StringOutput { return v.RoleARN }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationRecorderInput)(nil)).Elem(), &ConfigurationRecorder{})
	pulumi.RegisterOutputType(ConfigurationRecorderOutput{})
}
