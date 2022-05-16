// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecretsManager.Inputs
{

    public sealed class RotationScheduleHostedRotationLambdaArgs : Pulumi.ResourceArgs
    {
        [Input("excludeCharacters")]
        public Input<string>? ExcludeCharacters { get; set; }

        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        [Input("masterSecretArn")]
        public Input<string>? MasterSecretArn { get; set; }

        [Input("masterSecretKmsKeyArn")]
        public Input<string>? MasterSecretKmsKeyArn { get; set; }

        [Input("rotationLambdaName")]
        public Input<string>? RotationLambdaName { get; set; }

        [Input("rotationType", required: true)]
        public Input<string> RotationType { get; set; } = null!;

        [Input("superuserSecretArn")]
        public Input<string>? SuperuserSecretArn { get; set; }

        [Input("superuserSecretKmsKeyArn")]
        public Input<string>? SuperuserSecretKmsKeyArn { get; set; }

        [Input("vpcSecurityGroupIds")]
        public Input<string>? VpcSecurityGroupIds { get; set; }

        [Input("vpcSubnetIds")]
        public Input<string>? VpcSubnetIds { get; set; }

        public RotationScheduleHostedRotationLambdaArgs()
        {
        }
    }
}