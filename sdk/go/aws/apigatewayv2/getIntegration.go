// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGatewayV2::Integration
func LookupIntegration(ctx *pulumi.Context, args *LookupIntegrationArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationResult, error) {
	var rv LookupIntegrationResult
	err := ctx.Invoke("aws-native:apigatewayv2:getIntegration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationArgs struct {
	Id string `pulumi:"id"`
}

type LookupIntegrationResult struct {
	ConnectionId                *string               `pulumi:"connectionId"`
	ConnectionType              *string               `pulumi:"connectionType"`
	ContentHandlingStrategy     *string               `pulumi:"contentHandlingStrategy"`
	CredentialsArn              *string               `pulumi:"credentialsArn"`
	Description                 *string               `pulumi:"description"`
	Id                          *string               `pulumi:"id"`
	IntegrationMethod           *string               `pulumi:"integrationMethod"`
	IntegrationSubtype          *string               `pulumi:"integrationSubtype"`
	IntegrationType             *string               `pulumi:"integrationType"`
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

func LookupIntegrationOutput(ctx *pulumi.Context, args LookupIntegrationOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationResult, error) {
			args := v.(LookupIntegrationArgs)
			r, err := LookupIntegration(ctx, &args, opts...)
			var s LookupIntegrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationResultOutput)
}

type LookupIntegrationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupIntegrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationArgs)(nil)).Elem()
}

type LookupIntegrationResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationResult)(nil)).Elem()
}

func (o LookupIntegrationResultOutput) ToLookupIntegrationResultOutput() LookupIntegrationResultOutput {
	return o
}

func (o LookupIntegrationResultOutput) ToLookupIntegrationResultOutputWithContext(ctx context.Context) LookupIntegrationResultOutput {
	return o
}

func (o LookupIntegrationResultOutput) ConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.ConnectionId }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationResultOutput) ConnectionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.ConnectionType }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationResultOutput) ContentHandlingStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.ContentHandlingStrategy }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationResultOutput) CredentialsArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.CredentialsArn }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationResultOutput) IntegrationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.IntegrationMethod }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationResultOutput) IntegrationSubtype() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.IntegrationSubtype }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationResultOutput) IntegrationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.IntegrationType }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationResultOutput) IntegrationUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.IntegrationUri }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationResultOutput) PassthroughBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.PassthroughBehavior }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationResultOutput) PayloadFormatVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.PayloadFormatVersion }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationResultOutput) RequestParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupIntegrationResult) interface{} { return v.RequestParameters }).(pulumi.AnyOutput)
}

func (o LookupIntegrationResultOutput) RequestTemplates() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupIntegrationResult) interface{} { return v.RequestTemplates }).(pulumi.AnyOutput)
}

func (o LookupIntegrationResultOutput) ResponseParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupIntegrationResult) interface{} { return v.ResponseParameters }).(pulumi.AnyOutput)
}

func (o LookupIntegrationResultOutput) TemplateSelectionExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *string { return v.TemplateSelectionExpression }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationResultOutput) TimeoutInMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *int { return v.TimeoutInMillis }).(pulumi.IntPtrOutput)
}

func (o LookupIntegrationResultOutput) TlsConfig() IntegrationTlsConfigPtrOutput {
	return o.ApplyT(func(v LookupIntegrationResult) *IntegrationTlsConfig { return v.TlsConfig }).(IntegrationTlsConfigPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationResultOutput{})
}
