// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LakeFormation.Inputs
{

    public sealed class PermissionsColumnWildcardArgs : Pulumi.ResourceArgs
    {
        [Input("excludedColumnNames")]
        private InputList<string>? _excludedColumnNames;
        public InputList<string> ExcludedColumnNames
        {
            get => _excludedColumnNames ?? (_excludedColumnNames = new InputList<string>());
            set => _excludedColumnNames = value;
        }

        public PermissionsColumnWildcardArgs()
        {
        }
    }
}
