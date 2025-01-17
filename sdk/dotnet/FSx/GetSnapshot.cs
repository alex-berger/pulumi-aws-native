// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FSx
{
    public static class GetSnapshot
    {
        /// <summary>
        /// Resource Type definition for AWS::FSx::Snapshot
        /// </summary>
        public static Task<GetSnapshotResult> InvokeAsync(GetSnapshotArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotResult>("aws-native:fsx:getSnapshot", args ?? new GetSnapshotArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::FSx::Snapshot
        /// </summary>
        public static Output<GetSnapshotResult> Invoke(GetSnapshotInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSnapshotResult>("aws-native:fsx:getSnapshot", args ?? new GetSnapshotInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSnapshotArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetSnapshotArgs()
        {
        }
    }

    public sealed class GetSnapshotInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetSnapshotInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSnapshotResult
    {
        public readonly string? Id;
        public readonly string? Name;
        public readonly string? ResourceARN;
        public readonly ImmutableArray<Outputs.SnapshotTag> Tags;

        [OutputConstructor]
        private GetSnapshotResult(
            string? id,

            string? name,

            string? resourceARN,

            ImmutableArray<Outputs.SnapshotTag> tags)
        {
            Id = id;
            Name = name;
            ResourceARN = resourceARN;
            Tags = tags;
        }
    }
}
