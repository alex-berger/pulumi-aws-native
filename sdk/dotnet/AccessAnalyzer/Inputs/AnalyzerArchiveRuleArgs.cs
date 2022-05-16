// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AccessAnalyzer.Inputs
{

    /// <summary>
    /// An Access Analyzer archive rule. Archive rules automatically archive new findings that meet the criteria you define when you create the rule.
    /// </summary>
    public sealed class AnalyzerArchiveRuleArgs : Pulumi.ResourceArgs
    {
        [Input("filter", required: true)]
        private InputList<Inputs.AnalyzerFilterArgs>? _filter;
        public InputList<Inputs.AnalyzerFilterArgs> Filter
        {
            get => _filter ?? (_filter = new InputList<Inputs.AnalyzerFilterArgs>());
            set => _filter = value;
        }

        /// <summary>
        /// The archive rule name
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        public AnalyzerArchiveRuleArgs()
        {
        }
    }
}
