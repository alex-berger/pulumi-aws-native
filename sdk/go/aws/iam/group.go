// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IAM::Group
//
// Deprecated: Group is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Group struct {
	pulumi.CustomResourceState

	Arn               pulumi.StringOutput      `pulumi:"arn"`
	GroupName         pulumi.StringPtrOutput   `pulumi:"groupName"`
	ManagedPolicyArns pulumi.StringArrayOutput `pulumi:"managedPolicyArns"`
	Path              pulumi.StringPtrOutput   `pulumi:"path"`
	Policies          GroupPolicyArrayOutput   `pulumi:"policies"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		args = &GroupArgs{}
	}

	var resource Group
	err := ctx.RegisterResource("aws-native:iam:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("aws-native:iam:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
}

type GroupState struct {
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	GroupName         *string       `pulumi:"groupName"`
	ManagedPolicyArns []string      `pulumi:"managedPolicyArns"`
	Path              *string       `pulumi:"path"`
	Policies          []GroupPolicy `pulumi:"policies"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	GroupName         pulumi.StringPtrInput
	ManagedPolicyArns pulumi.StringArrayInput
	Path              pulumi.StringPtrInput
	Policies          GroupPolicyArrayInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterOutputType(GroupOutput{})
}
