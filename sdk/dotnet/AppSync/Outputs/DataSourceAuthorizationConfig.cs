// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-authorizationconfig.html
    /// </summary>
    [OutputType]
    public sealed class DataSourceAuthorizationConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-authorizationconfig.html#cfn-appsync-datasource-authorizationconfig-authorizationtype
        /// </summary>
        public readonly string AuthorizationType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-authorizationconfig.html#cfn-appsync-datasource-authorizationconfig-awsiamconfig
        /// </summary>
        public readonly Outputs.DataSourceAwsIamConfig? AwsIamConfig;

        [OutputConstructor]
        private DataSourceAuthorizationConfig(
            string authorizationType,

            Outputs.DataSourceAwsIamConfig? awsIamConfig)
        {
            AuthorizationType = authorizationType;
            AwsIamConfig = awsIamConfig;
        }
    }
}
