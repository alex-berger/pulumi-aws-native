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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.html
    /// </summary>
    public sealed class DeploymentGroupLoadBalancerInfoArgs : Pulumi.ResourceArgs
    {
        [Input("elbInfoList")]
        private InputList<Inputs.DeploymentGroupELBInfoArgs>? _elbInfoList;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.html#cfn-codedeploy-deploymentgroup-loadbalancerinfo-elbinfolist
        /// </summary>
        public InputList<Inputs.DeploymentGroupELBInfoArgs> ElbInfoList
        {
            get => _elbInfoList ?? (_elbInfoList = new InputList<Inputs.DeploymentGroupELBInfoArgs>());
            set => _elbInfoList = value;
        }

        [Input("targetGroupInfoList")]
        private InputList<Inputs.DeploymentGroupTargetGroupInfoArgs>? _targetGroupInfoList;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.html#cfn-codedeploy-deploymentgroup-loadbalancerinfo-targetgroupinfolist
        /// </summary>
        public InputList<Inputs.DeploymentGroupTargetGroupInfoArgs> TargetGroupInfoList
        {
            get => _targetGroupInfoList ?? (_targetGroupInfoList = new InputList<Inputs.DeploymentGroupTargetGroupInfoArgs>());
            set => _targetGroupInfoList = value;
        }

        public DeploymentGroupLoadBalancerInfoArgs()
        {
        }
    }
}
