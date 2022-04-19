// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IAM
{
    public static class GetRole
    {
        /// <summary>
        /// Resource Type definition for AWS::IAM::Role
        /// </summary>
        public static Task<GetRoleResult> InvokeAsync(GetRoleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRoleResult>("aws-native:iam:getRole", args ?? new GetRoleArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IAM::Role
        /// </summary>
        public static Output<GetRoleResult> Invoke(GetRoleInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRoleResult>("aws-native:iam:getRole", args ?? new GetRoleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRoleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A name for the IAM role, up to 64 characters in length.
        /// </summary>
        [Input("roleName", required: true)]
        public string RoleName { get; set; } = null!;

        public GetRoleArgs()
        {
        }
    }

    public sealed class GetRoleInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A name for the IAM role, up to 64 characters in length.
        /// </summary>
        [Input("roleName", required: true)]
        public Input<string> RoleName { get; set; } = null!;

        public GetRoleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRoleResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the role.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The trust policy that is associated with this role.
        /// </summary>
        public readonly object? AssumeRolePolicyDocument;
        /// <summary>
        /// A description of the role that you provide.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role. 
        /// </summary>
        public readonly ImmutableArray<string> ManagedPolicyArns;
        /// <summary>
        /// The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours. 
        /// </summary>
        public readonly int? MaxSessionDuration;
        /// <summary>
        /// The ARN of the policy used to set the permissions boundary for the role.
        /// </summary>
        public readonly string? PermissionsBoundary;
        /// <summary>
        /// Adds or updates an inline policy document that is embedded in the specified IAM role. 
        /// </summary>
        public readonly ImmutableArray<Outputs.RolePolicy> Policies;
        /// <summary>
        /// The stable and unique string identifying the role.
        /// </summary>
        public readonly string? RoleId;
        /// <summary>
        /// A list of tags that are attached to the role.
        /// </summary>
        public readonly ImmutableArray<Outputs.RoleTag> Tags;

        [OutputConstructor]
        private GetRoleResult(
            string? arn,

            object? assumeRolePolicyDocument,

            string? description,

            ImmutableArray<string> managedPolicyArns,

            int? maxSessionDuration,

            string? permissionsBoundary,

            ImmutableArray<Outputs.RolePolicy> policies,

            string? roleId,

            ImmutableArray<Outputs.RoleTag> tags)
        {
            Arn = arn;
            AssumeRolePolicyDocument = assumeRolePolicyDocument;
            Description = description;
            ManagedPolicyArns = managedPolicyArns;
            MaxSessionDuration = maxSessionDuration;
            PermissionsBoundary = permissionsBoundary;
            Policies = policies;
            RoleId = roleId;
            Tags = tags;
        }
    }
}
