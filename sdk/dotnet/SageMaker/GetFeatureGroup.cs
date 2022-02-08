// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetFeatureGroup
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::FeatureGroup
        /// </summary>
        public static Task<GetFeatureGroupResult> InvokeAsync(GetFeatureGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFeatureGroupResult>("aws-native:sagemaker:getFeatureGroup", args ?? new GetFeatureGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::FeatureGroup
        /// </summary>
        public static Output<GetFeatureGroupResult> Invoke(GetFeatureGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFeatureGroupResult>("aws-native:sagemaker:getFeatureGroup", args ?? new GetFeatureGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFeatureGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Name of the FeatureGroup.
        /// </summary>
        [Input("featureGroupName", required: true)]
        public string FeatureGroupName { get; set; } = null!;

        public GetFeatureGroupArgs()
        {
        }
    }

    public sealed class GetFeatureGroupInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Name of the FeatureGroup.
        /// </summary>
        [Input("featureGroupName", required: true)]
        public Input<string> FeatureGroupName { get; set; } = null!;

        public GetFeatureGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFeatureGroupResult
    {
        [OutputConstructor]
        private GetFeatureGroupResult()
        {
        }
    }
}