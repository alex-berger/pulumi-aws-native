// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Outputs
{

    [OutputType]
    public sealed class GraphQLApiLogConfig
    {
        public readonly string? CloudWatchLogsRoleArn;
        public readonly bool? ExcludeVerboseContent;
        public readonly string? FieldLogLevel;

        [OutputConstructor]
        private GraphQLApiLogConfig(
            string? cloudWatchLogsRoleArn,

            bool? excludeVerboseContent,

            string? fieldLogLevel)
        {
            CloudWatchLogsRoleArn = cloudWatchLogsRoleArn;
            ExcludeVerboseContent = excludeVerboseContent;
            FieldLogLevel = fieldLogLevel;
        }
    }
}