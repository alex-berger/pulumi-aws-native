// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// Baseline configuration used to validate that the data conforms to the specified constraints and statistics.
    /// </summary>
    public sealed class ModelQualityJobDefinitionModelQualityBaselineConfigArgs : Pulumi.ResourceArgs
    {
        [Input("baseliningJobName")]
        public Input<string>? BaseliningJobName { get; set; }

        [Input("constraintsResource")]
        public Input<Inputs.ModelQualityJobDefinitionConstraintsResourceArgs>? ConstraintsResource { get; set; }

        public ModelQualityJobDefinitionModelQualityBaselineConfigArgs()
        {
        }
    }
}
