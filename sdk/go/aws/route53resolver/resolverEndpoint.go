// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53resolver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Route53Resolver::ResolverEndpoint
//
// Deprecated: ResolverEndpoint is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ResolverEndpoint struct {
	pulumi.CustomResourceState

	Arn                pulumi.StringOutput                         `pulumi:"arn"`
	Direction          pulumi.StringOutput                         `pulumi:"direction"`
	HostVPCId          pulumi.StringOutput                         `pulumi:"hostVPCId"`
	IpAddressCount     pulumi.StringOutput                         `pulumi:"ipAddressCount"`
	IpAddresses        ResolverEndpointIpAddressRequestArrayOutput `pulumi:"ipAddresses"`
	Name               pulumi.StringPtrOutput                      `pulumi:"name"`
	ResolverEndpointId pulumi.StringOutput                         `pulumi:"resolverEndpointId"`
	SecurityGroupIds   pulumi.StringArrayOutput                    `pulumi:"securityGroupIds"`
	Tags               ResolverEndpointTagArrayOutput              `pulumi:"tags"`
}

// NewResolverEndpoint registers a new resource with the given unique name, arguments, and options.
func NewResolverEndpoint(ctx *pulumi.Context,
	name string, args *ResolverEndpointArgs, opts ...pulumi.ResourceOption) (*ResolverEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
	}
	if args.IpAddresses == nil {
		return nil, errors.New("invalid value for required argument 'IpAddresses'")
	}
	if args.SecurityGroupIds == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupIds'")
	}
	var resource ResolverEndpoint
	err := ctx.RegisterResource("aws-native:route53resolver:ResolverEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverEndpoint gets an existing ResolverEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverEndpointState, opts ...pulumi.ResourceOption) (*ResolverEndpoint, error) {
	var resource ResolverEndpoint
	err := ctx.ReadResource("aws-native:route53resolver:ResolverEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverEndpoint resources.
type resolverEndpointState struct {
}

type ResolverEndpointState struct {
}

func (ResolverEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverEndpointState)(nil)).Elem()
}

type resolverEndpointArgs struct {
	Direction        string                             `pulumi:"direction"`
	IpAddresses      []ResolverEndpointIpAddressRequest `pulumi:"ipAddresses"`
	Name             *string                            `pulumi:"name"`
	SecurityGroupIds []string                           `pulumi:"securityGroupIds"`
	Tags             []ResolverEndpointTag              `pulumi:"tags"`
}

// The set of arguments for constructing a ResolverEndpoint resource.
type ResolverEndpointArgs struct {
	Direction        pulumi.StringInput
	IpAddresses      ResolverEndpointIpAddressRequestArrayInput
	Name             pulumi.StringPtrInput
	SecurityGroupIds pulumi.StringArrayInput
	Tags             ResolverEndpointTagArrayInput
}

func (ResolverEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverEndpointArgs)(nil)).Elem()
}

type ResolverEndpointInput interface {
	pulumi.Input

	ToResolverEndpointOutput() ResolverEndpointOutput
	ToResolverEndpointOutputWithContext(ctx context.Context) ResolverEndpointOutput
}

func (*ResolverEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverEndpoint)(nil)).Elem()
}

func (i *ResolverEndpoint) ToResolverEndpointOutput() ResolverEndpointOutput {
	return i.ToResolverEndpointOutputWithContext(context.Background())
}

func (i *ResolverEndpoint) ToResolverEndpointOutputWithContext(ctx context.Context) ResolverEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverEndpointOutput)
}

type ResolverEndpointOutput struct{ *pulumi.OutputState }

func (ResolverEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverEndpoint)(nil)).Elem()
}

func (o ResolverEndpointOutput) ToResolverEndpointOutput() ResolverEndpointOutput {
	return o
}

func (o ResolverEndpointOutput) ToResolverEndpointOutputWithContext(ctx context.Context) ResolverEndpointOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverEndpointInput)(nil)).Elem(), &ResolverEndpoint{})
	pulumi.RegisterOutputType(ResolverEndpointOutput{})
}
