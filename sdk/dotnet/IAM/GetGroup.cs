// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IAM
{
    public static class GetGroup
    {
        /// <summary>
        /// Resource Type definition for AWS::IAM::Group
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("aws-native:iam:getGroup", args ?? new GetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IAM::Group
        /// </summary>
        public static Output<GetGroupResult> Invoke(GetGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGroupResult>("aws-native:iam:getGroup", args ?? new GetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetGroupArgs()
        {
        }
    }

    public sealed class GetGroupInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        public readonly string? Arn;
        public readonly string? Id;
        public readonly ImmutableArray<string> ManagedPolicyArns;
        public readonly string? Path;
        public readonly ImmutableArray<Outputs.GroupPolicy> Policies;

        [OutputConstructor]
        private GetGroupResult(
            string? arn,

            string? id,

            ImmutableArray<string> managedPolicyArns,

            string? path,

            ImmutableArray<Outputs.GroupPolicy> policies)
        {
            Arn = arn;
            Id = id;
            ManagedPolicyArns = managedPolicyArns;
            Path = path;
            Policies = policies;
        }
    }
}
