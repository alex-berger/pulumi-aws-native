// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::EC2::NetworkInsightsAccessScope
type NetworkInsightsAccessScope struct {
	pulumi.CustomResourceState

	CreatedDate                   pulumi.StringOutput                                         `pulumi:"createdDate"`
	ExcludePaths                  NetworkInsightsAccessScopeAccessScopePathRequestArrayOutput `pulumi:"excludePaths"`
	MatchPaths                    NetworkInsightsAccessScopeAccessScopePathRequestArrayOutput `pulumi:"matchPaths"`
	NetworkInsightsAccessScopeArn pulumi.StringOutput                                         `pulumi:"networkInsightsAccessScopeArn"`
	NetworkInsightsAccessScopeId  pulumi.StringOutput                                         `pulumi:"networkInsightsAccessScopeId"`
	Tags                          NetworkInsightsAccessScopeTagArrayOutput                    `pulumi:"tags"`
	UpdatedDate                   pulumi.StringOutput                                         `pulumi:"updatedDate"`
}

// NewNetworkInsightsAccessScope registers a new resource with the given unique name, arguments, and options.
func NewNetworkInsightsAccessScope(ctx *pulumi.Context,
	name string, args *NetworkInsightsAccessScopeArgs, opts ...pulumi.ResourceOption) (*NetworkInsightsAccessScope, error) {
	if args == nil {
		args = &NetworkInsightsAccessScopeArgs{}
	}

	var resource NetworkInsightsAccessScope
	err := ctx.RegisterResource("aws-native:ec2:NetworkInsightsAccessScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInsightsAccessScope gets an existing NetworkInsightsAccessScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInsightsAccessScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInsightsAccessScopeState, opts ...pulumi.ResourceOption) (*NetworkInsightsAccessScope, error) {
	var resource NetworkInsightsAccessScope
	err := ctx.ReadResource("aws-native:ec2:NetworkInsightsAccessScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInsightsAccessScope resources.
type networkInsightsAccessScopeState struct {
}

type NetworkInsightsAccessScopeState struct {
}

func (NetworkInsightsAccessScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInsightsAccessScopeState)(nil)).Elem()
}

type networkInsightsAccessScopeArgs struct {
	ExcludePaths []NetworkInsightsAccessScopeAccessScopePathRequest `pulumi:"excludePaths"`
	MatchPaths   []NetworkInsightsAccessScopeAccessScopePathRequest `pulumi:"matchPaths"`
	Tags         []NetworkInsightsAccessScopeTag                    `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkInsightsAccessScope resource.
type NetworkInsightsAccessScopeArgs struct {
	ExcludePaths NetworkInsightsAccessScopeAccessScopePathRequestArrayInput
	MatchPaths   NetworkInsightsAccessScopeAccessScopePathRequestArrayInput
	Tags         NetworkInsightsAccessScopeTagArrayInput
}

func (NetworkInsightsAccessScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInsightsAccessScopeArgs)(nil)).Elem()
}

type NetworkInsightsAccessScopeInput interface {
	pulumi.Input

	ToNetworkInsightsAccessScopeOutput() NetworkInsightsAccessScopeOutput
	ToNetworkInsightsAccessScopeOutputWithContext(ctx context.Context) NetworkInsightsAccessScopeOutput
}

func (*NetworkInsightsAccessScope) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInsightsAccessScope)(nil))
}

func (i *NetworkInsightsAccessScope) ToNetworkInsightsAccessScopeOutput() NetworkInsightsAccessScopeOutput {
	return i.ToNetworkInsightsAccessScopeOutputWithContext(context.Background())
}

func (i *NetworkInsightsAccessScope) ToNetworkInsightsAccessScopeOutputWithContext(ctx context.Context) NetworkInsightsAccessScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInsightsAccessScopeOutput)
}

type NetworkInsightsAccessScopeOutput struct{ *pulumi.OutputState }

func (NetworkInsightsAccessScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInsightsAccessScope)(nil))
}

func (o NetworkInsightsAccessScopeOutput) ToNetworkInsightsAccessScopeOutput() NetworkInsightsAccessScopeOutput {
	return o
}

func (o NetworkInsightsAccessScopeOutput) ToNetworkInsightsAccessScopeOutputWithContext(ctx context.Context) NetworkInsightsAccessScopeOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInsightsAccessScopeInput)(nil)).Elem(), &NetworkInsightsAccessScope{})
	pulumi.RegisterOutputType(NetworkInsightsAccessScopeOutput{})
}