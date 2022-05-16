// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelGlobalConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("initialAudioGain")]
        public Input<int>? InitialAudioGain { get; set; }

        [Input("inputEndAction")]
        public Input<string>? InputEndAction { get; set; }

        [Input("inputLossBehavior")]
        public Input<Inputs.ChannelInputLossBehaviorArgs>? InputLossBehavior { get; set; }

        [Input("outputLockingMode")]
        public Input<string>? OutputLockingMode { get; set; }

        [Input("outputTimingSource")]
        public Input<string>? OutputTimingSource { get; set; }

        [Input("supportLowFramerateInputs")]
        public Input<string>? SupportLowFramerateInputs { get; set; }

        public ChannelGlobalConfigurationArgs()
        {
        }
    }
}
