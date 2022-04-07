// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::SSM::Document resource is an SSM document in AWS Systems Manager that defines the actions that Systems Manager performs, which can be used to set up and run commands on your instances.
func LookupDocument(ctx *pulumi.Context, args *LookupDocumentArgs, opts ...pulumi.InvokeOption) (*LookupDocumentResult, error) {
	var rv LookupDocumentResult
	err := ctx.Invoke("aws-native:ssm:getDocument", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDocumentArgs struct {
	// A name for the Systems Manager document.
	Name string `pulumi:"name"`
}

type LookupDocumentResult struct {
	// Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment.
	Tags []DocumentTag `pulumi:"tags"`
}

func LookupDocumentOutput(ctx *pulumi.Context, args LookupDocumentOutputArgs, opts ...pulumi.InvokeOption) LookupDocumentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDocumentResult, error) {
			args := v.(LookupDocumentArgs)
			r, err := LookupDocument(ctx, &args, opts...)
			var s LookupDocumentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDocumentResultOutput)
}

type LookupDocumentOutputArgs struct {
	// A name for the Systems Manager document.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupDocumentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDocumentArgs)(nil)).Elem()
}

type LookupDocumentResultOutput struct{ *pulumi.OutputState }

func (LookupDocumentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDocumentResult)(nil)).Elem()
}

func (o LookupDocumentResultOutput) ToLookupDocumentResultOutput() LookupDocumentResultOutput {
	return o
}

func (o LookupDocumentResultOutput) ToLookupDocumentResultOutputWithContext(ctx context.Context) LookupDocumentResultOutput {
	return o
}

// Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment.
func (o LookupDocumentResultOutput) Tags() DocumentTagArrayOutput {
	return o.ApplyT(func(v LookupDocumentResult) []DocumentTag { return v.Tags }).(DocumentTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDocumentResultOutput{})
}
