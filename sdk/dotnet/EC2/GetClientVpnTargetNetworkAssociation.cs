// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetClientVpnTargetNetworkAssociation
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::ClientVpnTargetNetworkAssociation
        /// </summary>
        public static Task<GetClientVpnTargetNetworkAssociationResult> InvokeAsync(GetClientVpnTargetNetworkAssociationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClientVpnTargetNetworkAssociationResult>("aws-native:ec2:getClientVpnTargetNetworkAssociation", args ?? new GetClientVpnTargetNetworkAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::ClientVpnTargetNetworkAssociation
        /// </summary>
        public static Output<GetClientVpnTargetNetworkAssociationResult> Invoke(GetClientVpnTargetNetworkAssociationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetClientVpnTargetNetworkAssociationResult>("aws-native:ec2:getClientVpnTargetNetworkAssociation", args ?? new GetClientVpnTargetNetworkAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClientVpnTargetNetworkAssociationArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetClientVpnTargetNetworkAssociationArgs()
        {
        }
    }

    public sealed class GetClientVpnTargetNetworkAssociationInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetClientVpnTargetNetworkAssociationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClientVpnTargetNetworkAssociationResult
    {
        public readonly string? Id;

        [OutputConstructor]
        private GetClientVpnTargetNetworkAssociationResult(string? id)
        {
            Id = id;
        }
    }
}
