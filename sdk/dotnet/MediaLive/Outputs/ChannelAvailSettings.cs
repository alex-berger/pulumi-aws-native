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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-availsettings.html
    /// </summary>
    [OutputType]
    public sealed class ChannelAvailSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-availsettings.html#cfn-medialive-channel-availsettings-scte35spliceinsert
        /// </summary>
        public readonly Outputs.ChannelScte35SpliceInsert? Scte35SpliceInsert;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-availsettings.html#cfn-medialive-channel-availsettings-scte35timesignalapos
        /// </summary>
        public readonly Outputs.ChannelScte35TimeSignalApos? Scte35TimeSignalApos;

        [OutputConstructor]
        private ChannelAvailSettings(
            Outputs.ChannelScte35SpliceInsert? scte35SpliceInsert,

            Outputs.ChannelScte35TimeSignalApos? scte35TimeSignalApos)
        {
            Scte35SpliceInsert = scte35SpliceInsert;
            Scte35TimeSignalApos = scte35TimeSignalApos;
        }
    }
}