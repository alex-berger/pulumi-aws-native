// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Redshift.Outputs
{

    /// <summary>
    /// The connection endpoint for connecting to an Amazon Redshift cluster through the proxy.
    /// </summary>
    [OutputType]
    public sealed class VpcEndpointProperties
    {
        /// <summary>
        /// One or more network interfaces of the endpoint. Also known as an interface endpoint.
        /// </summary>
        public readonly ImmutableArray<Outputs.EndpointAccessNetworkInterface> NetworkInterfaces;
        /// <summary>
        /// The connection endpoint ID for connecting an Amazon Redshift cluster through the proxy.
        /// </summary>
        public readonly string? VpcEndpointId;
        /// <summary>
        /// The VPC identifier that the endpoint is associated.
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private VpcEndpointProperties(
            ImmutableArray<Outputs.EndpointAccessNetworkInterface> networkInterfaces,

            string? vpcEndpointId,

            string? vpcId)
        {
            NetworkInterfaces = networkInterfaces;
            VpcEndpointId = vpcEndpointId;
            VpcId = vpcId;
        }
    }
}
