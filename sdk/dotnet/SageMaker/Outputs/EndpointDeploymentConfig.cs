// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-deploymentconfig.html
    /// </summary>
    [OutputType]
    public sealed class EndpointDeploymentConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-deploymentconfig.html#cfn-sagemaker-endpoint-deploymentconfig-autorollbackconfiguration
        /// </summary>
        public readonly Outputs.EndpointAutoRollbackConfig? AutoRollbackConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-deploymentconfig.html#cfn-sagemaker-endpoint-deploymentconfig-bluegreenupdatepolicy
        /// </summary>
        public readonly Outputs.EndpointBlueGreenUpdatePolicy BlueGreenUpdatePolicy;

        [OutputConstructor]
        private EndpointDeploymentConfig(
            Outputs.EndpointAutoRollbackConfig? autoRollbackConfiguration,

            Outputs.EndpointBlueGreenUpdatePolicy blueGreenUpdatePolicy)
        {
            AutoRollbackConfiguration = autoRollbackConfiguration;
            BlueGreenUpdatePolicy = blueGreenUpdatePolicy;
        }
    }
}