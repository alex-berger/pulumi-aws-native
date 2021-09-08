// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppStream
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackfleetassociation.html
    /// </summary>
    [AwsNativeResourceType("aws-native:appstream:StackFleetAssociation")]
    public partial class StackFleetAssociation : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackfleetassociation.html#cfn-appstream-stackfleetassociation-fleetname
        /// </summary>
        [Output("fleetName")]
        public Output<string> FleetName { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackfleetassociation.html#cfn-appstream-stackfleetassociation-stackname
        /// </summary>
        [Output("stackName")]
        public Output<string> StackName { get; private set; } = null!;


        /// <summary>
        /// Create a StackFleetAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StackFleetAssociation(string name, StackFleetAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:appstream:StackFleetAssociation", name, args ?? new StackFleetAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StackFleetAssociation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:appstream:StackFleetAssociation", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing StackFleetAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StackFleetAssociation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StackFleetAssociation(name, id, options);
        }
    }

    public sealed class StackFleetAssociationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackfleetassociation.html#cfn-appstream-stackfleetassociation-fleetname
        /// </summary>
        [Input("fleetName", required: true)]
        public Input<string> FleetName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackfleetassociation.html#cfn-appstream-stackfleetassociation-stackname
        /// </summary>
        [Input("stackName", required: true)]
        public Input<string> StackName { get; set; } = null!;

        public StackFleetAssociationArgs()
        {
        }
    }
}
