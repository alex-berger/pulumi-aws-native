// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio.Inputs
{

    /// <summary>
    /// &lt;p&gt;The upload storage root location (folder) on streaming workstations where files are
    ///             uploaded.&lt;/p&gt;
    /// </summary>
    public sealed class LaunchProfileStreamingSessionStorageRootArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The folder path in Linux workstations where files are uploaded.&lt;/p&gt;
        /// </summary>
        [Input("linux")]
        public Input<string>? Linux { get; set; }

        /// <summary>
        /// &lt;p&gt;The folder path in Windows workstations where files are uploaded.&lt;/p&gt;
        /// </summary>
        [Input("windows")]
        public Input<string>? Windows { get; set; }

        public LaunchProfileStreamingSessionStorageRootArgs()
        {
        }
    }
}
