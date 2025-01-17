// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::Model
func LookupModel(ctx *pulumi.Context, args *LookupModelArgs, opts ...pulumi.InvokeOption) (*LookupModelResult, error) {
	var rv LookupModelResult
	err := ctx.Invoke("aws-native:apigateway:getModel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupModelArgs struct {
	// A name for the model. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name.
	Name string `pulumi:"name"`
	// The ID of a REST API with which to associate this model.
	RestApiId string `pulumi:"restApiId"`
}

type LookupModelResult struct {
	// A description that identifies this model.
	Description *string `pulumi:"description"`
	// The schema to use to transform data to one or more output formats. Specify null ({}) if you don't want to specify a schema.
	Schema interface{} `pulumi:"schema"`
}

func LookupModelOutput(ctx *pulumi.Context, args LookupModelOutputArgs, opts ...pulumi.InvokeOption) LookupModelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupModelResult, error) {
			args := v.(LookupModelArgs)
			r, err := LookupModel(ctx, &args, opts...)
			var s LookupModelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupModelResultOutput)
}

type LookupModelOutputArgs struct {
	// A name for the model. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name.
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of a REST API with which to associate this model.
	RestApiId pulumi.StringInput `pulumi:"restApiId"`
}

func (LookupModelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelArgs)(nil)).Elem()
}

type LookupModelResultOutput struct{ *pulumi.OutputState }

func (LookupModelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelResult)(nil)).Elem()
}

func (o LookupModelResultOutput) ToLookupModelResultOutput() LookupModelResultOutput {
	return o
}

func (o LookupModelResultOutput) ToLookupModelResultOutputWithContext(ctx context.Context) LookupModelResultOutput {
	return o
}

// A description that identifies this model.
func (o LookupModelResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModelResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The schema to use to transform data to one or more output formats. Specify null ({}) if you don't want to specify a schema.
func (o LookupModelResultOutput) Schema() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupModelResult) interface{} { return v.Schema }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupModelResultOutput{})
}
