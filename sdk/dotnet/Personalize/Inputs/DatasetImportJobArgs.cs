// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Personalize.Inputs
{

    /// <summary>
    /// Initial DatasetImportJob for the created dataset
    /// </summary>
    public sealed class DatasetImportJobArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon S3 bucket that contains the training data to import.
        /// </summary>
        [Input("dataSource")]
        public Input<Inputs.DatasetImportJobDataSourcePropertiesArgs>? DataSource { get; set; }

        /// <summary>
        /// The ARN of the dataset that receives the imported data
        /// </summary>
        [Input("datasetArn")]
        public Input<string>? DatasetArn { get; set; }

        /// <summary>
        /// The ARN of the dataset import job
        /// </summary>
        [Input("datasetImportJobArn")]
        public Input<string>? DatasetImportJobArn { get; set; }

        /// <summary>
        /// The name for the dataset import job.
        /// </summary>
        [Input("jobName")]
        public Input<string>? JobName { get; set; }

        /// <summary>
        /// The ARN of the IAM role that has permissions to read from the Amazon S3 data source.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        public DatasetImportJobArgs()
        {
        }
    }
}
