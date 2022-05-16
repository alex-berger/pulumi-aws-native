// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ManagedBlockchain
{
    /// <summary>
    /// Resource Type definition for AWS::ManagedBlockchain::Member
    /// </summary>
    [Obsolete(@"Member is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:managedblockchain:Member")]
    public partial class Member : Pulumi.CustomResource
    {
        [Output("invitationId")]
        public Output<string?> InvitationId { get; private set; } = null!;

        [Output("memberConfiguration")]
        public Output<Outputs.MemberConfiguration> MemberConfiguration { get; private set; } = null!;

        [Output("memberId")]
        public Output<string> MemberId { get; private set; } = null!;

        [Output("networkConfiguration")]
        public Output<Outputs.MemberNetworkConfiguration?> NetworkConfiguration { get; private set; } = null!;

        [Output("networkId")]
        public Output<string?> NetworkId { get; private set; } = null!;


        /// <summary>
        /// Create a Member resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Member(string name, MemberArgs args, CustomResourceOptions? options = null)
            : base("aws-native:managedblockchain:Member", name, args ?? new MemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Member(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:managedblockchain:Member", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Member resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Member Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Member(name, id, options);
        }
    }

    public sealed class MemberArgs : Pulumi.ResourceArgs
    {
        [Input("invitationId")]
        public Input<string>? InvitationId { get; set; }

        [Input("memberConfiguration", required: true)]
        public Input<Inputs.MemberConfigurationArgs> MemberConfiguration { get; set; } = null!;

        [Input("networkConfiguration")]
        public Input<Inputs.MemberNetworkConfigurationArgs>? NetworkConfiguration { get; set; }

        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        public MemberArgs()
        {
        }
    }
}
