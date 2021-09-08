// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columnlevelpermissionrule.html
    /// </summary>
    [OutputType]
    public sealed class DataSetColumnLevelPermissionRule
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columnlevelpermissionrule.html#cfn-quicksight-dataset-columnlevelpermissionrule-columnnames
        /// </summary>
        public readonly ImmutableArray<string> ColumnNames;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columnlevelpermissionrule.html#cfn-quicksight-dataset-columnlevelpermissionrule-principals
        /// </summary>
        public readonly ImmutableArray<string> Principals;

        [OutputConstructor]
        private DataSetColumnLevelPermissionRule(
            ImmutableArray<string> columnNames,

            ImmutableArray<string> principals)
        {
            ColumnNames = columnNames;
            Principals = principals;
        }
    }
}