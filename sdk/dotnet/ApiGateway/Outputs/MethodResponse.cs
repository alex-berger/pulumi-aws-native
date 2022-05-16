// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway.Outputs
{

    [OutputType]
    public sealed class MethodResponse
    {
        /// <summary>
        /// The resources used for the response's content type. Specify response models as key-value pairs (string-to-string maps), with a content type as the key and a Model resource name as the value.
        /// </summary>
        public readonly object? ResponseModels;
        /// <summary>
        /// Response parameters that API Gateway sends to the client that called a method. Specify response parameters as key-value pairs (string-to-Boolean maps), with a destination as the key and a Boolean as the value.
        /// </summary>
        public readonly object? ResponseParameters;
        /// <summary>
        /// The method response's status code, which you map to an IntegrationResponse.
        /// </summary>
        public readonly string StatusCode;

        [OutputConstructor]
        private MethodResponse(
            object? responseModels,

            object? responseParameters,

            string statusCode)
        {
            ResponseModels = responseModels;
            ResponseParameters = responseParameters;
            StatusCode = statusCode;
        }
    }
}