// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.StepFunctions
{
    public static class GetActivity
    {
        /// <summary>
        /// Resource schema for Activity
        /// </summary>
        public static Task<GetActivityResult> InvokeAsync(GetActivityArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetActivityResult>("aws-native:stepfunctions:getActivity", args ?? new GetActivityArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for Activity
        /// </summary>
        public static Output<GetActivityResult> Invoke(GetActivityInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetActivityResult>("aws-native:stepfunctions:getActivity", args ?? new GetActivityInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetActivityArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetActivityArgs()
        {
        }
    }

    public sealed class GetActivityInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetActivityInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetActivityResult
    {
        public readonly string? Arn;
        public readonly ImmutableArray<Outputs.ActivityTagsEntry> Tags;

        [OutputConstructor]
        private GetActivityResult(
            string? arn,

            ImmutableArray<Outputs.ActivityTagsEntry> tags)
        {
            Arn = arn;
            Tags = tags;
        }
    }
}
