// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Location
{
    public static class GetGeofenceCollection
    {
        /// <summary>
        /// Definition of AWS::Location::GeofenceCollection Resource Type
        /// </summary>
        public static Task<GetGeofenceCollectionResult> InvokeAsync(GetGeofenceCollectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGeofenceCollectionResult>("aws-native:location:getGeofenceCollection", args ?? new GetGeofenceCollectionArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Location::GeofenceCollection Resource Type
        /// </summary>
        public static Output<GetGeofenceCollectionResult> Invoke(GetGeofenceCollectionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGeofenceCollectionResult>("aws-native:location:getGeofenceCollection", args ?? new GetGeofenceCollectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGeofenceCollectionArgs : Pulumi.InvokeArgs
    {
        [Input("collectionName", required: true)]
        public string CollectionName { get; set; } = null!;

        public GetGeofenceCollectionArgs()
        {
        }
    }

    public sealed class GetGeofenceCollectionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("collectionName", required: true)]
        public Input<string> CollectionName { get; set; } = null!;

        public GetGeofenceCollectionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGeofenceCollectionResult
    {
        public readonly string? Arn;
        public readonly string? CollectionArn;
        public readonly string? CreateTime;
        public readonly string? KmsKeyId;
        public readonly string? UpdateTime;

        [OutputConstructor]
        private GetGeofenceCollectionResult(
            string? arn,

            string? collectionArn,

            string? createTime,

            string? kmsKeyId,

            string? updateTime)
        {
            Arn = arn;
            CollectionArn = collectionArn;
            CreateTime = createTime;
            KmsKeyId = kmsKeyId;
            UpdateTime = updateTime;
        }
    }
}