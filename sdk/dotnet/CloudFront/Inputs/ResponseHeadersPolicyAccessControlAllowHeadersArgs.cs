// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Inputs
{

    public sealed class ResponseHeadersPolicyAccessControlAllowHeadersArgs : Pulumi.ResourceArgs
    {
        [Input("items", required: true)]
        private InputList<string>? _items;
        public InputList<string> Items
        {
            get => _items ?? (_items = new InputList<string>());
            set => _items = value;
        }

        public ResponseHeadersPolicyAccessControlAllowHeadersArgs()
        {
        }
    }
}
