// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelNielsenConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("distributorId")]
        public Input<string>? DistributorId { get; set; }

        [Input("nielsenPcmToId3Tagging")]
        public Input<string>? NielsenPcmToId3Tagging { get; set; }

        public ChannelNielsenConfigurationArgs()
        {
        }
    }
}
