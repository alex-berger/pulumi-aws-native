// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DevOpsGuru
{
    public static class GetNotificationChannel
    {
        /// <summary>
        /// This resource schema represents the NotificationChannel resource in the Amazon DevOps Guru.
        /// </summary>
        public static Task<GetNotificationChannelResult> InvokeAsync(GetNotificationChannelArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNotificationChannelResult>("aws-native:devopsguru:getNotificationChannel", args ?? new GetNotificationChannelArgs(), options.WithDefaults());

        /// <summary>
        /// This resource schema represents the NotificationChannel resource in the Amazon DevOps Guru.
        /// </summary>
        public static Output<GetNotificationChannelResult> Invoke(GetNotificationChannelInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNotificationChannelResult>("aws-native:devopsguru:getNotificationChannel", args ?? new GetNotificationChannelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNotificationChannelArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of a notification channel.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetNotificationChannelArgs()
        {
        }
    }

    public sealed class GetNotificationChannelInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of a notification channel.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetNotificationChannelInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNotificationChannelResult
    {
        /// <summary>
        /// The ID of a notification channel.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private GetNotificationChannelResult(string? id)
        {
            Id = id;
        }
    }
}
