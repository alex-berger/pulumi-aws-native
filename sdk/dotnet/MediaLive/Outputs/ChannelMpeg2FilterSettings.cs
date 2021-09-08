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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2filtersettings.html
    /// </summary>
    [OutputType]
    public sealed class ChannelMpeg2FilterSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2filtersettings.html#cfn-medialive-channel-mpeg2filtersettings-temporalfiltersettings
        /// </summary>
        public readonly Outputs.ChannelTemporalFilterSettings? TemporalFilterSettings;

        [OutputConstructor]
        private ChannelMpeg2FilterSettings(Outputs.ChannelTemporalFilterSettings? temporalFilterSettings)
        {
            TemporalFilterSettings = temporalFilterSettings;
        }
    }
}