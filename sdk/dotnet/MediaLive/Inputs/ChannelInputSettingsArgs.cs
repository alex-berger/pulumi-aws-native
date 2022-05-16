// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelInputSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("audioSelectors")]
        private InputList<Inputs.ChannelAudioSelectorArgs>? _audioSelectors;
        public InputList<Inputs.ChannelAudioSelectorArgs> AudioSelectors
        {
            get => _audioSelectors ?? (_audioSelectors = new InputList<Inputs.ChannelAudioSelectorArgs>());
            set => _audioSelectors = value;
        }

        [Input("captionSelectors")]
        private InputList<Inputs.ChannelCaptionSelectorArgs>? _captionSelectors;
        public InputList<Inputs.ChannelCaptionSelectorArgs> CaptionSelectors
        {
            get => _captionSelectors ?? (_captionSelectors = new InputList<Inputs.ChannelCaptionSelectorArgs>());
            set => _captionSelectors = value;
        }

        [Input("deblockFilter")]
        public Input<string>? DeblockFilter { get; set; }

        [Input("denoiseFilter")]
        public Input<string>? DenoiseFilter { get; set; }

        [Input("filterStrength")]
        public Input<int>? FilterStrength { get; set; }

        [Input("inputFilter")]
        public Input<string>? InputFilter { get; set; }

        [Input("networkInputSettings")]
        public Input<Inputs.ChannelNetworkInputSettingsArgs>? NetworkInputSettings { get; set; }

        [Input("scte35Pid")]
        public Input<int>? Scte35Pid { get; set; }

        [Input("smpte2038DataPreference")]
        public Input<string>? Smpte2038DataPreference { get; set; }

        [Input("sourceEndBehavior")]
        public Input<string>? SourceEndBehavior { get; set; }

        [Input("videoSelector")]
        public Input<Inputs.ChannelVideoSelectorArgs>? VideoSelector { get; set; }

        public ChannelInputSettingsArgs()
        {
        }
    }
}