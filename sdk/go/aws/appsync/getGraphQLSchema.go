// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppSync::GraphQLSchema
func LookupGraphQLSchema(ctx *pulumi.Context, args *LookupGraphQLSchemaArgs, opts ...pulumi.InvokeOption) (*LookupGraphQLSchemaResult, error) {
	var rv LookupGraphQLSchemaResult
	err := ctx.Invoke("aws-native:appsync:getGraphQLSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGraphQLSchemaArgs struct {
	Id string `pulumi:"id"`
}

type LookupGraphQLSchemaResult struct {
	Definition           *string `pulumi:"definition"`
	DefinitionS3Location *string `pulumi:"definitionS3Location"`
	Id                   *string `pulumi:"id"`
}

func LookupGraphQLSchemaOutput(ctx *pulumi.Context, args LookupGraphQLSchemaOutputArgs, opts ...pulumi.InvokeOption) LookupGraphQLSchemaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGraphQLSchemaResult, error) {
			args := v.(LookupGraphQLSchemaArgs)
			r, err := LookupGraphQLSchema(ctx, &args, opts...)
			var s LookupGraphQLSchemaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGraphQLSchemaResultOutput)
}

type LookupGraphQLSchemaOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupGraphQLSchemaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGraphQLSchemaArgs)(nil)).Elem()
}

type LookupGraphQLSchemaResultOutput struct{ *pulumi.OutputState }

func (LookupGraphQLSchemaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGraphQLSchemaResult)(nil)).Elem()
}

func (o LookupGraphQLSchemaResultOutput) ToLookupGraphQLSchemaResultOutput() LookupGraphQLSchemaResultOutput {
	return o
}

func (o LookupGraphQLSchemaResultOutput) ToLookupGraphQLSchemaResultOutputWithContext(ctx context.Context) LookupGraphQLSchemaResultOutput {
	return o
}

func (o LookupGraphQLSchemaResultOutput) Definition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLSchemaResult) *string { return v.Definition }).(pulumi.StringPtrOutput)
}

func (o LookupGraphQLSchemaResultOutput) DefinitionS3Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLSchemaResult) *string { return v.DefinitionS3Location }).(pulumi.StringPtrOutput)
}

func (o LookupGraphQLSchemaResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLSchemaResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGraphQLSchemaResultOutput{})
}
