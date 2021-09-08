// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html
    /// </summary>
    public sealed class MethodMethodResponseArgs : Pulumi.ResourceArgs
    {
        [Input("responseModels")]
        private InputMap<string>? _responseModels;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html#cfn-apigateway-method-methodresponse-responsemodels
        /// </summary>
        public InputMap<string> ResponseModels
        {
            get => _responseModels ?? (_responseModels = new InputMap<string>());
            set => _responseModels = value;
        }

        [Input("responseParameters")]
        private InputMap<bool>? _responseParameters;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html#cfn-apigateway-method-methodresponse-responseparameters
        /// </summary>
        public InputMap<bool> ResponseParameters
        {
            get => _responseParameters ?? (_responseParameters = new InputMap<bool>());
            set => _responseParameters = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html#cfn-apigateway-method-methodresponse-statuscode
        /// </summary>
        [Input("statusCode", required: true)]
        public Input<string> StatusCode { get; set; } = null!;

        public MethodMethodResponseArgs()
        {
        }
    }
}