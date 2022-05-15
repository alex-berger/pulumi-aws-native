// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iottwinmaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::IoTTwinMaker::Workspace
type Workspace struct {
	pulumi.CustomResourceState

	// The ARN of the workspace.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The date and time when the workspace was created.
	CreationDateTime pulumi.StringOutput `pulumi:"creationDateTime"`
	// The description of the workspace.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ARN of the execution role associated with the workspace.
	Role pulumi.StringOutput `pulumi:"role"`
	// The ARN of the S3 bucket where resources associated with the workspace are stored.
	S3Location pulumi.StringOutput `pulumi:"s3Location"`
	// A map of key-value pairs to associate with a resource.
	Tags pulumi.AnyOutput `pulumi:"tags"`
	// The date and time of the current update.
	UpdateDateTime pulumi.StringOutput `pulumi:"updateDateTime"`
	// The ID of the workspace.
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.S3Location == nil {
		return nil, errors.New("invalid value for required argument 'S3Location'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	var resource Workspace
	err := ctx.RegisterResource("aws-native:iottwinmaker:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspace gets an existing Workspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("aws-native:iottwinmaker:Workspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workspace resources.
type workspaceState struct {
}

type WorkspaceState struct {
}

func (WorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceState)(nil)).Elem()
}

type workspaceArgs struct {
	// The description of the workspace.
	Description *string `pulumi:"description"`
	// The ARN of the execution role associated with the workspace.
	Role string `pulumi:"role"`
	// The ARN of the S3 bucket where resources associated with the workspace are stored.
	S3Location string `pulumi:"s3Location"`
	// A map of key-value pairs to associate with a resource.
	Tags interface{} `pulumi:"tags"`
	// The ID of the workspace.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// The description of the workspace.
	Description pulumi.StringPtrInput
	// The ARN of the execution role associated with the workspace.
	Role pulumi.StringInput
	// The ARN of the S3 bucket where resources associated with the workspace are stored.
	S3Location pulumi.StringInput
	// A map of key-value pairs to associate with a resource.
	Tags pulumi.Input
	// The ID of the workspace.
	WorkspaceId pulumi.StringInput
}

func (WorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceArgs)(nil)).Elem()
}

type WorkspaceInput interface {
	pulumi.Input

	ToWorkspaceOutput() WorkspaceOutput
	ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput
}

func (*Workspace) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (i *Workspace) ToWorkspaceOutput() WorkspaceOutput {
	return i.ToWorkspaceOutputWithContext(context.Background())
}

func (i *Workspace) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceOutput)
}

type WorkspaceOutput struct{ *pulumi.OutputState }

func (WorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (o WorkspaceOutput) ToWorkspaceOutput() WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return o
}

// The ARN of the workspace.
func (o WorkspaceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The date and time when the workspace was created.
func (o WorkspaceOutput) CreationDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.CreationDateTime }).(pulumi.StringOutput)
}

// The description of the workspace.
func (o WorkspaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ARN of the execution role associated with the workspace.
func (o WorkspaceOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// The ARN of the S3 bucket where resources associated with the workspace are stored.
func (o WorkspaceOutput) S3Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.S3Location }).(pulumi.StringOutput)
}

// A map of key-value pairs to associate with a resource.
func (o WorkspaceOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *Workspace) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

// The date and time of the current update.
func (o WorkspaceOutput) UpdateDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.UpdateDateTime }).(pulumi.StringOutput)
}

// The ID of the workspace.
func (o WorkspaceOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceInput)(nil)).Elem(), &Workspace{})
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
