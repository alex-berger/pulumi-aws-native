// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FIS.Inputs
{

    public sealed class ExperimentTemplateLogConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("cloudWatchLogsConfiguration")]
        public Input<Inputs.ExperimentTemplateLogConfigurationCloudWatchLogsConfigurationPropertiesArgs>? CloudWatchLogsConfiguration { get; set; }

        [Input("logSchemaVersion", required: true)]
        public Input<int> LogSchemaVersion { get; set; } = null!;

        [Input("s3Configuration")]
        public Input<Inputs.ExperimentTemplateLogConfigurationS3ConfigurationPropertiesArgs>? S3Configuration { get; set; }

        public ExperimentTemplateLogConfigurationArgs()
        {
        }
    }
}
