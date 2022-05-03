// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// Specifies a time limit for how long the monitoring job is allowed to run.
    /// </summary>
    [OutputType]
    public sealed class ModelExplainabilityJobDefinitionStoppingCondition
    {
        /// <summary>
        /// The maximum runtime allowed in seconds.
        /// </summary>
        public readonly int MaxRuntimeInSeconds;

        [OutputConstructor]
        private ModelExplainabilityJobDefinitionStoppingCondition(int maxRuntimeInSeconds)
        {
            MaxRuntimeInSeconds = maxRuntimeInSeconds;
        }
    }
}
