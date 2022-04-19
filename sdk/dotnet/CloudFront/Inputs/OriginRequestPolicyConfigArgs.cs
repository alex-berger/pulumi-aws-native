// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Inputs
{

    public sealed class OriginRequestPolicyConfigArgs : Pulumi.ResourceArgs
    {
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("cookiesConfig", required: true)]
        public Input<Inputs.OriginRequestPolicyCookiesConfigArgs> CookiesConfig { get; set; } = null!;

        [Input("headersConfig", required: true)]
        public Input<Inputs.OriginRequestPolicyHeadersConfigArgs> HeadersConfig { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("queryStringsConfig", required: true)]
        public Input<Inputs.OriginRequestPolicyQueryStringsConfigArgs> QueryStringsConfig { get; set; } = null!;

        public OriginRequestPolicyConfigArgs()
        {
        }
    }
}
