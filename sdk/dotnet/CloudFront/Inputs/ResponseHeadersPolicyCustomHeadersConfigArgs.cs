// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Inputs
{

    public sealed class ResponseHeadersPolicyCustomHeadersConfigArgs : Pulumi.ResourceArgs
    {
        [Input("items", required: true)]
        private InputList<Inputs.ResponseHeadersPolicyCustomHeaderArgs>? _items;
        public InputList<Inputs.ResponseHeadersPolicyCustomHeaderArgs> Items
        {
            get => _items ?? (_items = new InputList<Inputs.ResponseHeadersPolicyCustomHeaderArgs>());
            set => _items = value;
        }

        public ResponseHeadersPolicyCustomHeadersConfigArgs()
        {
        }
    }
}
