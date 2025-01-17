// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElastiCache
{
    public static class GetUser
    {
        /// <summary>
        /// Resource Type definition for AWS::ElastiCache::User
        /// </summary>
        public static Task<GetUserResult> InvokeAsync(GetUserArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("aws-native:elasticache:getUser", args ?? new GetUserArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ElastiCache::User
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetUserResult>("aws-native:elasticache:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the user.
        /// </summary>
        [Input("userId", required: true)]
        public string UserId { get; set; } = null!;

        public GetUserArgs()
        {
        }
    }

    public sealed class GetUserInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the user.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public GetUserInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the user account.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Indicates the user status. Can be "active", "modifying" or "deleting".
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private GetUserResult(
            string? arn,

            string? status)
        {
            Arn = arn;
            Status = status;
        }
    }
}
