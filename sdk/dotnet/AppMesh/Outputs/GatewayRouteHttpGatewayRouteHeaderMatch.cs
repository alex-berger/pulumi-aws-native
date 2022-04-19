// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Outputs
{

    [OutputType]
    public sealed class GatewayRouteHttpGatewayRouteHeaderMatch
    {
        public readonly string? Exact;
        public readonly string? Prefix;
        public readonly Outputs.GatewayRouteRangeMatch? Range;
        public readonly string? Regex;
        public readonly string? Suffix;

        [OutputConstructor]
        private GatewayRouteHttpGatewayRouteHeaderMatch(
            string? exact,

            string? prefix,

            Outputs.GatewayRouteRangeMatch? range,

            string? regex,

            string? suffix)
        {
            Exact = exact;
            Prefix = prefix;
            Range = range;
            Regex = regex;
            Suffix = suffix;
        }
    }
}
