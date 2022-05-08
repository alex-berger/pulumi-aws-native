// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Outputs
{

    [OutputType]
    public sealed class DatasetVariable
    {
        public readonly Outputs.DatasetContentVersionValue? DatasetContentVersionValue;
        public readonly double? DoubleValue;
        public readonly Outputs.DatasetOutputFileUriValue? OutputFileUriValue;
        public readonly string? StringValue;
        public readonly string VariableName;

        [OutputConstructor]
        private DatasetVariable(
            Outputs.DatasetContentVersionValue? datasetContentVersionValue,

            double? doubleValue,

            Outputs.DatasetOutputFileUriValue? outputFileUriValue,

            string? stringValue,

            string variableName)
        {
            DatasetContentVersionValue = datasetContentVersionValue;
            DoubleValue = doubleValue;
            OutputFileUriValue = outputFileUriValue;
            StringValue = stringValue;
            VariableName = variableName;
        }
    }
}
