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
    /// Defines when your alarm is invoked.
    /// </summary>
    public sealed class AlarmModelAlarmRuleArgs : Pulumi.ResourceArgs
    {
        [Input("simpleRule")]
        public Input<Inputs.AlarmModelSimpleRuleArgs>? SimpleRule { get; set; }

        public AlarmModelAlarmRuleArgs()
        {
        }
    }
}