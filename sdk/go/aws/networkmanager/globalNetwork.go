// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networkmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::NetworkManager::GlobalNetwork type specifies a global network of the user's account
type GlobalNetwork struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the global network.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the global network.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The tags for the global network.
	Tags GlobalNetworkTagArrayOutput `pulumi:"tags"`
}

// NewGlobalNetwork registers a new resource with the given unique name, arguments, and options.
func NewGlobalNetwork(ctx *pulumi.Context,
	name string, args *GlobalNetworkArgs, opts ...pulumi.ResourceOption) (*GlobalNetwork, error) {
	if args == nil {
		args = &GlobalNetworkArgs{}
	}

	var resource GlobalNetwork
	err := ctx.RegisterResource("aws-native:networkmanager:GlobalNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalNetwork gets an existing GlobalNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalNetworkState, opts ...pulumi.ResourceOption) (*GlobalNetwork, error) {
	var resource GlobalNetwork
	err := ctx.ReadResource("aws-native:networkmanager:GlobalNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalNetwork resources.
type globalNetworkState struct {
}

type GlobalNetworkState struct {
}

func (GlobalNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalNetworkState)(nil)).Elem()
}

type globalNetworkArgs struct {
	// The description of the global network.
	Description *string `pulumi:"description"`
	// The tags for the global network.
	Tags []GlobalNetworkTag `pulumi:"tags"`
}

// The set of arguments for constructing a GlobalNetwork resource.
type GlobalNetworkArgs struct {
	// The description of the global network.
	Description pulumi.StringPtrInput
	// The tags for the global network.
	Tags GlobalNetworkTagArrayInput
}

func (GlobalNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalNetworkArgs)(nil)).Elem()
}

type GlobalNetworkInput interface {
	pulumi.Input

	ToGlobalNetworkOutput() GlobalNetworkOutput
	ToGlobalNetworkOutputWithContext(ctx context.Context) GlobalNetworkOutput
}

func (*GlobalNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalNetwork)(nil)).Elem()
}

func (i *GlobalNetwork) ToGlobalNetworkOutput() GlobalNetworkOutput {
	return i.ToGlobalNetworkOutputWithContext(context.Background())
}

func (i *GlobalNetwork) ToGlobalNetworkOutputWithContext(ctx context.Context) GlobalNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalNetworkOutput)
}

type GlobalNetworkOutput struct{ *pulumi.OutputState }

func (GlobalNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalNetwork)(nil)).Elem()
}

func (o GlobalNetworkOutput) ToGlobalNetworkOutput() GlobalNetworkOutput {
	return o
}

func (o GlobalNetworkOutput) ToGlobalNetworkOutputWithContext(ctx context.Context) GlobalNetworkOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalNetworkInput)(nil)).Elem(), &GlobalNetwork{})
	pulumi.RegisterOutputType(GlobalNetworkOutput{})
}
