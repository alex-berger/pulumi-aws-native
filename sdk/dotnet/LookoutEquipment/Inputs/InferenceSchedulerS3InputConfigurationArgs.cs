// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutEquipment.Inputs
{

    /// <summary>
    /// Specifies configuration information for the input data for the inference, including input data S3 location.
    /// </summary>
    public sealed class InferenceSchedulerS3InputConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public InferenceSchedulerS3InputConfigurationArgs()
        {
        }
    }
}
