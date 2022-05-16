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
    /// Specifies configuration information for the input data for the inference, including timestamp format and delimiter.
    /// </summary>
    public sealed class InferenceSchedulerInputNameConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates the delimiter character used between items in the data.
        /// </summary>
        [Input("componentTimestampDelimiter")]
        public Input<string>? ComponentTimestampDelimiter { get; set; }

        /// <summary>
        /// The format of the timestamp, whether Epoch time, or standard, with or without hyphens (-).
        /// </summary>
        [Input("timestampFormat")]
        public Input<string>? TimestampFormat { get; set; }

        public InferenceSchedulerInputNameConfigurationArgs()
        {
        }
    }
}