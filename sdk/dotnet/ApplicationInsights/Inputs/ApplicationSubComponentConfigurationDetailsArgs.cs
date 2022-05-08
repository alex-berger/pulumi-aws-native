// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationInsights.Inputs
{

    /// <summary>
    /// The configuration settings of sub components.
    /// </summary>
    public sealed class ApplicationSubComponentConfigurationDetailsArgs : Pulumi.ResourceArgs
    {
        [Input("alarmMetrics")]
        private InputList<Inputs.ApplicationAlarmMetricArgs>? _alarmMetrics;

        /// <summary>
        /// A list of metrics to monitor for the component.
        /// </summary>
        public InputList<Inputs.ApplicationAlarmMetricArgs> AlarmMetrics
        {
            get => _alarmMetrics ?? (_alarmMetrics = new InputList<Inputs.ApplicationAlarmMetricArgs>());
            set => _alarmMetrics = value;
        }

        [Input("logs")]
        private InputList<Inputs.ApplicationLogArgs>? _logs;

        /// <summary>
        /// A list of logs to monitor for the component.
        /// </summary>
        public InputList<Inputs.ApplicationLogArgs> Logs
        {
            get => _logs ?? (_logs = new InputList<Inputs.ApplicationLogArgs>());
            set => _logs = value;
        }

        [Input("windowsEvents")]
        private InputList<Inputs.ApplicationWindowsEventArgs>? _windowsEvents;

        /// <summary>
        /// A list of Windows Events to log.
        /// </summary>
        public InputList<Inputs.ApplicationWindowsEventArgs> WindowsEvents
        {
            get => _windowsEvents ?? (_windowsEvents = new InputList<Inputs.ApplicationWindowsEventArgs>());
            set => _windowsEvents = value;
        }

        public ApplicationSubComponentConfigurationDetailsArgs()
        {
        }
    }
}
