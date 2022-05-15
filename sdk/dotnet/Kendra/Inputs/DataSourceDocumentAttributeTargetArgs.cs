// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    public sealed class DataSourceDocumentAttributeTargetArgs : Pulumi.ResourceArgs
    {
        [Input("targetDocumentAttributeKey", required: true)]
        public Input<string> TargetDocumentAttributeKey { get; set; } = null!;

        [Input("targetDocumentAttributeValue")]
        public Input<Inputs.DataSourceDocumentAttributeValueArgs>? TargetDocumentAttributeValue { get; set; }

        [Input("targetDocumentAttributeValueDeletion")]
        public Input<bool>? TargetDocumentAttributeValueDeletion { get; set; }

        public DataSourceDocumentAttributeTargetArgs()
        {
        }
    }
}
