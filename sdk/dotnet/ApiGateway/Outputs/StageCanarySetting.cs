// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway.Outputs
{

    /// <summary>
    /// Specifies settings for the canary deployment in this stage.
    /// </summary>
    [OutputType]
    public sealed class StageCanarySetting
    {
        /// <summary>
        /// The identifier of the deployment that the stage points to.
        /// </summary>
        public readonly string? DeploymentId;
        /// <summary>
        /// The percentage (0-100) of traffic diverted to a canary deployment.
        /// </summary>
        public readonly double? PercentTraffic;
        /// <summary>
        /// Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values.
        /// </summary>
        public readonly object? StageVariableOverrides;
        /// <summary>
        /// Whether the canary deployment uses the stage cache or not.
        /// </summary>
        public readonly bool? UseStageCache;

        [OutputConstructor]
        private StageCanarySetting(
            string? deploymentId,

            double? percentTraffic,

            object? stageVariableOverrides,

            bool? useStageCache)
        {
            DeploymentId = deploymentId;
            PercentTraffic = percentTraffic;
            StageVariableOverrides = stageVariableOverrides;
            UseStageCache = useStageCache;
        }
    }
}
