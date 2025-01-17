// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DLM
{
    public static class GetLifecyclePolicy
    {
        /// <summary>
        /// Resource Type definition for AWS::DLM::LifecyclePolicy
        /// </summary>
        public static Task<GetLifecyclePolicyResult> InvokeAsync(GetLifecyclePolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLifecyclePolicyResult>("aws-native:dlm:getLifecyclePolicy", args ?? new GetLifecyclePolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::DLM::LifecyclePolicy
        /// </summary>
        public static Output<GetLifecyclePolicyResult> Invoke(GetLifecyclePolicyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLifecyclePolicyResult>("aws-native:dlm:getLifecyclePolicy", args ?? new GetLifecyclePolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLifecyclePolicyArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetLifecyclePolicyArgs()
        {
        }
    }

    public sealed class GetLifecyclePolicyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetLifecyclePolicyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLifecyclePolicyResult
    {
        public readonly string? Arn;
        public readonly string? Description;
        public readonly string? ExecutionRoleArn;
        public readonly string? Id;
        public readonly Outputs.LifecyclePolicyPolicyDetails? PolicyDetails;
        public readonly string? State;
        public readonly ImmutableArray<Outputs.LifecyclePolicyTag> Tags;

        [OutputConstructor]
        private GetLifecyclePolicyResult(
            string? arn,

            string? description,

            string? executionRoleArn,

            string? id,

            Outputs.LifecyclePolicyPolicyDetails? policyDetails,

            string? state,

            ImmutableArray<Outputs.LifecyclePolicyTag> tags)
        {
            Arn = arn;
            Description = description;
            ExecutionRoleArn = executionRoleArn;
            Id = id;
            PolicyDetails = policyDetails;
            State = state;
            Tags = tags;
        }
    }
}
