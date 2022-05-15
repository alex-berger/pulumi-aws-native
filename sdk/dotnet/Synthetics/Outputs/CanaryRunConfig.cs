// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Synthetics.Outputs
{

    [OutputType]
    public sealed class CanaryRunConfig
    {
        /// <summary>
        /// Enable active tracing if set to true
        /// </summary>
        public readonly bool? ActiveTracing;
        /// <summary>
        /// Environment variable key-value pairs.
        /// </summary>
        public readonly object? EnvironmentVariables;
        /// <summary>
        /// Provide maximum memory available for canary in MB
        /// </summary>
        public readonly int? MemoryInMB;
        /// <summary>
        /// Provide maximum canary timeout per run in seconds
        /// </summary>
        public readonly int? TimeoutInSeconds;

        [OutputConstructor]
        private CanaryRunConfig(
            bool? activeTracing,

            object? environmentVariables,

            int? memoryInMB,

            int? timeoutInSeconds)
        {
            ActiveTracing = activeTracing;
            EnvironmentVariables = environmentVariables;
            MemoryInMB = memoryInMB;
            TimeoutInSeconds = timeoutInSeconds;
        }
    }
}
