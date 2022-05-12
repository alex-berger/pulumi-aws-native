// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sso

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for SSO assignmet
type Assignment struct {
	pulumi.CustomResourceState

	// The sso instance that the permission set is owned.
	InstanceArn pulumi.StringOutput `pulumi:"instanceArn"`
	// The permission set that the assignemt will be assigned
	PermissionSetArn pulumi.StringOutput `pulumi:"permissionSetArn"`
	// The assignee's identifier, user id/group id
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// The assignee's type, user/group
	PrincipalType AssignmentPrincipalTypeOutput `pulumi:"principalType"`
	// The account id to be provisioned.
	TargetId pulumi.StringOutput `pulumi:"targetId"`
	// The type of resource to be provsioned to, only aws account now
	TargetType AssignmentTargetTypeOutput `pulumi:"targetType"`
}

// NewAssignment registers a new resource with the given unique name, arguments, and options.
func NewAssignment(ctx *pulumi.Context,
	name string, args *AssignmentArgs, opts ...pulumi.ResourceOption) (*Assignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceArn == nil {
		return nil, errors.New("invalid value for required argument 'InstanceArn'")
	}
	if args.PermissionSetArn == nil {
		return nil, errors.New("invalid value for required argument 'PermissionSetArn'")
	}
	if args.PrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalId'")
	}
	if args.PrincipalType == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalType'")
	}
	if args.TargetId == nil {
		return nil, errors.New("invalid value for required argument 'TargetId'")
	}
	if args.TargetType == nil {
		return nil, errors.New("invalid value for required argument 'TargetType'")
	}
	var resource Assignment
	err := ctx.RegisterResource("aws-native:sso:Assignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssignment gets an existing Assignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssignmentState, opts ...pulumi.ResourceOption) (*Assignment, error) {
	var resource Assignment
	err := ctx.ReadResource("aws-native:sso:Assignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Assignment resources.
type assignmentState struct {
}

type AssignmentState struct {
}

func (AssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentState)(nil)).Elem()
}

type assignmentArgs struct {
	// The sso instance that the permission set is owned.
	InstanceArn string `pulumi:"instanceArn"`
	// The permission set that the assignemt will be assigned
	PermissionSetArn string `pulumi:"permissionSetArn"`
	// The assignee's identifier, user id/group id
	PrincipalId string `pulumi:"principalId"`
	// The assignee's type, user/group
	PrincipalType AssignmentPrincipalType `pulumi:"principalType"`
	// The account id to be provisioned.
	TargetId string `pulumi:"targetId"`
	// The type of resource to be provsioned to, only aws account now
	TargetType AssignmentTargetType `pulumi:"targetType"`
}

// The set of arguments for constructing a Assignment resource.
type AssignmentArgs struct {
	// The sso instance that the permission set is owned.
	InstanceArn pulumi.StringInput
	// The permission set that the assignemt will be assigned
	PermissionSetArn pulumi.StringInput
	// The assignee's identifier, user id/group id
	PrincipalId pulumi.StringInput
	// The assignee's type, user/group
	PrincipalType AssignmentPrincipalTypeInput
	// The account id to be provisioned.
	TargetId pulumi.StringInput
	// The type of resource to be provsioned to, only aws account now
	TargetType AssignmentTargetTypeInput
}

func (AssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentArgs)(nil)).Elem()
}

type AssignmentInput interface {
	pulumi.Input

	ToAssignmentOutput() AssignmentOutput
	ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput
}

func (*Assignment) ElementType() reflect.Type {
	return reflect.TypeOf((**Assignment)(nil)).Elem()
}

func (i *Assignment) ToAssignmentOutput() AssignmentOutput {
	return i.ToAssignmentOutputWithContext(context.Background())
}

func (i *Assignment) ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentOutput)
}

type AssignmentOutput struct{ *pulumi.OutputState }

func (AssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Assignment)(nil)).Elem()
}

func (o AssignmentOutput) ToAssignmentOutput() AssignmentOutput {
	return o
}

func (o AssignmentOutput) ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput {
	return o
}

// The sso instance that the permission set is owned.
func (o AssignmentOutput) InstanceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringOutput { return v.InstanceArn }).(pulumi.StringOutput)
}

// The permission set that the assignemt will be assigned
func (o AssignmentOutput) PermissionSetArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringOutput { return v.PermissionSetArn }).(pulumi.StringOutput)
}

// The assignee's identifier, user id/group id
func (o AssignmentOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

// The assignee's type, user/group
func (o AssignmentOutput) PrincipalType() AssignmentPrincipalTypeOutput {
	return o.ApplyT(func(v *Assignment) AssignmentPrincipalTypeOutput { return v.PrincipalType }).(AssignmentPrincipalTypeOutput)
}

// The account id to be provisioned.
func (o AssignmentOutput) TargetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringOutput { return v.TargetId }).(pulumi.StringOutput)
}

// The type of resource to be provsioned to, only aws account now
func (o AssignmentOutput) TargetType() AssignmentTargetTypeOutput {
	return o.ApplyT(func(v *Assignment) AssignmentTargetTypeOutput { return v.TargetType }).(AssignmentTargetTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssignmentInput)(nil)).Elem(), &Assignment{})
	pulumi.RegisterOutputType(AssignmentOutput{})
}
