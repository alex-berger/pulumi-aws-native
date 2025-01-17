// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Outputs
{

    [OutputType]
    public sealed class ChannelEmbeddedSourceSettings
    {
        public readonly string? Convert608To708;
        public readonly string? Scte20Detection;
        public readonly int? Source608ChannelNumber;
        public readonly int? Source608TrackNumber;

        [OutputConstructor]
        private ChannelEmbeddedSourceSettings(
            string? convert608To708,

            string? scte20Detection,

            int? source608ChannelNumber,

            int? source608TrackNumber)
        {
            Convert608To708 = convert608To708;
            Scte20Detection = scte20Detection;
            Source608ChannelNumber = source608ChannelNumber;
            Source608TrackNumber = source608TrackNumber;
        }
    }
}
