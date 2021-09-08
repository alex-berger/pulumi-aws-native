// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElastiCache
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html
    /// </summary>
    [AwsNativeResourceType("aws-native:elasticache:SecurityGroupIngress")]
    public partial class SecurityGroupIngress : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-cachesecuritygroupname
        /// </summary>
        [Output("cacheSecurityGroupName")]
        public Output<string> CacheSecurityGroupName { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-ec2securitygroupname
        /// </summary>
        [Output("eC2SecurityGroupName")]
        public Output<string> EC2SecurityGroupName { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-ec2securitygroupownerid
        /// </summary>
        [Output("eC2SecurityGroupOwnerId")]
        public Output<string?> EC2SecurityGroupOwnerId { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityGroupIngress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityGroupIngress(string name, SecurityGroupIngressArgs args, CustomResourceOptions? options = null)
            : base("aws-native:elasticache:SecurityGroupIngress", name, args ?? new SecurityGroupIngressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityGroupIngress(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:elasticache:SecurityGroupIngress", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecurityGroupIngress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityGroupIngress Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SecurityGroupIngress(name, id, options);
        }
    }

    public sealed class SecurityGroupIngressArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-cachesecuritygroupname
        /// </summary>
        [Input("cacheSecurityGroupName", required: true)]
        public Input<string> CacheSecurityGroupName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-ec2securitygroupname
        /// </summary>
        [Input("eC2SecurityGroupName", required: true)]
        public Input<string> EC2SecurityGroupName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-ec2securitygroupownerid
        /// </summary>
        [Input("eC2SecurityGroupOwnerId")]
        public Input<string>? EC2SecurityGroupOwnerId { get; set; }

        public SecurityGroupIngressArgs()
        {
        }
    }
}
