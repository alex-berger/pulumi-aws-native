// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGatewayV2::Integration
//
// Deprecated: Integration is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Integration struct {
	pulumi.CustomResourceState

	ApiId                       pulumi.StringOutput           `pulumi:"apiId"`
	ConnectionId                pulumi.StringPtrOutput        `pulumi:"connectionId"`
	ConnectionType              pulumi.StringPtrOutput        `pulumi:"connectionType"`
	ContentHandlingStrategy     pulumi.StringPtrOutput        `pulumi:"contentHandlingStrategy"`
	CredentialsArn              pulumi.StringPtrOutput        `pulumi:"credentialsArn"`
	Description                 pulumi.StringPtrOutput        `pulumi:"description"`
	IntegrationMethod           pulumi.StringPtrOutput        `pulumi:"integrationMethod"`
	IntegrationSubtype          pulumi.StringPtrOutput        `pulumi:"integrationSubtype"`
	IntegrationType             pulumi.StringOutput           `pulumi:"integrationType"`
	IntegrationUri              pulumi.StringPtrOutput        `pulumi:"integrationUri"`
	PassthroughBehavior         pulumi.StringPtrOutput        `pulumi:"passthroughBehavior"`
	PayloadFormatVersion        pulumi.StringPtrOutput        `pulumi:"payloadFormatVersion"`
	RequestParameters           pulumi.AnyOutput              `pulumi:"requestParameters"`
	RequestTemplates            pulumi.AnyOutput              `pulumi:"requestTemplates"`
	ResponseParameters          pulumi.AnyOutput              `pulumi:"responseParameters"`
	TemplateSelectionExpression pulumi.StringPtrOutput        `pulumi:"templateSelectionExpression"`
	TimeoutInMillis             pulumi.IntPtrOutput           `pulumi:"timeoutInMillis"`
	TlsConfig                   IntegrationTlsConfigPtrOutput `pulumi:"tlsConfig"`
}

// NewIntegration registers a new resource with the given unique name, arguments, and options.
func NewIntegration(ctx *pulumi.Context,
	name string, args *IntegrationArgs, opts ...pulumi.ResourceOption) (*Integration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.IntegrationType == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationType'")
	}
	var resource Integration
	err := ctx.RegisterResource("aws-native:apigatewayv2:Integration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegration gets an existing Integration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationState, opts ...pulumi.ResourceOption) (*Integration, error) {
	var resource Integration
	err := ctx.ReadResource("aws-native:apigatewayv2:Integration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Integration resources.
type integrationState struct {
}

type IntegrationState struct {
}

func (IntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationState)(nil)).Elem()
}

type integrationArgs struct {
	ApiId                       string                `pulumi:"apiId"`
	ConnectionId                *string               `pulumi:"connectionId"`
	ConnectionType              *string               `pulumi:"connectionType"`
	ContentHandlingStrategy     *string               `pulumi:"contentHandlingStrategy"`
	CredentialsArn              *string               `pulumi:"credentialsArn"`
	Description                 *string               `pulumi:"description"`
	IntegrationMethod           *string               `pulumi:"integrationMethod"`
	IntegrationSubtype          *string               `pulumi:"integrationSubtype"`
	IntegrationType             string                `pulumi:"integrationType"`
	IntegrationUri              *string               `pulumi:"integrationUri"`
	PassthroughBehavior         *string               `pulumi:"passthroughBehavior"`
	PayloadFormatVersion        *string               `pulumi:"payloadFormatVersion"`
	RequestParameters           interface{}           `pulumi:"requestParameters"`
	RequestTemplates            interface{}           `pulumi:"requestTemplates"`
	ResponseParameters          interface{}           `pulumi:"responseParameters"`
	TemplateSelectionExpression *string               `pulumi:"templateSelectionExpression"`
	TimeoutInMillis             *int                  `pulumi:"timeoutInMillis"`
	TlsConfig                   *IntegrationTlsConfig `pulumi:"tlsConfig"`
}

// The set of arguments for constructing a Integration resource.
type IntegrationArgs struct {
	ApiId                       pulumi.StringInput
	ConnectionId                pulumi.StringPtrInput
	ConnectionType              pulumi.StringPtrInput
	ContentHandlingStrategy     pulumi.StringPtrInput
	CredentialsArn              pulumi.StringPtrInput
	Description                 pulumi.StringPtrInput
	IntegrationMethod           pulumi.StringPtrInput
	IntegrationSubtype          pulumi.StringPtrInput
	IntegrationType             pulumi.StringInput
	IntegrationUri              pulumi.StringPtrInput
	PassthroughBehavior         pulumi.StringPtrInput
	PayloadFormatVersion        pulumi.StringPtrInput
	RequestParameters           pulumi.Input
	RequestTemplates            pulumi.Input
	ResponseParameters          pulumi.Input
	TemplateSelectionExpression pulumi.StringPtrInput
	TimeoutInMillis             pulumi.IntPtrInput
	TlsConfig                   IntegrationTlsConfigPtrInput
}

func (IntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationArgs)(nil)).Elem()
}

type IntegrationInput interface {
	pulumi.Input

	ToIntegrationOutput() IntegrationOutput
	ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput
}

func (*Integration) ElementType() reflect.Type {
	return reflect.TypeOf((**Integration)(nil)).Elem()
}

func (i *Integration) ToIntegrationOutput() IntegrationOutput {
	return i.ToIntegrationOutputWithContext(context.Background())
}

func (i *Integration) ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationOutput)
}

type IntegrationOutput struct{ *pulumi.OutputState }

func (IntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Integration)(nil)).Elem()
}

func (o IntegrationOutput) ToIntegrationOutput() IntegrationOutput {
	return o
}

func (o IntegrationOutput) ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationInput)(nil)).Elem(), &Integration{})
	pulumi.RegisterOutputType(IntegrationOutput{})
}
