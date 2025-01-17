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
    public sealed class ChannelAc3Settings
    {
        public readonly double? Bitrate;
        public readonly string? BitstreamMode;
        public readonly string? CodingMode;
        public readonly int? Dialnorm;
        public readonly string? DrcProfile;
        public readonly string? LfeFilter;
        public readonly string? MetadataControl;

        [OutputConstructor]
        private ChannelAc3Settings(
            double? bitrate,

            string? bitstreamMode,

            string? codingMode,

            int? dialnorm,

            string? drcProfile,

            string? lfeFilter,

            string? metadataControl)
        {
            Bitrate = bitrate;
            BitstreamMode = bitstreamMode;
            CodingMode = codingMode;
            Dialnorm = dialnorm;
            DrcProfile = drcProfile;
            LfeFilter = lfeFilter;
            MetadataControl = metadataControl;
        }
    }
}
