// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway
{
    public static class GetMethod
    {
        /// <summary>
        /// Resource Type definition for AWS::ApiGateway::Method
        /// </summary>
        public static Task<GetMethodResult> InvokeAsync(GetMethodArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMethodResult>("aws-native:apigateway:getMethod", args ?? new GetMethodArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ApiGateway::Method
        /// </summary>
        public static Output<GetMethodResult> Invoke(GetMethodInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMethodResult>("aws-native:apigateway:getMethod", args ?? new GetMethodInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMethodArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The backend system that the method calls when it receives a request.
        /// </summary>
        [Input("httpMethod", required: true)]
        public string HttpMethod { get; set; } = null!;

        /// <summary>
        /// The ID of an API Gateway resource.
        /// </summary>
        [Input("resourceId", required: true)]
        public string ResourceId { get; set; } = null!;

        /// <summary>
        /// The ID of the RestApi resource in which API Gateway creates the method.
        /// </summary>
        [Input("restApiId", required: true)]
        public string RestApiId { get; set; } = null!;

        public GetMethodArgs()
        {
        }
    }

    public sealed class GetMethodInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The backend system that the method calls when it receives a request.
        /// </summary>
        [Input("httpMethod", required: true)]
        public Input<string> HttpMethod { get; set; } = null!;

        /// <summary>
        /// The ID of an API Gateway resource.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// The ID of the RestApi resource in which API Gateway creates the method.
        /// </summary>
        [Input("restApiId", required: true)]
        public Input<string> RestApiId { get; set; } = null!;

        public GetMethodInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMethodResult
    {
        /// <summary>
        /// Indicates whether the method requires clients to submit a valid API key.
        /// </summary>
        public readonly bool? ApiKeyRequired;
        /// <summary>
        /// A list of authorization scopes configured on the method.
        /// </summary>
        public readonly ImmutableArray<string> AuthorizationScopes;
        /// <summary>
        /// The method's authorization type.
        /// </summary>
        public readonly Pulumi.AwsNative.ApiGateway.MethodAuthorizationType? AuthorizationType;
        /// <summary>
        /// The identifier of the authorizer to use on this method.
        /// </summary>
        public readonly string? AuthorizerId;
        /// <summary>
        /// The backend system that the method calls when it receives a request.
        /// </summary>
        public readonly Outputs.MethodIntegration? Integration;
        /// <summary>
        /// The responses that can be sent to the client who calls the method.
        /// </summary>
        public readonly ImmutableArray<Outputs.MethodResponse> MethodResponses;
        /// <summary>
        /// A friendly operation name for the method.
        /// </summary>
        public readonly string? OperationName;
        /// <summary>
        /// The resources that are used for the request's content type. Specify request models as key-value pairs (string-to-string mapping), with a content type as the key and a Model resource name as the value.
        /// </summary>
        public readonly object? RequestModels;
        /// <summary>
        /// The request parameters that API Gateway accepts. Specify request parameters as key-value pairs (string-to-Boolean mapping), with a source as the key and a Boolean as the value.
        /// </summary>
        public readonly object? RequestParameters;
        /// <summary>
        /// The ID of the associated request validator.
        /// </summary>
        public readonly string? RequestValidatorId;

        [OutputConstructor]
        private GetMethodResult(
            bool? apiKeyRequired,

            ImmutableArray<string> authorizationScopes,

            Pulumi.AwsNative.ApiGateway.MethodAuthorizationType? authorizationType,

            string? authorizerId,

            Outputs.MethodIntegration? integration,

            ImmutableArray<Outputs.MethodResponse> methodResponses,

            string? operationName,

            object? requestModels,

            object? requestParameters,

            string? requestValidatorId)
        {
            ApiKeyRequired = apiKeyRequired;
            AuthorizationScopes = authorizationScopes;
            AuthorizationType = authorizationType;
            AuthorizerId = authorizerId;
            Integration = integration;
            MethodResponses = methodResponses;
            OperationName = operationName;
            RequestModels = requestModels;
            RequestParameters = requestParameters;
            RequestValidatorId = requestValidatorId;
        }
    }
}
