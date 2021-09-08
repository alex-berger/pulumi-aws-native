// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-remediationconfiguration-ssmcontrols.html
    /// </summary>
    [OutputType]
    public sealed class RemediationConfigurationSsmControls
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-remediationconfiguration-ssmcontrols.html#cfn-config-remediationconfiguration-ssmcontrols-concurrentexecutionratepercentage
        /// </summary>
        public readonly int? ConcurrentExecutionRatePercentage;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-remediationconfiguration-ssmcontrols.html#cfn-config-remediationconfiguration-ssmcontrols-errorpercentage
        /// </summary>
        public readonly int? ErrorPercentage;

        [OutputConstructor]
        private RemediationConfigurationSsmControls(
            int? concurrentExecutionRatePercentage,

            int? errorPercentage)
        {
            ConcurrentExecutionRatePercentage = concurrentExecutionRatePercentage;
            ErrorPercentage = errorPercentage;
        }
    }
}
