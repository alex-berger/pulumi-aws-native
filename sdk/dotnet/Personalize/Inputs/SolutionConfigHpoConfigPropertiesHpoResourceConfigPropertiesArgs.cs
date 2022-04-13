// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Personalize.Inputs
{

    /// <summary>
    /// Describes the resource configuration for hyperparameter optimization (HPO).
    /// </summary>
    public sealed class SolutionConfigHpoConfigPropertiesHpoResourceConfigPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of training jobs when you create a solution version. The maximum value for maxNumberOfTrainingJobs is 40.
        /// </summary>
        [Input("maxNumberOfTrainingJobs")]
        public Input<string>? MaxNumberOfTrainingJobs { get; set; }

        /// <summary>
        /// The maximum number of parallel training jobs when you create a solution version. The maximum value for maxParallelTrainingJobs is 10.
        /// </summary>
        [Input("maxParallelTrainingJobs")]
        public Input<string>? MaxParallelTrainingJobs { get; set; }

        public SolutionConfigHpoConfigPropertiesHpoResourceConfigPropertiesArgs()
        {
        }
    }
}
