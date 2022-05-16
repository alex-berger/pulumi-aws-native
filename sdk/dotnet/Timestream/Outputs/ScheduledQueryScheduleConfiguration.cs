// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Timestream.Outputs
{

    /// <summary>
    /// Configuration for when the scheduled query is executed.
    /// </summary>
    [OutputType]
    public sealed class ScheduledQueryScheduleConfiguration
    {
        public readonly string ScheduleExpression;

        [OutputConstructor]
        private ScheduledQueryScheduleConfiguration(string scheduleExpression)
        {
            ScheduleExpression = scheduleExpression;
        }
    }
}