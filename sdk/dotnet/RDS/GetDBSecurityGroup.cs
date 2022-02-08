// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RDS
{
    public static class GetDBSecurityGroup
    {
        /// <summary>
        /// Resource Type definition for AWS::RDS::DBSecurityGroup
        /// </summary>
        public static Task<GetDBSecurityGroupResult> InvokeAsync(GetDBSecurityGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDBSecurityGroupResult>("aws-native:rds:getDBSecurityGroup", args ?? new GetDBSecurityGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::RDS::DBSecurityGroup
        /// </summary>
        public static Output<GetDBSecurityGroupResult> Invoke(GetDBSecurityGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDBSecurityGroupResult>("aws-native:rds:getDBSecurityGroup", args ?? new GetDBSecurityGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDBSecurityGroupArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDBSecurityGroupArgs()
        {
        }
    }

    public sealed class GetDBSecurityGroupInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDBSecurityGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDBSecurityGroupResult
    {
        public readonly ImmutableArray<Outputs.DBSecurityGroupIngress> DBSecurityGroupIngress;
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.DBSecurityGroupTag> Tags;

        [OutputConstructor]
        private GetDBSecurityGroupResult(
            ImmutableArray<Outputs.DBSecurityGroupIngress> dBSecurityGroupIngress,

            string? id,

            ImmutableArray<Outputs.DBSecurityGroupTag> tags)
        {
            DBSecurityGroupIngress = dBSecurityGroupIngress;
            Id = id;
            Tags = tags;
        }
    }
}