// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetDHCPOptions
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::DHCPOptions
        /// </summary>
        public static Task<GetDHCPOptionsResult> InvokeAsync(GetDHCPOptionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDHCPOptionsResult>("aws-native:ec2:getDHCPOptions", args ?? new GetDHCPOptionsArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::DHCPOptions
        /// </summary>
        public static Output<GetDHCPOptionsResult> Invoke(GetDHCPOptionsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDHCPOptionsResult>("aws-native:ec2:getDHCPOptions", args ?? new GetDHCPOptionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDHCPOptionsArgs : Pulumi.InvokeArgs
    {
        [Input("dhcpOptionsId", required: true)]
        public string DhcpOptionsId { get; set; } = null!;

        public GetDHCPOptionsArgs()
        {
        }
    }

    public sealed class GetDHCPOptionsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("dhcpOptionsId", required: true)]
        public Input<string> DhcpOptionsId { get; set; } = null!;

        public GetDHCPOptionsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDHCPOptionsResult
    {
        public readonly string? DhcpOptionsId;
        /// <summary>
        /// Any tags assigned to the DHCP options set.
        /// </summary>
        public readonly ImmutableArray<Outputs.DHCPOptionsTag> Tags;

        [OutputConstructor]
        private GetDHCPOptionsResult(
            string? dhcpOptionsId,

            ImmutableArray<Outputs.DHCPOptionsTag> tags)
        {
            DhcpOptionsId = dhcpOptionsId;
            Tags = tags;
        }
    }
}
