// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDB.Inputs
{

    public sealed class TableAttributeDefinitionArgs : Pulumi.ResourceArgs
    {
        [Input("attributeName", required: true)]
        public Input<string> AttributeName { get; set; } = null!;

        [Input("attributeType", required: true)]
        public Input<string> AttributeType { get; set; } = null!;

        public TableAttributeDefinitionArgs()
        {
        }
    }
}
