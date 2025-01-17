// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataPipeline
{
    public static class GetPipeline
    {
        /// <summary>
        /// Resource Type definition for AWS::DataPipeline::Pipeline
        /// </summary>
        public static Task<GetPipelineResult> InvokeAsync(GetPipelineArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPipelineResult>("aws-native:datapipeline:getPipeline", args ?? new GetPipelineArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::DataPipeline::Pipeline
        /// </summary>
        public static Output<GetPipelineResult> Invoke(GetPipelineInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPipelineResult>("aws-native:datapipeline:getPipeline", args ?? new GetPipelineInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPipelineArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetPipelineArgs()
        {
        }
    }

    public sealed class GetPipelineInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetPipelineInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPipelineResult
    {
        public readonly bool? Activate;
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.PipelineParameterObject> ParameterObjects;
        public readonly ImmutableArray<Outputs.PipelineParameterValue> ParameterValues;
        public readonly ImmutableArray<Outputs.PipelineObject> PipelineObjects;
        public readonly ImmutableArray<Outputs.PipelineTag> PipelineTags;

        [OutputConstructor]
        private GetPipelineResult(
            bool? activate,

            string? id,

            ImmutableArray<Outputs.PipelineParameterObject> parameterObjects,

            ImmutableArray<Outputs.PipelineParameterValue> parameterValues,

            ImmutableArray<Outputs.PipelineObject> pipelineObjects,

            ImmutableArray<Outputs.PipelineTag> pipelineTags)
        {
            Activate = activate;
            Id = id;
            ParameterObjects = parameterObjects;
            ParameterValues = parameterValues;
            PipelineObjects = pipelineObjects;
            PipelineTags = pipelineTags;
        }
    }
}
