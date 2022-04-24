// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wisdom

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Wisdom::AssistantAssociation Resource Type
func LookupAssistantAssociation(ctx *pulumi.Context, args *LookupAssistantAssociationArgs, opts ...pulumi.InvokeOption) (*LookupAssistantAssociationResult, error) {
	var rv LookupAssistantAssociationResult
	err := ctx.Invoke("aws-native:wisdom:getAssistantAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssistantAssociationArgs struct {
	AssistantAssociationId string `pulumi:"assistantAssociationId"`
	AssistantId            string `pulumi:"assistantId"`
}

type LookupAssistantAssociationResult struct {
	AssistantArn            *string `pulumi:"assistantArn"`
	AssistantAssociationArn *string `pulumi:"assistantAssociationArn"`
	AssistantAssociationId  *string `pulumi:"assistantAssociationId"`
}

func LookupAssistantAssociationOutput(ctx *pulumi.Context, args LookupAssistantAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupAssistantAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAssistantAssociationResult, error) {
			args := v.(LookupAssistantAssociationArgs)
			r, err := LookupAssistantAssociation(ctx, &args, opts...)
			var s LookupAssistantAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAssistantAssociationResultOutput)
}

type LookupAssistantAssociationOutputArgs struct {
	AssistantAssociationId pulumi.StringInput `pulumi:"assistantAssociationId"`
	AssistantId            pulumi.StringInput `pulumi:"assistantId"`
}

func (LookupAssistantAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssistantAssociationArgs)(nil)).Elem()
}

type LookupAssistantAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupAssistantAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssistantAssociationResult)(nil)).Elem()
}

func (o LookupAssistantAssociationResultOutput) ToLookupAssistantAssociationResultOutput() LookupAssistantAssociationResultOutput {
	return o
}

func (o LookupAssistantAssociationResultOutput) ToLookupAssistantAssociationResultOutputWithContext(ctx context.Context) LookupAssistantAssociationResultOutput {
	return o
}

func (o LookupAssistantAssociationResultOutput) AssistantArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssistantAssociationResult) *string { return v.AssistantArn }).(pulumi.StringPtrOutput)
}

func (o LookupAssistantAssociationResultOutput) AssistantAssociationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssistantAssociationResult) *string { return v.AssistantAssociationArn }).(pulumi.StringPtrOutput)
}

func (o LookupAssistantAssociationResultOutput) AssistantAssociationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssistantAssociationResult) *string { return v.AssistantAssociationId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssistantAssociationResultOutput{})
}
