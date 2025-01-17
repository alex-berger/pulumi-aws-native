// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway
{
    public static class GetModel
    {
        /// <summary>
        /// Resource Type definition for AWS::ApiGateway::Model
        /// </summary>
        public static Task<GetModelResult> InvokeAsync(GetModelArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetModelResult>("aws-native:apigateway:getModel", args ?? new GetModelArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ApiGateway::Model
        /// </summary>
        public static Output<GetModelResult> Invoke(GetModelInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetModelResult>("aws-native:apigateway:getModel", args ?? new GetModelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetModelArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A name for the model. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The ID of a REST API with which to associate this model.
        /// </summary>
        [Input("restApiId", required: true)]
        public string RestApiId { get; set; } = null!;

        public GetModelArgs()
        {
        }
    }

    public sealed class GetModelInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A name for the model. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of a REST API with which to associate this model.
        /// </summary>
        [Input("restApiId", required: true)]
        public Input<string> RestApiId { get; set; } = null!;

        public GetModelInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetModelResult
    {
        /// <summary>
        /// A description that identifies this model.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The schema to use to transform data to one or more output formats. Specify null ({}) if you don't want to specify a schema.
        /// </summary>
        public readonly object? Schema;

        [OutputConstructor]
        private GetModelResult(
            string? description,

            object? schema)
        {
            Description = description;
            Schema = schema;
        }
    }
}
