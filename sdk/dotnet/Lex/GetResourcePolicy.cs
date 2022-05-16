// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex
{
    public static class GetResourcePolicy
    {
        /// <summary>
        /// A resource policy with specified policy statements that attaches to a Lex bot or bot alias.
        /// </summary>
        public static Task<GetResourcePolicyResult> InvokeAsync(GetResourcePolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourcePolicyResult>("aws-native:lex:getResourcePolicy", args ?? new GetResourcePolicyArgs(), options.WithDefaults());

        /// <summary>
        /// A resource policy with specified policy statements that attaches to a Lex bot or bot alias.
        /// </summary>
        public static Output<GetResourcePolicyResult> Invoke(GetResourcePolicyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetResourcePolicyResult>("aws-native:lex:getResourcePolicy", args ?? new GetResourcePolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourcePolicyArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetResourcePolicyArgs()
        {
        }
    }

    public sealed class GetResourcePolicyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetResourcePolicyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResourcePolicyResult
    {
        public readonly string? Id;
        public readonly Outputs.ResourcePolicyPolicy? Policy;
        public readonly string? ResourceArn;
        public readonly string? RevisionId;

        [OutputConstructor]
        private GetResourcePolicyResult(
            string? id,

            Outputs.ResourcePolicyPolicy? policy,

            string? resourceArn,

            string? revisionId)
        {
            Id = id;
            Policy = policy;
            ResourceArn = resourceArn;
            RevisionId = revisionId;
        }
    }
}
