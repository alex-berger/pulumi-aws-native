// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass.Inputs
{

    public sealed class FunctionDefinitionRunAsArgs : Pulumi.ResourceArgs
    {
        [Input("gid")]
        public Input<int>? Gid { get; set; }

        [Input("uid")]
        public Input<int>? Uid { get; set; }

        public FunctionDefinitionRunAsArgs()
        {
        }
    }
}
