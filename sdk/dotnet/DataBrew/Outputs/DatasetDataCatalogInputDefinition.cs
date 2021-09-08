// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datacataloginputdefinition.html
    /// </summary>
    [OutputType]
    public sealed class DatasetDataCatalogInputDefinition
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datacataloginputdefinition.html#cfn-databrew-dataset-datacataloginputdefinition-catalogid
        /// </summary>
        public readonly string? CatalogId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datacataloginputdefinition.html#cfn-databrew-dataset-datacataloginputdefinition-databasename
        /// </summary>
        public readonly string? DatabaseName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datacataloginputdefinition.html#cfn-databrew-dataset-datacataloginputdefinition-tablename
        /// </summary>
        public readonly string? TableName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datacataloginputdefinition.html#cfn-databrew-dataset-datacataloginputdefinition-tempdirectory
        /// </summary>
        public readonly Outputs.DatasetS3Location? TempDirectory;

        [OutputConstructor]
        private DatasetDataCatalogInputDefinition(
            string? catalogId,

            string? databaseName,

            string? tableName,

            Outputs.DatasetS3Location? tempDirectory)
        {
            CatalogId = catalogId;
            DatabaseName = databaseName;
            TableName = tableName;
            TempDirectory = tempDirectory;
        }
    }
}