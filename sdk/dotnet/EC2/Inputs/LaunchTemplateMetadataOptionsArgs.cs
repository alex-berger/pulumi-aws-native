// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class LaunchTemplateMetadataOptionsArgs : Pulumi.ResourceArgs
    {
        [Input("httpEndpoint")]
        public Input<string>? HttpEndpoint { get; set; }

        [Input("httpProtocolIpv6")]
        public Input<string>? HttpProtocolIpv6 { get; set; }

        [Input("httpPutResponseHopLimit")]
        public Input<int>? HttpPutResponseHopLimit { get; set; }

        [Input("httpTokens")]
        public Input<string>? HttpTokens { get; set; }

        [Input("instanceMetadataTags")]
        public Input<string>? InstanceMetadataTags { get; set; }

        public LaunchTemplateMetadataOptionsArgs()
        {
        }
    }
}
