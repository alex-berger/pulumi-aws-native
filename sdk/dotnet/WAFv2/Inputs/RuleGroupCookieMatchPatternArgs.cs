// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Inputs
{

    /// <summary>
    /// The pattern to look for in the request cookies.
    /// </summary>
    public sealed class RuleGroupCookieMatchPatternArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Inspect all parts of the web request cookies.
        /// </summary>
        [Input("all")]
        public Input<object>? All { get; set; }

        [Input("excludedCookies")]
        private InputList<string>? _excludedCookies;
        public InputList<string> ExcludedCookies
        {
            get => _excludedCookies ?? (_excludedCookies = new InputList<string>());
            set => _excludedCookies = value;
        }

        [Input("includedCookies")]
        private InputList<string>? _includedCookies;
        public InputList<string> IncludedCookies
        {
            get => _includedCookies ?? (_includedCookies = new InputList<string>());
            set => _includedCookies = value;
        }

        public RuleGroupCookieMatchPatternArgs()
        {
        }
    }
}
