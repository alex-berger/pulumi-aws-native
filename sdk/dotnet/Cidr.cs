// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative
{
    public static class Cidr
    {
        public static Task<CidrResult> InvokeAsync(CidrArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<CidrResult>("aws-native:index:cidr", args ?? new CidrArgs(), options.WithDefaults());

        public static Output<CidrResult> Invoke(CidrInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<CidrResult>("aws-native:index:cidr", args ?? new CidrInvokeArgs(), options.WithDefaults());
    }


    public sealed class CidrArgs : Pulumi.InvokeArgs
    {
        [Input("cidrBits", required: true)]
        public int CidrBits { get; set; }

        [Input("count", required: true)]
        public int Count { get; set; }

        [Input("ipBlock", required: true)]
        public string IpBlock { get; set; } = null!;

        public CidrArgs()
        {
        }
    }

    public sealed class CidrInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("cidrBits", required: true)]
        public Input<int> CidrBits { get; set; } = null!;

        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        [Input("ipBlock", required: true)]
        public Input<string> IpBlock { get; set; } = null!;

        public CidrInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class CidrResult
    {
        public readonly ImmutableArray<string> Subnets;

        [OutputConstructor]
        private CidrResult(ImmutableArray<string> subnets)
        {
            Subnets = subnets;
        }
    }
}
