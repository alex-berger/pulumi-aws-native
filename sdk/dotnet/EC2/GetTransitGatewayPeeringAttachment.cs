// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetTransitGatewayPeeringAttachment
    {
        /// <summary>
        /// The AWS::EC2::TransitGatewayPeeringAttachment type
        /// </summary>
        public static Task<GetTransitGatewayPeeringAttachmentResult> InvokeAsync(GetTransitGatewayPeeringAttachmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTransitGatewayPeeringAttachmentResult>("aws-native:ec2:getTransitGatewayPeeringAttachment", args ?? new GetTransitGatewayPeeringAttachmentArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::EC2::TransitGatewayPeeringAttachment type
        /// </summary>
        public static Output<GetTransitGatewayPeeringAttachmentResult> Invoke(GetTransitGatewayPeeringAttachmentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTransitGatewayPeeringAttachmentResult>("aws-native:ec2:getTransitGatewayPeeringAttachment", args ?? new GetTransitGatewayPeeringAttachmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTransitGatewayPeeringAttachmentArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the transit gateway peering attachment.
        /// </summary>
        [Input("transitGatewayAttachmentId", required: true)]
        public string TransitGatewayAttachmentId { get; set; } = null!;

        public GetTransitGatewayPeeringAttachmentArgs()
        {
        }
    }

    public sealed class GetTransitGatewayPeeringAttachmentInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the transit gateway peering attachment.
        /// </summary>
        [Input("transitGatewayAttachmentId", required: true)]
        public Input<string> TransitGatewayAttachmentId { get; set; } = null!;

        public GetTransitGatewayPeeringAttachmentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTransitGatewayPeeringAttachmentResult
    {
        /// <summary>
        /// The time the transit gateway peering attachment was created.
        /// </summary>
        public readonly string? CreationTime;
        /// <summary>
        /// The state of the transit gateway peering attachment. Note that the initiating state has been deprecated.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// The status of the transit gateway peering attachment.
        /// </summary>
        public readonly Outputs.TransitGatewayPeeringAttachmentPeeringAttachmentStatus? Status;
        /// <summary>
        /// The tags for the transit gateway peering attachment.
        /// </summary>
        public readonly ImmutableArray<Outputs.TransitGatewayPeeringAttachmentTag> Tags;
        /// <summary>
        /// The ID of the transit gateway peering attachment.
        /// </summary>
        public readonly string? TransitGatewayAttachmentId;

        [OutputConstructor]
        private GetTransitGatewayPeeringAttachmentResult(
            string? creationTime,

            string? state,

            Outputs.TransitGatewayPeeringAttachmentPeeringAttachmentStatus? status,

            ImmutableArray<Outputs.TransitGatewayPeeringAttachmentTag> tags,

            string? transitGatewayAttachmentId)
        {
            CreationTime = creationTime;
            State = state;
            Status = status;
            Tags = tags;
            TransitGatewayAttachmentId = transitGatewayAttachmentId;
        }
    }
}
