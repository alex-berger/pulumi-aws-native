// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFormation
{
    /// <summary>
    /// A module that has been registered in the CloudFormation registry.
    /// </summary>
    [AwsNativeResourceType("aws-native:cloudformation:ModuleVersion")]
    public partial class ModuleVersion : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the module.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The description of the registered module.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The URL of a page providing detailed documentation for this module.
        /// </summary>
        [Output("documentationUrl")]
        public Output<string> DocumentationUrl { get; private set; } = null!;

        /// <summary>
        /// Indicator of whether this module version is the current default version
        /// </summary>
        [Output("isDefaultVersion")]
        public Output<bool> IsDefaultVersion { get; private set; } = null!;

        /// <summary>
        /// The name of the module being registered.
        /// 
        /// Recommended module naming pattern: company_or_organization::service::type::MODULE.
        /// </summary>
        [Output("moduleName")]
        public Output<string> ModuleName { get; private set; } = null!;

        /// <summary>
        /// The url to the S3 bucket containing the schema and template fragment for the module you want to register.
        /// </summary>
        [Output("modulePackage")]
        public Output<string> ModulePackage { get; private set; } = null!;

        /// <summary>
        /// The schema defining input parameters to and resources generated by the module.
        /// </summary>
        [Output("schema")]
        public Output<string> Schema { get; private set; } = null!;

        /// <summary>
        /// The time that the specified module version was registered.
        /// </summary>
        [Output("timeCreated")]
        public Output<string> TimeCreated { get; private set; } = null!;

        /// <summary>
        /// The version ID of the module represented by this module instance.
        /// </summary>
        [Output("versionId")]
        public Output<string> VersionId { get; private set; } = null!;

        /// <summary>
        /// The scope at which the type is visible and usable in CloudFormation operations.
        /// 
        /// The only allowed value at present is:
        /// 
        /// PRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.
        /// </summary>
        [Output("visibility")]
        public Output<Pulumi.AwsNative.CloudFormation.ModuleVersionVisibility> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a ModuleVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ModuleVersion(string name, ModuleVersionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cloudformation:ModuleVersion", name, args ?? new ModuleVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ModuleVersion(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cloudformation:ModuleVersion", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ModuleVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ModuleVersion Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ModuleVersion(name, id, options);
        }
    }

    public sealed class ModuleVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the module being registered.
        /// 
        /// Recommended module naming pattern: company_or_organization::service::type::MODULE.
        /// </summary>
        [Input("moduleName", required: true)]
        public Input<string> ModuleName { get; set; } = null!;

        /// <summary>
        /// The url to the S3 bucket containing the schema and template fragment for the module you want to register.
        /// </summary>
        [Input("modulePackage", required: true)]
        public Input<string> ModulePackage { get; set; } = null!;

        public ModuleVersionArgs()
        {
        }
    }
}
