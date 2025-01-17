// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transfer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Transfer::Server
//
// Deprecated: Server is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Server struct {
	pulumi.CustomResourceState

	Arn                           pulumi.StringOutput                    `pulumi:"arn"`
	Certificate                   pulumi.StringPtrOutput                 `pulumi:"certificate"`
	Domain                        pulumi.StringPtrOutput                 `pulumi:"domain"`
	EndpointDetails               ServerEndpointDetailsPtrOutput         `pulumi:"endpointDetails"`
	EndpointType                  pulumi.StringPtrOutput                 `pulumi:"endpointType"`
	IdentityProviderDetails       ServerIdentityProviderDetailsPtrOutput `pulumi:"identityProviderDetails"`
	IdentityProviderType          pulumi.StringPtrOutput                 `pulumi:"identityProviderType"`
	LoggingRole                   pulumi.StringPtrOutput                 `pulumi:"loggingRole"`
	PostAuthenticationLoginBanner pulumi.StringPtrOutput                 `pulumi:"postAuthenticationLoginBanner"`
	PreAuthenticationLoginBanner  pulumi.StringPtrOutput                 `pulumi:"preAuthenticationLoginBanner"`
	ProtocolDetails               ServerProtocolDetailsPtrOutput         `pulumi:"protocolDetails"`
	Protocols                     ServerProtocolArrayOutput              `pulumi:"protocols"`
	SecurityPolicyName            pulumi.StringPtrOutput                 `pulumi:"securityPolicyName"`
	ServerId                      pulumi.StringOutput                    `pulumi:"serverId"`
	Tags                          ServerTagArrayOutput                   `pulumi:"tags"`
	WorkflowDetails               ServerWorkflowDetailsPtrOutput         `pulumi:"workflowDetails"`
}

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil {
		args = &ServerArgs{}
	}

	var resource Server
	err := ctx.RegisterResource("aws-native:transfer:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServer gets an existing Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("aws-native:transfer:Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Server resources.
type serverState struct {
}

type ServerState struct {
}

func (ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverState)(nil)).Elem()
}

type serverArgs struct {
	Certificate                   *string                        `pulumi:"certificate"`
	Domain                        *string                        `pulumi:"domain"`
	EndpointDetails               *ServerEndpointDetails         `pulumi:"endpointDetails"`
	EndpointType                  *string                        `pulumi:"endpointType"`
	IdentityProviderDetails       *ServerIdentityProviderDetails `pulumi:"identityProviderDetails"`
	IdentityProviderType          *string                        `pulumi:"identityProviderType"`
	LoggingRole                   *string                        `pulumi:"loggingRole"`
	PostAuthenticationLoginBanner *string                        `pulumi:"postAuthenticationLoginBanner"`
	PreAuthenticationLoginBanner  *string                        `pulumi:"preAuthenticationLoginBanner"`
	ProtocolDetails               *ServerProtocolDetails         `pulumi:"protocolDetails"`
	Protocols                     []ServerProtocol               `pulumi:"protocols"`
	SecurityPolicyName            *string                        `pulumi:"securityPolicyName"`
	Tags                          []ServerTag                    `pulumi:"tags"`
	WorkflowDetails               *ServerWorkflowDetails         `pulumi:"workflowDetails"`
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	Certificate                   pulumi.StringPtrInput
	Domain                        pulumi.StringPtrInput
	EndpointDetails               ServerEndpointDetailsPtrInput
	EndpointType                  pulumi.StringPtrInput
	IdentityProviderDetails       ServerIdentityProviderDetailsPtrInput
	IdentityProviderType          pulumi.StringPtrInput
	LoggingRole                   pulumi.StringPtrInput
	PostAuthenticationLoginBanner pulumi.StringPtrInput
	PreAuthenticationLoginBanner  pulumi.StringPtrInput
	ProtocolDetails               ServerProtocolDetailsPtrInput
	Protocols                     ServerProtocolArrayInput
	SecurityPolicyName            pulumi.StringPtrInput
	Tags                          ServerTagArrayInput
	WorkflowDetails               ServerWorkflowDetailsPtrInput
}

func (ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverArgs)(nil)).Elem()
}

type ServerInput interface {
	pulumi.Input

	ToServerOutput() ServerOutput
	ToServerOutputWithContext(ctx context.Context) ServerOutput
}

func (*Server) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (i *Server) ToServerOutput() ServerOutput {
	return i.ToServerOutputWithContext(context.Background())
}

func (i *Server) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerOutput)
}

type ServerOutput struct{ *pulumi.OutputState }

func (ServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (o ServerOutput) ToServerOutput() ServerOutput {
	return o
}

func (o ServerOutput) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return o
}

func (o ServerOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ServerOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Certificate }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) EndpointDetails() ServerEndpointDetailsPtrOutput {
	return o.ApplyT(func(v *Server) ServerEndpointDetailsPtrOutput { return v.EndpointDetails }).(ServerEndpointDetailsPtrOutput)
}

func (o ServerOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.EndpointType }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) IdentityProviderDetails() ServerIdentityProviderDetailsPtrOutput {
	return o.ApplyT(func(v *Server) ServerIdentityProviderDetailsPtrOutput { return v.IdentityProviderDetails }).(ServerIdentityProviderDetailsPtrOutput)
}

func (o ServerOutput) IdentityProviderType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.IdentityProviderType }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) LoggingRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.LoggingRole }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) PostAuthenticationLoginBanner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.PostAuthenticationLoginBanner }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) PreAuthenticationLoginBanner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.PreAuthenticationLoginBanner }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) ProtocolDetails() ServerProtocolDetailsPtrOutput {
	return o.ApplyT(func(v *Server) ServerProtocolDetailsPtrOutput { return v.ProtocolDetails }).(ServerProtocolDetailsPtrOutput)
}

func (o ServerOutput) Protocols() ServerProtocolArrayOutput {
	return o.ApplyT(func(v *Server) ServerProtocolArrayOutput { return v.Protocols }).(ServerProtocolArrayOutput)
}

func (o ServerOutput) SecurityPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.SecurityPolicyName }).(pulumi.StringPtrOutput)
}

func (o ServerOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.ServerId }).(pulumi.StringOutput)
}

func (o ServerOutput) Tags() ServerTagArrayOutput {
	return o.ApplyT(func(v *Server) ServerTagArrayOutput { return v.Tags }).(ServerTagArrayOutput)
}

func (o ServerOutput) WorkflowDetails() ServerWorkflowDetailsPtrOutput {
	return o.ApplyT(func(v *Server) ServerWorkflowDetailsPtrOutput { return v.WorkflowDetails }).(ServerWorkflowDetailsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerInput)(nil)).Elem(), &Server{})
	pulumi.RegisterOutputType(ServerOutput{})
}
