// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LicenseManager
{
    /// <summary>
    /// An example resource schema demonstrating some basic constructs and validation rules.
    /// </summary>
    [AwsNativeResourceType("aws-native:licensemanager:Grant")]
    public partial class Grant : Pulumi.CustomResource
    {
        [Output("allowedOperations")]
        public Output<ImmutableArray<string>> AllowedOperations { get; private set; } = null!;

        /// <summary>
        /// Arn of the grant.
        /// </summary>
        [Output("grantArn")]
        public Output<string> GrantArn { get; private set; } = null!;

        /// <summary>
        /// Name for the created Grant.
        /// </summary>
        [Output("grantName")]
        public Output<string?> GrantName { get; private set; } = null!;

        /// <summary>
        /// Home region for the created grant.
        /// </summary>
        [Output("homeRegion")]
        public Output<string?> HomeRegion { get; private set; } = null!;

        /// <summary>
        /// License Arn for the grant.
        /// </summary>
        [Output("licenseArn")]
        public Output<string?> LicenseArn { get; private set; } = null!;

        [Output("principals")]
        public Output<ImmutableArray<string>> Principals { get; private set; } = null!;

        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// The version of the grant.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Grant resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Grant(string name, GrantArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:licensemanager:Grant", name, args ?? new GrantArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Grant(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:licensemanager:Grant", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Grant resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Grant Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Grant(name, id, options);
        }
    }

    public sealed class GrantArgs : Pulumi.ResourceArgs
    {
        [Input("allowedOperations")]
        private InputList<string>? _allowedOperations;
        public InputList<string> AllowedOperations
        {
            get => _allowedOperations ?? (_allowedOperations = new InputList<string>());
            set => _allowedOperations = value;
        }

        /// <summary>
        /// Name for the created Grant.
        /// </summary>
        [Input("grantName")]
        public Input<string>? GrantName { get; set; }

        /// <summary>
        /// Home region for the created grant.
        /// </summary>
        [Input("homeRegion")]
        public Input<string>? HomeRegion { get; set; }

        /// <summary>
        /// License Arn for the grant.
        /// </summary>
        [Input("licenseArn")]
        public Input<string>? LicenseArn { get; set; }

        [Input("principals")]
        private InputList<string>? _principals;
        public InputList<string> Principals
        {
            get => _principals ?? (_principals = new InputList<string>());
            set => _principals = value;
        }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public GrantArgs()
        {
        }
    }
}
