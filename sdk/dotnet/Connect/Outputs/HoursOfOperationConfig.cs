// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Outputs
{

    /// <summary>
    /// Contains information about the hours of operation.
    /// </summary>
    [OutputType]
    public sealed class HoursOfOperationConfig
    {
        /// <summary>
        /// The day that the hours of operation applies to.
        /// </summary>
        public readonly Pulumi.AwsNative.Connect.HoursOfOperationConfigDay Day;
        /// <summary>
        /// The end time that your contact center closes.
        /// </summary>
        public readonly Outputs.HoursOfOperationTimeSlice EndTime;
        /// <summary>
        /// The start time that your contact center opens.
        /// </summary>
        public readonly Outputs.HoursOfOperationTimeSlice StartTime;

        [OutputConstructor]
        private HoursOfOperationConfig(
            Pulumi.AwsNative.Connect.HoursOfOperationConfigDay day,

            Outputs.HoursOfOperationTimeSlice endTime,

            Outputs.HoursOfOperationTimeSlice startTime)
        {
            Day = day;
            EndTime = endTime;
            StartTime = startTime;
        }
    }
}
