// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECS
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html
    /// </summary>
    [AwsNativeResourceType("aws-native:ecs:ClusterCapacityProviderAssociations")]
    public partial class ClusterCapacityProviderAssociations : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-capacityproviders
        /// </summary>
        [Output("capacityProviders")]
        public Output<ImmutableArray<string>> CapacityProviders { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-cluster
        /// </summary>
        [Output("cluster")]
        public Output<string> Cluster { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-defaultcapacityproviderstrategy
        /// </summary>
        [Output("defaultCapacityProviderStrategy")]
        public Output<ImmutableArray<Outputs.ClusterCapacityProviderAssociationsCapacityProviderStrategy>> DefaultCapacityProviderStrategy { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterCapacityProviderAssociations resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterCapacityProviderAssociations(string name, ClusterCapacityProviderAssociationsArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ecs:ClusterCapacityProviderAssociations", name, args ?? new ClusterCapacityProviderAssociationsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterCapacityProviderAssociations(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ecs:ClusterCapacityProviderAssociations", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ClusterCapacityProviderAssociations resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterCapacityProviderAssociations Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ClusterCapacityProviderAssociations(name, id, options);
        }
    }

    public sealed class ClusterCapacityProviderAssociationsArgs : Pulumi.ResourceArgs
    {
        [Input("capacityProviders", required: true)]
        private InputList<string>? _capacityProviders;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-capacityproviders
        /// </summary>
        public InputList<string> CapacityProviders
        {
            get => _capacityProviders ?? (_capacityProviders = new InputList<string>());
            set => _capacityProviders = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-cluster
        /// </summary>
        [Input("cluster", required: true)]
        public Input<string> Cluster { get; set; } = null!;

        [Input("defaultCapacityProviderStrategy", required: true)]
        private InputList<Inputs.ClusterCapacityProviderAssociationsCapacityProviderStrategyArgs>? _defaultCapacityProviderStrategy;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-clustercapacityproviderassociations.html#cfn-ecs-clustercapacityproviderassociations-defaultcapacityproviderstrategy
        /// </summary>
        public InputList<Inputs.ClusterCapacityProviderAssociationsCapacityProviderStrategyArgs> DefaultCapacityProviderStrategy
        {
            get => _defaultCapacityProviderStrategy ?? (_defaultCapacityProviderStrategy = new InputList<Inputs.ClusterCapacityProviderAssociationsCapacityProviderStrategyArgs>());
            set => _defaultCapacityProviderStrategy = value;
        }

        public ClusterCapacityProviderAssociationsArgs()
        {
        }
    }
}