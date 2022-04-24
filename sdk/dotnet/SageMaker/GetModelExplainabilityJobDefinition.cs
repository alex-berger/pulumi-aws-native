// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetModelExplainabilityJobDefinition
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::ModelExplainabilityJobDefinition
        /// </summary>
        public static Task<GetModelExplainabilityJobDefinitionResult> InvokeAsync(GetModelExplainabilityJobDefinitionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetModelExplainabilityJobDefinitionResult>("aws-native:sagemaker:getModelExplainabilityJobDefinition", args ?? new GetModelExplainabilityJobDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::ModelExplainabilityJobDefinition
        /// </summary>
        public static Output<GetModelExplainabilityJobDefinitionResult> Invoke(GetModelExplainabilityJobDefinitionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetModelExplainabilityJobDefinitionResult>("aws-native:sagemaker:getModelExplainabilityJobDefinition", args ?? new GetModelExplainabilityJobDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetModelExplainabilityJobDefinitionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of job definition.
        /// </summary>
        [Input("jobDefinitionArn", required: true)]
        public string JobDefinitionArn { get; set; } = null!;

        public GetModelExplainabilityJobDefinitionArgs()
        {
        }
    }

    public sealed class GetModelExplainabilityJobDefinitionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of job definition.
        /// </summary>
        [Input("jobDefinitionArn", required: true)]
        public Input<string> JobDefinitionArn { get; set; } = null!;

        public GetModelExplainabilityJobDefinitionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetModelExplainabilityJobDefinitionResult
    {
        /// <summary>
        /// The time at which the job definition was created.
        /// </summary>
        public readonly string? CreationTime;
        /// <summary>
        /// The Amazon Resource Name (ARN) of job definition.
        /// </summary>
        public readonly string? JobDefinitionArn;

        [OutputConstructor]
        private GetModelExplainabilityJobDefinitionResult(
            string? creationTime,

            string? jobDefinitionArn)
        {
            CreationTime = creationTime;
            JobDefinitionArn = jobDefinitionArn;
        }
    }
}
