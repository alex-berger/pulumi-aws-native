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
    public sealed class ChannelFrameCaptureOutputSettings
    {
        public readonly string? NameModifier;

        [OutputConstructor]
        private ChannelFrameCaptureOutputSettings(string? nameModifier)
        {
            NameModifier = nameModifier;
        }
    }
}
