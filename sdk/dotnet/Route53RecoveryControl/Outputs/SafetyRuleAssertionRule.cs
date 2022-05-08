// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53RecoveryControl.Outputs
{

    /// <summary>
    /// An assertion rule enforces that, when a routing control state is changed, that the criteria set by the rule configuration is met. Otherwise, the change to the routing control is not accepted.
    /// </summary>
    [OutputType]
    public sealed class SafetyRuleAssertionRule
    {
        /// <summary>
        /// The routing controls that are part of transactions that are evaluated to determine if a request to change a routing control state is allowed. For example, you might include three routing controls, one for each of three AWS Regions.
        /// </summary>
        public readonly ImmutableArray<string> AssertedControls;
        /// <summary>
        /// An evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail. This helps prevent "flapping" of state. The wait period is 5000 ms by default, but you can choose a custom value.
        /// </summary>
        public readonly int WaitPeriodMs;

        [OutputConstructor]
        private SafetyRuleAssertionRule(
            ImmutableArray<string> assertedControls,

            int waitPeriodMs)
        {
            AssertedControls = assertedControls;
            WaitPeriodMs = waitPeriodMs;
        }
    }
}
