// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Inputs
{

    public sealed class JobDefinitionDeviceArgs : Pulumi.ResourceArgs
    {
        [Input("containerPath")]
        public Input<string>? ContainerPath { get; set; }

        [Input("hostPath")]
        public Input<string>? HostPath { get; set; }

        [Input("permissions")]
        private InputList<string>? _permissions;
        public InputList<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<string>());
            set => _permissions = value;
        }

        public JobDefinitionDeviceArgs()
        {
        }
    }
}