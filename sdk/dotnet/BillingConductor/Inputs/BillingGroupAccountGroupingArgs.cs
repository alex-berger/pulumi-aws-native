// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.BillingConductor.Inputs
{

    public sealed class BillingGroupAccountGroupingArgs : Pulumi.ResourceArgs
    {
        [Input("linkedAccountIds", required: true)]
        private InputList<string>? _linkedAccountIds;
        public InputList<string> LinkedAccountIds
        {
            get => _linkedAccountIds ?? (_linkedAccountIds = new InputList<string>());
            set => _linkedAccountIds = value;
        }

        public BillingGroupAccountGroupingArgs()
        {
        }
    }
}
