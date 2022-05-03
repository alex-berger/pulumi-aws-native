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
    /// &lt;p&gt;An integer parameter.&lt;/p&gt;
    /// </summary>
    public sealed class AnalysisIntegerParameterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The name of the integer parameter.&lt;/p&gt;
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<double>? _values;

        /// <summary>
        /// &lt;p&gt;The values for the integer parameter.&lt;/p&gt;
        /// </summary>
        public InputList<double> Values
        {
            get => _values ?? (_values = new InputList<double>());
            set => _values = value;
        }

        public AnalysisIntegerParameterArgs()
        {
        }
    }
}
