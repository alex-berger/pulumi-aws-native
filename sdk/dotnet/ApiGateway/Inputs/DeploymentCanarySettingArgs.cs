// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-canarysetting.html
    /// </summary>
    public sealed class DeploymentCanarySettingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-canarysetting.html#cfn-apigateway-deployment-canarysetting-percenttraffic
        /// </summary>
        [Input("percentTraffic")]
        public Input<double>? PercentTraffic { get; set; }

        [Input("stageVariableOverrides")]
        private InputMap<string>? _stageVariableOverrides;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-canarysetting.html#cfn-apigateway-deployment-canarysetting-stagevariableoverrides
        /// </summary>
        public InputMap<string> StageVariableOverrides
        {
            get => _stageVariableOverrides ?? (_stageVariableOverrides = new InputMap<string>());
            set => _stageVariableOverrides = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-canarysetting.html#cfn-apigateway-deployment-canarysetting-usestagecache
        /// </summary>
        [Input("useStageCache")]
        public Input<bool>? UseStageCache { get; set; }

        public DeploymentCanarySettingArgs()
        {
        }
    }
}
