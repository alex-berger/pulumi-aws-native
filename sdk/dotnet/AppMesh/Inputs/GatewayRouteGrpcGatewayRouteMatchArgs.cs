// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayroutematch.html
    /// </summary>
    public sealed class GatewayRouteGrpcGatewayRouteMatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayroutematch.html#cfn-appmesh-gatewayroute-grpcgatewayroutematch-hostname
        /// </summary>
        [Input("hostname")]
        public Input<Inputs.GatewayRouteGatewayRouteHostnameMatchArgs>? Hostname { get; set; }

        [Input("metadata")]
        private InputList<Inputs.GatewayRouteGrpcGatewayRouteMetadataArgs>? _metadata;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayroutematch.html#cfn-appmesh-gatewayroute-grpcgatewayroutematch-metadata
        /// </summary>
        public InputList<Inputs.GatewayRouteGrpcGatewayRouteMetadataArgs> Metadata
        {
            get => _metadata ?? (_metadata = new InputList<Inputs.GatewayRouteGrpcGatewayRouteMetadataArgs>());
            set => _metadata = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-grpcgatewayroutematch.html#cfn-appmesh-gatewayroute-grpcgatewayroutematch-servicename
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public GatewayRouteGrpcGatewayRouteMatchArgs()
        {
        }
    }
}