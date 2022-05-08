// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppStream.Inputs
{

    public sealed class StackStorageConnectorArgs : Pulumi.ResourceArgs
    {
        [Input("connectorType", required: true)]
        public Input<string> ConnectorType { get; set; } = null!;

        [Input("domains")]
        private InputList<string>? _domains;
        public InputList<string> Domains
        {
            get => _domains ?? (_domains = new InputList<string>());
            set => _domains = value;
        }

        [Input("resourceIdentifier")]
        public Input<string>? ResourceIdentifier { get; set; }

        public StackStorageConnectorArgs()
        {
        }
    }
}
