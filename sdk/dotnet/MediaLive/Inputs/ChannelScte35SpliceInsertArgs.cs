// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelScte35SpliceInsertArgs : Pulumi.ResourceArgs
    {
        [Input("adAvailOffset")]
        public Input<int>? AdAvailOffset { get; set; }

        [Input("noRegionalBlackoutFlag")]
        public Input<string>? NoRegionalBlackoutFlag { get; set; }

        [Input("webDeliveryAllowedFlag")]
        public Input<string>? WebDeliveryAllowedFlag { get; set; }

        public ChannelScte35SpliceInsertArgs()
        {
        }
    }
}
