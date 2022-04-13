// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Outputs
{

    /// <summary>
    /// A statistical ranking (percentile) which indicates a threshold value by which a behavior is determined to be in compliance or in violation of the behavior.
    /// </summary>
    [OutputType]
    public sealed class SecurityProfileStatisticalThreshold
    {
        /// <summary>
        /// The percentile which resolves to a threshold value by which compliance with a behavior is determined
        /// </summary>
        public readonly Pulumi.AwsNative.IoT.SecurityProfileStatisticalThresholdStatistic? Statistic;

        [OutputConstructor]
        private SecurityProfileStatisticalThreshold(Pulumi.AwsNative.IoT.SecurityProfileStatisticalThresholdStatistic? statistic)
        {
            Statistic = statistic;
        }
    }
}
