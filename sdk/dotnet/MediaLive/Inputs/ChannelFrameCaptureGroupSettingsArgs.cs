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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecapturegroupsettings.html
    /// </summary>
    public sealed class ChannelFrameCaptureGroupSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecapturegroupsettings.html#cfn-medialive-channel-framecapturegroupsettings-destination
        /// </summary>
        [Input("destination")]
        public Input<Inputs.ChannelOutputLocationRefArgs>? Destination { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecapturegroupsettings.html#cfn-medialive-channel-framecapturegroupsettings-framecapturecdnsettings
        /// </summary>
        [Input("frameCaptureCdnSettings")]
        public Input<Inputs.ChannelFrameCaptureCdnSettingsArgs>? FrameCaptureCdnSettings { get; set; }

        public ChannelFrameCaptureGroupSettingsArgs()
        {
        }
    }
}