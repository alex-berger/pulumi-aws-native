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
    /// Resource Type definition for AWS::IAM::UserToGroupAddition
    /// </summary>
    [Obsolete(@"UserToGroupAddition is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:iam:UserToGroupAddition")]
    public partial class UserToGroupAddition : Pulumi.CustomResource
    {
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        [Output("users")]
        public Output<ImmutableArray<string>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a UserToGroupAddition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserToGroupAddition(string name, UserToGroupAdditionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:iam:UserToGroupAddition", name, args ?? new UserToGroupAdditionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserToGroupAddition(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:iam:UserToGroupAddition", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing UserToGroupAddition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserToGroupAddition Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new UserToGroupAddition(name, id, options);
        }
    }

    public sealed class UserToGroupAdditionArgs : Pulumi.ResourceArgs
    {
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        [Input("users", required: true)]
        private InputList<string>? _users;
        public InputList<string> Users
        {
            get => _users ?? (_users = new InputList<string>());
            set => _users = value;
        }

        public UserToGroupAdditionArgs()
        {
        }
    }
}