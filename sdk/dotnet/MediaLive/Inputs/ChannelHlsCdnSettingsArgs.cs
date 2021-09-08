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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlscdnsettings.html
    /// </summary>
    public sealed class ChannelHlsCdnSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlscdnsettings.html#cfn-medialive-channel-hlscdnsettings-hlsakamaisettings
        /// </summary>
        [Input("hlsAkamaiSettings")]
        public Input<Inputs.ChannelHlsAkamaiSettingsArgs>? HlsAkamaiSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlscdnsettings.html#cfn-medialive-channel-hlscdnsettings-hlsbasicputsettings
        /// </summary>
        [Input("hlsBasicPutSettings")]
        public Input<Inputs.ChannelHlsBasicPutSettingsArgs>? HlsBasicPutSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlscdnsettings.html#cfn-medialive-channel-hlscdnsettings-hlsmediastoresettings
        /// </summary>
        [Input("hlsMediaStoreSettings")]
        public Input<Inputs.ChannelHlsMediaStoreSettingsArgs>? HlsMediaStoreSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlscdnsettings.html#cfn-medialive-channel-hlscdnsettings-hlss3settings
        /// </summary>
        [Input("hlsS3Settings")]
        public Input<Inputs.ChannelHlsS3SettingsArgs>? HlsS3Settings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlscdnsettings.html#cfn-medialive-channel-hlscdnsettings-hlswebdavsettings
        /// </summary>
        [Input("hlsWebdavSettings")]
        public Input<Inputs.ChannelHlsWebdavSettingsArgs>? HlsWebdavSettings { get; set; }

        public ChannelHlsCdnSettingsArgs()
        {
        }
    }
}
