// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53resolver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Route53Resolver::ResolverRuleAssociation
type ResolverRuleAssociation struct {
	pulumi.CustomResourceState

	// The name of an association between a Resolver rule and a VPC.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Primary Identifier for Resolver Rule Association
	ResolverRuleAssociationId pulumi.StringOutput `pulumi:"resolverRuleAssociationId"`
	// The ID of the Resolver rule that you associated with the VPC that is specified by VPCId.
	ResolverRuleId pulumi.StringOutput `pulumi:"resolverRuleId"`
	// The ID of the VPC that you associated the Resolver rule with.
	VPCId pulumi.StringOutput `pulumi:"vPCId"`
}

// NewResolverRuleAssociation registers a new resource with the given unique name, arguments, and options.
func NewResolverRuleAssociation(ctx *pulumi.Context,
	name string, args *ResolverRuleAssociationArgs, opts ...pulumi.ResourceOption) (*ResolverRuleAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResolverRuleId == nil {
		return nil, errors.New("invalid value for required argument 'ResolverRuleId'")
	}
	if args.VPCId == nil {
		return nil, errors.New("invalid value for required argument 'VPCId'")
	}
	var resource ResolverRuleAssociation
	err := ctx.RegisterResource("aws-native:route53resolver:ResolverRuleAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverRuleAssociation gets an existing ResolverRuleAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverRuleAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverRuleAssociationState, opts ...pulumi.ResourceOption) (*ResolverRuleAssociation, error) {
	var resource ResolverRuleAssociation
	err := ctx.ReadResource("aws-native:route53resolver:ResolverRuleAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverRuleAssociation resources.
type resolverRuleAssociationState struct {
}

type ResolverRuleAssociationState struct {
}

func (ResolverRuleAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverRuleAssociationState)(nil)).Elem()
}

type resolverRuleAssociationArgs struct {
	// The name of an association between a Resolver rule and a VPC.
	Name *string `pulumi:"name"`
	// The ID of the Resolver rule that you associated with the VPC that is specified by VPCId.
	ResolverRuleId string `pulumi:"resolverRuleId"`
	// The ID of the VPC that you associated the Resolver rule with.
	VPCId string `pulumi:"vPCId"`
}

// The set of arguments for constructing a ResolverRuleAssociation resource.
type ResolverRuleAssociationArgs struct {
	// The name of an association between a Resolver rule and a VPC.
	Name pulumi.StringPtrInput
	// The ID of the Resolver rule that you associated with the VPC that is specified by VPCId.
	ResolverRuleId pulumi.StringInput
	// The ID of the VPC that you associated the Resolver rule with.
	VPCId pulumi.StringInput
}

func (ResolverRuleAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverRuleAssociationArgs)(nil)).Elem()
}

type ResolverRuleAssociationInput interface {
	pulumi.Input

	ToResolverRuleAssociationOutput() ResolverRuleAssociationOutput
	ToResolverRuleAssociationOutputWithContext(ctx context.Context) ResolverRuleAssociationOutput
}

func (*ResolverRuleAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverRuleAssociation)(nil)).Elem()
}

func (i *ResolverRuleAssociation) ToResolverRuleAssociationOutput() ResolverRuleAssociationOutput {
	return i.ToResolverRuleAssociationOutputWithContext(context.Background())
}

func (i *ResolverRuleAssociation) ToResolverRuleAssociationOutputWithContext(ctx context.Context) ResolverRuleAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverRuleAssociationOutput)
}

type ResolverRuleAssociationOutput struct{ *pulumi.OutputState }

func (ResolverRuleAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverRuleAssociation)(nil)).Elem()
}

func (o ResolverRuleAssociationOutput) ToResolverRuleAssociationOutput() ResolverRuleAssociationOutput {
	return o
}

func (o ResolverRuleAssociationOutput) ToResolverRuleAssociationOutputWithContext(ctx context.Context) ResolverRuleAssociationOutput {
	return o
}

// The name of an association between a Resolver rule and a VPC.
func (o ResolverRuleAssociationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResolverRuleAssociation) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Primary Identifier for Resolver Rule Association
func (o ResolverRuleAssociationOutput) ResolverRuleAssociationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverRuleAssociation) pulumi.StringOutput { return v.ResolverRuleAssociationId }).(pulumi.StringOutput)
}

// The ID of the Resolver rule that you associated with the VPC that is specified by VPCId.
func (o ResolverRuleAssociationOutput) ResolverRuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverRuleAssociation) pulumi.StringOutput { return v.ResolverRuleId }).(pulumi.StringOutput)
}

// The ID of the VPC that you associated the Resolver rule with.
func (o ResolverRuleAssociationOutput) VPCId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverRuleAssociation) pulumi.StringOutput { return v.VPCId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverRuleAssociationInput)(nil)).Elem(), &ResolverRuleAssociation{})
	pulumi.RegisterOutputType(ResolverRuleAssociationOutput{})
}
