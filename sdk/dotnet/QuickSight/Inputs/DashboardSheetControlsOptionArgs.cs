// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    /// <summary>
    /// &lt;p&gt;Sheet controls option.&lt;/p&gt;
    /// </summary>
    public sealed class DashboardSheetControlsOptionArgs : Pulumi.ResourceArgs
    {
        [Input("visibilityState")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardUIState>? VisibilityState { get; set; }

        public DashboardSheetControlsOptionArgs()
        {
        }
    }
}
