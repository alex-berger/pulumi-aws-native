// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkManager.Outputs
{

    [OutputType]
    public sealed class CoreNetworkEdge
    {
        /// <summary>
        /// The ASN of a core network edge.
        /// </summary>
        public readonly double? Asn;
        /// <summary>
        /// The Region where a core network edge is located.
        /// </summary>
        public readonly string? EdgeLocation;
        public readonly ImmutableArray<string> InsideCidrBlocks;

        [OutputConstructor]
        private CoreNetworkEdge(
            double? asn,

            string? edgeLocation,

            ImmutableArray<string> insideCidrBlocks)
        {
            Asn = asn;
            EdgeLocation = edgeLocation;
            InsideCidrBlocks = insideCidrBlocks;
        }
    }
}