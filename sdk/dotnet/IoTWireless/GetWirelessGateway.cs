// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTWireless
{
    public static class GetWirelessGateway
    {
        /// <summary>
        /// Create and manage wireless gateways, including LoRa gateways.
        /// </summary>
        public static Task<GetWirelessGatewayResult> InvokeAsync(GetWirelessGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWirelessGatewayResult>("aws-native:iotwireless:getWirelessGateway", args ?? new GetWirelessGatewayArgs(), options.WithDefaults());

        /// <summary>
        /// Create and manage wireless gateways, including LoRa gateways.
        /// </summary>
        public static Output<GetWirelessGatewayResult> Invoke(GetWirelessGatewayInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetWirelessGatewayResult>("aws-native:iotwireless:getWirelessGateway", args ?? new GetWirelessGatewayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWirelessGatewayArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id for Wireless Gateway. Returned upon successful create.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetWirelessGatewayArgs()
        {
        }
    }

    public sealed class GetWirelessGatewayInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id for Wireless Gateway. Returned upon successful create.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetWirelessGatewayInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWirelessGatewayResult
    {
        /// <summary>
        /// Arn for Wireless Gateway. Returned upon successful create.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Description of Wireless Gateway.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Id for Wireless Gateway. Returned upon successful create.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The date and time when the most recent uplink was received.
        /// </summary>
        public readonly string? LastUplinkReceivedAt;
        /// <summary>
        /// The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Gateway.
        /// </summary>
        public readonly Outputs.WirelessGatewayLoRaWANGateway? LoRaWAN;
        /// <summary>
        /// Name of Wireless Gateway.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A list of key-value pairs that contain metadata for the gateway.
        /// </summary>
        public readonly ImmutableArray<Outputs.WirelessGatewayTag> Tags;
        /// <summary>
        /// Thing Arn. Passed into Update to associate a Thing with the Wireless Gateway.
        /// </summary>
        public readonly string? ThingArn;
        /// <summary>
        /// Thing Arn. If there is a Thing created, this can be returned with a Get call.
        /// </summary>
        public readonly string? ThingName;

        [OutputConstructor]
        private GetWirelessGatewayResult(
            string? arn,

            string? description,

            string? id,

            string? lastUplinkReceivedAt,

            Outputs.WirelessGatewayLoRaWANGateway? loRaWAN,

            string? name,

            ImmutableArray<Outputs.WirelessGatewayTag> tags,

            string? thingArn,

            string? thingName)
        {
            Arn = arn;
            Description = description;
            Id = id;
            LastUplinkReceivedAt = lastUplinkReceivedAt;
            LoRaWAN = loRaWAN;
            Name = name;
            Tags = tags;
            ThingArn = thingArn;
            ThingName = thingName;
        }
    }
}
