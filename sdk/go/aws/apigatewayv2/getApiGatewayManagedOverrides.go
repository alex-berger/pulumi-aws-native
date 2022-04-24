// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGatewayV2::ApiGatewayManagedOverrides
func LookupApiGatewayManagedOverrides(ctx *pulumi.Context, args *LookupApiGatewayManagedOverridesArgs, opts ...pulumi.InvokeOption) (*LookupApiGatewayManagedOverridesResult, error) {
	var rv LookupApiGatewayManagedOverridesResult
	err := ctx.Invoke("aws-native:apigatewayv2:getApiGatewayManagedOverrides", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiGatewayManagedOverridesArgs struct {
	Id string `pulumi:"id"`
}

type LookupApiGatewayManagedOverridesResult struct {
	Id          *string                                         `pulumi:"id"`
	Integration *ApiGatewayManagedOverridesIntegrationOverrides `pulumi:"integration"`
	Route       *ApiGatewayManagedOverridesRouteOverrides       `pulumi:"route"`
	Stage       *ApiGatewayManagedOverridesStageOverrides       `pulumi:"stage"`
}

func LookupApiGatewayManagedOverridesOutput(ctx *pulumi.Context, args LookupApiGatewayManagedOverridesOutputArgs, opts ...pulumi.InvokeOption) LookupApiGatewayManagedOverridesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiGatewayManagedOverridesResult, error) {
			args := v.(LookupApiGatewayManagedOverridesArgs)
			r, err := LookupApiGatewayManagedOverrides(ctx, &args, opts...)
			var s LookupApiGatewayManagedOverridesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiGatewayManagedOverridesResultOutput)
}

type LookupApiGatewayManagedOverridesOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupApiGatewayManagedOverridesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiGatewayManagedOverridesArgs)(nil)).Elem()
}

type LookupApiGatewayManagedOverridesResultOutput struct{ *pulumi.OutputState }

func (LookupApiGatewayManagedOverridesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiGatewayManagedOverridesResult)(nil)).Elem()
}

func (o LookupApiGatewayManagedOverridesResultOutput) ToLookupApiGatewayManagedOverridesResultOutput() LookupApiGatewayManagedOverridesResultOutput {
	return o
}

func (o LookupApiGatewayManagedOverridesResultOutput) ToLookupApiGatewayManagedOverridesResultOutputWithContext(ctx context.Context) LookupApiGatewayManagedOverridesResultOutput {
	return o
}

func (o LookupApiGatewayManagedOverridesResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiGatewayManagedOverridesResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupApiGatewayManagedOverridesResultOutput) Integration() ApiGatewayManagedOverridesIntegrationOverridesPtrOutput {
	return o.ApplyT(func(v LookupApiGatewayManagedOverridesResult) *ApiGatewayManagedOverridesIntegrationOverrides {
		return v.Integration
	}).(ApiGatewayManagedOverridesIntegrationOverridesPtrOutput)
}

func (o LookupApiGatewayManagedOverridesResultOutput) Route() ApiGatewayManagedOverridesRouteOverridesPtrOutput {
	return o.ApplyT(func(v LookupApiGatewayManagedOverridesResult) *ApiGatewayManagedOverridesRouteOverrides {
		return v.Route
	}).(ApiGatewayManagedOverridesRouteOverridesPtrOutput)
}

func (o LookupApiGatewayManagedOverridesResultOutput) Stage() ApiGatewayManagedOverridesStageOverridesPtrOutput {
	return o.ApplyT(func(v LookupApiGatewayManagedOverridesResult) *ApiGatewayManagedOverridesStageOverrides {
		return v.Stage
	}).(ApiGatewayManagedOverridesStageOverridesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiGatewayManagedOverridesResultOutput{})
}
