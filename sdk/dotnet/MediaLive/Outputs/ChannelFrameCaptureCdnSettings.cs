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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecapturecdnsettings.html
    /// </summary>
    [OutputType]
    public sealed class ChannelFrameCaptureCdnSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecapturecdnsettings.html#cfn-medialive-channel-framecapturecdnsettings-framecaptures3settings
        /// </summary>
        public readonly Outputs.ChannelFrameCaptureS3Settings? FrameCaptureS3Settings;

        [OutputConstructor]
        private ChannelFrameCaptureCdnSettings(Outputs.ChannelFrameCaptureS3Settings? frameCaptureS3Settings)
        {
            FrameCaptureS3Settings = frameCaptureS3Settings;
        }
    }
}