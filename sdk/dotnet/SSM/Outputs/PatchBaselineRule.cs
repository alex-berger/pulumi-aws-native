// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSM.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rule.html
    /// </summary>
    [OutputType]
    public sealed class PatchBaselineRule
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rule.html#cfn-ssm-patchbaseline-rule-approveafterdays
        /// </summary>
        public readonly int? ApproveAfterDays;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rule.html#cfn-ssm-patchbaseline-rule-approveuntildate
        /// </summary>
        public readonly Outputs.PatchBaselinePatchStringDate? ApproveUntilDate;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rule.html#cfn-ssm-patchbaseline-rule-compliancelevel
        /// </summary>
        public readonly string? ComplianceLevel;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rule.html#cfn-ssm-patchbaseline-rule-enablenonsecurity
        /// </summary>
        public readonly bool? EnableNonSecurity;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rule.html#cfn-ssm-patchbaseline-rule-patchfiltergroup
        /// </summary>
        public readonly Outputs.PatchBaselinePatchFilterGroup? PatchFilterGroup;

        [OutputConstructor]
        private PatchBaselineRule(
            int? approveAfterDays,

            Outputs.PatchBaselinePatchStringDate? approveUntilDate,

            string? complianceLevel,

            bool? enableNonSecurity,

            Outputs.PatchBaselinePatchFilterGroup? patchFilterGroup)
        {
            ApproveAfterDays = approveAfterDays;
            ApproveUntilDate = approveUntilDate;
            ComplianceLevel = complianceLevel;
            EnableNonSecurity = enableNonSecurity;
            PatchFilterGroup = patchFilterGroup;
        }
    }
}