// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElastiCache
{
    public static class GetSecurityGroup
    {
        /// <summary>
        /// Resource Type definition for AWS::ElastiCache::SecurityGroup
        /// </summary>
        public static Task<GetSecurityGroupResult> InvokeAsync(GetSecurityGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecurityGroupResult>("aws-native:elasticache:getSecurityGroup", args ?? new GetSecurityGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ElastiCache::SecurityGroup
        /// </summary>
        public static Output<GetSecurityGroupResult> Invoke(GetSecurityGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSecurityGroupResult>("aws-native:elasticache:getSecurityGroup", args ?? new GetSecurityGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityGroupArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetSecurityGroupArgs()
        {
        }
    }

    public sealed class GetSecurityGroupInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetSecurityGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSecurityGroupResult
    {
        public readonly string? Description;
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.SecurityGroupTag> Tags;

        [OutputConstructor]
        private GetSecurityGroupResult(
            string? description,

            string? id,

            ImmutableArray<Outputs.SecurityGroupTag> tags)
        {
            Description = description;
            Id = id;
            Tags = tags;
        }
    }
}