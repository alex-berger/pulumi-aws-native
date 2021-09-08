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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-output.html
    /// </summary>
    [OutputType]
    public sealed class ChannelOutput
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-output.html#cfn-medialive-channel-output-audiodescriptionnames
        /// </summary>
        public readonly ImmutableArray<string> AudioDescriptionNames;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-output.html#cfn-medialive-channel-output-captiondescriptionnames
        /// </summary>
        public readonly ImmutableArray<string> CaptionDescriptionNames;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-output.html#cfn-medialive-channel-output-outputname
        /// </summary>
        public readonly string? OutputName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-output.html#cfn-medialive-channel-output-outputsettings
        /// </summary>
        public readonly Outputs.ChannelOutputSettings? OutputSettings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-output.html#cfn-medialive-channel-output-videodescriptionname
        /// </summary>
        public readonly string? VideoDescriptionName;

        [OutputConstructor]
        private ChannelOutput(
            ImmutableArray<string> audioDescriptionNames,

            ImmutableArray<string> captionDescriptionNames,

            string? outputName,

            Outputs.ChannelOutputSettings? outputSettings,

            string? videoDescriptionName)
        {
            AudioDescriptionNames = audioDescriptionNames;
            CaptionDescriptionNames = captionDescriptionNames;
            OutputName = outputName;
            OutputSettings = outputSettings;
            VideoDescriptionName = videoDescriptionName;
        }
    }
}