// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGatewayV2::Authorizer
func LookupAuthorizer(ctx *pulumi.Context, args *LookupAuthorizerArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizerResult, error) {
	var rv LookupAuthorizerResult
	err := ctx.Invoke("aws-native:apigatewayv2:getAuthorizer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAuthorizerArgs struct {
	Id string `pulumi:"id"`
}

type LookupAuthorizerResult struct {
	AuthorizerCredentialsArn       *string                     `pulumi:"authorizerCredentialsArn"`
	AuthorizerPayloadFormatVersion *string                     `pulumi:"authorizerPayloadFormatVersion"`
	AuthorizerResultTtlInSeconds   *int                        `pulumi:"authorizerResultTtlInSeconds"`
	AuthorizerType                 *string                     `pulumi:"authorizerType"`
	AuthorizerUri                  *string                     `pulumi:"authorizerUri"`
	EnableSimpleResponses          *bool                       `pulumi:"enableSimpleResponses"`
	Id                             *string                     `pulumi:"id"`
	IdentitySource                 []string                    `pulumi:"identitySource"`
	IdentityValidationExpression   *string                     `pulumi:"identityValidationExpression"`
	JwtConfiguration               *AuthorizerJWTConfiguration `pulumi:"jwtConfiguration"`
	Name                           *string                     `pulumi:"name"`
}

func LookupAuthorizerOutput(ctx *pulumi.Context, args LookupAuthorizerOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAuthorizerResult, error) {
			args := v.(LookupAuthorizerArgs)
			r, err := LookupAuthorizer(ctx, &args, opts...)
			var s LookupAuthorizerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAuthorizerResultOutput)
}

type LookupAuthorizerOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupAuthorizerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizerArgs)(nil)).Elem()
}

type LookupAuthorizerResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizerResult)(nil)).Elem()
}

func (o LookupAuthorizerResultOutput) ToLookupAuthorizerResultOutput() LookupAuthorizerResultOutput {
	return o
}

func (o LookupAuthorizerResultOutput) ToLookupAuthorizerResultOutputWithContext(ctx context.Context) LookupAuthorizerResultOutput {
	return o
}

func (o LookupAuthorizerResultOutput) AuthorizerCredentialsArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *string { return v.AuthorizerCredentialsArn }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizerResultOutput) AuthorizerPayloadFormatVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *string { return v.AuthorizerPayloadFormatVersion }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizerResultOutput) AuthorizerResultTtlInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *int { return v.AuthorizerResultTtlInSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupAuthorizerResultOutput) AuthorizerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *string { return v.AuthorizerType }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizerResultOutput) AuthorizerUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *string { return v.AuthorizerUri }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizerResultOutput) EnableSimpleResponses() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *bool { return v.EnableSimpleResponses }).(pulumi.BoolPtrOutput)
}

func (o LookupAuthorizerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizerResultOutput) IdentitySource() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) []string { return v.IdentitySource }).(pulumi.StringArrayOutput)
}

func (o LookupAuthorizerResultOutput) IdentityValidationExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *string { return v.IdentityValidationExpression }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizerResultOutput) JwtConfiguration() AuthorizerJWTConfigurationPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *AuthorizerJWTConfiguration { return v.JwtConfiguration }).(AuthorizerJWTConfigurationPtrOutput)
}

func (o LookupAuthorizerResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizerResultOutput{})
}
