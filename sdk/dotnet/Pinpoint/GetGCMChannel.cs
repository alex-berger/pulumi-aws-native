// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint
{
    public static class GetGCMChannel
    {
        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::GCMChannel
        /// </summary>
        public static Task<GetGCMChannelResult> InvokeAsync(GetGCMChannelArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGCMChannelResult>("aws-native:pinpoint:getGCMChannel", args ?? new GetGCMChannelArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::GCMChannel
        /// </summary>
        public static Output<GetGCMChannelResult> Invoke(GetGCMChannelInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGCMChannelResult>("aws-native:pinpoint:getGCMChannel", args ?? new GetGCMChannelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGCMChannelArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetGCMChannelArgs()
        {
        }
    }

    public sealed class GetGCMChannelInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetGCMChannelInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGCMChannelResult
    {
        public readonly string? ApiKey;
        public readonly bool? Enabled;
        public readonly string? Id;

        [OutputConstructor]
        private GetGCMChannelResult(
            string? apiKey,

            bool? enabled,

            string? id)
        {
            ApiKey = apiKey;
            Enabled = enabled;
            Id = id;
        }
    }
}
