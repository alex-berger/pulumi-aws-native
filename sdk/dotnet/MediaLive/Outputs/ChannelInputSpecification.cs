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
    public sealed class ChannelInputSpecification
    {
        public readonly string? Codec;
        public readonly string? MaximumBitrate;
        public readonly string? Resolution;

        [OutputConstructor]
        private ChannelInputSpecification(
            string? codec,

            string? maximumBitrate,

            string? resolution)
        {
            Codec = codec;
            MaximumBitrate = maximumBitrate;
            Resolution = resolution;
        }
    }
}
