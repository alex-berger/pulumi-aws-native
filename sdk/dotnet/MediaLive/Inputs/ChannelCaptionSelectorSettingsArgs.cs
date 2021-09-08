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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselectorsettings.html
    /// </summary>
    public sealed class ChannelCaptionSelectorSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselectorsettings.html#cfn-medialive-channel-captionselectorsettings-ancillarysourcesettings
        /// </summary>
        [Input("ancillarySourceSettings")]
        public Input<Inputs.ChannelAncillarySourceSettingsArgs>? AncillarySourceSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselectorsettings.html#cfn-medialive-channel-captionselectorsettings-aribsourcesettings
        /// </summary>
        [Input("aribSourceSettings")]
        public Input<Inputs.ChannelAribSourceSettingsArgs>? AribSourceSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselectorsettings.html#cfn-medialive-channel-captionselectorsettings-dvbsubsourcesettings
        /// </summary>
        [Input("dvbSubSourceSettings")]
        public Input<Inputs.ChannelDvbSubSourceSettingsArgs>? DvbSubSourceSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselectorsettings.html#cfn-medialive-channel-captionselectorsettings-embeddedsourcesettings
        /// </summary>
        [Input("embeddedSourceSettings")]
        public Input<Inputs.ChannelEmbeddedSourceSettingsArgs>? EmbeddedSourceSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselectorsettings.html#cfn-medialive-channel-captionselectorsettings-scte20sourcesettings
        /// </summary>
        [Input("scte20SourceSettings")]
        public Input<Inputs.ChannelScte20SourceSettingsArgs>? Scte20SourceSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselectorsettings.html#cfn-medialive-channel-captionselectorsettings-scte27sourcesettings
        /// </summary>
        [Input("scte27SourceSettings")]
        public Input<Inputs.ChannelScte27SourceSettingsArgs>? Scte27SourceSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselectorsettings.html#cfn-medialive-channel-captionselectorsettings-teletextsourcesettings
        /// </summary>
        [Input("teletextSourceSettings")]
        public Input<Inputs.ChannelTeletextSourceSettingsArgs>? TeletextSourceSettings { get; set; }

        public ChannelCaptionSelectorSettingsArgs()
        {
        }
    }
}
