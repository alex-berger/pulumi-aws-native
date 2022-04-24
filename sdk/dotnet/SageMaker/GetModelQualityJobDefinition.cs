// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetModelQualityJobDefinition
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::ModelQualityJobDefinition
        /// </summary>
        public static Task<GetModelQualityJobDefinitionResult> InvokeAsync(GetModelQualityJobDefinitionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetModelQualityJobDefinitionResult>("aws-native:sagemaker:getModelQualityJobDefinition", args ?? new GetModelQualityJobDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::ModelQualityJobDefinition
        /// </summary>
        public static Output<GetModelQualityJobDefinitionResult> Invoke(GetModelQualityJobDefinitionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetModelQualityJobDefinitionResult>("aws-native:sagemaker:getModelQualityJobDefinition", args ?? new GetModelQualityJobDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetModelQualityJobDefinitionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of job definition.
        /// </summary>
        [Input("jobDefinitionArn", required: true)]
        public string JobDefinitionArn { get; set; } = null!;

        public GetModelQualityJobDefinitionArgs()
        {
        }
    }

    public sealed class GetModelQualityJobDefinitionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of job definition.
        /// </summary>
        [Input("jobDefinitionArn", required: true)]
        public Input<string> JobDefinitionArn { get; set; } = null!;

        public GetModelQualityJobDefinitionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetModelQualityJobDefinitionResult
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
        private GetModelQualityJobDefinitionResult(
            string? creationTime,

            string? jobDefinitionArn)
        {
            CreationTime = creationTime;
            JobDefinitionArn = jobDefinitionArn;
        }
    }
}
