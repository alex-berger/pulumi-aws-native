// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeBuild.Outputs
{

    [OutputType]
    public sealed class ProjectCloudWatchLogsConfig
    {
        public readonly string? GroupName;
        public readonly string Status;
        public readonly string? StreamName;

        [OutputConstructor]
        private ProjectCloudWatchLogsConfig(
            string? groupName,

            string status,

            string? streamName)
        {
            GroupName = groupName;
            Status = status;
            StreamName = streamName;
        }
    }
}
