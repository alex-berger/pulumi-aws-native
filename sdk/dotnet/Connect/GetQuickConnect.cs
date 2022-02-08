// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect
{
    public static class GetQuickConnect
    {
        /// <summary>
        /// Resource Type definition for AWS::Connect::QuickConnect
        /// </summary>
        public static Task<GetQuickConnectResult> InvokeAsync(GetQuickConnectArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetQuickConnectResult>("aws-native:connect:getQuickConnect", args ?? new GetQuickConnectArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Connect::QuickConnect
        /// </summary>
        public static Output<GetQuickConnectResult> Invoke(GetQuickConnectInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetQuickConnectResult>("aws-native:connect:getQuickConnect", args ?? new GetQuickConnectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetQuickConnectArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the quick connect.
        /// </summary>
        [Input("quickConnectArn", required: true)]
        public string QuickConnectArn { get; set; } = null!;

        public GetQuickConnectArgs()
        {
        }
    }

    public sealed class GetQuickConnectInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the quick connect.
        /// </summary>
        [Input("quickConnectArn", required: true)]
        public Input<string> QuickConnectArn { get; set; } = null!;

        public GetQuickConnectInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetQuickConnectResult
    {
        /// <summary>
        /// The description of the quick connect.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The identifier of the Amazon Connect instance.
        /// </summary>
        public readonly string? InstanceArn;
        /// <summary>
        /// The name of the quick connect.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The Amazon Resource Name (ARN) for the quick connect.
        /// </summary>
        public readonly string? QuickConnectArn;
        /// <summary>
        /// Configuration settings for the quick connect.
        /// </summary>
        public readonly Outputs.QuickConnectConfig? QuickConnectConfig;
        /// <summary>
        /// One or more tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.QuickConnectTag> Tags;

        [OutputConstructor]
        private GetQuickConnectResult(
            string? description,

            string? instanceArn,

            string? name,

            string? quickConnectArn,

            Outputs.QuickConnectConfig? quickConnectConfig,

            ImmutableArray<Outputs.QuickConnectTag> tags)
        {
            Description = description;
            InstanceArn = instanceArn;
            Name = name;
            QuickConnectArn = quickConnectArn;
            QuickConnectConfig = quickConnectConfig;
            Tags = tags;
        }
    }
}