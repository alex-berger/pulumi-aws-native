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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-profileconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class JobProfileConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-profileconfiguration.html#cfn-databrew-job-profileconfiguration-columnstatisticsconfigurations
        /// </summary>
        public readonly ImmutableArray<Outputs.JobColumnStatisticsConfiguration> ColumnStatisticsConfigurations;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-profileconfiguration.html#cfn-databrew-job-profileconfiguration-datasetstatisticsconfiguration
        /// </summary>
        public readonly Outputs.JobStatisticsConfiguration? DatasetStatisticsConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-profileconfiguration.html#cfn-databrew-job-profileconfiguration-profilecolumns
        /// </summary>
        public readonly ImmutableArray<Outputs.JobColumnSelector> ProfileColumns;

        [OutputConstructor]
        private JobProfileConfiguration(
            ImmutableArray<Outputs.JobColumnStatisticsConfiguration> columnStatisticsConfigurations,

            Outputs.JobStatisticsConfiguration? datasetStatisticsConfiguration,

            ImmutableArray<Outputs.JobColumnSelector> profileColumns)
        {
            ColumnStatisticsConfigurations = columnStatisticsConfigurations;
            DatasetStatisticsConfiguration = datasetStatisticsConfiguration;
            ProfileColumns = profileColumns;
        }
    }
}