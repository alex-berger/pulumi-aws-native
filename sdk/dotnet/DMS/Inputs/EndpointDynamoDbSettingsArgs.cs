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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-dynamodbsettings.html
    /// </summary>
    public sealed class EndpointDynamoDbSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-dynamodbsettings.html#cfn-dms-endpoint-dynamodbsettings-serviceaccessrolearn
        /// </summary>
        [Input("serviceAccessRoleArn")]
        public Input<string>? ServiceAccessRoleArn { get; set; }

        public EndpointDynamoDbSettingsArgs()
        {
        }
    }
}
