// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeDeploy.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html
    /// </summary>
    [OutputType]
    public sealed class DeploymentGroupRevisionLocation
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-githublocation
        /// </summary>
        public readonly Outputs.DeploymentGroupGitHubLocation? GitHubLocation;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-revisiontype
        /// </summary>
        public readonly string? RevisionType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-s3location
        /// </summary>
        public readonly Outputs.DeploymentGroupS3Location? S3Location;

        [OutputConstructor]
        private DeploymentGroupRevisionLocation(
            Outputs.DeploymentGroupGitHubLocation? gitHubLocation,

            string? revisionType,

            Outputs.DeploymentGroupS3Location? s3Location)
        {
            GitHubLocation = gitHubLocation;
            RevisionType = revisionType;
            S3Location = s3Location;
        }
    }
}
