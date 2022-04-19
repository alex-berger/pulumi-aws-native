// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFRegional
{
    /// <summary>
    /// Resource Type definition for AWS::WAFRegional::GeoMatchSet
    /// </summary>
    [Obsolete(@"GeoMatchSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:wafregional:GeoMatchSet")]
    public partial class GeoMatchSet : Pulumi.CustomResource
    {
        [Output("geoMatchConstraints")]
        public Output<ImmutableArray<Outputs.GeoMatchSetGeoMatchConstraint>> GeoMatchConstraints { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a GeoMatchSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GeoMatchSet(string name, GeoMatchSetArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:wafregional:GeoMatchSet", name, args ?? new GeoMatchSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GeoMatchSet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:wafregional:GeoMatchSet", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing GeoMatchSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GeoMatchSet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GeoMatchSet(name, id, options);
        }
    }

    public sealed class GeoMatchSetArgs : Pulumi.ResourceArgs
    {
        [Input("geoMatchConstraints")]
        private InputList<Inputs.GeoMatchSetGeoMatchConstraintArgs>? _geoMatchConstraints;
        public InputList<Inputs.GeoMatchSetGeoMatchConstraintArgs> GeoMatchConstraints
        {
            get => _geoMatchConstraints ?? (_geoMatchConstraints = new InputList<Inputs.GeoMatchSetGeoMatchConstraintArgs>());
            set => _geoMatchConstraints = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GeoMatchSetArgs()
        {
        }
    }
}
