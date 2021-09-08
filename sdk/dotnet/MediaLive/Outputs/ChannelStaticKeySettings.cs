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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-statickeysettings.html
    /// </summary>
    [OutputType]
    public sealed class ChannelStaticKeySettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-statickeysettings.html#cfn-medialive-channel-statickeysettings-keyproviderserver
        /// </summary>
        public readonly Outputs.ChannelInputLocation? KeyProviderServer;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-statickeysettings.html#cfn-medialive-channel-statickeysettings-statickeyvalue
        /// </summary>
        public readonly string? StaticKeyValue;

        [OutputConstructor]
        private ChannelStaticKeySettings(
            Outputs.ChannelInputLocation? keyProviderServer,

            string? staticKeyValue)
        {
            KeyProviderServer = keyProviderServer;
            StaticKeyValue = staticKeyValue;
        }
    }
}
