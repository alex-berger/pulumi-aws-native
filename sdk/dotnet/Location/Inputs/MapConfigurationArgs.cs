// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Location.Inputs
{

    public sealed class MapConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("style", required: true)]
        public Input<string> Style { get; set; } = null!;

        public MapConfigurationArgs()
        {
        }
    }
}
