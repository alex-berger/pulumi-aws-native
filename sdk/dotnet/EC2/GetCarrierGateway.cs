// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetCarrierGateway
    {
        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Task<GetCarrierGatewayResult> InvokeAsync(GetCarrierGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCarrierGatewayResult>("aws-native:ec2:getCarrierGateway", args ?? new GetCarrierGatewayArgs(), options.WithDefaults());

        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Output<GetCarrierGatewayResult> Invoke(GetCarrierGatewayInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCarrierGatewayResult>("aws-native:ec2:getCarrierGateway", args ?? new GetCarrierGatewayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCarrierGatewayArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the carrier gateway.
        /// </summary>
        [Input("carrierGatewayId", required: true)]
        public string CarrierGatewayId { get; set; } = null!;

        public GetCarrierGatewayArgs()
        {
        }
    }

    public sealed class GetCarrierGatewayInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the carrier gateway.
        /// </summary>
        [Input("carrierGatewayId", required: true)]
        public Input<string> CarrierGatewayId { get; set; } = null!;

        public GetCarrierGatewayInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCarrierGatewayResult
    {
        /// <summary>
        /// The ID of the carrier gateway.
        /// </summary>
        public readonly string? CarrierGatewayId;
        /// <summary>
        /// The ID of the owner.
        /// </summary>
        public readonly string? OwnerId;
        /// <summary>
        /// The state of the carrier gateway.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// The tags for the carrier gateway.
        /// </summary>
        public readonly ImmutableArray<Outputs.CarrierGatewayTag> Tags;

        [OutputConstructor]
        private GetCarrierGatewayResult(
            string? carrierGatewayId,

            string? ownerId,

            string? state,

            ImmutableArray<Outputs.CarrierGatewayTag> tags)
        {
            CarrierGatewayId = carrierGatewayId;
            OwnerId = ownerId;
            State = state;
            Tags = tags;
        }
    }
}
