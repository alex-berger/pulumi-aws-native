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
    /// The inputs for a monitoring job.
    /// </summary>
    [OutputType]
    public sealed class ModelQualityJobDefinitionModelQualityJobInput
    {
        public readonly Outputs.ModelQualityJobDefinitionEndpointInput EndpointInput;
        public readonly Outputs.ModelQualityJobDefinitionMonitoringGroundTruthS3Input GroundTruthS3Input;

        [OutputConstructor]
        private ModelQualityJobDefinitionModelQualityJobInput(
            Outputs.ModelQualityJobDefinitionEndpointInput endpointInput,

            Outputs.ModelQualityJobDefinitionMonitoringGroundTruthS3Input groundTruthS3Input)
        {
            EndpointInput = endpointInput;
            GroundTruthS3Input = groundTruthS3Input;
        }
    }
}
