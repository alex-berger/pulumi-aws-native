// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Inputs
{

    public sealed class GatewayRouteSpecArgs : Pulumi.ResourceArgs
    {
        [Input("grpcRoute")]
        public Input<Inputs.GatewayRouteGrpcGatewayRouteArgs>? GrpcRoute { get; set; }

        [Input("http2Route")]
        public Input<Inputs.GatewayRouteHttpGatewayRouteArgs>? Http2Route { get; set; }

        [Input("httpRoute")]
        public Input<Inputs.GatewayRouteHttpGatewayRouteArgs>? HttpRoute { get; set; }

        [Input("priority")]
        public Input<int>? Priority { get; set; }

        public GatewayRouteSpecArgs()
        {
        }
    }
}
