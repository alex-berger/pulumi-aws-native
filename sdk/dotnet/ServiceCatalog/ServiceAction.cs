// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceCatalog
{
    /// <summary>
    /// Resource Schema for AWS::ServiceCatalog::ServiceAction
    /// </summary>
    [AwsNativeResourceType("aws-native:servicecatalog:ServiceAction")]
    public partial class ServiceAction : Pulumi.CustomResource
    {
        [Output("acceptLanguage")]
        public Output<Pulumi.AwsNative.ServiceCatalog.ServiceActionAcceptLanguage?> AcceptLanguage { get; private set; } = null!;

        [Output("definition")]
        public Output<ImmutableArray<Outputs.ServiceActionDefinitionParameter>> Definition { get; private set; } = null!;

        [Output("definitionType")]
        public Output<Pulumi.AwsNative.ServiceCatalog.ServiceActionDefinitionType> DefinitionType { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceAction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceAction(string name, ServiceActionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:servicecatalog:ServiceAction", name, args ?? new ServiceActionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceAction(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:servicecatalog:ServiceAction", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceAction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceAction Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServiceAction(name, id, options);
        }
    }

    public sealed class ServiceActionArgs : Pulumi.ResourceArgs
    {
        [Input("acceptLanguage")]
        public Input<Pulumi.AwsNative.ServiceCatalog.ServiceActionAcceptLanguage>? AcceptLanguage { get; set; }

        [Input("definition", required: true)]
        private InputList<Inputs.ServiceActionDefinitionParameterArgs>? _definition;
        public InputList<Inputs.ServiceActionDefinitionParameterArgs> Definition
        {
            get => _definition ?? (_definition = new InputList<Inputs.ServiceActionDefinitionParameterArgs>());
            set => _definition = value;
        }

        [Input("definitionType", required: true)]
        public Input<Pulumi.AwsNative.ServiceCatalog.ServiceActionDefinitionType> DefinitionType { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public ServiceActionArgs()
        {
        }
    }
}