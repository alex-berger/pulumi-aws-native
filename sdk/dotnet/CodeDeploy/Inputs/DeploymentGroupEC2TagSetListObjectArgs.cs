// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeDeploy.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagsetlistobject.html
    /// </summary>
    public sealed class DeploymentGroupEC2TagSetListObjectArgs : Pulumi.ResourceArgs
    {
        [Input("ec2TagGroup")]
        private InputList<Inputs.DeploymentGroupEC2TagFilterArgs>? _ec2TagGroup;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagsetlistobject.html#cfn-codedeploy-deploymentgroup-ec2tagsetlistobject-ec2taggroup
        /// </summary>
        public InputList<Inputs.DeploymentGroupEC2TagFilterArgs> Ec2TagGroup
        {
            get => _ec2TagGroup ?? (_ec2TagGroup = new InputList<Inputs.DeploymentGroupEC2TagFilterArgs>());
            set => _ec2TagGroup = value;
        }

        public DeploymentGroupEC2TagSetListObjectArgs()
        {
        }
    }
}
