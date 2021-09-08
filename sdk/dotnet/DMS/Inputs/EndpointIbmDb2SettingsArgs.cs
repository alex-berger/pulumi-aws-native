// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DMS.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-ibmdb2settings.html
    /// </summary>
    public sealed class EndpointIbmDb2SettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-ibmdb2settings.html#cfn-dms-endpoint-ibmdb2settings-secretsmanageraccessrolearn
        /// </summary>
        [Input("secretsManagerAccessRoleArn")]
        public Input<string>? SecretsManagerAccessRoleArn { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-ibmdb2settings.html#cfn-dms-endpoint-ibmdb2settings-secretsmanagersecretid
        /// </summary>
        [Input("secretsManagerSecretId")]
        public Input<string>? SecretsManagerSecretId { get; set; }

        public EndpointIbmDb2SettingsArgs()
        {
        }
    }
}
