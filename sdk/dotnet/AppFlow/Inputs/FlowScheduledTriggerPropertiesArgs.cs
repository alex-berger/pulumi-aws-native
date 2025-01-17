// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    /// <summary>
    /// Details required for scheduled trigger type
    /// </summary>
    public sealed class FlowScheduledTriggerPropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("dataPullMode")]
        public Input<Pulumi.AwsNative.AppFlow.FlowScheduledTriggerPropertiesDataPullMode>? DataPullMode { get; set; }

        [Input("scheduleEndTime")]
        public Input<double>? ScheduleEndTime { get; set; }

        [Input("scheduleExpression", required: true)]
        public Input<string> ScheduleExpression { get; set; } = null!;

        [Input("scheduleOffset")]
        public Input<double>? ScheduleOffset { get; set; }

        [Input("scheduleStartTime")]
        public Input<double>? ScheduleStartTime { get; set; }

        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public FlowScheduledTriggerPropertiesArgs()
        {
        }
    }
}
