// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelRtmpGroupSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("adMarkers")]
        private InputList<string>? _adMarkers;
        public InputList<string> AdMarkers
        {
            get => _adMarkers ?? (_adMarkers = new InputList<string>());
            set => _adMarkers = value;
        }

        [Input("authenticationScheme")]
        public Input<string>? AuthenticationScheme { get; set; }

        [Input("cacheFullBehavior")]
        public Input<string>? CacheFullBehavior { get; set; }

        [Input("cacheLength")]
        public Input<int>? CacheLength { get; set; }

        [Input("captionData")]
        public Input<string>? CaptionData { get; set; }

        [Input("inputLossAction")]
        public Input<string>? InputLossAction { get; set; }

        [Input("restartDelay")]
        public Input<int>? RestartDelay { get; set; }

        public ChannelRtmpGroupSettingsArgs()
        {
        }
    }
}
