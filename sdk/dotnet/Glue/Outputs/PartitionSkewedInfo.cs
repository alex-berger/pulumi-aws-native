// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-skewedinfo.html
    /// </summary>
    [OutputType]
    public sealed class PartitionSkewedInfo
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-skewedinfo.html#cfn-glue-partition-skewedinfo-skewedcolumnnames
        /// </summary>
        public readonly ImmutableArray<string> SkewedColumnNames;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-skewedinfo.html#cfn-glue-partition-skewedinfo-skewedcolumnvaluelocationmaps
        /// </summary>
        public readonly Union<System.Text.Json.JsonElement, string>? SkewedColumnValueLocationMaps;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-skewedinfo.html#cfn-glue-partition-skewedinfo-skewedcolumnvalues
        /// </summary>
        public readonly ImmutableArray<string> SkewedColumnValues;

        [OutputConstructor]
        private PartitionSkewedInfo(
            ImmutableArray<string> skewedColumnNames,

            Union<System.Text.Json.JsonElement, string>? skewedColumnValueLocationMaps,

            ImmutableArray<string> skewedColumnValues)
        {
            SkewedColumnNames = skewedColumnNames;
            SkewedColumnValueLocationMaps = skewedColumnValueLocationMaps;
            SkewedColumnValues = skewedColumnValues;
        }
    }
}
