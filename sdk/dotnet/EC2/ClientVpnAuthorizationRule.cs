// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    /// <summary>
    /// Resource Type definition for AWS::EC2::ClientVpnAuthorizationRule
    /// </summary>
    [Obsolete(@"ClientVpnAuthorizationRule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:ec2:ClientVpnAuthorizationRule")]
    public partial class ClientVpnAuthorizationRule : Pulumi.CustomResource
    {
        [Output("accessGroupId")]
        public Output<string?> AccessGroupId { get; private set; } = null!;

        [Output("authorizeAllGroups")]
        public Output<bool?> AuthorizeAllGroups { get; private set; } = null!;

        [Output("clientVpnEndpointId")]
        public Output<string> ClientVpnEndpointId { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("targetNetworkCidr")]
        public Output<string> TargetNetworkCidr { get; private set; } = null!;


        /// <summary>
        /// Create a ClientVpnAuthorizationRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientVpnAuthorizationRule(string name, ClientVpnAuthorizationRuleArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:ClientVpnAuthorizationRule", name, args ?? new ClientVpnAuthorizationRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClientVpnAuthorizationRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:ClientVpnAuthorizationRule", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ClientVpnAuthorizationRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientVpnAuthorizationRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ClientVpnAuthorizationRule(name, id, options);
        }
    }

    public sealed class ClientVpnAuthorizationRuleArgs : Pulumi.ResourceArgs
    {
        [Input("accessGroupId")]
        public Input<string>? AccessGroupId { get; set; }

        [Input("authorizeAllGroups")]
        public Input<bool>? AuthorizeAllGroups { get; set; }

        [Input("clientVpnEndpointId", required: true)]
        public Input<string> ClientVpnEndpointId { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("targetNetworkCidr", required: true)]
        public Input<string> TargetNetworkCidr { get; set; } = null!;

        public ClientVpnAuthorizationRuleArgs()
        {
        }
    }
}
