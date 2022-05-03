// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package robomaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::RoboMaker::RobotApplicationVersion resource creates an AWS RoboMaker RobotApplicationVersion. This helps you control which code your robot uses.
type RobotApplicationVersion struct {
	pulumi.CustomResourceState

	Application        pulumi.StringOutput `pulumi:"application"`
	ApplicationVersion pulumi.StringOutput `pulumi:"applicationVersion"`
	Arn                pulumi.StringOutput `pulumi:"arn"`
	// The revision ID of robot application.
	CurrentRevisionId pulumi.StringPtrOutput `pulumi:"currentRevisionId"`
}

// NewRobotApplicationVersion registers a new resource with the given unique name, arguments, and options.
func NewRobotApplicationVersion(ctx *pulumi.Context,
	name string, args *RobotApplicationVersionArgs, opts ...pulumi.ResourceOption) (*RobotApplicationVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Application == nil {
		return nil, errors.New("invalid value for required argument 'Application'")
	}
	var resource RobotApplicationVersion
	err := ctx.RegisterResource("aws-native:robomaker:RobotApplicationVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRobotApplicationVersion gets an existing RobotApplicationVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRobotApplicationVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RobotApplicationVersionState, opts ...pulumi.ResourceOption) (*RobotApplicationVersion, error) {
	var resource RobotApplicationVersion
	err := ctx.ReadResource("aws-native:robomaker:RobotApplicationVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RobotApplicationVersion resources.
type robotApplicationVersionState struct {
}

type RobotApplicationVersionState struct {
}

func (RobotApplicationVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*robotApplicationVersionState)(nil)).Elem()
}

type robotApplicationVersionArgs struct {
	Application string `pulumi:"application"`
	// The revision ID of robot application.
	CurrentRevisionId *string `pulumi:"currentRevisionId"`
}

// The set of arguments for constructing a RobotApplicationVersion resource.
type RobotApplicationVersionArgs struct {
	Application pulumi.StringInput
	// The revision ID of robot application.
	CurrentRevisionId pulumi.StringPtrInput
}

func (RobotApplicationVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*robotApplicationVersionArgs)(nil)).Elem()
}

type RobotApplicationVersionInput interface {
	pulumi.Input

	ToRobotApplicationVersionOutput() RobotApplicationVersionOutput
	ToRobotApplicationVersionOutputWithContext(ctx context.Context) RobotApplicationVersionOutput
}

func (*RobotApplicationVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**RobotApplicationVersion)(nil)).Elem()
}

func (i *RobotApplicationVersion) ToRobotApplicationVersionOutput() RobotApplicationVersionOutput {
	return i.ToRobotApplicationVersionOutputWithContext(context.Background())
}

func (i *RobotApplicationVersion) ToRobotApplicationVersionOutputWithContext(ctx context.Context) RobotApplicationVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RobotApplicationVersionOutput)
}

type RobotApplicationVersionOutput struct{ *pulumi.OutputState }

func (RobotApplicationVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RobotApplicationVersion)(nil)).Elem()
}

func (o RobotApplicationVersionOutput) ToRobotApplicationVersionOutput() RobotApplicationVersionOutput {
	return o
}

func (o RobotApplicationVersionOutput) ToRobotApplicationVersionOutputWithContext(ctx context.Context) RobotApplicationVersionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RobotApplicationVersionInput)(nil)).Elem(), &RobotApplicationVersion{})
	pulumi.RegisterOutputType(RobotApplicationVersionOutput{})
}
