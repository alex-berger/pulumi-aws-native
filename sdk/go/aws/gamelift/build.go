// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GameLift::Build
//
// Deprecated: Build is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Build struct {
	pulumi.CustomResourceState

	Name            pulumi.StringPtrOutput   `pulumi:"name"`
	OperatingSystem pulumi.StringPtrOutput   `pulumi:"operatingSystem"`
	StorageLocation BuildS3LocationPtrOutput `pulumi:"storageLocation"`
	Version         pulumi.StringPtrOutput   `pulumi:"version"`
}

// NewBuild registers a new resource with the given unique name, arguments, and options.
func NewBuild(ctx *pulumi.Context,
	name string, args *BuildArgs, opts ...pulumi.ResourceOption) (*Build, error) {
	if args == nil {
		args = &BuildArgs{}
	}

	var resource Build
	err := ctx.RegisterResource("aws-native:gamelift:Build", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBuild gets an existing Build resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBuild(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildState, opts ...pulumi.ResourceOption) (*Build, error) {
	var resource Build
	err := ctx.ReadResource("aws-native:gamelift:Build", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Build resources.
type buildState struct {
}

type BuildState struct {
}

func (BuildState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildState)(nil)).Elem()
}

type buildArgs struct {
	Name            *string          `pulumi:"name"`
	OperatingSystem *string          `pulumi:"operatingSystem"`
	StorageLocation *BuildS3Location `pulumi:"storageLocation"`
	Version         *string          `pulumi:"version"`
}

// The set of arguments for constructing a Build resource.
type BuildArgs struct {
	Name            pulumi.StringPtrInput
	OperatingSystem pulumi.StringPtrInput
	StorageLocation BuildS3LocationPtrInput
	Version         pulumi.StringPtrInput
}

func (BuildArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildArgs)(nil)).Elem()
}

type BuildInput interface {
	pulumi.Input

	ToBuildOutput() BuildOutput
	ToBuildOutputWithContext(ctx context.Context) BuildOutput
}

func (*Build) ElementType() reflect.Type {
	return reflect.TypeOf((**Build)(nil)).Elem()
}

func (i *Build) ToBuildOutput() BuildOutput {
	return i.ToBuildOutputWithContext(context.Background())
}

func (i *Build) ToBuildOutputWithContext(ctx context.Context) BuildOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildOutput)
}

type BuildOutput struct{ *pulumi.OutputState }

func (BuildOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Build)(nil)).Elem()
}

func (o BuildOutput) ToBuildOutput() BuildOutput {
	return o
}

func (o BuildOutput) ToBuildOutputWithContext(ctx context.Context) BuildOutput {
	return o
}

func (o BuildOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Build) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o BuildOutput) OperatingSystem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Build) pulumi.StringPtrOutput { return v.OperatingSystem }).(pulumi.StringPtrOutput)
}

func (o BuildOutput) StorageLocation() BuildS3LocationPtrOutput {
	return o.ApplyT(func(v *Build) BuildS3LocationPtrOutput { return v.StorageLocation }).(BuildS3LocationPtrOutput)
}

func (o BuildOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Build) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BuildInput)(nil)).Elem(), &Build{})
	pulumi.RegisterOutputType(BuildOutput{})
}
