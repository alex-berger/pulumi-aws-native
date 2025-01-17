// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync
{
    public static class GetLocationObjectStorage
    {
        /// <summary>
        /// Resource schema for AWS::DataSync::LocationObjectStorage.
        /// </summary>
        public static Task<GetLocationObjectStorageResult> InvokeAsync(GetLocationObjectStorageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLocationObjectStorageResult>("aws-native:datasync:getLocationObjectStorage", args ?? new GetLocationObjectStorageArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::DataSync::LocationObjectStorage.
        /// </summary>
        public static Output<GetLocationObjectStorageResult> Invoke(GetLocationObjectStorageInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLocationObjectStorageResult>("aws-native:datasync:getLocationObjectStorage", args ?? new GetLocationObjectStorageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLocationObjectStorageArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the location that is created.
        /// </summary>
        [Input("locationArn", required: true)]
        public string LocationArn { get; set; } = null!;

        public GetLocationObjectStorageArgs()
        {
        }
    }

    public sealed class GetLocationObjectStorageInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the location that is created.
        /// </summary>
        [Input("locationArn", required: true)]
        public Input<string> LocationArn { get; set; } = null!;

        public GetLocationObjectStorageInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLocationObjectStorageResult
    {
        /// <summary>
        /// Optional. The access key is used if credentials are required to access the self-managed object storage server.
        /// </summary>
        public readonly string? AccessKey;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.
        /// </summary>
        public readonly ImmutableArray<string> AgentArns;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the location that is created.
        /// </summary>
        public readonly string? LocationArn;
        /// <summary>
        /// The URL of the object storage location that was described.
        /// </summary>
        public readonly string? LocationUri;
        /// <summary>
        /// The port that your self-managed server accepts inbound network traffic on.
        /// </summary>
        public readonly int? ServerPort;
        /// <summary>
        /// The protocol that the object storage server uses to communicate.
        /// </summary>
        public readonly Pulumi.AwsNative.DataSync.LocationObjectStorageServerProtocol? ServerProtocol;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.LocationObjectStorageTag> Tags;

        [OutputConstructor]
        private GetLocationObjectStorageResult(
            string? accessKey,

            ImmutableArray<string> agentArns,

            string? locationArn,

            string? locationUri,

            int? serverPort,

            Pulumi.AwsNative.DataSync.LocationObjectStorageServerProtocol? serverProtocol,

            ImmutableArray<Outputs.LocationObjectStorageTag> tags)
        {
            AccessKey = accessKey;
            AgentArns = agentArns;
            LocationArn = locationArn;
            LocationUri = locationUri;
            ServerPort = serverPort;
            ServerProtocol = serverProtocol;
            Tags = tags;
        }
    }
}
