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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-bluegreendeploymentconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class DeploymentGroupBlueGreenDeploymentConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-bluegreendeploymentconfiguration.html#cfn-codedeploy-deploymentgroup-bluegreendeploymentconfiguration-deploymentreadyoption
        /// </summary>
        public readonly Outputs.DeploymentGroupDeploymentReadyOption? DeploymentReadyOption;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-bluegreendeploymentconfiguration.html#cfn-codedeploy-deploymentgroup-bluegreendeploymentconfiguration-greenfleetprovisioningoption
        /// </summary>
        public readonly Outputs.DeploymentGroupGreenFleetProvisioningOption? GreenFleetProvisioningOption;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-bluegreendeploymentconfiguration.html#cfn-codedeploy-deploymentgroup-bluegreendeploymentconfiguration-terminateblueinstancesondeploymentsuccess
        /// </summary>
        public readonly Outputs.DeploymentGroupBlueInstanceTerminationOption? TerminateBlueInstancesOnDeploymentSuccess;

        [OutputConstructor]
        private DeploymentGroupBlueGreenDeploymentConfiguration(
            Outputs.DeploymentGroupDeploymentReadyOption? deploymentReadyOption,

            Outputs.DeploymentGroupGreenFleetProvisioningOption? greenFleetProvisioningOption,

            Outputs.DeploymentGroupBlueInstanceTerminationOption? terminateBlueInstancesOnDeploymentSuccess)
        {
            DeploymentReadyOption = deploymentReadyOption;
            GreenFleetProvisioningOption = greenFleetProvisioningOption;
            TerminateBlueInstancesOnDeploymentSuccess = terminateBlueInstancesOnDeploymentSuccess;
        }
    }
}