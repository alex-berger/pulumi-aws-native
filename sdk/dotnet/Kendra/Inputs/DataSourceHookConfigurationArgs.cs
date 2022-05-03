// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    public sealed class DataSourceHookConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("invocationCondition")]
        public Input<Inputs.DataSourceDocumentAttributeConditionArgs>? InvocationCondition { get; set; }

        [Input("lambdaArn", required: true)]
        public Input<string> LambdaArn { get; set; } = null!;

        [Input("s3Bucket", required: true)]
        public Input<string> S3Bucket { get; set; } = null!;

        public DataSourceHookConfigurationArgs()
        {
        }
    }
}
