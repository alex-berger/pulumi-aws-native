// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeBuild.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html
    /// </summary>
    [OutputType]
    public sealed class ProjectProjectTriggers
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html#cfn-codebuild-project-projecttriggers-buildtype
        /// </summary>
        public readonly string? BuildType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html#cfn-codebuild-project-projecttriggers-filtergroups
        /// </summary>
        public readonly ImmutableArray<Outputs.ProjectFilterGroup> FilterGroups;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html#cfn-codebuild-project-projecttriggers-webhook
        /// </summary>
        public readonly bool? Webhook;

        [OutputConstructor]
        private ProjectProjectTriggers(
            string? buildType,

            ImmutableArray<Outputs.ProjectFilterGroup> filterGroups,

            bool? webhook)
        {
            BuildType = buildType;
            FilterGroups = filterGroups;
            Webhook = webhook;
        }
    }
}
