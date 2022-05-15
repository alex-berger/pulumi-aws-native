// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Outputs
{

    [OutputType]
    public sealed class TableSkewedInfo
    {
        public readonly ImmutableArray<string> SkewedColumnNames;
        public readonly object? SkewedColumnValueLocationMaps;
        public readonly ImmutableArray<string> SkewedColumnValues;

        [OutputConstructor]
        private TableSkewedInfo(
            ImmutableArray<string> skewedColumnNames,

            object? skewedColumnValueLocationMaps,

            ImmutableArray<string> skewedColumnValues)
        {
            SkewedColumnNames = skewedColumnNames;
            SkewedColumnValueLocationMaps = skewedColumnValueLocationMaps;
            SkewedColumnValues = skewedColumnValues;
        }
    }
}
