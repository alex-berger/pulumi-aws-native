// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSO.Inputs
{

    public sealed class InstanceAccessControlAttributeConfigurationAccessControlAttributeValueArgs : Pulumi.ResourceArgs
    {
        [Input("source", required: true)]
        private InputList<string>? _source;
        public InputList<string> Source
        {
            get => _source ?? (_source = new InputList<string>());
            set => _source = value;
        }

        public InstanceAccessControlAttributeConfigurationAccessControlAttributeValueArgs()
        {
        }
    }
}
