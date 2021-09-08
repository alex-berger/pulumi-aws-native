// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodePipeline
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html
    /// </summary>
    [AwsNativeResourceType("aws-native:codepipeline:Webhook")]
    public partial class Webhook : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-authentication
        /// </summary>
        [Output("authentication")]
        public Output<string> Authentication { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-authenticationconfiguration
        /// </summary>
        [Output("authenticationConfiguration")]
        public Output<Outputs.WebhookWebhookAuthConfiguration> AuthenticationConfiguration { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-filters
        /// </summary>
        [Output("filters")]
        public Output<ImmutableArray<Outputs.WebhookWebhookFilterRule>> Filters { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-name
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-registerwiththirdparty
        /// </summary>
        [Output("registerWithThirdParty")]
        public Output<bool?> RegisterWithThirdParty { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-targetaction
        /// </summary>
        [Output("targetAction")]
        public Output<string> TargetAction { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-targetpipeline
        /// </summary>
        [Output("targetPipeline")]
        public Output<string> TargetPipeline { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-targetpipelineversion
        /// </summary>
        [Output("targetPipelineVersion")]
        public Output<int> TargetPipelineVersion { get; private set; } = null!;

        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Webhook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Webhook(string name, WebhookArgs args, CustomResourceOptions? options = null)
            : base("aws-native:codepipeline:Webhook", name, args ?? new WebhookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Webhook(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:codepipeline:Webhook", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Webhook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Webhook Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Webhook(name, id, options);
        }
    }

    public sealed class WebhookArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-authentication
        /// </summary>
        [Input("authentication", required: true)]
        public Input<string> Authentication { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-authenticationconfiguration
        /// </summary>
        [Input("authenticationConfiguration", required: true)]
        public Input<Inputs.WebhookWebhookAuthConfigurationArgs> AuthenticationConfiguration { get; set; } = null!;

        [Input("filters", required: true)]
        private InputList<Inputs.WebhookWebhookFilterRuleArgs>? _filters;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-filters
        /// </summary>
        public InputList<Inputs.WebhookWebhookFilterRuleArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.WebhookWebhookFilterRuleArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-registerwiththirdparty
        /// </summary>
        [Input("registerWithThirdParty")]
        public Input<bool>? RegisterWithThirdParty { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-targetaction
        /// </summary>
        [Input("targetAction", required: true)]
        public Input<string> TargetAction { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-targetpipeline
        /// </summary>
        [Input("targetPipeline", required: true)]
        public Input<string> TargetPipeline { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-targetpipelineversion
        /// </summary>
        [Input("targetPipelineVersion", required: true)]
        public Input<int> TargetPipelineVersion { get; set; } = null!;

        public WebhookArgs()
        {
        }
    }
}