// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelFailoverConditionSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("audioSilenceSettings")]
        public Input<Inputs.ChannelAudioSilenceFailoverSettingsArgs>? AudioSilenceSettings { get; set; }

        [Input("inputLossSettings")]
        public Input<Inputs.ChannelInputLossFailoverSettingsArgs>? InputLossSettings { get; set; }

        [Input("videoBlackSettings")]
        public Input<Inputs.ChannelVideoBlackFailoverSettingsArgs>? VideoBlackSettings { get; set; }

        public ChannelFailoverConditionSettingsArgs()
        {
        }
    }
}
