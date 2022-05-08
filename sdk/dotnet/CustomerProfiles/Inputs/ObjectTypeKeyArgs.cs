// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles.Inputs
{

    /// <summary>
    /// An object that defines the Key element of a ProfileObject. A Key is a special element that can be used to search for a customer profile.
    /// </summary>
    public sealed class ObjectTypeKeyArgs : Pulumi.ResourceArgs
    {
        [Input("fieldNames")]
        private InputList<string>? _fieldNames;

        /// <summary>
        /// The reference for the key name of the fields map. 
        /// </summary>
        public InputList<string> FieldNames
        {
            get => _fieldNames ?? (_fieldNames = new InputList<string>());
            set => _fieldNames = value;
        }

        [Input("standardIdentifiers")]
        private InputList<Pulumi.AwsNative.CustomerProfiles.ObjectTypeKeyStandardIdentifiersItem>? _standardIdentifiers;

        /// <summary>
        /// The types of keys that a ProfileObject can have. Each ProfileObject can have only 1 UNIQUE key but multiple PROFILE keys. PROFILE means that this key can be used to tie an object to a PROFILE. UNIQUE means that it can be used to uniquely identify an object. If a key a is marked as SECONDARY, it will be used to search for profiles after all other PROFILE keys have been searched. A LOOKUP_ONLY key is only used to match a profile but is not persisted to be used for searching of the profile. A NEW_ONLY key is only used if the profile does not already exist before the object is ingested, otherwise it is only used for matching objects to profiles.
        /// </summary>
        public InputList<Pulumi.AwsNative.CustomerProfiles.ObjectTypeKeyStandardIdentifiersItem> StandardIdentifiers
        {
            get => _standardIdentifiers ?? (_standardIdentifiers = new InputList<Pulumi.AwsNative.CustomerProfiles.ObjectTypeKeyStandardIdentifiersItem>());
            set => _standardIdentifiers = value;
        }

        public ObjectTypeKeyArgs()
        {
        }
    }
}
