// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Inputs
{

    public sealed class ListenerRuleRuleConditionArgs : Pulumi.ResourceArgs
    {
        [Input("field")]
        public Input<string>? Field { get; set; }

        [Input("hostHeaderConfig")]
        public Input<Inputs.ListenerRuleHostHeaderConfigArgs>? HostHeaderConfig { get; set; }

        [Input("httpHeaderConfig")]
        public Input<Inputs.ListenerRuleHttpHeaderConfigArgs>? HttpHeaderConfig { get; set; }

        [Input("httpRequestMethodConfig")]
        public Input<Inputs.ListenerRuleHttpRequestMethodConfigArgs>? HttpRequestMethodConfig { get; set; }

        [Input("pathPatternConfig")]
        public Input<Inputs.ListenerRulePathPatternConfigArgs>? PathPatternConfig { get; set; }

        [Input("queryStringConfig")]
        public Input<Inputs.ListenerRuleQueryStringConfigArgs>? QueryStringConfig { get; set; }

        [Input("sourceIpConfig")]
        public Input<Inputs.ListenerRuleSourceIpConfigArgs>? SourceIpConfig { get; set; }

        [Input("values")]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public ListenerRuleRuleConditionArgs()
        {
        }
    }
}
