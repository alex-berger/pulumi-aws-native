// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DMS.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html
    /// </summary>
    [OutputType]
    public sealed class EndpointMicrosoftSqlServerSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-secretsmanageraccessrolearn
        /// </summary>
        public readonly string? SecretsManagerAccessRoleArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-secretsmanagersecretid
        /// </summary>
        public readonly string? SecretsManagerSecretId;

        [OutputConstructor]
        private EndpointMicrosoftSqlServerSettings(
            string? secretsManagerAccessRoleArn,

            string? secretsManagerSecretId)
        {
            SecretsManagerAccessRoleArn = secretsManagerAccessRoleArn;
            SecretsManagerSecretId = secretsManagerSecretId;
        }
    }
}
