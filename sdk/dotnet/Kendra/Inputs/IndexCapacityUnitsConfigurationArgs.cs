// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    public sealed class IndexCapacityUnitsConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("queryCapacityUnits", required: true)]
        public Input<int> QueryCapacityUnits { get; set; } = null!;

        [Input("storageCapacityUnits", required: true)]
        public Input<int> StorageCapacityUnits { get; set; } = null!;

        public IndexCapacityUnitsConfigurationArgs()
        {
        }
    }
}
