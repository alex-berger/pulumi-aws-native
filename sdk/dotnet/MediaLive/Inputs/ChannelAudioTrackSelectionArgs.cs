// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelAudioTrackSelectionArgs : Pulumi.ResourceArgs
    {
        [Input("tracks")]
        private InputList<Inputs.ChannelAudioTrackArgs>? _tracks;
        public InputList<Inputs.ChannelAudioTrackArgs> Tracks
        {
            get => _tracks ?? (_tracks = new InputList<Inputs.ChannelAudioTrackArgs>());
            set => _tracks = value;
        }

        public ChannelAudioTrackSelectionArgs()
        {
        }
    }
}
