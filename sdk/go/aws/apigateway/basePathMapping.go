// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::BasePathMapping
type BasePathMapping struct {
	pulumi.CustomResourceState

	// The base path name that callers of the API must provide in the URL after the domain name.
	BasePath pulumi.StringPtrOutput `pulumi:"basePath"`
	// The DomainName of an AWS::ApiGateway::DomainName resource.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The ID of the API.
	RestApiId pulumi.StringPtrOutput `pulumi:"restApiId"`
	// The name of the API's stage.
	Stage pulumi.StringPtrOutput `pulumi:"stage"`
}

// NewBasePathMapping registers a new resource with the given unique name, arguments, and options.
func NewBasePathMapping(ctx *pulumi.Context,
	name string, args *BasePathMappingArgs, opts ...pulumi.ResourceOption) (*BasePathMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	var resource BasePathMapping
	err := ctx.RegisterResource("aws-native:apigateway:BasePathMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBasePathMapping gets an existing BasePathMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBasePathMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BasePathMappingState, opts ...pulumi.ResourceOption) (*BasePathMapping, error) {
	var resource BasePathMapping
	err := ctx.ReadResource("aws-native:apigateway:BasePathMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BasePathMapping resources.
type basePathMappingState struct {
}

type BasePathMappingState struct {
}

func (BasePathMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*basePathMappingState)(nil)).Elem()
}

type basePathMappingArgs struct {
	// The base path name that callers of the API must provide in the URL after the domain name.
	BasePath *string `pulumi:"basePath"`
	// The DomainName of an AWS::ApiGateway::DomainName resource.
	DomainName string `pulumi:"domainName"`
	// The ID of the API.
	RestApiId *string `pulumi:"restApiId"`
	// The name of the API's stage.
	Stage *string `pulumi:"stage"`
}

// The set of arguments for constructing a BasePathMapping resource.
type BasePathMappingArgs struct {
	// The base path name that callers of the API must provide in the URL after the domain name.
	BasePath pulumi.StringPtrInput
	// The DomainName of an AWS::ApiGateway::DomainName resource.
	DomainName pulumi.StringInput
	// The ID of the API.
	RestApiId pulumi.StringPtrInput
	// The name of the API's stage.
	Stage pulumi.StringPtrInput
}

func (BasePathMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*basePathMappingArgs)(nil)).Elem()
}

type BasePathMappingInput interface {
	pulumi.Input

	ToBasePathMappingOutput() BasePathMappingOutput
	ToBasePathMappingOutputWithContext(ctx context.Context) BasePathMappingOutput
}

func (*BasePathMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**BasePathMapping)(nil)).Elem()
}

func (i *BasePathMapping) ToBasePathMappingOutput() BasePathMappingOutput {
	return i.ToBasePathMappingOutputWithContext(context.Background())
}

func (i *BasePathMapping) ToBasePathMappingOutputWithContext(ctx context.Context) BasePathMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasePathMappingOutput)
}

type BasePathMappingOutput struct{ *pulumi.OutputState }

func (BasePathMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BasePathMapping)(nil)).Elem()
}

func (o BasePathMappingOutput) ToBasePathMappingOutput() BasePathMappingOutput {
	return o
}

func (o BasePathMappingOutput) ToBasePathMappingOutputWithContext(ctx context.Context) BasePathMappingOutput {
	return o
}

// The base path name that callers of the API must provide in the URL after the domain name.
func (o BasePathMappingOutput) BasePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasePathMapping) pulumi.StringPtrOutput { return v.BasePath }).(pulumi.StringPtrOutput)
}

// The DomainName of an AWS::ApiGateway::DomainName resource.
func (o BasePathMappingOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *BasePathMapping) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The ID of the API.
func (o BasePathMappingOutput) RestApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasePathMapping) pulumi.StringPtrOutput { return v.RestApiId }).(pulumi.StringPtrOutput)
}

// The name of the API's stage.
func (o BasePathMappingOutput) Stage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasePathMapping) pulumi.StringPtrOutput { return v.Stage }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BasePathMappingInput)(nil)).Elem(), &BasePathMapping{})
	pulumi.RegisterOutputType(BasePathMappingOutput{})
}
