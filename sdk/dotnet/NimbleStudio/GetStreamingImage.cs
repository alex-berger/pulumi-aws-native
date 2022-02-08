// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio
{
    public static class GetStreamingImage
    {
        /// <summary>
        /// Represents a streaming session machine image that can be used to launch a streaming session
        /// </summary>
        public static Task<GetStreamingImageResult> InvokeAsync(GetStreamingImageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStreamingImageResult>("aws-native:nimblestudio:getStreamingImage", args ?? new GetStreamingImageArgs(), options.WithDefaults());

        /// <summary>
        /// Represents a streaming session machine image that can be used to launch a streaming session
        /// </summary>
        public static Output<GetStreamingImageResult> Invoke(GetStreamingImageInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetStreamingImageResult>("aws-native:nimblestudio:getStreamingImage", args ?? new GetStreamingImageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStreamingImageArgs : Pulumi.InvokeArgs
    {
        [Input("streamingImageId", required: true)]
        public string StreamingImageId { get; set; } = null!;

        /// <summary>
        /// &lt;p&gt;The studioId. &lt;/p&gt;
        /// </summary>
        [Input("studioId", required: true)]
        public string StudioId { get; set; } = null!;

        public GetStreamingImageArgs()
        {
        }
    }

    public sealed class GetStreamingImageInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("streamingImageId", required: true)]
        public Input<string> StreamingImageId { get; set; } = null!;

        /// <summary>
        /// &lt;p&gt;The studioId. &lt;/p&gt;
        /// </summary>
        [Input("studioId", required: true)]
        public Input<string> StudioId { get; set; } = null!;

        public GetStreamingImageInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStreamingImageResult
    {
        /// <summary>
        /// &lt;p&gt;A human-readable description of the streaming image.&lt;/p&gt;
        /// </summary>
        public readonly string? Description;
        public readonly Outputs.StreamingImageEncryptionConfiguration? EncryptionConfiguration;
        /// <summary>
        /// &lt;p&gt;The list of EULAs that must be accepted before a Streaming Session can be started using this streaming image.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<string> EulaIds;
        /// <summary>
        /// &lt;p&gt;A friendly name for a streaming image resource.&lt;/p&gt;
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// &lt;p&gt;The owner of the streaming image, either the studioId that contains the streaming image, or 'amazon' for images that are provided by Amazon Nimble Studio.&lt;/p&gt;
        /// </summary>
        public readonly string? Owner;
        /// <summary>
        /// &lt;p&gt;The platform of the streaming image, either WINDOWS or LINUX.&lt;/p&gt;
        /// </summary>
        public readonly string? Platform;
        public readonly string? StreamingImageId;

        [OutputConstructor]
        private GetStreamingImageResult(
            string? description,

            Outputs.StreamingImageEncryptionConfiguration? encryptionConfiguration,

            ImmutableArray<string> eulaIds,

            string? name,

            string? owner,

            string? platform,

            string? streamingImageId)
        {
            Description = description;
            EncryptionConfiguration = encryptionConfiguration;
            EulaIds = eulaIds;
            Name = name;
            Owner = owner;
            Platform = platform;
            StreamingImageId = streamingImageId;
        }
    }
}