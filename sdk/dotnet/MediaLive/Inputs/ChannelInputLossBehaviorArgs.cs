// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelInputLossBehaviorArgs : Pulumi.ResourceArgs
    {
        [Input("blackFrameMsec")]
        public Input<int>? BlackFrameMsec { get; set; }

        [Input("inputLossImageColor")]
        public Input<string>? InputLossImageColor { get; set; }

        [Input("inputLossImageSlate")]
        public Input<Inputs.ChannelInputLocationArgs>? InputLossImageSlate { get; set; }

        [Input("inputLossImageType")]
        public Input<string>? InputLossImageType { get; set; }

        [Input("repeatFrameMsec")]
        public Input<int>? RepeatFrameMsec { get; set; }

        public ChannelInputLossBehaviorArgs()
        {
        }
    }
}
