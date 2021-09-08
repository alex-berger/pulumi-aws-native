// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistener.html
    /// </summary>
    [OutputType]
    public sealed class VirtualGatewayVirtualGatewayListener
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistener.html#cfn-appmesh-virtualgateway-virtualgatewaylistener-connectionpool
        /// </summary>
        public readonly Outputs.VirtualGatewayVirtualGatewayConnectionPool? ConnectionPool;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistener.html#cfn-appmesh-virtualgateway-virtualgatewaylistener-healthcheck
        /// </summary>
        public readonly Outputs.VirtualGatewayVirtualGatewayHealthCheckPolicy? HealthCheck;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistener.html#cfn-appmesh-virtualgateway-virtualgatewaylistener-portmapping
        /// </summary>
        public readonly Outputs.VirtualGatewayVirtualGatewayPortMapping PortMapping;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistener.html#cfn-appmesh-virtualgateway-virtualgatewaylistener-tls
        /// </summary>
        public readonly Outputs.VirtualGatewayVirtualGatewayListenerTls? TLS;

        [OutputConstructor]
        private VirtualGatewayVirtualGatewayListener(
            Outputs.VirtualGatewayVirtualGatewayConnectionPool? connectionPool,

            Outputs.VirtualGatewayVirtualGatewayHealthCheckPolicy? healthCheck,

            Outputs.VirtualGatewayVirtualGatewayPortMapping portMapping,

            Outputs.VirtualGatewayVirtualGatewayListenerTls? tLS)
        {
            ConnectionPool = connectionPool;
            HealthCheck = healthCheck;
            PortMapping = portMapping;
            TLS = tLS;
        }
    }
}
