// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datapipeline

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DataPipeline::Pipeline
func LookupPipeline(ctx *pulumi.Context, args *LookupPipelineArgs, opts ...pulumi.InvokeOption) (*LookupPipelineResult, error) {
	var rv LookupPipelineResult
	err := ctx.Invoke("aws-native:datapipeline:getPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPipelineArgs struct {
	Id string `pulumi:"id"`
}

type LookupPipelineResult struct {
	Activate         *bool                     `pulumi:"activate"`
	Id               *string                   `pulumi:"id"`
	ParameterObjects []PipelineParameterObject `pulumi:"parameterObjects"`
	ParameterValues  []PipelineParameterValue  `pulumi:"parameterValues"`
	PipelineObjects  []PipelineObject          `pulumi:"pipelineObjects"`
	PipelineTags     []PipelineTag             `pulumi:"pipelineTags"`
}

func LookupPipelineOutput(ctx *pulumi.Context, args LookupPipelineOutputArgs, opts ...pulumi.InvokeOption) LookupPipelineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPipelineResult, error) {
			args := v.(LookupPipelineArgs)
			r, err := LookupPipeline(ctx, &args, opts...)
			var s LookupPipelineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPipelineResultOutput)
}

type LookupPipelineOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupPipelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineArgs)(nil)).Elem()
}

type LookupPipelineResultOutput struct{ *pulumi.OutputState }

func (LookupPipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineResult)(nil)).Elem()
}

func (o LookupPipelineResultOutput) ToLookupPipelineResultOutput() LookupPipelineResultOutput {
	return o
}

func (o LookupPipelineResultOutput) ToLookupPipelineResultOutputWithContext(ctx context.Context) LookupPipelineResultOutput {
	return o
}

func (o LookupPipelineResultOutput) Activate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *bool { return v.Activate }).(pulumi.BoolPtrOutput)
}

func (o LookupPipelineResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPipelineResultOutput) ParameterObjects() PipelineParameterObjectArrayOutput {
	return o.ApplyT(func(v LookupPipelineResult) []PipelineParameterObject { return v.ParameterObjects }).(PipelineParameterObjectArrayOutput)
}

func (o LookupPipelineResultOutput) ParameterValues() PipelineParameterValueArrayOutput {
	return o.ApplyT(func(v LookupPipelineResult) []PipelineParameterValue { return v.ParameterValues }).(PipelineParameterValueArrayOutput)
}

func (o LookupPipelineResultOutput) PipelineObjects() PipelineObjectArrayOutput {
	return o.ApplyT(func(v LookupPipelineResult) []PipelineObject { return v.PipelineObjects }).(PipelineObjectArrayOutput)
}

func (o LookupPipelineResultOutput) PipelineTags() PipelineTagArrayOutput {
	return o.ApplyT(func(v LookupPipelineResult) []PipelineTag { return v.PipelineTags }).(PipelineTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPipelineResultOutput{})
}
