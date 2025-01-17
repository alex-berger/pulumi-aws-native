// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Redshift
{
    /// <summary>
    /// Resource Type definition for AWS::Redshift::ClusterParameterGroup
    /// </summary>
    [Obsolete(@"ClusterParameterGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:redshift:ClusterParameterGroup")]
    public partial class ClusterParameterGroup : Pulumi.CustomResource
    {
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("parameterGroupFamily")]
        public Output<string> ParameterGroupFamily { get; private set; } = null!;

        [Output("parameters")]
        public Output<ImmutableArray<Outputs.ClusterParameterGroupParameter>> Parameters { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.ClusterParameterGroupTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterParameterGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterParameterGroup(string name, ClusterParameterGroupArgs args, CustomResourceOptions? options = null)
            : base("aws-native:redshift:ClusterParameterGroup", name, args ?? new ClusterParameterGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterParameterGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:redshift:ClusterParameterGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ClusterParameterGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterParameterGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ClusterParameterGroup(name, id, options);
        }
    }

    public sealed class ClusterParameterGroupArgs : Pulumi.ResourceArgs
    {
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("parameterGroupFamily", required: true)]
        public Input<string> ParameterGroupFamily { get; set; } = null!;

        [Input("parameters")]
        private InputList<Inputs.ClusterParameterGroupParameterArgs>? _parameters;
        public InputList<Inputs.ClusterParameterGroupParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ClusterParameterGroupParameterArgs>());
            set => _parameters = value;
        }

        [Input("tags")]
        private InputList<Inputs.ClusterParameterGroupTagArgs>? _tags;
        public InputList<Inputs.ClusterParameterGroupTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ClusterParameterGroupTagArgs>());
            set => _tags = value;
        }

        public ClusterParameterGroupArgs()
        {
        }
    }
}
