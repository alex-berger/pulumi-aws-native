// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutMetrics.Inputs
{

    public sealed class AnomalyDetectorFileFormatDescriptorArgs : Pulumi.ResourceArgs
    {
        [Input("csvFormatDescriptor")]
        public Input<Inputs.AnomalyDetectorCsvFormatDescriptorArgs>? CsvFormatDescriptor { get; set; }

        [Input("jsonFormatDescriptor")]
        public Input<Inputs.AnomalyDetectorJsonFormatDescriptorArgs>? JsonFormatDescriptor { get; set; }

        public AnomalyDetectorFileFormatDescriptorArgs()
        {
        }
    }
}