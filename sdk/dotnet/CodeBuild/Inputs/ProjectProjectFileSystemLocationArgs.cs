// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeBuild.Inputs
{

    public sealed class ProjectProjectFileSystemLocationArgs : Pulumi.ResourceArgs
    {
        [Input("identifier", required: true)]
        public Input<string> Identifier { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("mountOptions")]
        public Input<string>? MountOptions { get; set; }

        [Input("mountPoint", required: true)]
        public Input<string> MountPoint { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ProjectProjectFileSystemLocationArgs()
        {
        }
    }
}