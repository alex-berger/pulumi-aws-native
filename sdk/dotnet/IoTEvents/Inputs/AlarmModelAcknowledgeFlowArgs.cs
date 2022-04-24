// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Inputs
{

    /// <summary>
    /// Specifies whether to get notified for alarm state changes.
    /// </summary>
    public sealed class AlarmModelAcknowledgeFlowArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The value must be TRUE or FALSE. If TRUE, you receive a notification when the alarm state changes. You must choose to acknowledge the notification before the alarm state can return to NORMAL. If FALSE, you won't receive notifications. The alarm automatically changes to the NORMAL state when the input property value returns to the specified range.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public AlarmModelAcknowledgeFlowArgs()
        {
        }
    }
}
