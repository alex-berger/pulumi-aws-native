// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-versionweight.html
    /// </summary>
    public sealed class AliasVersionWeightArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-versionweight.html#cfn-lambda-alias-versionweight-functionversion
        /// </summary>
        [Input("functionVersion", required: true)]
        public Input<string> FunctionVersion { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-versionweight.html#cfn-lambda-alias-versionweight-functionweight
        /// </summary>
        [Input("functionWeight", required: true)]
        public Input<double> FunctionWeight { get; set; } = null!;

        public AliasVersionWeightArgs()
        {
        }
    }
}
