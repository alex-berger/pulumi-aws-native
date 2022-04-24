// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetSubnetRouteTableAssociation
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::SubnetRouteTableAssociation
        /// </summary>
        public static Task<GetSubnetRouteTableAssociationResult> InvokeAsync(GetSubnetRouteTableAssociationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubnetRouteTableAssociationResult>("aws-native:ec2:getSubnetRouteTableAssociation", args ?? new GetSubnetRouteTableAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::SubnetRouteTableAssociation
        /// </summary>
        public static Output<GetSubnetRouteTableAssociationResult> Invoke(GetSubnetRouteTableAssociationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSubnetRouteTableAssociationResult>("aws-native:ec2:getSubnetRouteTableAssociation", args ?? new GetSubnetRouteTableAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSubnetRouteTableAssociationArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetSubnetRouteTableAssociationArgs()
        {
        }
    }

    public sealed class GetSubnetRouteTableAssociationInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetSubnetRouteTableAssociationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSubnetRouteTableAssociationResult
    {
        public readonly string? Id;

        [OutputConstructor]
        private GetSubnetRouteTableAssociationResult(string? id)
        {
            Id = id;
        }
    }
}
