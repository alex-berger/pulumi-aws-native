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
    public sealed class ChannelHdr10Settings
    {
        public readonly int? MaxCll;
        public readonly int? MaxFall;

        [OutputConstructor]
        private ChannelHdr10Settings(
            int? maxCll,

            int? maxFall)
        {
            MaxCll = maxCll;
            MaxFall = maxFall;
        }
    }
}
