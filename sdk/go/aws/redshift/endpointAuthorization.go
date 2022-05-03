// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes an endpoint authorization for authorizing Redshift-managed VPC endpoint access to a cluster across AWS accounts.
type EndpointAuthorization struct {
	pulumi.CustomResourceState

	// The target AWS account ID to grant or revoke access for.
	Account pulumi.StringOutput `pulumi:"account"`
	// Indicates whether all VPCs in the grantee account are allowed access to the cluster.
	AllowedAllVPCs pulumi.BoolOutput `pulumi:"allowedAllVPCs"`
	// The VPCs allowed access to the cluster.
	AllowedVPCs pulumi.StringArrayOutput `pulumi:"allowedVPCs"`
	// The time (UTC) when the authorization was created.
	AuthorizeTime pulumi.StringOutput `pulumi:"authorizeTime"`
	// The cluster identifier.
	ClusterIdentifier pulumi.StringOutput `pulumi:"clusterIdentifier"`
	// The status of the cluster.
	ClusterStatus pulumi.StringOutput `pulumi:"clusterStatus"`
	// The number of Redshift-managed VPC endpoints created for the authorization.
	EndpointCount pulumi.IntOutput `pulumi:"endpointCount"`
	//  Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// The AWS account ID of the grantee of the cluster.
	Grantee pulumi.StringOutput `pulumi:"grantee"`
	// The AWS account ID of the cluster owner.
	Grantor pulumi.StringOutput `pulumi:"grantor"`
	// The status of the authorization action.
	Status pulumi.StringOutput `pulumi:"status"`
	// The virtual private cloud (VPC) identifiers to grant or revoke access to.
	VpcIds pulumi.StringArrayOutput `pulumi:"vpcIds"`
}

// NewEndpointAuthorization registers a new resource with the given unique name, arguments, and options.
func NewEndpointAuthorization(ctx *pulumi.Context,
	name string, args *EndpointAuthorizationArgs, opts ...pulumi.ResourceOption) (*EndpointAuthorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Account == nil {
		return nil, errors.New("invalid value for required argument 'Account'")
	}
	if args.ClusterIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'ClusterIdentifier'")
	}
	var resource EndpointAuthorization
	err := ctx.RegisterResource("aws-native:redshift:EndpointAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointAuthorization gets an existing EndpointAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointAuthorizationState, opts ...pulumi.ResourceOption) (*EndpointAuthorization, error) {
	var resource EndpointAuthorization
	err := ctx.ReadResource("aws-native:redshift:EndpointAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointAuthorization resources.
type endpointAuthorizationState struct {
}

type EndpointAuthorizationState struct {
}

func (EndpointAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointAuthorizationState)(nil)).Elem()
}

type endpointAuthorizationArgs struct {
	// The target AWS account ID to grant or revoke access for.
	Account string `pulumi:"account"`
	// The cluster identifier.
	ClusterIdentifier string `pulumi:"clusterIdentifier"`
	//  Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted.
	Force *bool `pulumi:"force"`
	// The virtual private cloud (VPC) identifiers to grant or revoke access to.
	VpcIds []string `pulumi:"vpcIds"`
}

// The set of arguments for constructing a EndpointAuthorization resource.
type EndpointAuthorizationArgs struct {
	// The target AWS account ID to grant or revoke access for.
	Account pulumi.StringInput
	// The cluster identifier.
	ClusterIdentifier pulumi.StringInput
	//  Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted.
	Force pulumi.BoolPtrInput
	// The virtual private cloud (VPC) identifiers to grant or revoke access to.
	VpcIds pulumi.StringArrayInput
}

func (EndpointAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointAuthorizationArgs)(nil)).Elem()
}

type EndpointAuthorizationInput interface {
	pulumi.Input

	ToEndpointAuthorizationOutput() EndpointAuthorizationOutput
	ToEndpointAuthorizationOutputWithContext(ctx context.Context) EndpointAuthorizationOutput
}

func (*EndpointAuthorization) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointAuthorization)(nil)).Elem()
}

func (i *EndpointAuthorization) ToEndpointAuthorizationOutput() EndpointAuthorizationOutput {
	return i.ToEndpointAuthorizationOutputWithContext(context.Background())
}

func (i *EndpointAuthorization) ToEndpointAuthorizationOutputWithContext(ctx context.Context) EndpointAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointAuthorizationOutput)
}

type EndpointAuthorizationOutput struct{ *pulumi.OutputState }

func (EndpointAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointAuthorization)(nil)).Elem()
}

func (o EndpointAuthorizationOutput) ToEndpointAuthorizationOutput() EndpointAuthorizationOutput {
	return o
}

func (o EndpointAuthorizationOutput) ToEndpointAuthorizationOutputWithContext(ctx context.Context) EndpointAuthorizationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointAuthorizationInput)(nil)).Elem(), &EndpointAuthorization{})
	pulumi.RegisterOutputType(EndpointAuthorizationOutput{})
}
