// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    /// <summary>
    /// Custom request handling.
    /// </summary>
    [OutputType]
    public sealed class RuleGroupCustomRequestHandling
    {
        /// <summary>
        /// Collection of HTTP headers.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleGroupCustomHTTPHeader> InsertHeaders;

        [OutputConstructor]
        private RuleGroupCustomRequestHandling(ImmutableArray<Outputs.RuleGroupCustomHTTPHeader> insertHeaders)
        {
            InsertHeaders = insertHeaders;
        }
    }
}
