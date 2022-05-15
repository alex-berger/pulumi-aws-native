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
    /// The output object for a monitoring job.
    /// </summary>
    public sealed class ModelExplainabilityJobDefinitionMonitoringOutputArgs : Pulumi.ResourceArgs
    {
        [Input("s3Output", required: true)]
        public Input<Inputs.ModelExplainabilityJobDefinitionS3OutputArgs> S3Output { get; set; } = null!;

        public ModelExplainabilityJobDefinitionMonitoringOutputArgs()
        {
        }
    }
}
