// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchService.Inputs
{

    public sealed class DomainEndpointOptionsArgs : Pulumi.ResourceArgs
    {
        [Input("customEndpoint")]
        public Input<string>? CustomEndpoint { get; set; }

        [Input("customEndpointCertificateArn")]
        public Input<string>? CustomEndpointCertificateArn { get; set; }

        [Input("customEndpointEnabled")]
        public Input<bool>? CustomEndpointEnabled { get; set; }

        [Input("enforceHTTPS")]
        public Input<bool>? EnforceHTTPS { get; set; }

        [Input("tLSSecurityPolicy")]
        public Input<string>? TLSSecurityPolicy { get; set; }

        public DomainEndpointOptionsArgs()
        {
        }
    }
}
