// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway
{
    public static class GetGatewayResponse
    {
        /// <summary>
        /// Resource Type definition for AWS::ApiGateway::GatewayResponse
        /// </summary>
        public static Task<GetGatewayResponseResult> InvokeAsync(GetGatewayResponseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGatewayResponseResult>("aws-native:apigateway:getGatewayResponse", args ?? new GetGatewayResponseArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ApiGateway::GatewayResponse
        /// </summary>
        public static Output<GetGatewayResponseResult> Invoke(GetGatewayResponseInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGatewayResponseResult>("aws-native:apigateway:getGatewayResponse", args ?? new GetGatewayResponseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A Cloudformation auto generated ID.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetGatewayResponseArgs()
        {
        }
    }

    public sealed class GetGatewayResponseInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A Cloudformation auto generated ID.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetGatewayResponseInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGatewayResponseResult
    {
        /// <summary>
        /// A Cloudformation auto generated ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The response parameters (paths, query strings, and headers) for the response.
        /// </summary>
        public readonly object? ResponseParameters;
        /// <summary>
        /// The response templates for the response.
        /// </summary>
        public readonly object? ResponseTemplates;
        /// <summary>
        /// The HTTP status code for the response.
        /// </summary>
        public readonly string? StatusCode;

        [OutputConstructor]
        private GetGatewayResponseResult(
            string? id,

            object? responseParameters,

            object? responseTemplates,

            string? statusCode)
        {
            Id = id;
            ResponseParameters = responseParameters;
            ResponseTemplates = responseTemplates;
            StatusCode = statusCode;
        }
    }
}