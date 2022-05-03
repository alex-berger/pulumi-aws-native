// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Timestream.Inputs
{

    /// <summary>
    /// Configuration for when the scheduled query is executed.
    /// </summary>
    public sealed class ScheduledQueryScheduleConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("scheduleExpression", required: true)]
        public Input<string> ScheduleExpression { get; set; } = null!;

        public ScheduledQueryScheduleConfigurationArgs()
        {
        }
    }
}
