// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QLDB
{
    public static class GetStream
    {
        /// <summary>
        /// Resource schema for AWS::QLDB::Stream.
        /// </summary>
        public static Task<GetStreamResult> InvokeAsync(GetStreamArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStreamResult>("aws-native:qldb:getStream", args ?? new GetStreamArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::QLDB::Stream.
        /// </summary>
        public static Output<GetStreamResult> Invoke(GetStreamInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetStreamResult>("aws-native:qldb:getStream", args ?? new GetStreamInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStreamArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("ledgerName", required: true)]
        public string LedgerName { get; set; } = null!;

        public GetStreamArgs()
        {
        }
    }

    public sealed class GetStreamInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("ledgerName", required: true)]
        public Input<string> LedgerName { get; set; } = null!;

        public GetStreamInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStreamResult
    {
        public readonly string? Arn;
        public readonly string? Id;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.StreamTag> Tags;

        [OutputConstructor]
        private GetStreamResult(
            string? arn,

            string? id,

            ImmutableArray<Outputs.StreamTag> tags)
        {
            Arn = arn;
            Id = id;
            Tags = tags;
        }
    }
}