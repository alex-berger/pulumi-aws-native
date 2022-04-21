// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kinesisanalyticsv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::KinesisAnalyticsV2::Application
//
// Deprecated: Application is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Application struct {
	pulumi.CustomResourceState

	ApplicationConfiguration ApplicationConfigurationPtrOutput `pulumi:"applicationConfiguration"`
	ApplicationDescription   pulumi.StringPtrOutput            `pulumi:"applicationDescription"`
	ApplicationMode          pulumi.StringPtrOutput            `pulumi:"applicationMode"`
	ApplicationName          pulumi.StringPtrOutput            `pulumi:"applicationName"`
	RuntimeEnvironment       pulumi.StringOutput               `pulumi:"runtimeEnvironment"`
	ServiceExecutionRole     pulumi.StringOutput               `pulumi:"serviceExecutionRole"`
	Tags                     ApplicationTagArrayOutput         `pulumi:"tags"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RuntimeEnvironment == nil {
		return nil, errors.New("invalid value for required argument 'RuntimeEnvironment'")
	}
	if args.ServiceExecutionRole == nil {
		return nil, errors.New("invalid value for required argument 'ServiceExecutionRole'")
	}
	var resource Application
	err := ctx.RegisterResource("aws-native:kinesisanalyticsv2:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("aws-native:kinesisanalyticsv2:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
}

type ApplicationState struct {
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	ApplicationConfiguration *ApplicationConfiguration `pulumi:"applicationConfiguration"`
	ApplicationDescription   *string                   `pulumi:"applicationDescription"`
	ApplicationMode          *string                   `pulumi:"applicationMode"`
	ApplicationName          *string                   `pulumi:"applicationName"`
	RuntimeEnvironment       string                    `pulumi:"runtimeEnvironment"`
	ServiceExecutionRole     string                    `pulumi:"serviceExecutionRole"`
	Tags                     []ApplicationTag          `pulumi:"tags"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	ApplicationConfiguration ApplicationConfigurationPtrInput
	ApplicationDescription   pulumi.StringPtrInput
	ApplicationMode          pulumi.StringPtrInput
	ApplicationName          pulumi.StringPtrInput
	RuntimeEnvironment       pulumi.StringInput
	ServiceExecutionRole     pulumi.StringInput
	Tags                     ApplicationTagArrayInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInput)(nil)).Elem(), &Application{})
	pulumi.RegisterOutputType(ApplicationOutput{})
}
