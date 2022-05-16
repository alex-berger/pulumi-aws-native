// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Outputs
{

    [OutputType]
    public sealed class NetworkInsightsAccessScopePathStatementRequest
    {
        public readonly Outputs.NetworkInsightsAccessScopePacketHeaderStatementRequest? PacketHeaderStatement;
        public readonly Outputs.NetworkInsightsAccessScopeResourceStatementRequest? ResourceStatement;

        [OutputConstructor]
        private NetworkInsightsAccessScopePathStatementRequest(
            Outputs.NetworkInsightsAccessScopePacketHeaderStatementRequest? packetHeaderStatement,

            Outputs.NetworkInsightsAccessScopeResourceStatementRequest? resourceStatement)
        {
            PacketHeaderStatement = packetHeaderStatement;
            ResourceStatement = resourceStatement;
        }
    }
}