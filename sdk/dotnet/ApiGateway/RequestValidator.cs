// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway
{
    /// <summary>
    /// Resource Type definition for AWS::ApiGateway::RequestValidator
    /// </summary>
    [AwsNativeResourceType("aws-native:apigateway:RequestValidator")]
    public partial class RequestValidator : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the request validator.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the request validator.
        /// </summary>
        [Output("requestValidatorId")]
        public Output<string> RequestValidatorId { get; private set; } = null!;

        /// <summary>
        /// The identifier of the targeted API entity.
        /// </summary>
        [Output("restApiId")]
        public Output<string> RestApiId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to validate the request body according to the configured schema for the targeted API and method. 
        /// </summary>
        [Output("validateRequestBody")]
        public Output<bool?> ValidateRequestBody { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to validate request parameters.
        /// </summary>
        [Output("validateRequestParameters")]
        public Output<bool?> ValidateRequestParameters { get; private set; } = null!;


        /// <summary>
        /// Create a RequestValidator resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RequestValidator(string name, RequestValidatorArgs args, CustomResourceOptions? options = null)
            : base("aws-native:apigateway:RequestValidator", name, args ?? new RequestValidatorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RequestValidator(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:apigateway:RequestValidator", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing RequestValidator resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RequestValidator Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RequestValidator(name, id, options);
        }
    }

    public sealed class RequestValidatorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the request validator.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The identifier of the targeted API entity.
        /// </summary>
        [Input("restApiId", required: true)]
        public Input<string> RestApiId { get; set; } = null!;

        /// <summary>
        /// Indicates whether to validate the request body according to the configured schema for the targeted API and method. 
        /// </summary>
        [Input("validateRequestBody")]
        public Input<bool>? ValidateRequestBody { get; set; }

        /// <summary>
        /// Indicates whether to validate request parameters.
        /// </summary>
        [Input("validateRequestParameters")]
        public Input<bool>? ValidateRequestParameters { get; set; }

        public RequestValidatorArgs()
        {
        }
    }
}
