// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::Workteam
//
// Deprecated: Workteam is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Workteam struct {
	pulumi.CustomResourceState

	Description               pulumi.StringPtrOutput                     `pulumi:"description"`
	MemberDefinitions         WorkteamMemberDefinitionArrayOutput        `pulumi:"memberDefinitions"`
	NotificationConfiguration WorkteamNotificationConfigurationPtrOutput `pulumi:"notificationConfiguration"`
	Tags                      WorkteamTagArrayOutput                     `pulumi:"tags"`
	WorkteamName              pulumi.StringPtrOutput                     `pulumi:"workteamName"`
}

// NewWorkteam registers a new resource with the given unique name, arguments, and options.
func NewWorkteam(ctx *pulumi.Context,
	name string, args *WorkteamArgs, opts ...pulumi.ResourceOption) (*Workteam, error) {
	if args == nil {
		args = &WorkteamArgs{}
	}

	var resource Workteam
	err := ctx.RegisterResource("aws-native:sagemaker:Workteam", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkteam gets an existing Workteam resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkteam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkteamState, opts ...pulumi.ResourceOption) (*Workteam, error) {
	var resource Workteam
	err := ctx.ReadResource("aws-native:sagemaker:Workteam", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workteam resources.
type workteamState struct {
}

type WorkteamState struct {
}

func (WorkteamState) ElementType() reflect.Type {
	return reflect.TypeOf((*workteamState)(nil)).Elem()
}

type workteamArgs struct {
	Description               *string                            `pulumi:"description"`
	MemberDefinitions         []WorkteamMemberDefinition         `pulumi:"memberDefinitions"`
	NotificationConfiguration *WorkteamNotificationConfiguration `pulumi:"notificationConfiguration"`
	Tags                      []WorkteamTag                      `pulumi:"tags"`
	WorkteamName              *string                            `pulumi:"workteamName"`
}

// The set of arguments for constructing a Workteam resource.
type WorkteamArgs struct {
	Description               pulumi.StringPtrInput
	MemberDefinitions         WorkteamMemberDefinitionArrayInput
	NotificationConfiguration WorkteamNotificationConfigurationPtrInput
	Tags                      WorkteamTagArrayInput
	WorkteamName              pulumi.StringPtrInput
}

func (WorkteamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workteamArgs)(nil)).Elem()
}

type WorkteamInput interface {
	pulumi.Input

	ToWorkteamOutput() WorkteamOutput
	ToWorkteamOutputWithContext(ctx context.Context) WorkteamOutput
}

func (*Workteam) ElementType() reflect.Type {
	return reflect.TypeOf((**Workteam)(nil)).Elem()
}

func (i *Workteam) ToWorkteamOutput() WorkteamOutput {
	return i.ToWorkteamOutputWithContext(context.Background())
}

func (i *Workteam) ToWorkteamOutputWithContext(ctx context.Context) WorkteamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkteamOutput)
}

type WorkteamOutput struct{ *pulumi.OutputState }

func (WorkteamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workteam)(nil)).Elem()
}

func (o WorkteamOutput) ToWorkteamOutput() WorkteamOutput {
	return o
}

func (o WorkteamOutput) ToWorkteamOutputWithContext(ctx context.Context) WorkteamOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkteamInput)(nil)).Elem(), &Workteam{})
	pulumi.RegisterOutputType(WorkteamOutput{})
}
