// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DLM.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-sharerule.html
    /// </summary>
    [OutputType]
    public sealed class LifecyclePolicyShareRule
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-sharerule.html#cfn-dlm-lifecyclepolicy-sharerule-targetaccounts
        /// </summary>
        public readonly ImmutableArray<string> TargetAccounts;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-sharerule.html#cfn-dlm-lifecyclepolicy-sharerule-unshareinterval
        /// </summary>
        public readonly int? UnshareInterval;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-sharerule.html#cfn-dlm-lifecyclepolicy-sharerule-unshareintervalunit
        /// </summary>
        public readonly string? UnshareIntervalUnit;

        [OutputConstructor]
        private LifecyclePolicyShareRule(
            ImmutableArray<string> targetAccounts,

            int? unshareInterval,

            string? unshareIntervalUnit)
        {
            TargetAccounts = targetAccounts;
            UnshareInterval = unshareInterval;
            UnshareIntervalUnit = unshareIntervalUnit;
        }
    }
}
