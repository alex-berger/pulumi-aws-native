// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFormation
{
    public static class GetStack
    {
        /// <summary>
        /// Resource Type definition for AWS::CloudFormation::Stack
        /// </summary>
        public static Task<GetStackResult> InvokeAsync(GetStackArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStackResult>("aws-native:cloudformation:getStack", args ?? new GetStackArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::CloudFormation::Stack
        /// </summary>
        public static Output<GetStackResult> Invoke(GetStackInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetStackResult>("aws-native:cloudformation:getStack", args ?? new GetStackInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStackArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetStackArgs()
        {
        }
    }

    public sealed class GetStackInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetStackInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStackResult
    {
        public readonly string? Id;
        public readonly ImmutableArray<string> NotificationARNs;
        public readonly object? Parameters;
        public readonly ImmutableArray<Outputs.StackTag> Tags;
        public readonly string? TemplateURL;
        public readonly int? TimeoutInMinutes;

        [OutputConstructor]
        private GetStackResult(
            string? id,

            ImmutableArray<string> notificationARNs,

            object? parameters,

            ImmutableArray<Outputs.StackTag> tags,

            string? templateURL,

            int? timeoutInMinutes)
        {
            Id = id;
            NotificationARNs = notificationARNs;
            Parameters = parameters;
            Tags = tags;
            TemplateURL = templateURL;
            TimeoutInMinutes = timeoutInMinutes;
        }
    }
}
