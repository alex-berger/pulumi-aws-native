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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-blueinstanceterminationoption.html
    /// </summary>
    [OutputType]
    public sealed class DeploymentGroupBlueInstanceTerminationOption
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-blueinstanceterminationoption.html#cfn-codedeploy-deploymentgroup-bluegreendeploymentconfiguration-blueinstanceterminationoption-action
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-blueinstanceterminationoption.html#cfn-codedeploy-deploymentgroup-bluegreendeploymentconfiguration-blueinstanceterminationoption-terminationwaittimeinminutes
        /// </summary>
        public readonly int? TerminationWaitTimeInMinutes;

        [OutputConstructor]
        private DeploymentGroupBlueInstanceTerminationOption(
            string? action,

            int? terminationWaitTimeInMinutes)
        {
            Action = action;
            TerminationWaitTimeInMinutes = terminationWaitTimeInMinutes;
        }
    }
}