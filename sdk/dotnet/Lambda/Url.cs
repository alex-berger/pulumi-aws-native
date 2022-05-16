// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda
{
    /// <summary>
    /// Resource Type definition for AWS::Lambda::Url
    /// </summary>
    [AwsNativeResourceType("aws-native:lambda:Url")]
    public partial class Url : Pulumi.CustomResource
    {
        /// <summary>
        /// Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
        /// </summary>
        [Output("authType")]
        public Output<Pulumi.AwsNative.Lambda.UrlAuthType> AuthType { get; private set; } = null!;

        [Output("cors")]
        public Output<Outputs.UrlCors?> Cors { get; private set; } = null!;

        /// <summary>
        /// The full Amazon Resource Name (ARN) of the function associated with the Function URL.
        /// </summary>
        [Output("functionArn")]
        public Output<string> FunctionArn { get; private set; } = null!;

        /// <summary>
        /// The generated url for this resource.
        /// </summary>
        [Output("functionUrl")]
        public Output<string> FunctionUrl { get; private set; } = null!;

        /// <summary>
        /// The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed.
        /// </summary>
        [Output("qualifier")]
        public Output<string?> Qualifier { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the function associated with the Function URL.
        /// </summary>
        [Output("targetFunctionArn")]
        public Output<string> TargetFunctionArn { get; private set; } = null!;


        /// <summary>
        /// Create a Url resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Url(string name, UrlArgs args, CustomResourceOptions? options = null)
            : base("aws-native:lambda:Url", name, args ?? new UrlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Url(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:lambda:Url", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Url resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Url Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Url(name, id, options);
        }
    }

    public sealed class UrlArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
        /// </summary>
        [Input("authType", required: true)]
        public Input<Pulumi.AwsNative.Lambda.UrlAuthType> AuthType { get; set; } = null!;

        [Input("cors")]
        public Input<Inputs.UrlCorsArgs>? Cors { get; set; }

        /// <summary>
        /// The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed.
        /// </summary>
        [Input("qualifier")]
        public Input<string>? Qualifier { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the function associated with the Function URL.
        /// </summary>
        [Input("targetFunctionArn", required: true)]
        public Input<string> TargetFunctionArn { get; set; } = null!;

        public UrlArgs()
        {
        }
    }
}