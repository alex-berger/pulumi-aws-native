// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IAM
{
    public static class GetPolicy
    {
        /// <summary>
        /// Resource Type definition for AWS::IAM::Policy
        /// </summary>
        public static Task<GetPolicyResult> InvokeAsync(GetPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyResult>("aws-native:iam:getPolicy", args ?? new GetPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IAM::Policy
        /// </summary>
        public static Output<GetPolicyResult> Invoke(GetPolicyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPolicyResult>("aws-native:iam:getPolicy", args ?? new GetPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPolicyArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetPolicyArgs()
        {
        }
    }

    public sealed class GetPolicyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetPolicyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPolicyResult
    {
        public readonly ImmutableArray<string> Groups;
        public readonly string? Id;
        public readonly object? PolicyDocument;
        public readonly string? PolicyName;
        public readonly ImmutableArray<string> Roles;
        public readonly ImmutableArray<string> Users;

        [OutputConstructor]
        private GetPolicyResult(
            ImmutableArray<string> groups,

            string? id,

            object? policyDocument,

            string? policyName,

            ImmutableArray<string> roles,

            ImmutableArray<string> users)
        {
            Groups = groups;
            Id = id;
            PolicyDocument = policyDocument;
            PolicyName = policyName;
            Roles = roles;
            Users = users;
        }
    }
}
