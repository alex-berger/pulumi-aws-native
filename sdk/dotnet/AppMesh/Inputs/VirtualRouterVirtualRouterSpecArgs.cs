// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualrouter-virtualrouterspec.html
    /// </summary>
    public sealed class VirtualRouterVirtualRouterSpecArgs : Pulumi.ResourceArgs
    {
        [Input("listeners", required: true)]
        private InputList<Inputs.VirtualRouterVirtualRouterListenerArgs>? _listeners;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualrouter-virtualrouterspec.html#cfn-appmesh-virtualrouter-virtualrouterspec-listeners
        /// </summary>
        public InputList<Inputs.VirtualRouterVirtualRouterListenerArgs> Listeners
        {
            get => _listeners ?? (_listeners = new InputList<Inputs.VirtualRouterVirtualRouterListenerArgs>());
            set => _listeners = value;
        }

        public VirtualRouterVirtualRouterSpecArgs()
        {
        }
    }
}