// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetGatewayRouteTableAssociation
    {
        /// <summary>
        /// Associates a gateway with a route table. The gateway and route table must be in the same VPC. This association causes the incoming traffic to the gateway to be routed according to the routes in the route table.
        /// </summary>
        public static Task<GetGatewayRouteTableAssociationResult> InvokeAsync(GetGatewayRouteTableAssociationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGatewayRouteTableAssociationResult>("aws-native:ec2:getGatewayRouteTableAssociation", args ?? new GetGatewayRouteTableAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// Associates a gateway with a route table. The gateway and route table must be in the same VPC. This association causes the incoming traffic to the gateway to be routed according to the routes in the route table.
        /// </summary>
        public static Output<GetGatewayRouteTableAssociationResult> Invoke(GetGatewayRouteTableAssociationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGatewayRouteTableAssociationResult>("aws-native:ec2:getGatewayRouteTableAssociation", args ?? new GetGatewayRouteTableAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayRouteTableAssociationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the gateway.
        /// </summary>
        [Input("gatewayId", required: true)]
        public string GatewayId { get; set; } = null!;

        public GetGatewayRouteTableAssociationArgs()
        {
        }
    }

    public sealed class GetGatewayRouteTableAssociationInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the gateway.
        /// </summary>
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        public GetGatewayRouteTableAssociationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGatewayRouteTableAssociationResult
    {
        /// <summary>
        /// The route table association ID.
        /// </summary>
        public readonly string? AssociationId;
        /// <summary>
        /// The ID of the route table.
        /// </summary>
        public readonly string? RouteTableId;

        [OutputConstructor]
        private GetGatewayRouteTableAssociationResult(
            string? associationId,

            string? routeTableId)
        {
            AssociationId = associationId;
            RouteTableId = routeTableId;
        }
    }
}
