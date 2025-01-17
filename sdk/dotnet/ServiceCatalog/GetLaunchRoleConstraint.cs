// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceCatalog
{
    public static class GetLaunchRoleConstraint
    {
        /// <summary>
        /// Resource Type definition for AWS::ServiceCatalog::LaunchRoleConstraint
        /// </summary>
        public static Task<GetLaunchRoleConstraintResult> InvokeAsync(GetLaunchRoleConstraintArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLaunchRoleConstraintResult>("aws-native:servicecatalog:getLaunchRoleConstraint", args ?? new GetLaunchRoleConstraintArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ServiceCatalog::LaunchRoleConstraint
        /// </summary>
        public static Output<GetLaunchRoleConstraintResult> Invoke(GetLaunchRoleConstraintInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLaunchRoleConstraintResult>("aws-native:servicecatalog:getLaunchRoleConstraint", args ?? new GetLaunchRoleConstraintInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLaunchRoleConstraintArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetLaunchRoleConstraintArgs()
        {
        }
    }

    public sealed class GetLaunchRoleConstraintInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetLaunchRoleConstraintInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLaunchRoleConstraintResult
    {
        public readonly string? AcceptLanguage;
        public readonly string? Description;
        public readonly string? Id;
        public readonly string? LocalRoleName;
        public readonly string? RoleArn;

        [OutputConstructor]
        private GetLaunchRoleConstraintResult(
            string? acceptLanguage,

            string? description,

            string? id,

            string? localRoleName,

            string? roleArn)
        {
            AcceptLanguage = acceptLanguage;
            Description = description;
            Id = id;
            LocalRoleName = localRoleName;
            RoleArn = roleArn;
        }
    }
}
