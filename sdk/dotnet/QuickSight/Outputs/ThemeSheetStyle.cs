// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;The theme display options for sheets. &lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class ThemeSheetStyle
    {
        public readonly Outputs.ThemeTileStyle? Tile;
        public readonly Outputs.ThemeTileLayoutStyle? TileLayout;

        [OutputConstructor]
        private ThemeSheetStyle(
            Outputs.ThemeTileStyle? tile,

            Outputs.ThemeTileLayoutStyle? tileLayout)
        {
            Tile = tile;
            TileLayout = tileLayout;
        }
    }
}