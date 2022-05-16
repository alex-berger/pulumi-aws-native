// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceCatalog
{
    /// <summary>
    /// Resource Type definition for AWS::ServiceCatalog::LaunchTemplateConstraint
    /// </summary>
    [Obsolete(@"LaunchTemplateConstraint is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:servicecatalog:LaunchTemplateConstraint")]
    public partial class LaunchTemplateConstraint : Pulumi.CustomResource
    {
        [Output("acceptLanguage")]
        public Output<string?> AcceptLanguage { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("portfolioId")]
        public Output<string> PortfolioId { get; private set; } = null!;

        [Output("productId")]
        public Output<string> ProductId { get; private set; } = null!;

        [Output("rules")]
        public Output<string> Rules { get; private set; } = null!;


        /// <summary>
        /// Create a LaunchTemplateConstraint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LaunchTemplateConstraint(string name, LaunchTemplateConstraintArgs args, CustomResourceOptions? options = null)
            : base("aws-native:servicecatalog:LaunchTemplateConstraint", name, args ?? new LaunchTemplateConstraintArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LaunchTemplateConstraint(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:servicecatalog:LaunchTemplateConstraint", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing LaunchTemplateConstraint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LaunchTemplateConstraint Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LaunchTemplateConstraint(name, id, options);
        }
    }

    public sealed class LaunchTemplateConstraintArgs : Pulumi.ResourceArgs
    {
        [Input("acceptLanguage")]
        public Input<string>? AcceptLanguage { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("portfolioId", required: true)]
        public Input<string> PortfolioId { get; set; } = null!;

        [Input("productId", required: true)]
        public Input<string> ProductId { get; set; } = null!;

        [Input("rules", required: true)]
        public Input<string> Rules { get; set; } = null!;

        public LaunchTemplateConstraintArgs()
        {
        }
    }
}