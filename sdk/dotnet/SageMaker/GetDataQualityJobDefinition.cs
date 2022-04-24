// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetDataQualityJobDefinition
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::DataQualityJobDefinition
        /// </summary>
        public static Task<GetDataQualityJobDefinitionResult> InvokeAsync(GetDataQualityJobDefinitionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDataQualityJobDefinitionResult>("aws-native:sagemaker:getDataQualityJobDefinition", args ?? new GetDataQualityJobDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::DataQualityJobDefinition
        /// </summary>
        public static Output<GetDataQualityJobDefinitionResult> Invoke(GetDataQualityJobDefinitionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDataQualityJobDefinitionResult>("aws-native:sagemaker:getDataQualityJobDefinition", args ?? new GetDataQualityJobDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataQualityJobDefinitionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of job definition.
        /// </summary>
        [Input("jobDefinitionArn", required: true)]
        public string JobDefinitionArn { get; set; } = null!;

        public GetDataQualityJobDefinitionArgs()
        {
        }
    }

    public sealed class GetDataQualityJobDefinitionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of job definition.
        /// </summary>
        [Input("jobDefinitionArn", required: true)]
        public Input<string> JobDefinitionArn { get; set; } = null!;

        public GetDataQualityJobDefinitionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDataQualityJobDefinitionResult
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
        private GetDataQualityJobDefinitionResult(
            string? creationTime,

            string? jobDefinitionArn)
        {
            CreationTime = creationTime;
            JobDefinitionArn = jobDefinitionArn;
        }
    }
}
