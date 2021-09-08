// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html
    /// </summary>
    public sealed class ChannelInputSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("audioSelectors")]
        private InputList<Inputs.ChannelAudioSelectorArgs>? _audioSelectors;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-audioselectors
        /// </summary>
        public InputList<Inputs.ChannelAudioSelectorArgs> AudioSelectors
        {
            get => _audioSelectors ?? (_audioSelectors = new InputList<Inputs.ChannelAudioSelectorArgs>());
            set => _audioSelectors = value;
        }

        [Input("captionSelectors")]
        private InputList<Inputs.ChannelCaptionSelectorArgs>? _captionSelectors;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-captionselectors
        /// </summary>
        public InputList<Inputs.ChannelCaptionSelectorArgs> CaptionSelectors
        {
            get => _captionSelectors ?? (_captionSelectors = new InputList<Inputs.ChannelCaptionSelectorArgs>());
            set => _captionSelectors = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-deblockfilter
        /// </summary>
        [Input("deblockFilter")]
        public Input<string>? DeblockFilter { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-denoisefilter
        /// </summary>
        [Input("denoiseFilter")]
        public Input<string>? DenoiseFilter { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-filterstrength
        /// </summary>
        [Input("filterStrength")]
        public Input<int>? FilterStrength { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-inputfilter
        /// </summary>
        [Input("inputFilter")]
        public Input<string>? InputFilter { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-networkinputsettings
        /// </summary>
        [Input("networkInputSettings")]
        public Input<Inputs.ChannelNetworkInputSettingsArgs>? NetworkInputSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-smpte2038datapreference
        /// </summary>
        [Input("smpte2038DataPreference")]
        public Input<string>? Smpte2038DataPreference { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-sourceendbehavior
        /// </summary>
        [Input("sourceEndBehavior")]
        public Input<string>? SourceEndBehavior { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-videoselector
        /// </summary>
        [Input("videoSelector")]
        public Input<Inputs.ChannelVideoSelectorArgs>? VideoSelector { get; set; }

        public ChannelInputSettingsArgs()
        {
        }
    }
}
