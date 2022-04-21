// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53resolver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Route53Resolver::ResolverConfig.
type ResolverConfig struct {
	pulumi.CustomResourceState

	// ResolverAutodefinedReverseStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.
	AutodefinedReverse ResolverConfigAutodefinedReverseOutput `pulumi:"autodefinedReverse"`
	// Represents the desired status of AutodefinedReverse. The only supported value on creation is DISABLE. Deletion of this resource will return AutodefinedReverse to its default value (ENABLED).
	AutodefinedReverseFlag ResolverConfigAutodefinedReverseFlagOutput `pulumi:"autodefinedReverseFlag"`
	// AccountId
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// ResourceId
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
}

// NewResolverConfig registers a new resource with the given unique name, arguments, and options.
func NewResolverConfig(ctx *pulumi.Context,
	name string, args *ResolverConfigArgs, opts ...pulumi.ResourceOption) (*ResolverConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutodefinedReverseFlag == nil {
		return nil, errors.New("invalid value for required argument 'AutodefinedReverseFlag'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	var resource ResolverConfig
	err := ctx.RegisterResource("aws-native:route53resolver:ResolverConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverConfig gets an existing ResolverConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverConfigState, opts ...pulumi.ResourceOption) (*ResolverConfig, error) {
	var resource ResolverConfig
	err := ctx.ReadResource("aws-native:route53resolver:ResolverConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverConfig resources.
type resolverConfigState struct {
}

type ResolverConfigState struct {
}

func (ResolverConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverConfigState)(nil)).Elem()
}

type resolverConfigArgs struct {
	// Represents the desired status of AutodefinedReverse. The only supported value on creation is DISABLE. Deletion of this resource will return AutodefinedReverse to its default value (ENABLED).
	AutodefinedReverseFlag ResolverConfigAutodefinedReverseFlag `pulumi:"autodefinedReverseFlag"`
	// ResourceId
	ResourceId string `pulumi:"resourceId"`
}

// The set of arguments for constructing a ResolverConfig resource.
type ResolverConfigArgs struct {
	// Represents the desired status of AutodefinedReverse. The only supported value on creation is DISABLE. Deletion of this resource will return AutodefinedReverse to its default value (ENABLED).
	AutodefinedReverseFlag ResolverConfigAutodefinedReverseFlagInput
	// ResourceId
	ResourceId pulumi.StringInput
}

func (ResolverConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverConfigArgs)(nil)).Elem()
}

type ResolverConfigInput interface {
	pulumi.Input

	ToResolverConfigOutput() ResolverConfigOutput
	ToResolverConfigOutputWithContext(ctx context.Context) ResolverConfigOutput
}

func (*ResolverConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverConfig)(nil)).Elem()
}

func (i *ResolverConfig) ToResolverConfigOutput() ResolverConfigOutput {
	return i.ToResolverConfigOutputWithContext(context.Background())
}

func (i *ResolverConfig) ToResolverConfigOutputWithContext(ctx context.Context) ResolverConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverConfigOutput)
}

type ResolverConfigOutput struct{ *pulumi.OutputState }

func (ResolverConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverConfig)(nil)).Elem()
}

func (o ResolverConfigOutput) ToResolverConfigOutput() ResolverConfigOutput {
	return o
}

func (o ResolverConfigOutput) ToResolverConfigOutputWithContext(ctx context.Context) ResolverConfigOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverConfigInput)(nil)).Elem(), &ResolverConfig{})
	pulumi.RegisterOutputType(ResolverConfigOutput{})
}
