// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DMS.Inputs
{

    public sealed class EndpointRedshiftSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("acceptAnyDate")]
        public Input<bool>? AcceptAnyDate { get; set; }

        [Input("afterConnectScript")]
        public Input<string>? AfterConnectScript { get; set; }

        [Input("bucketFolder")]
        public Input<string>? BucketFolder { get; set; }

        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        [Input("caseSensitiveNames")]
        public Input<bool>? CaseSensitiveNames { get; set; }

        [Input("compUpdate")]
        public Input<bool>? CompUpdate { get; set; }

        [Input("connectionTimeout")]
        public Input<int>? ConnectionTimeout { get; set; }

        [Input("dateFormat")]
        public Input<string>? DateFormat { get; set; }

        [Input("emptyAsNull")]
        public Input<bool>? EmptyAsNull { get; set; }

        [Input("encryptionMode")]
        public Input<string>? EncryptionMode { get; set; }

        [Input("explicitIds")]
        public Input<bool>? ExplicitIds { get; set; }

        [Input("fileTransferUploadStreams")]
        public Input<int>? FileTransferUploadStreams { get; set; }

        [Input("loadTimeout")]
        public Input<int>? LoadTimeout { get; set; }

        [Input("maxFileSize")]
        public Input<int>? MaxFileSize { get; set; }

        [Input("removeQuotes")]
        public Input<bool>? RemoveQuotes { get; set; }

        [Input("replaceChars")]
        public Input<string>? ReplaceChars { get; set; }

        [Input("replaceInvalidChars")]
        public Input<string>? ReplaceInvalidChars { get; set; }

        [Input("secretsManagerAccessRoleArn")]
        public Input<string>? SecretsManagerAccessRoleArn { get; set; }

        [Input("secretsManagerSecretId")]
        public Input<string>? SecretsManagerSecretId { get; set; }

        [Input("serverSideEncryptionKmsKeyId")]
        public Input<string>? ServerSideEncryptionKmsKeyId { get; set; }

        [Input("serviceAccessRoleArn")]
        public Input<string>? ServiceAccessRoleArn { get; set; }

        [Input("timeFormat")]
        public Input<string>? TimeFormat { get; set; }

        [Input("trimBlanks")]
        public Input<bool>? TrimBlanks { get; set; }

        [Input("truncateColumns")]
        public Input<bool>? TruncateColumns { get; set; }

        [Input("writeBufferSize")]
        public Input<int>? WriteBufferSize { get; set; }

        public EndpointRedshiftSettingsArgs()
        {
        }
    }
}
