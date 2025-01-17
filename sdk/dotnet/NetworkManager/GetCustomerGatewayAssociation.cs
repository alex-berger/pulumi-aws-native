// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkManager
{
    public static class GetCustomerGatewayAssociation
    {
        /// <summary>
        /// The AWS::NetworkManager::CustomerGatewayAssociation type associates a customer gateway with a device and optionally, with a link.
        /// </summary>
        public static Task<GetCustomerGatewayAssociationResult> InvokeAsync(GetCustomerGatewayAssociationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCustomerGatewayAssociationResult>("aws-native:networkmanager:getCustomerGatewayAssociation", args ?? new GetCustomerGatewayAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::NetworkManager::CustomerGatewayAssociation type associates a customer gateway with a device and optionally, with a link.
        /// </summary>
        public static Output<GetCustomerGatewayAssociationResult> Invoke(GetCustomerGatewayAssociationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCustomerGatewayAssociationResult>("aws-native:networkmanager:getCustomerGatewayAssociation", args ?? new GetCustomerGatewayAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCustomerGatewayAssociationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the customer gateway.
        /// </summary>
        [Input("customerGatewayArn", required: true)]
        public string CustomerGatewayArn { get; set; } = null!;

        /// <summary>
        /// The ID of the global network.
        /// </summary>
        [Input("globalNetworkId", required: true)]
        public string GlobalNetworkId { get; set; } = null!;

        public GetCustomerGatewayAssociationArgs()
        {
        }
    }

    public sealed class GetCustomerGatewayAssociationInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the customer gateway.
        /// </summary>
        [Input("customerGatewayArn", required: true)]
        public Input<string> CustomerGatewayArn { get; set; } = null!;

        /// <summary>
        /// The ID of the global network.
        /// </summary>
        [Input("globalNetworkId", required: true)]
        public Input<string> GlobalNetworkId { get; set; } = null!;

        public GetCustomerGatewayAssociationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCustomerGatewayAssociationResult
    {
        [OutputConstructor]
        private GetCustomerGatewayAssociationResult()
        {
        }
    }
}
