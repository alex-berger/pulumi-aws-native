// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight
{
    public static class GetDataSet
    {
        /// <summary>
        /// Definition of the AWS::QuickSight::DataSet Resource Type.
        /// </summary>
        public static Task<GetDataSetResult> InvokeAsync(GetDataSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDataSetResult>("aws-native:quicksight:getDataSet", args ?? new GetDataSetArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of the AWS::QuickSight::DataSet Resource Type.
        /// </summary>
        public static Output<GetDataSetResult> Invoke(GetDataSetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDataSetResult>("aws-native:quicksight:getDataSet", args ?? new GetDataSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataSetArgs : Pulumi.InvokeArgs
    {
        [Input("awsAccountId", required: true)]
        public string AwsAccountId { get; set; } = null!;

        [Input("dataSetId", required: true)]
        public string DataSetId { get; set; } = null!;

        public GetDataSetArgs()
        {
        }
    }

    public sealed class GetDataSetInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("awsAccountId", required: true)]
        public Input<string> AwsAccountId { get; set; } = null!;

        [Input("dataSetId", required: true)]
        public Input<string> DataSetId { get; set; } = null!;

        public GetDataSetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDataSetResult
    {
        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) of the resource.&lt;/p&gt;
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// &lt;p&gt;Groupings of columns that work together in certain QuickSight features. Currently, only geospatial hierarchy is supported.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSetColumnGroup> ColumnGroups;
        public readonly ImmutableArray<Outputs.DataSetColumnLevelPermissionRule> ColumnLevelPermissionRules;
        /// <summary>
        /// &lt;p&gt;The amount of SPICE capacity used by this dataset. This is 0 if the dataset isn't
        ///             imported into SPICE.&lt;/p&gt;
        /// </summary>
        public readonly double? ConsumedSpiceCapacityInBytes;
        /// <summary>
        /// &lt;p&gt;The time that this dataset was created.&lt;/p&gt;
        /// </summary>
        public readonly string? CreatedTime;
        public readonly Pulumi.AwsNative.QuickSight.DataSetImportMode? ImportMode;
        /// <summary>
        /// &lt;p&gt;The last time that this dataset was updated.&lt;/p&gt;
        /// </summary>
        public readonly string? LastUpdatedTime;
        public readonly Outputs.DataSetLogicalTableMap? LogicalTableMap;
        /// <summary>
        /// &lt;p&gt;The display name for the dataset.&lt;/p&gt;
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// &lt;p&gt;The list of columns after all transforms. These columns are available in templates,
        ///             analyses, and dashboards.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSetOutputColumn> OutputColumns;
        /// <summary>
        /// &lt;p&gt;A list of resource permissions on the dataset.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSetResourcePermission> Permissions;
        public readonly Outputs.DataSetPhysicalTableMap? PhysicalTableMap;
        public readonly Outputs.DataSetRowLevelPermissionDataSet? RowLevelPermissionDataSet;
        /// <summary>
        /// &lt;p&gt;Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSetTag> Tags;

        [OutputConstructor]
        private GetDataSetResult(
            string? arn,

            ImmutableArray<Outputs.DataSetColumnGroup> columnGroups,

            ImmutableArray<Outputs.DataSetColumnLevelPermissionRule> columnLevelPermissionRules,

            double? consumedSpiceCapacityInBytes,

            string? createdTime,

            Pulumi.AwsNative.QuickSight.DataSetImportMode? importMode,

            string? lastUpdatedTime,

            Outputs.DataSetLogicalTableMap? logicalTableMap,

            string? name,

            ImmutableArray<Outputs.DataSetOutputColumn> outputColumns,

            ImmutableArray<Outputs.DataSetResourcePermission> permissions,

            Outputs.DataSetPhysicalTableMap? physicalTableMap,

            Outputs.DataSetRowLevelPermissionDataSet? rowLevelPermissionDataSet,

            ImmutableArray<Outputs.DataSetTag> tags)
        {
            Arn = arn;
            ColumnGroups = columnGroups;
            ColumnLevelPermissionRules = columnLevelPermissionRules;
            ConsumedSpiceCapacityInBytes = consumedSpiceCapacityInBytes;
            CreatedTime = createdTime;
            ImportMode = importMode;
            LastUpdatedTime = lastUpdatedTime;
            LogicalTableMap = logicalTableMap;
            Name = name;
            OutputColumns = outputColumns;
            Permissions = permissions;
            PhysicalTableMap = physicalTableMap;
            RowLevelPermissionDataSet = rowLevelPermissionDataSet;
            Tags = tags;
        }
    }
}
