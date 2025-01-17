// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    [OutputType]
    public sealed class RuleGroupForwardedIPConfiguration
    {
        public readonly Pulumi.AwsNative.WAFv2.RuleGroupForwardedIPConfigurationFallbackBehavior FallbackBehavior;
        public readonly string HeaderName;

        [OutputConstructor]
        private RuleGroupForwardedIPConfiguration(
            Pulumi.AwsNative.WAFv2.RuleGroupForwardedIPConfigurationFallbackBehavior fallbackBehavior,

            string headerName)
        {
            FallbackBehavior = fallbackBehavior;
            HeaderName = headerName;
        }
    }
}
