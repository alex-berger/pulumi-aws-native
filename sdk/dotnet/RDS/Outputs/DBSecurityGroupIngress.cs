// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RDS.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html
    /// </summary>
    [OutputType]
    public sealed class DBSecurityGroupIngress
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-cidrip
        /// </summary>
        public readonly string? CIDRIP;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-ec2securitygroupid
        /// </summary>
        public readonly string? EC2SecurityGroupId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-ec2securitygroupname
        /// </summary>
        public readonly string? EC2SecurityGroupName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-ec2securitygroupownerid
        /// </summary>
        public readonly string? EC2SecurityGroupOwnerId;

        [OutputConstructor]
        private DBSecurityGroupIngress(
            string? cIDRIP,

            string? eC2SecurityGroupId,

            string? eC2SecurityGroupName,

            string? eC2SecurityGroupOwnerId)
        {
            CIDRIP = cIDRIP;
            EC2SecurityGroupId = eC2SecurityGroupId;
            EC2SecurityGroupName = eC2SecurityGroupName;
            EC2SecurityGroupOwnerId = eC2SecurityGroupOwnerId;
        }
    }
}
