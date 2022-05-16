// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelHlsOutputSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("h265PackagingType")]
        public Input<string>? H265PackagingType { get; set; }

        [Input("hlsSettings")]
        public Input<Inputs.ChannelHlsSettingsArgs>? HlsSettings { get; set; }

        [Input("nameModifier")]
        public Input<string>? NameModifier { get; set; }

        [Input("segmentModifier")]
        public Input<string>? SegmentModifier { get; set; }

        public ChannelHlsOutputSettingsArgs()
        {
        }
    }
}
