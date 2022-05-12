// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutMetrics.Inputs
{

    public sealed class AnomalyDetectorTimestampColumnArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A timestamp format for the timestamps in the dataset
        /// </summary>
        [Input("columnFormat")]
        public Input<string>? ColumnFormat { get; set; }

        [Input("columnName")]
        public Input<string>? ColumnName { get; set; }

        public AnomalyDetectorTimestampColumnArgs()
        {
        }
    }
}
