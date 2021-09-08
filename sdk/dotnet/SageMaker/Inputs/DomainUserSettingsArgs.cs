// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-usersettings.html
    /// </summary>
    public sealed class DomainUserSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-usersettings.html#cfn-sagemaker-domain-usersettings-executionrole
        /// </summary>
        [Input("executionRole")]
        public Input<string>? ExecutionRole { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-usersettings.html#cfn-sagemaker-domain-usersettings-jupyterserverappsettings
        /// </summary>
        [Input("jupyterServerAppSettings")]
        public Input<Inputs.DomainJupyterServerAppSettingsArgs>? JupyterServerAppSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-usersettings.html#cfn-sagemaker-domain-usersettings-kernelgatewayappsettings
        /// </summary>
        [Input("kernelGatewayAppSettings")]
        public Input<Inputs.DomainKernelGatewayAppSettingsArgs>? KernelGatewayAppSettings { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-usersettings.html#cfn-sagemaker-domain-usersettings-securitygroups
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-usersettings.html#cfn-sagemaker-domain-usersettings-sharingsettings
        /// </summary>
        [Input("sharingSettings")]
        public Input<Inputs.DomainSharingSettingsArgs>? SharingSettings { get; set; }

        public DomainUserSettingsArgs()
        {
        }
    }
}