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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlswebdavsettings.html
    /// </summary>
    public sealed class ChannelHlsWebdavSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlswebdavsettings.html#cfn-medialive-channel-hlswebdavsettings-connectionretryinterval
        /// </summary>
        [Input("connectionRetryInterval")]
        public Input<int>? ConnectionRetryInterval { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlswebdavsettings.html#cfn-medialive-channel-hlswebdavsettings-filecacheduration
        /// </summary>
        [Input("filecacheDuration")]
        public Input<int>? FilecacheDuration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlswebdavsettings.html#cfn-medialive-channel-hlswebdavsettings-httptransfermode
        /// </summary>
        [Input("httpTransferMode")]
        public Input<string>? HttpTransferMode { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlswebdavsettings.html#cfn-medialive-channel-hlswebdavsettings-numretries
        /// </summary>
        [Input("numRetries")]
        public Input<int>? NumRetries { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlswebdavsettings.html#cfn-medialive-channel-hlswebdavsettings-restartdelay
        /// </summary>
        [Input("restartDelay")]
        public Input<int>? RestartDelay { get; set; }

        public ChannelHlsWebdavSettingsArgs()
        {
        }
    }
}
