// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceDiscovery.Inputs
{

    public sealed class PublicDnsNamespaceSOAArgs : Pulumi.ResourceArgs
    {
        [Input("tTL")]
        public Input<double>? TTL { get; set; }

        public PublicDnsNamespaceSOAArgs()
        {
        }
    }
}
