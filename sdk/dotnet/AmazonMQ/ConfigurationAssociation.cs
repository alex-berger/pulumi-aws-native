// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmazonMQ
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configurationassociation.html
    /// </summary>
    [AwsNativeResourceType("aws-native:amazonmq:ConfigurationAssociation")]
    public partial class ConfigurationAssociation : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configurationassociation.html#cfn-amazonmq-configurationassociation-broker
        /// </summary>
        [Output("broker")]
        public Output<string> Broker { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configurationassociation.html#cfn-amazonmq-configurationassociation-configuration
        /// </summary>
        [Output("configuration")]
        public Output<Outputs.ConfigurationAssociationConfigurationId> Configuration { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigurationAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigurationAssociation(string name, ConfigurationAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:amazonmq:ConfigurationAssociation", name, args ?? new ConfigurationAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigurationAssociation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:amazonmq:ConfigurationAssociation", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigurationAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigurationAssociation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConfigurationAssociation(name, id, options);
        }
    }

    public sealed class ConfigurationAssociationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configurationassociation.html#cfn-amazonmq-configurationassociation-broker
        /// </summary>
        [Input("broker", required: true)]
        public Input<string> Broker { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configurationassociation.html#cfn-amazonmq-configurationassociation-configuration
        /// </summary>
        [Input("configuration", required: true)]
        public Input<Inputs.ConfigurationAssociationConfigurationIdArgs> Configuration { get; set; } = null!;

        public ConfigurationAssociationArgs()
        {
        }
    }
}
