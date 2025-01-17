// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetVPCGatewayAttachment
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::VPCGatewayAttachment
        /// </summary>
        public static Task<GetVPCGatewayAttachmentResult> InvokeAsync(GetVPCGatewayAttachmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVPCGatewayAttachmentResult>("aws-native:ec2:getVPCGatewayAttachment", args ?? new GetVPCGatewayAttachmentArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::VPCGatewayAttachment
        /// </summary>
        public static Output<GetVPCGatewayAttachmentResult> Invoke(GetVPCGatewayAttachmentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVPCGatewayAttachmentResult>("aws-native:ec2:getVPCGatewayAttachment", args ?? new GetVPCGatewayAttachmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVPCGatewayAttachmentArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetVPCGatewayAttachmentArgs()
        {
        }
    }

    public sealed class GetVPCGatewayAttachmentInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetVPCGatewayAttachmentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVPCGatewayAttachmentResult
    {
        public readonly string? Id;
        public readonly string? InternetGatewayId;
        public readonly string? VpcId;
        public readonly string? VpnGatewayId;

        [OutputConstructor]
        private GetVPCGatewayAttachmentResult(
            string? id,

            string? internetGatewayId,

            string? vpcId,

            string? vpnGatewayId)
        {
            Id = id;
            InternetGatewayId = internetGatewayId;
            VpcId = vpcId;
            VpnGatewayId = vpnGatewayId;
        }
    }
}
