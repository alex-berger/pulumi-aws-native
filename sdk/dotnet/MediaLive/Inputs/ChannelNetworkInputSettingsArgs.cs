// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelNetworkInputSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("hlsInputSettings")]
        public Input<Inputs.ChannelHlsInputSettingsArgs>? HlsInputSettings { get; set; }

        [Input("serverValidation")]
        public Input<string>? ServerValidation { get; set; }

        public ChannelNetworkInputSettingsArgs()
        {
        }
    }
}