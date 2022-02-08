// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage
{
    public static class GetChannel
    {
        /// <summary>
        /// Resource schema for AWS::MediaPackage::Channel
        /// </summary>
        public static Task<GetChannelResult> InvokeAsync(GetChannelArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetChannelResult>("aws-native:mediapackage:getChannel", args ?? new GetChannelArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::MediaPackage::Channel
        /// </summary>
        public static Output<GetChannelResult> Invoke(GetChannelInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetChannelResult>("aws-native:mediapackage:getChannel", args ?? new GetChannelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetChannelArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Channel.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetChannelArgs()
        {
        }
    }

    public sealed class GetChannelInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Channel.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetChannelInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetChannelResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) assigned to the Channel.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// A short text description of the Channel.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The configuration parameters for egress access logging.
        /// </summary>
        public readonly Outputs.ChannelLogConfiguration? EgressAccessLogs;
        /// <summary>
        /// A short text description of the Channel.
        /// </summary>
        public readonly Outputs.ChannelHlsIngest? HlsIngest;
        /// <summary>
        /// The configuration parameters for egress access logging.
        /// </summary>
        public readonly Outputs.ChannelLogConfiguration? IngressAccessLogs;

        [OutputConstructor]
        private GetChannelResult(
            string? arn,

            string? description,

            Outputs.ChannelLogConfiguration? egressAccessLogs,

            Outputs.ChannelHlsIngest? hlsIngest,

            Outputs.ChannelLogConfiguration? ingressAccessLogs)
        {
            Arn = arn;
            Description = description;
            EgressAccessLogs = egressAccessLogs;
            HlsIngest = hlsIngest;
            IngressAccessLogs = ingressAccessLogs;
        }
    }
}