// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpsWorks.Inputs
{

    public sealed class LayerRecipesArgs : Pulumi.ResourceArgs
    {
        [Input("configure")]
        private InputList<string>? _configure;
        public InputList<string> Configure
        {
            get => _configure ?? (_configure = new InputList<string>());
            set => _configure = value;
        }

        [Input("deploy")]
        private InputList<string>? _deploy;
        public InputList<string> Deploy
        {
            get => _deploy ?? (_deploy = new InputList<string>());
            set => _deploy = value;
        }

        [Input("setup")]
        private InputList<string>? _setup;
        public InputList<string> Setup
        {
            get => _setup ?? (_setup = new InputList<string>());
            set => _setup = value;
        }

        [Input("shutdown")]
        private InputList<string>? _shutdown;
        public InputList<string> Shutdown
        {
            get => _shutdown ?? (_shutdown = new InputList<string>());
            set => _shutdown = value;
        }

        [Input("undeploy")]
        private InputList<string>? _undeploy;
        public InputList<string> Undeploy
        {
            get => _undeploy ?? (_undeploy = new InputList<string>());
            set => _undeploy = value;
        }

        public LayerRecipesArgs()
        {
        }
    }
}
