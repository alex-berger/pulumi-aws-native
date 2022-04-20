// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package evidently

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Evidently::Launch.
type Launch struct {
	pulumi.CustomResourceState

	Arn         pulumi.StringOutput    `pulumi:"arn"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Start or Stop Launch Launch. Default is not started.
	ExecutionStatus       LaunchExecutionStatusObjectPtrOutput    `pulumi:"executionStatus"`
	Groups                LaunchGroupObjectArrayOutput            `pulumi:"groups"`
	MetricMonitors        LaunchMetricDefinitionObjectArrayOutput `pulumi:"metricMonitors"`
	Name                  pulumi.StringOutput                     `pulumi:"name"`
	Project               pulumi.StringOutput                     `pulumi:"project"`
	RandomizationSalt     pulumi.StringPtrOutput                  `pulumi:"randomizationSalt"`
	ScheduledSplitsConfig LaunchStepConfigArrayOutput             `pulumi:"scheduledSplitsConfig"`
	// An array of key-value pairs to apply to this resource.
	Tags LaunchTagArrayOutput `pulumi:"tags"`
}

// NewLaunch registers a new resource with the given unique name, arguments, and options.
func NewLaunch(ctx *pulumi.Context,
	name string, args *LaunchArgs, opts ...pulumi.ResourceOption) (*Launch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Groups == nil {
		return nil, errors.New("invalid value for required argument 'Groups'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.ScheduledSplitsConfig == nil {
		return nil, errors.New("invalid value for required argument 'ScheduledSplitsConfig'")
	}
	var resource Launch
	err := ctx.RegisterResource("aws-native:evidently:Launch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLaunch gets an existing Launch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLaunch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LaunchState, opts ...pulumi.ResourceOption) (*Launch, error) {
	var resource Launch
	err := ctx.ReadResource("aws-native:evidently:Launch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Launch resources.
type launchState struct {
}

type LaunchState struct {
}

func (LaunchState) ElementType() reflect.Type {
	return reflect.TypeOf((*launchState)(nil)).Elem()
}

type launchArgs struct {
	Description *string `pulumi:"description"`
	// Start or Stop Launch Launch. Default is not started.
	ExecutionStatus       *LaunchExecutionStatusObject   `pulumi:"executionStatus"`
	Groups                []LaunchGroupObject            `pulumi:"groups"`
	MetricMonitors        []LaunchMetricDefinitionObject `pulumi:"metricMonitors"`
	Name                  *string                        `pulumi:"name"`
	Project               string                         `pulumi:"project"`
	RandomizationSalt     *string                        `pulumi:"randomizationSalt"`
	ScheduledSplitsConfig []LaunchStepConfig             `pulumi:"scheduledSplitsConfig"`
	// An array of key-value pairs to apply to this resource.
	Tags []LaunchTag `pulumi:"tags"`
}

// The set of arguments for constructing a Launch resource.
type LaunchArgs struct {
	Description pulumi.StringPtrInput
	// Start or Stop Launch Launch. Default is not started.
	ExecutionStatus       LaunchExecutionStatusObjectPtrInput
	Groups                LaunchGroupObjectArrayInput
	MetricMonitors        LaunchMetricDefinitionObjectArrayInput
	Name                  pulumi.StringPtrInput
	Project               pulumi.StringInput
	RandomizationSalt     pulumi.StringPtrInput
	ScheduledSplitsConfig LaunchStepConfigArrayInput
	// An array of key-value pairs to apply to this resource.
	Tags LaunchTagArrayInput
}

func (LaunchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*launchArgs)(nil)).Elem()
}

type LaunchInput interface {
	pulumi.Input

	ToLaunchOutput() LaunchOutput
	ToLaunchOutputWithContext(ctx context.Context) LaunchOutput
}

func (*Launch) ElementType() reflect.Type {
	return reflect.TypeOf((**Launch)(nil)).Elem()
}

func (i *Launch) ToLaunchOutput() LaunchOutput {
	return i.ToLaunchOutputWithContext(context.Background())
}

func (i *Launch) ToLaunchOutputWithContext(ctx context.Context) LaunchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchOutput)
}

type LaunchOutput struct{ *pulumi.OutputState }

func (LaunchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Launch)(nil)).Elem()
}

func (o LaunchOutput) ToLaunchOutput() LaunchOutput {
	return o
}

func (o LaunchOutput) ToLaunchOutputWithContext(ctx context.Context) LaunchOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LaunchInput)(nil)).Elem(), &Launch{})
	pulumi.RegisterOutputType(LaunchOutput{})
}
