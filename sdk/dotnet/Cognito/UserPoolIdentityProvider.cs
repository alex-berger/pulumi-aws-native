// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito
{
    /// <summary>
    /// Resource Type definition for AWS::Cognito::UserPoolIdentityProvider
    /// </summary>
    [Obsolete(@"UserPoolIdentityProvider is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:cognito:UserPoolIdentityProvider")]
    public partial class UserPoolIdentityProvider : Pulumi.CustomResource
    {
        [Output("attributeMapping")]
        public Output<object?> AttributeMapping { get; private set; } = null!;

        [Output("idpIdentifiers")]
        public Output<ImmutableArray<string>> IdpIdentifiers { get; private set; } = null!;

        [Output("providerDetails")]
        public Output<object?> ProviderDetails { get; private set; } = null!;

        [Output("providerName")]
        public Output<string> ProviderName { get; private set; } = null!;

        [Output("providerType")]
        public Output<string> ProviderType { get; private set; } = null!;

        [Output("userPoolId")]
        public Output<string> UserPoolId { get; private set; } = null!;


        /// <summary>
        /// Create a UserPoolIdentityProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserPoolIdentityProvider(string name, UserPoolIdentityProviderArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cognito:UserPoolIdentityProvider", name, args ?? new UserPoolIdentityProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserPoolIdentityProvider(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cognito:UserPoolIdentityProvider", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing UserPoolIdentityProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserPoolIdentityProvider Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new UserPoolIdentityProvider(name, id, options);
        }
    }

    public sealed class UserPoolIdentityProviderArgs : Pulumi.ResourceArgs
    {
        [Input("attributeMapping")]
        public Input<object>? AttributeMapping { get; set; }

        [Input("idpIdentifiers")]
        private InputList<string>? _idpIdentifiers;
        public InputList<string> IdpIdentifiers
        {
            get => _idpIdentifiers ?? (_idpIdentifiers = new InputList<string>());
            set => _idpIdentifiers = value;
        }

        [Input("providerDetails")]
        public Input<object>? ProviderDetails { get; set; }

        [Input("providerName", required: true)]
        public Input<string> ProviderName { get; set; } = null!;

        [Input("providerType", required: true)]
        public Input<string> ProviderType { get; set; } = null!;

        [Input("userPoolId", required: true)]
        public Input<string> UserPoolId { get; set; } = null!;

        public UserPoolIdentityProviderArgs()
        {
        }
    }
}
