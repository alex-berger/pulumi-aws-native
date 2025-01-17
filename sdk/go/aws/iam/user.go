// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IAM::User
//
// Deprecated: User is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type User struct {
	pulumi.CustomResourceState

	Arn                 pulumi.StringOutput       `pulumi:"arn"`
	Groups              pulumi.StringArrayOutput  `pulumi:"groups"`
	LoginProfile        UserLoginProfilePtrOutput `pulumi:"loginProfile"`
	ManagedPolicyArns   pulumi.StringArrayOutput  `pulumi:"managedPolicyArns"`
	Path                pulumi.StringPtrOutput    `pulumi:"path"`
	PermissionsBoundary pulumi.StringPtrOutput    `pulumi:"permissionsBoundary"`
	Policies            UserPolicyArrayOutput     `pulumi:"policies"`
	Tags                UserTagArrayOutput        `pulumi:"tags"`
	UserName            pulumi.StringPtrOutput    `pulumi:"userName"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		args = &UserArgs{}
	}

	var resource User
	err := ctx.RegisterResource("aws-native:iam:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("aws-native:iam:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
}

type UserState struct {
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	Groups              []string          `pulumi:"groups"`
	LoginProfile        *UserLoginProfile `pulumi:"loginProfile"`
	ManagedPolicyArns   []string          `pulumi:"managedPolicyArns"`
	Path                *string           `pulumi:"path"`
	PermissionsBoundary *string           `pulumi:"permissionsBoundary"`
	Policies            []UserPolicy      `pulumi:"policies"`
	Tags                []UserTag         `pulumi:"tags"`
	UserName            *string           `pulumi:"userName"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	Groups              pulumi.StringArrayInput
	LoginProfile        UserLoginProfilePtrInput
	ManagedPolicyArns   pulumi.StringArrayInput
	Path                pulumi.StringPtrInput
	PermissionsBoundary pulumi.StringPtrInput
	Policies            UserPolicyArrayInput
	Tags                UserTagArrayInput
	UserName            pulumi.StringPtrInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

func (o UserOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o UserOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *User) pulumi.StringArrayOutput { return v.Groups }).(pulumi.StringArrayOutput)
}

func (o UserOutput) LoginProfile() UserLoginProfilePtrOutput {
	return o.ApplyT(func(v *User) UserLoginProfilePtrOutput { return v.LoginProfile }).(UserLoginProfilePtrOutput)
}

func (o UserOutput) ManagedPolicyArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *User) pulumi.StringArrayOutput { return v.ManagedPolicyArns }).(pulumi.StringArrayOutput)
}

func (o UserOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

func (o UserOutput) PermissionsBoundary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.PermissionsBoundary }).(pulumi.StringPtrOutput)
}

func (o UserOutput) Policies() UserPolicyArrayOutput {
	return o.ApplyT(func(v *User) UserPolicyArrayOutput { return v.Policies }).(UserPolicyArrayOutput)
}

func (o UserOutput) Tags() UserTagArrayOutput {
	return o.ApplyT(func(v *User) UserTagArrayOutput { return v.Tags }).(UserTagArrayOutput)
}

func (o UserOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterOutputType(UserOutput{})
}
