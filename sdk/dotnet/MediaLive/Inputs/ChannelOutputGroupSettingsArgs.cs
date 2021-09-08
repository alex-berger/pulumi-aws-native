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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html
    /// </summary>
    public sealed class ChannelOutputGroupSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html#cfn-medialive-channel-outputgroupsettings-archivegroupsettings
        /// </summary>
        [Input("archiveGroupSettings")]
        public Input<Inputs.ChannelArchiveGroupSettingsArgs>? ArchiveGroupSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html#cfn-medialive-channel-outputgroupsettings-framecapturegroupsettings
        /// </summary>
        [Input("frameCaptureGroupSettings")]
        public Input<Inputs.ChannelFrameCaptureGroupSettingsArgs>? FrameCaptureGroupSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html#cfn-medialive-channel-outputgroupsettings-hlsgroupsettings
        /// </summary>
        [Input("hlsGroupSettings")]
        public Input<Inputs.ChannelHlsGroupSettingsArgs>? HlsGroupSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html#cfn-medialive-channel-outputgroupsettings-mediapackagegroupsettings
        /// </summary>
        [Input("mediaPackageGroupSettings")]
        public Input<Inputs.ChannelMediaPackageGroupSettingsArgs>? MediaPackageGroupSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html#cfn-medialive-channel-outputgroupsettings-mssmoothgroupsettings
        /// </summary>
        [Input("msSmoothGroupSettings")]
        public Input<Inputs.ChannelMsSmoothGroupSettingsArgs>? MsSmoothGroupSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html#cfn-medialive-channel-outputgroupsettings-multiplexgroupsettings
        /// </summary>
        [Input("multiplexGroupSettings")]
        public Input<Inputs.ChannelMultiplexGroupSettingsArgs>? MultiplexGroupSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html#cfn-medialive-channel-outputgroupsettings-rtmpgroupsettings
        /// </summary>
        [Input("rtmpGroupSettings")]
        public Input<Inputs.ChannelRtmpGroupSettingsArgs>? RtmpGroupSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroupsettings.html#cfn-medialive-channel-outputgroupsettings-udpgroupsettings
        /// </summary>
        [Input("udpGroupSettings")]
        public Input<Inputs.ChannelUdpGroupSettingsArgs>? UdpGroupSettings { get; set; }

        public ChannelOutputGroupSettingsArgs()
        {
        }
    }
}
