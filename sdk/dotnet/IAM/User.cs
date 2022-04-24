// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IAM
{
    /// <summary>
    /// Resource Type definition for AWS::IAM::User
    /// </summary>
    [Obsolete(@"User is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:iam:User")]
    public partial class User : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("groups")]
        public Output<ImmutableArray<string>> Groups { get; private set; } = null!;

        [Output("loginProfile")]
        public Output<Outputs.UserLoginProfile?> LoginProfile { get; private set; } = null!;

        [Output("managedPolicyArns")]
        public Output<ImmutableArray<string>> ManagedPolicyArns { get; private set; } = null!;

        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        [Output("permissionsBoundary")]
        public Output<string?> PermissionsBoundary { get; private set; } = null!;

        [Output("policies")]
        public Output<ImmutableArray<Outputs.UserPolicy>> Policies { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.UserTag>> Tags { get; private set; } = null!;

        [Output("userName")]
        public Output<string?> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:iam:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:iam:User", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new User(name, id, options);
        }
    }

    public sealed class UserArgs : Pulumi.ResourceArgs
    {
        [Input("groups")]
        private InputList<string>? _groups;
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        [Input("loginProfile")]
        public Input<Inputs.UserLoginProfileArgs>? LoginProfile { get; set; }

        [Input("managedPolicyArns")]
        private InputList<string>? _managedPolicyArns;
        public InputList<string> ManagedPolicyArns
        {
            get => _managedPolicyArns ?? (_managedPolicyArns = new InputList<string>());
            set => _managedPolicyArns = value;
        }

        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("permissionsBoundary")]
        public Input<string>? PermissionsBoundary { get; set; }

        [Input("policies")]
        private InputList<Inputs.UserPolicyArgs>? _policies;
        public InputList<Inputs.UserPolicyArgs> Policies
        {
            get => _policies ?? (_policies = new InputList<Inputs.UserPolicyArgs>());
            set => _policies = value;
        }

        [Input("tags")]
        private InputList<Inputs.UserTagArgs>? _tags;
        public InputList<Inputs.UserTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.UserTagArgs>());
            set => _tags = value;
        }

        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public UserArgs()
        {
        }
    }
}
