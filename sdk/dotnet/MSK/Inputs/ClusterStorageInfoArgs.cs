// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MSK.Inputs
{

    public sealed class ClusterStorageInfoArgs : Pulumi.ResourceArgs
    {
        [Input("eBSStorageInfo")]
        public Input<Inputs.ClusterEBSStorageInfoArgs>? EBSStorageInfo { get; set; }

        public ClusterStorageInfoArgs()
        {
        }
    }
}
