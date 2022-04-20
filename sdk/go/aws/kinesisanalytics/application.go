// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kinesisanalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::KinesisAnalytics::Application
//
// Deprecated: Application is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Application struct {
	pulumi.CustomResourceState

	ApplicationCode        pulumi.StringPtrOutput          `pulumi:"applicationCode"`
	ApplicationDescription pulumi.StringPtrOutput          `pulumi:"applicationDescription"`
	ApplicationName        pulumi.StringPtrOutput          `pulumi:"applicationName"`
	Inputs                 ApplicationInputTypeArrayOutput `pulumi:"inputs"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Inputs == nil {
		return nil, errors.New("invalid value for required argument 'Inputs'")
	}
	var resource Application
	err := ctx.RegisterResource("aws-native:kinesisanalytics:Application", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:kinesisanalytics:Application", name, id, state, &resource, opts...)
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
	ApplicationCode        *string                `pulumi:"applicationCode"`
	ApplicationDescription *string                `pulumi:"applicationDescription"`
	ApplicationName        *string                `pulumi:"applicationName"`
	Inputs                 []ApplicationInputType `pulumi:"inputs"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	ApplicationCode        pulumi.StringPtrInput
	ApplicationDescription pulumi.StringPtrInput
	ApplicationName        pulumi.StringPtrInput
	Inputs                 ApplicationInputTypeArrayInput
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
