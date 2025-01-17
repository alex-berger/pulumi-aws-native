// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codepipeline

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CodePipeline::Pipeline
func LookupPipeline(ctx *pulumi.Context, args *LookupPipelineArgs, opts ...pulumi.InvokeOption) (*LookupPipelineResult, error) {
	var rv LookupPipelineResult
	err := ctx.Invoke("aws-native:codepipeline:getPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPipelineArgs struct {
	Id string `pulumi:"id"`
}

type LookupPipelineResult struct {
	ArtifactStore                  *PipelineArtifactStore     `pulumi:"artifactStore"`
	ArtifactStores                 []PipelineArtifactStoreMap `pulumi:"artifactStores"`
	DisableInboundStageTransitions []PipelineStageTransition  `pulumi:"disableInboundStageTransitions"`
	Id                             *string                    `pulumi:"id"`
	RestartExecutionOnUpdate       *bool                      `pulumi:"restartExecutionOnUpdate"`
	RoleArn                        *string                    `pulumi:"roleArn"`
	Stages                         []PipelineStageDeclaration `pulumi:"stages"`
	Tags                           []PipelineTag              `pulumi:"tags"`
	Version                        *string                    `pulumi:"version"`
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

func (o LookupPipelineResultOutput) ArtifactStore() PipelineArtifactStorePtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *PipelineArtifactStore { return v.ArtifactStore }).(PipelineArtifactStorePtrOutput)
}

func (o LookupPipelineResultOutput) ArtifactStores() PipelineArtifactStoreMapArrayOutput {
	return o.ApplyT(func(v LookupPipelineResult) []PipelineArtifactStoreMap { return v.ArtifactStores }).(PipelineArtifactStoreMapArrayOutput)
}

func (o LookupPipelineResultOutput) DisableInboundStageTransitions() PipelineStageTransitionArrayOutput {
	return o.ApplyT(func(v LookupPipelineResult) []PipelineStageTransition { return v.DisableInboundStageTransitions }).(PipelineStageTransitionArrayOutput)
}

func (o LookupPipelineResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPipelineResultOutput) RestartExecutionOnUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *bool { return v.RestartExecutionOnUpdate }).(pulumi.BoolPtrOutput)
}

func (o LookupPipelineResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupPipelineResultOutput) Stages() PipelineStageDeclarationArrayOutput {
	return o.ApplyT(func(v LookupPipelineResult) []PipelineStageDeclaration { return v.Stages }).(PipelineStageDeclarationArrayOutput)
}

func (o LookupPipelineResultOutput) Tags() PipelineTagArrayOutput {
	return o.ApplyT(func(v LookupPipelineResult) []PipelineTag { return v.Tags }).(PipelineTagArrayOutput)
}

func (o LookupPipelineResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPipelineResultOutput{})
}
