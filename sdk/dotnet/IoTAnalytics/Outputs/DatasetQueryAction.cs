// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-queryaction.html
    /// </summary>
    [OutputType]
    public sealed class DatasetQueryAction
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-queryaction.html#cfn-iotanalytics-dataset-queryaction-filters
        /// </summary>
        public readonly ImmutableArray<Outputs.DatasetFilter> Filters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-queryaction.html#cfn-iotanalytics-dataset-queryaction-sqlquery
        /// </summary>
        public readonly string SqlQuery;

        [OutputConstructor]
        private DatasetQueryAction(
            ImmutableArray<Outputs.DatasetFilter> filters,

            string sqlQuery)
        {
            Filters = filters;
            SqlQuery = sqlQuery;
        }
    }
}