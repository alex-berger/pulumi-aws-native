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
    public sealed class ProjectBuildBatchConfig
    {
        public readonly string? BatchReportMode;
        public readonly bool? CombineArtifacts;
        public readonly Outputs.ProjectBatchRestrictions? Restrictions;
        public readonly string? ServiceRole;
        public readonly int? TimeoutInMins;

        [OutputConstructor]
        private ProjectBuildBatchConfig(
            string? batchReportMode,

            bool? combineArtifacts,

            Outputs.ProjectBatchRestrictions? restrictions,

            string? serviceRole,

            int? timeoutInMins)
        {
            BatchReportMode = batchReportMode;
            CombineArtifacts = combineArtifacts;
            Restrictions = restrictions;
            ServiceRole = serviceRole;
            TimeoutInMins = timeoutInMins;
        }
    }
}