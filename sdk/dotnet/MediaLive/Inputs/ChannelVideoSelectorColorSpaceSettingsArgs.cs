// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelVideoSelectorColorSpaceSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("hdr10Settings")]
        public Input<Inputs.ChannelHdr10SettingsArgs>? Hdr10Settings { get; set; }

        public ChannelVideoSelectorColorSpaceSettingsArgs()
        {
        }
    }
}
