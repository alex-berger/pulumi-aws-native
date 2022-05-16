// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    /// <summary>
    /// Specifies data related to access patterns to be collected and made available to analyze the tradeoffs between different storage classes for an Amazon S3 bucket.
    /// </summary>
    public sealed class BucketStorageClassAnalysisArgs : Pulumi.ResourceArgs
    {
        [Input("dataExport")]
        public Input<Inputs.BucketDataExportArgs>? DataExport { get; set; }

        public BucketStorageClassAnalysisArgs()
        {
        }
    }
}