// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint
{
    public static class GetEventStream
    {
        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::EventStream
        /// </summary>
        public static Task<GetEventStreamResult> InvokeAsync(GetEventStreamArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEventStreamResult>("aws-native:pinpoint:getEventStream", args ?? new GetEventStreamArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::EventStream
        /// </summary>
        public static Output<GetEventStreamResult> Invoke(GetEventStreamInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEventStreamResult>("aws-native:pinpoint:getEventStream", args ?? new GetEventStreamInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEventStreamArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetEventStreamArgs()
        {
        }
    }

    public sealed class GetEventStreamInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetEventStreamInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEventStreamResult
    {
        public readonly string? DestinationStreamArn;
        public readonly string? Id;
        public readonly string? RoleArn;

        [OutputConstructor]
        private GetEventStreamResult(
            string? destinationStreamArn,

            string? id,

            string? roleArn)
        {
            DestinationStreamArn = destinationStreamArn;
            Id = id;
            RoleArn = roleArn;
        }
    }
}
