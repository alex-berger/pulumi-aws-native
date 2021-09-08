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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-wavsettings.html
    /// </summary>
    public sealed class ChannelWavSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-wavsettings.html#cfn-medialive-channel-wavsettings-bitdepth
        /// </summary>
        [Input("bitDepth")]
        public Input<double>? BitDepth { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-wavsettings.html#cfn-medialive-channel-wavsettings-codingmode
        /// </summary>
        [Input("codingMode")]
        public Input<string>? CodingMode { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-wavsettings.html#cfn-medialive-channel-wavsettings-samplerate
        /// </summary>
        [Input("sampleRate")]
        public Input<double>? SampleRate { get; set; }

        public ChannelWavSettingsArgs()
        {
        }
    }
}