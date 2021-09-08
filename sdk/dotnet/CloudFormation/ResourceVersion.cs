// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFormation
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-resourceversion.html
    /// </summary>
    [AwsNativeResourceType("aws-native:cloudformation:ResourceVersion")]
    public partial class ResourceVersion : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-resourceversion.html#cfn-cloudformation-resourceversion-executionrolearn
        /// </summary>
        [Output("executionRoleArn")]
        public Output<string?> ExecutionRoleArn { get; private set; } = null!;

        [Output("isDefaultVersion")]
        public Output<bool> IsDefaultVersion { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-resourceversion.html#cfn-cloudformation-resourceversion-loggingconfig
        /// </summary>
        [Output("loggingConfig")]
        public Output<Outputs.ResourceVersionLoggingConfig?> LoggingConfig { get; private set; } = null!;

        [Output("provisioningType")]
        public Output<string> ProvisioningType { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-resourceversion.html#cfn-cloudformation-resourceversion-schemahandlerpackage
        /// </summary>
        [Output("schemaHandlerPackage")]
        public Output<string> SchemaHandlerPackage { get; private set; } = null!;

        [Output("typeArn")]
        public Output<string> TypeArn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-resourceversion.html#cfn-cloudformation-resourceversion-typename
        /// </summary>
        [Output("typeName")]
        public Output<string> TypeName { get; private set; } = null!;

        [Output("versionId")]
        public Output<string> VersionId { get; private set; } = null!;

        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceVersion(string name, ResourceVersionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cloudformation:ResourceVersion", name, args ?? new ResourceVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceVersion(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cloudformation:ResourceVersion", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ResourceVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceVersion Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ResourceVersion(name, id, options);
        }
    }

    public sealed class ResourceVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-resourceversion.html#cfn-cloudformation-resourceversion-executionrolearn
        /// </summary>
        [Input("executionRoleArn")]
        public Input<string>? ExecutionRoleArn { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-resourceversion.html#cfn-cloudformation-resourceversion-loggingconfig
        /// </summary>
        [Input("loggingConfig")]
        public Input<Inputs.ResourceVersionLoggingConfigArgs>? LoggingConfig { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-resourceversion.html#cfn-cloudformation-resourceversion-schemahandlerpackage
        /// </summary>
        [Input("schemaHandlerPackage", required: true)]
        public Input<string> SchemaHandlerPackage { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-resourceversion.html#cfn-cloudformation-resourceversion-typename
        /// </summary>
        [Input("typeName", required: true)]
        public Input<string> TypeName { get; set; } = null!;

        public ResourceVersionArgs()
        {
        }
    }
}