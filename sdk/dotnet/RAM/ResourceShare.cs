// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RAM
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html
    /// </summary>
    [AwsNativeResourceType("aws-native:ram:ResourceShare")]
    public partial class ResourceShare : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-allowexternalprincipals
        /// </summary>
        [Output("allowExternalPrincipals")]
        public Output<bool?> AllowExternalPrincipals { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-permissionarns
        /// </summary>
        [Output("permissionArns")]
        public Output<ImmutableArray<string>> PermissionArns { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-principals
        /// </summary>
        [Output("principals")]
        public Output<ImmutableArray<string>> Principals { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-resourcearns
        /// </summary>
        [Output("resourceArns")]
        public Output<ImmutableArray<string>> ResourceArns { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceShare resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceShare(string name, ResourceShareArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ram:ResourceShare", name, args ?? new ResourceShareArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceShare(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ram:ResourceShare", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ResourceShare resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceShare Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ResourceShare(name, id, options);
        }
    }

    public sealed class ResourceShareArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-allowexternalprincipals
        /// </summary>
        [Input("allowExternalPrincipals")]
        public Input<bool>? AllowExternalPrincipals { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("permissionArns")]
        private InputList<string>? _permissionArns;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-permissionarns
        /// </summary>
        public InputList<string> PermissionArns
        {
            get => _permissionArns ?? (_permissionArns = new InputList<string>());
            set => _permissionArns = value;
        }

        [Input("principals")]
        private InputList<string>? _principals;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-principals
        /// </summary>
        public InputList<string> Principals
        {
            get => _principals ?? (_principals = new InputList<string>());
            set => _principals = value;
        }

        [Input("resourceArns")]
        private InputList<string>? _resourceArns;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-resourcearns
        /// </summary>
        public InputList<string> ResourceArns
        {
            get => _resourceArns ?? (_resourceArns = new InputList<string>());
            set => _resourceArns = value;
        }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-tags
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public ResourceShareArgs()
        {
        }
    }
}
