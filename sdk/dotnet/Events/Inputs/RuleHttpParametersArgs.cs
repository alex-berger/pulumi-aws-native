// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-httpparameters.html
    /// </summary>
    public sealed class RuleHttpParametersArgs : Pulumi.ResourceArgs
    {
        [Input("headerParameters")]
        private InputMap<string>? _headerParameters;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-httpparameters.html#cfn-events-rule-httpparameters-headerparameters
        /// </summary>
        public InputMap<string> HeaderParameters
        {
            get => _headerParameters ?? (_headerParameters = new InputMap<string>());
            set => _headerParameters = value;
        }

        [Input("pathParameterValues")]
        private InputList<string>? _pathParameterValues;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-httpparameters.html#cfn-events-rule-httpparameters-pathparametervalues
        /// </summary>
        public InputList<string> PathParameterValues
        {
            get => _pathParameterValues ?? (_pathParameterValues = new InputList<string>());
            set => _pathParameterValues = value;
        }

        [Input("queryStringParameters")]
        private InputMap<string>? _queryStringParameters;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-httpparameters.html#cfn-events-rule-httpparameters-querystringparameters
        /// </summary>
        public InputMap<string> QueryStringParameters
        {
            get => _queryStringParameters ?? (_queryStringParameters = new InputMap<string>());
            set => _queryStringParameters = value;
        }

        public RuleHttpParametersArgs()
        {
        }
    }
}
