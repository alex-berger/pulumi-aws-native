// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Lambda::Permission
//
// Deprecated: Permission is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Permission struct {
	pulumi.CustomResourceState

	Action              pulumi.StringOutput    `pulumi:"action"`
	EventSourceToken    pulumi.StringPtrOutput `pulumi:"eventSourceToken"`
	FunctionName        pulumi.StringOutput    `pulumi:"functionName"`
	FunctionUrlAuthType pulumi.StringPtrOutput `pulumi:"functionUrlAuthType"`
	Principal           pulumi.StringOutput    `pulumi:"principal"`
	PrincipalOrgID      pulumi.StringPtrOutput `pulumi:"principalOrgID"`
	SourceAccount       pulumi.StringPtrOutput `pulumi:"sourceAccount"`
	SourceArn           pulumi.StringPtrOutput `pulumi:"sourceArn"`
}

// NewPermission registers a new resource with the given unique name, arguments, and options.
func NewPermission(ctx *pulumi.Context,
	name string, args *PermissionArgs, opts ...pulumi.ResourceOption) (*Permission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.FunctionName == nil {
		return nil, errors.New("invalid value for required argument 'FunctionName'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	var resource Permission
	err := ctx.RegisterResource("aws-native:lambda:Permission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPermission gets an existing Permission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PermissionState, opts ...pulumi.ResourceOption) (*Permission, error) {
	var resource Permission
	err := ctx.ReadResource("aws-native:lambda:Permission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Permission resources.
type permissionState struct {
}

type PermissionState struct {
}

func (PermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionState)(nil)).Elem()
}

type permissionArgs struct {
	Action              string  `pulumi:"action"`
	EventSourceToken    *string `pulumi:"eventSourceToken"`
	FunctionName        string  `pulumi:"functionName"`
	FunctionUrlAuthType *string `pulumi:"functionUrlAuthType"`
	Principal           string  `pulumi:"principal"`
	PrincipalOrgID      *string `pulumi:"principalOrgID"`
	SourceAccount       *string `pulumi:"sourceAccount"`
	SourceArn           *string `pulumi:"sourceArn"`
}

// The set of arguments for constructing a Permission resource.
type PermissionArgs struct {
	Action              pulumi.StringInput
	EventSourceToken    pulumi.StringPtrInput
	FunctionName        pulumi.StringInput
	FunctionUrlAuthType pulumi.StringPtrInput
	Principal           pulumi.StringInput
	PrincipalOrgID      pulumi.StringPtrInput
	SourceAccount       pulumi.StringPtrInput
	SourceArn           pulumi.StringPtrInput
}

func (PermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionArgs)(nil)).Elem()
}

type PermissionInput interface {
	pulumi.Input

	ToPermissionOutput() PermissionOutput
	ToPermissionOutputWithContext(ctx context.Context) PermissionOutput
}

func (*Permission) ElementType() reflect.Type {
	return reflect.TypeOf((**Permission)(nil)).Elem()
}

func (i *Permission) ToPermissionOutput() PermissionOutput {
	return i.ToPermissionOutputWithContext(context.Background())
}

func (i *Permission) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionOutput)
}

type PermissionOutput struct{ *pulumi.OutputState }

func (PermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Permission)(nil)).Elem()
}

func (o PermissionOutput) ToPermissionOutput() PermissionOutput {
	return o
}

func (o PermissionOutput) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionInput)(nil)).Elem(), &Permission{})
	pulumi.RegisterOutputType(PermissionOutput{})
}
