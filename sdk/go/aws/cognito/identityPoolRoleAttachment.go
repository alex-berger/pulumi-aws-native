// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Cognito::IdentityPoolRoleAttachment
//
// Deprecated: IdentityPoolRoleAttachment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type IdentityPoolRoleAttachment struct {
	pulumi.CustomResourceState

	IdentityPoolId pulumi.StringOutput `pulumi:"identityPoolId"`
	RoleMappings   pulumi.AnyOutput    `pulumi:"roleMappings"`
	Roles          pulumi.AnyOutput    `pulumi:"roles"`
}

// NewIdentityPoolRoleAttachment registers a new resource with the given unique name, arguments, and options.
func NewIdentityPoolRoleAttachment(ctx *pulumi.Context,
	name string, args *IdentityPoolRoleAttachmentArgs, opts ...pulumi.ResourceOption) (*IdentityPoolRoleAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IdentityPoolId == nil {
		return nil, errors.New("invalid value for required argument 'IdentityPoolId'")
	}
	var resource IdentityPoolRoleAttachment
	err := ctx.RegisterResource("aws-native:cognito:IdentityPoolRoleAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityPoolRoleAttachment gets an existing IdentityPoolRoleAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityPoolRoleAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityPoolRoleAttachmentState, opts ...pulumi.ResourceOption) (*IdentityPoolRoleAttachment, error) {
	var resource IdentityPoolRoleAttachment
	err := ctx.ReadResource("aws-native:cognito:IdentityPoolRoleAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityPoolRoleAttachment resources.
type identityPoolRoleAttachmentState struct {
}

type IdentityPoolRoleAttachmentState struct {
}

func (IdentityPoolRoleAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityPoolRoleAttachmentState)(nil)).Elem()
}

type identityPoolRoleAttachmentArgs struct {
	IdentityPoolId string      `pulumi:"identityPoolId"`
	RoleMappings   interface{} `pulumi:"roleMappings"`
	Roles          interface{} `pulumi:"roles"`
}

// The set of arguments for constructing a IdentityPoolRoleAttachment resource.
type IdentityPoolRoleAttachmentArgs struct {
	IdentityPoolId pulumi.StringInput
	RoleMappings   pulumi.Input
	Roles          pulumi.Input
}

func (IdentityPoolRoleAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityPoolRoleAttachmentArgs)(nil)).Elem()
}

type IdentityPoolRoleAttachmentInput interface {
	pulumi.Input

	ToIdentityPoolRoleAttachmentOutput() IdentityPoolRoleAttachmentOutput
	ToIdentityPoolRoleAttachmentOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentOutput
}

func (*IdentityPoolRoleAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityPoolRoleAttachment)(nil)).Elem()
}

func (i *IdentityPoolRoleAttachment) ToIdentityPoolRoleAttachmentOutput() IdentityPoolRoleAttachmentOutput {
	return i.ToIdentityPoolRoleAttachmentOutputWithContext(context.Background())
}

func (i *IdentityPoolRoleAttachment) ToIdentityPoolRoleAttachmentOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPoolRoleAttachmentOutput)
}

type IdentityPoolRoleAttachmentOutput struct{ *pulumi.OutputState }

func (IdentityPoolRoleAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityPoolRoleAttachment)(nil)).Elem()
}

func (o IdentityPoolRoleAttachmentOutput) ToIdentityPoolRoleAttachmentOutput() IdentityPoolRoleAttachmentOutput {
	return o
}

func (o IdentityPoolRoleAttachmentOutput) ToIdentityPoolRoleAttachmentOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityPoolRoleAttachmentInput)(nil)).Elem(), &IdentityPoolRoleAttachment{})
	pulumi.RegisterOutputType(IdentityPoolRoleAttachmentOutput{})
}
