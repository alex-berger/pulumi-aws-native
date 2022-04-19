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
    /// Container image configuration object for the monitoring job.
    /// </summary>
    public sealed class ModelBiasJobDefinitionModelBiasAppSpecificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The S3 URI to an analysis configuration file
        /// </summary>
        [Input("configUri", required: true)]
        public Input<string> ConfigUri { get; set; } = null!;

        /// <summary>
        /// Sets the environment variables in the Docker container
        /// </summary>
        [Input("environment")]
        public Input<object>? Environment { get; set; }

        /// <summary>
        /// The container image to be run by the monitoring job.
        /// </summary>
        [Input("imageUri", required: true)]
        public Input<string> ImageUri { get; set; } = null!;

        public ModelBiasJobDefinitionModelBiasAppSpecificationArgs()
        {
        }
    }
}
