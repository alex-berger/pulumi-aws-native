// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetImageVersion
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::ImageVersion
        /// </summary>
        public static Task<GetImageVersionResult> InvokeAsync(GetImageVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImageVersionResult>("aws-native:sagemaker:getImageVersion", args ?? new GetImageVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::ImageVersion
        /// </summary>
        public static Output<GetImageVersionResult> Invoke(GetImageVersionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetImageVersionResult>("aws-native:sagemaker:getImageVersion", args ?? new GetImageVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetImageVersionArgs : Pulumi.InvokeArgs
    {
        [Input("imageVersionArn", required: true)]
        public string ImageVersionArn { get; set; } = null!;

        public GetImageVersionArgs()
        {
        }
    }

    public sealed class GetImageVersionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("imageVersionArn", required: true)]
        public Input<string> ImageVersionArn { get; set; } = null!;

        public GetImageVersionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetImageVersionResult
    {
        public readonly string? ContainerImage;
        public readonly string? ImageArn;
        public readonly string? ImageVersionArn;
        public readonly int? Version;

        [OutputConstructor]
        private GetImageVersionResult(
            string? containerImage,

            string? imageArn,

            string? imageVersionArn,

            int? version)
        {
            ContainerImage = containerImage;
            ImageArn = imageArn;
            ImageVersionArn = imageVersionArn;
            Version = version;
        }
    }
}
