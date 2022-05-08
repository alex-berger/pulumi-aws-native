// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect
{
    /// <summary>
    /// Resource Type definition for AWS::Connect::UserHierarchyGroup
    /// </summary>
    [AwsNativeResourceType("aws-native:connect:UserHierarchyGroup")]
    public partial class UserHierarchyGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The identifier of the Amazon Connect instance.
        /// </summary>
        [Output("instanceArn")]
        public Output<string> InstanceArn { get; private set; } = null!;

        /// <summary>
        /// The name of the user hierarchy group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the parent user hierarchy group.
        /// </summary>
        [Output("parentGroupArn")]
        public Output<string?> ParentGroupArn { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the user hierarchy group.
        /// </summary>
        [Output("userHierarchyGroupArn")]
        public Output<string> UserHierarchyGroupArn { get; private set; } = null!;


        /// <summary>
        /// Create a UserHierarchyGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserHierarchyGroup(string name, UserHierarchyGroupArgs args, CustomResourceOptions? options = null)
            : base("aws-native:connect:UserHierarchyGroup", name, args ?? new UserHierarchyGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserHierarchyGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:connect:UserHierarchyGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing UserHierarchyGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserHierarchyGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new UserHierarchyGroup(name, id, options);
        }
    }

    public sealed class UserHierarchyGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identifier of the Amazon Connect instance.
        /// </summary>
        [Input("instanceArn", required: true)]
        public Input<string> InstanceArn { get; set; } = null!;

        /// <summary>
        /// The name of the user hierarchy group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) for the parent user hierarchy group.
        /// </summary>
        [Input("parentGroupArn")]
        public Input<string>? ParentGroupArn { get; set; }

        public UserHierarchyGroupArgs()
        {
        }
    }
}
