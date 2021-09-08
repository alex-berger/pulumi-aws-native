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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html
    /// </summary>
    [OutputType]
    public sealed class ChannelEncoderSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-audiodescriptions
        /// </summary>
        public readonly ImmutableArray<Outputs.ChannelAudioDescription> AudioDescriptions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-availblanking
        /// </summary>
        public readonly Outputs.ChannelAvailBlanking? AvailBlanking;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-availconfiguration
        /// </summary>
        public readonly Outputs.ChannelAvailConfiguration? AvailConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-blackoutslate
        /// </summary>
        public readonly Outputs.ChannelBlackoutSlate? BlackoutSlate;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-captiondescriptions
        /// </summary>
        public readonly ImmutableArray<Outputs.ChannelCaptionDescription> CaptionDescriptions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-featureactivations
        /// </summary>
        public readonly Outputs.ChannelFeatureActivations? FeatureActivations;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-globalconfiguration
        /// </summary>
        public readonly Outputs.ChannelGlobalConfiguration? GlobalConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-motiongraphicsconfiguration
        /// </summary>
        public readonly Outputs.ChannelMotionGraphicsConfiguration? MotionGraphicsConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-nielsenconfiguration
        /// </summary>
        public readonly Outputs.ChannelNielsenConfiguration? NielsenConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-outputgroups
        /// </summary>
        public readonly ImmutableArray<Outputs.ChannelOutputGroup> OutputGroups;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-timecodeconfig
        /// </summary>
        public readonly Outputs.ChannelTimecodeConfig? TimecodeConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-videodescriptions
        /// </summary>
        public readonly ImmutableArray<Outputs.ChannelVideoDescription> VideoDescriptions;

        [OutputConstructor]
        private ChannelEncoderSettings(
            ImmutableArray<Outputs.ChannelAudioDescription> audioDescriptions,

            Outputs.ChannelAvailBlanking? availBlanking,

            Outputs.ChannelAvailConfiguration? availConfiguration,

            Outputs.ChannelBlackoutSlate? blackoutSlate,

            ImmutableArray<Outputs.ChannelCaptionDescription> captionDescriptions,

            Outputs.ChannelFeatureActivations? featureActivations,

            Outputs.ChannelGlobalConfiguration? globalConfiguration,

            Outputs.ChannelMotionGraphicsConfiguration? motionGraphicsConfiguration,

            Outputs.ChannelNielsenConfiguration? nielsenConfiguration,

            ImmutableArray<Outputs.ChannelOutputGroup> outputGroups,

            Outputs.ChannelTimecodeConfig? timecodeConfig,

            ImmutableArray<Outputs.ChannelVideoDescription> videoDescriptions)
        {
            AudioDescriptions = audioDescriptions;
            AvailBlanking = availBlanking;
            AvailConfiguration = availConfiguration;
            BlackoutSlate = blackoutSlate;
            CaptionDescriptions = captionDescriptions;
            FeatureActivations = featureActivations;
            GlobalConfiguration = globalConfiguration;
            MotionGraphicsConfiguration = motionGraphicsConfiguration;
            NielsenConfiguration = nielsenConfiguration;
            OutputGroups = outputGroups;
            TimecodeConfig = timecodeConfig;
            VideoDescriptions = videoDescriptions;
        }
    }
}
