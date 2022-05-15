// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IAM::ManagedPolicy
//
// Deprecated: ManagedPolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ManagedPolicy struct {
	pulumi.CustomResourceState

	Description       pulumi.StringPtrOutput   `pulumi:"description"`
	Groups            pulumi.StringArrayOutput `pulumi:"groups"`
	ManagedPolicyName pulumi.StringPtrOutput   `pulumi:"managedPolicyName"`
	Path              pulumi.StringPtrOutput   `pulumi:"path"`
	PolicyDocument    pulumi.AnyOutput         `pulumi:"policyDocument"`
	Roles             pulumi.StringArrayOutput `pulumi:"roles"`
	Users             pulumi.StringArrayOutput `pulumi:"users"`
}

// NewManagedPolicy registers a new resource with the given unique name, arguments, and options.
func NewManagedPolicy(ctx *pulumi.Context,
	name string, args *ManagedPolicyArgs, opts ...pulumi.ResourceOption) (*ManagedPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyDocument == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDocument'")
	}
	var resource ManagedPolicy
	err := ctx.RegisterResource("aws-native:iam:ManagedPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedPolicy gets an existing ManagedPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedPolicyState, opts ...pulumi.ResourceOption) (*ManagedPolicy, error) {
	var resource ManagedPolicy
	err := ctx.ReadResource("aws-native:iam:ManagedPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedPolicy resources.
type managedPolicyState struct {
}

type ManagedPolicyState struct {
}

func (ManagedPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPolicyState)(nil)).Elem()
}

type managedPolicyArgs struct {
	Description       *string     `pulumi:"description"`
	Groups            []string    `pulumi:"groups"`
	ManagedPolicyName *string     `pulumi:"managedPolicyName"`
	Path              *string     `pulumi:"path"`
	PolicyDocument    interface{} `pulumi:"policyDocument"`
	Roles             []string    `pulumi:"roles"`
	Users             []string    `pulumi:"users"`
}

// The set of arguments for constructing a ManagedPolicy resource.
type ManagedPolicyArgs struct {
	Description       pulumi.StringPtrInput
	Groups            pulumi.StringArrayInput
	ManagedPolicyName pulumi.StringPtrInput
	Path              pulumi.StringPtrInput
	PolicyDocument    pulumi.Input
	Roles             pulumi.StringArrayInput
	Users             pulumi.StringArrayInput
}

func (ManagedPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPolicyArgs)(nil)).Elem()
}

type ManagedPolicyInput interface {
	pulumi.Input

	ToManagedPolicyOutput() ManagedPolicyOutput
	ToManagedPolicyOutputWithContext(ctx context.Context) ManagedPolicyOutput
}

func (*ManagedPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPolicy)(nil)).Elem()
}

func (i *ManagedPolicy) ToManagedPolicyOutput() ManagedPolicyOutput {
	return i.ToManagedPolicyOutputWithContext(context.Background())
}

func (i *ManagedPolicy) ToManagedPolicyOutputWithContext(ctx context.Context) ManagedPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPolicyOutput)
}

type ManagedPolicyOutput struct{ *pulumi.OutputState }

func (ManagedPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPolicy)(nil)).Elem()
}

func (o ManagedPolicyOutput) ToManagedPolicyOutput() ManagedPolicyOutput {
	return o
}

func (o ManagedPolicyOutput) ToManagedPolicyOutputWithContext(ctx context.Context) ManagedPolicyOutput {
	return o
}

func (o ManagedPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ManagedPolicyOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.StringArrayOutput { return v.Groups }).(pulumi.StringArrayOutput)
}

func (o ManagedPolicyOutput) ManagedPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.StringPtrOutput { return v.ManagedPolicyName }).(pulumi.StringPtrOutput)
}

func (o ManagedPolicyOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ManagedPolicyOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.AnyOutput { return v.PolicyDocument }).(pulumi.AnyOutput)
}

func (o ManagedPolicyOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o ManagedPolicyOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedPolicy) pulumi.StringArrayOutput { return v.Users }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedPolicyInput)(nil)).Elem(), &ManagedPolicy{})
	pulumi.RegisterOutputType(ManagedPolicyOutput{})
}
