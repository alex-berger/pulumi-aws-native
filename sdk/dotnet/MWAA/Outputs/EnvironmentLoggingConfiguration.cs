// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MWAA.Outputs
{

    /// <summary>
    /// Logging configuration for the environment.
    /// </summary>
    [OutputType]
    public sealed class EnvironmentLoggingConfiguration
    {
        public readonly Outputs.EnvironmentModuleLoggingConfiguration? DagProcessingLogs;
        public readonly Outputs.EnvironmentModuleLoggingConfiguration? SchedulerLogs;
        public readonly Outputs.EnvironmentModuleLoggingConfiguration? TaskLogs;
        public readonly Outputs.EnvironmentModuleLoggingConfiguration? WebserverLogs;
        public readonly Outputs.EnvironmentModuleLoggingConfiguration? WorkerLogs;

        [OutputConstructor]
        private EnvironmentLoggingConfiguration(
            Outputs.EnvironmentModuleLoggingConfiguration? dagProcessingLogs,

            Outputs.EnvironmentModuleLoggingConfiguration? schedulerLogs,

            Outputs.EnvironmentModuleLoggingConfiguration? taskLogs,

            Outputs.EnvironmentModuleLoggingConfiguration? webserverLogs,

            Outputs.EnvironmentModuleLoggingConfiguration? workerLogs)
        {
            DagProcessingLogs = dagProcessingLogs;
            SchedulerLogs = schedulerLogs;
            TaskLogs = taskLogs;
            WebserverLogs = webserverLogs;
            WorkerLogs = workerLogs;
        }
    }
}
