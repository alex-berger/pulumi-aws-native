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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archivecontainersettings.html
    /// </summary>
    public sealed class ChannelArchiveContainerSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archivecontainersettings.html#cfn-medialive-channel-archivecontainersettings-m2tssettings
        /// </summary>
        [Input("m2tsSettings")]
        public Input<Inputs.ChannelM2tsSettingsArgs>? M2tsSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archivecontainersettings.html#cfn-medialive-channel-archivecontainersettings-rawsettings
        /// </summary>
        [Input("rawSettings")]
        public Input<Inputs.ChannelRawSettingsArgs>? RawSettings { get; set; }

        public ChannelArchiveContainerSettingsArgs()
        {
        }
    }
}
