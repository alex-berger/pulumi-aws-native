// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FMS.Inputs
{

    /// <summary>
    /// An FMS includeMap or excludeMap.
    /// </summary>
    public sealed class PolicyIEMapArgs : Pulumi.ResourceArgs
    {
        [Input("aCCOUNT")]
        private InputList<string>? _aCCOUNT;
        public InputList<string> ACCOUNT
        {
            get => _aCCOUNT ?? (_aCCOUNT = new InputList<string>());
            set => _aCCOUNT = value;
        }

        [Input("oRGUNIT")]
        private InputList<string>? _oRGUNIT;
        public InputList<string> ORGUNIT
        {
            get => _oRGUNIT ?? (_oRGUNIT = new InputList<string>());
            set => _oRGUNIT = value;
        }

        public PolicyIEMapArgs()
        {
        }
    }
}
