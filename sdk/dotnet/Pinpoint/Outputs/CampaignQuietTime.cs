// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Outputs
{

    [OutputType]
    public sealed class CampaignQuietTime
    {
        public readonly string End;
        public readonly string Start;

        [OutputConstructor]
        private CampaignQuietTime(
            string end,

            string start)
        {
            End = end;
            Start = start;
        }
    }
}
