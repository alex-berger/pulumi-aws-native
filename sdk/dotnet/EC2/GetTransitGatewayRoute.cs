// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetTransitGatewayRoute
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::TransitGatewayRoute
        /// </summary>
        public static Task<GetTransitGatewayRouteResult> InvokeAsync(GetTransitGatewayRouteArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTransitGatewayRouteResult>("aws-native:ec2:getTransitGatewayRoute", args ?? new GetTransitGatewayRouteArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::TransitGatewayRoute
        /// </summary>
        public static Output<GetTransitGatewayRouteResult> Invoke(GetTransitGatewayRouteInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTransitGatewayRouteResult>("aws-native:ec2:getTransitGatewayRoute", args ?? new GetTransitGatewayRouteInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTransitGatewayRouteArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetTransitGatewayRouteArgs()
        {
        }
    }

    public sealed class GetTransitGatewayRouteInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetTransitGatewayRouteInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTransitGatewayRouteResult
    {
        public readonly string? Id;

        [OutputConstructor]
        private GetTransitGatewayRouteResult(string? id)
        {
            Id = id;
        }
    }
}
