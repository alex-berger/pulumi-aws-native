// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-wavsettings.html
    /// </summary>
    [OutputType]
    public sealed class ChannelWavSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-wavsettings.html#cfn-medialive-channel-wavsettings-bitdepth
        /// </summary>
        public readonly double? BitDepth;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-wavsettings.html#cfn-medialive-channel-wavsettings-codingmode
        /// </summary>
        public readonly string? CodingMode;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-wavsettings.html#cfn-medialive-channel-wavsettings-samplerate
        /// </summary>
        public readonly double? SampleRate;

        [OutputConstructor]
        private ChannelWavSettings(
            double? bitDepth,

            string? codingMode,

            double? sampleRate)
        {
            BitDepth = bitDepth;
            CodingMode = codingMode;
            SampleRate = sampleRate;
        }
    }
}
