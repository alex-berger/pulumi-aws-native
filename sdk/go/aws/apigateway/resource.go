// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::Resource
type Resource struct {
	pulumi.CustomResourceState

	// The parent resource's identifier.
	ParentId pulumi.StringOutput `pulumi:"parentId"`
	// The last path segment for this resource.
	PathPart pulumi.StringOutput `pulumi:"pathPart"`
	// A unique primary identifier for a Resource
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The ID of the RestApi resource in which you want to create this resource..
	RestApiId pulumi.StringOutput `pulumi:"restApiId"`
}

// NewResource registers a new resource with the given unique name, arguments, and options.
func NewResource(ctx *pulumi.Context,
	name string, args *ResourceArgs, opts ...pulumi.ResourceOption) (*Resource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentId == nil {
		return nil, errors.New("invalid value for required argument 'ParentId'")
	}
	if args.PathPart == nil {
		return nil, errors.New("invalid value for required argument 'PathPart'")
	}
	if args.RestApiId == nil {
		return nil, errors.New("invalid value for required argument 'RestApiId'")
	}
	var resource Resource
	err := ctx.RegisterResource("aws-native:apigateway:Resource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResource gets an existing Resource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceState, opts ...pulumi.ResourceOption) (*Resource, error) {
	var resource Resource
	err := ctx.ReadResource("aws-native:apigateway:Resource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Resource resources.
type resourceState struct {
}

type ResourceState struct {
}

func (ResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceState)(nil)).Elem()
}

type resourceArgs struct {
	// The parent resource's identifier.
	ParentId string `pulumi:"parentId"`
	// The last path segment for this resource.
	PathPart string `pulumi:"pathPart"`
	// The ID of the RestApi resource in which you want to create this resource..
	RestApiId string `pulumi:"restApiId"`
}

// The set of arguments for constructing a Resource resource.
type ResourceArgs struct {
	// The parent resource's identifier.
	ParentId pulumi.StringInput
	// The last path segment for this resource.
	PathPart pulumi.StringInput
	// The ID of the RestApi resource in which you want to create this resource..
	RestApiId pulumi.StringInput
}

func (ResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceArgs)(nil)).Elem()
}

type ResourceInput interface {
	pulumi.Input

	ToResourceOutput() ResourceOutput
	ToResourceOutputWithContext(ctx context.Context) ResourceOutput
}

func (*Resource) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (i *Resource) ToResourceOutput() ResourceOutput {
	return i.ToResourceOutputWithContext(context.Background())
}

func (i *Resource) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceOutput)
}

type ResourceOutput struct{ *pulumi.OutputState }

func (ResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (o ResourceOutput) ToResourceOutput() ResourceOutput {
	return o
}

func (o ResourceOutput) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return o
}

// The parent resource's identifier.
func (o ResourceOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringOutput { return v.ParentId }).(pulumi.StringOutput)
}

// The last path segment for this resource.
func (o ResourceOutput) PathPart() pulumi.StringOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringOutput { return v.PathPart }).(pulumi.StringOutput)
}

// A unique primary identifier for a Resource
func (o ResourceOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// The ID of the RestApi resource in which you want to create this resource..
func (o ResourceOutput) RestApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringOutput { return v.RestApiId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceInput)(nil)).Elem(), &Resource{})
	pulumi.RegisterOutputType(ResourceOutput{})
}
