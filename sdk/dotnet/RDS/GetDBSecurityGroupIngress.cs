// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RDS
{
    public static class GetDBSecurityGroupIngress
    {
        /// <summary>
        /// Resource Type definition for AWS::RDS::DBSecurityGroupIngress
        /// </summary>
        public static Task<GetDBSecurityGroupIngressResult> InvokeAsync(GetDBSecurityGroupIngressArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDBSecurityGroupIngressResult>("aws-native:rds:getDBSecurityGroupIngress", args ?? new GetDBSecurityGroupIngressArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::RDS::DBSecurityGroupIngress
        /// </summary>
        public static Output<GetDBSecurityGroupIngressResult> Invoke(GetDBSecurityGroupIngressInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDBSecurityGroupIngressResult>("aws-native:rds:getDBSecurityGroupIngress", args ?? new GetDBSecurityGroupIngressInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDBSecurityGroupIngressArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDBSecurityGroupIngressArgs()
        {
        }
    }

    public sealed class GetDBSecurityGroupIngressInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDBSecurityGroupIngressInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDBSecurityGroupIngressResult
    {
        public readonly string? CIDRIP;
        public readonly string? DBSecurityGroupName;
        public readonly string? EC2SecurityGroupId;
        public readonly string? EC2SecurityGroupName;
        public readonly string? EC2SecurityGroupOwnerId;
        public readonly string? Id;

        [OutputConstructor]
        private GetDBSecurityGroupIngressResult(
            string? cIDRIP,

            string? dBSecurityGroupName,

            string? eC2SecurityGroupId,

            string? eC2SecurityGroupName,

            string? eC2SecurityGroupOwnerId,

            string? id)
        {
            CIDRIP = cIDRIP;
            DBSecurityGroupName = dBSecurityGroupName;
            EC2SecurityGroupId = eC2SecurityGroupId;
            EC2SecurityGroupName = eC2SecurityGroupName;
            EC2SecurityGroupOwnerId = eC2SecurityGroupOwnerId;
            Id = id;
        }
    }
}
