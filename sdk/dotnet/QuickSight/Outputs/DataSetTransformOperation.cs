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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperation.html
    /// </summary>
    [OutputType]
    public sealed class DataSetTransformOperation
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperation.html#cfn-quicksight-dataset-transformoperation-castcolumntypeoperation
        /// </summary>
        public readonly Outputs.DataSetCastColumnTypeOperation? CastColumnTypeOperation;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperation.html#cfn-quicksight-dataset-transformoperation-createcolumnsoperation
        /// </summary>
        public readonly Outputs.DataSetCreateColumnsOperation? CreateColumnsOperation;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperation.html#cfn-quicksight-dataset-transformoperation-filteroperation
        /// </summary>
        public readonly Outputs.DataSetFilterOperation? FilterOperation;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperation.html#cfn-quicksight-dataset-transformoperation-projectoperation
        /// </summary>
        public readonly Outputs.DataSetProjectOperation? ProjectOperation;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperation.html#cfn-quicksight-dataset-transformoperation-renamecolumnoperation
        /// </summary>
        public readonly Outputs.DataSetRenameColumnOperation? RenameColumnOperation;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperation.html#cfn-quicksight-dataset-transformoperation-tagcolumnoperation
        /// </summary>
        public readonly Outputs.DataSetTagColumnOperation? TagColumnOperation;

        [OutputConstructor]
        private DataSetTransformOperation(
            Outputs.DataSetCastColumnTypeOperation? castColumnTypeOperation,

            Outputs.DataSetCreateColumnsOperation? createColumnsOperation,

            Outputs.DataSetFilterOperation? filterOperation,

            Outputs.DataSetProjectOperation? projectOperation,

            Outputs.DataSetRenameColumnOperation? renameColumnOperation,

            Outputs.DataSetTagColumnOperation? tagColumnOperation)
        {
            CastColumnTypeOperation = castColumnTypeOperation;
            CreateColumnsOperation = createColumnsOperation;
            FilterOperation = filterOperation;
            ProjectOperation = projectOperation;
            RenameColumnOperation = renameColumnOperation;
            TagColumnOperation = tagColumnOperation;
        }
    }
}