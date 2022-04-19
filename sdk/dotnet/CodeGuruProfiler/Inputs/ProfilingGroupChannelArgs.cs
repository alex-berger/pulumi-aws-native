// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeGuruProfiler.Inputs
{

    /// <summary>
    /// Notification medium for users to get alerted for events that occur in application profile. We support SNS topic as a notification channel.
    /// </summary>
    public sealed class ProfilingGroupChannelArgs : Pulumi.ResourceArgs
    {
        [Input("channelId")]
        public Input<string>? ChannelId { get; set; }

        [Input("channelUri", required: true)]
        public Input<string> ChannelUri { get; set; } = null!;

        public ProfilingGroupChannelArgs()
        {
        }
    }
}
