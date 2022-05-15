// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Inputs
{

    public sealed class VirtualGatewaySpecArgs : Pulumi.ResourceArgs
    {
        [Input("backendDefaults")]
        public Input<Inputs.VirtualGatewayBackendDefaultsArgs>? BackendDefaults { get; set; }

        [Input("listeners", required: true)]
        private InputList<Inputs.VirtualGatewayListenerArgs>? _listeners;
        public InputList<Inputs.VirtualGatewayListenerArgs> Listeners
        {
            get => _listeners ?? (_listeners = new InputList<Inputs.VirtualGatewayListenerArgs>());
            set => _listeners = value;
        }

        [Input("logging")]
        public Input<Inputs.VirtualGatewayLoggingArgs>? Logging { get; set; }

        public VirtualGatewaySpecArgs()
        {
        }
    }
}
