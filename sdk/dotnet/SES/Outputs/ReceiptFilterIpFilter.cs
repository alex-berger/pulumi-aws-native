// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SES.Outputs
{

    [OutputType]
    public sealed class ReceiptFilterIpFilter
    {
        public readonly string Cidr;
        public readonly string Policy;

        [OutputConstructor]
        private ReceiptFilterIpFilter(
            string cidr,

            string policy)
        {
            Cidr = cidr;
            Policy = policy;
        }
    }
}
