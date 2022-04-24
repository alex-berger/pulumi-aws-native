// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FSx.Inputs
{

    public sealed class StorageVirtualMachineSelfManagedActiveDirectoryConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("dnsIps")]
        private InputList<string>? _dnsIps;
        public InputList<string> DnsIps
        {
            get => _dnsIps ?? (_dnsIps = new InputList<string>());
            set => _dnsIps = value;
        }

        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        [Input("fileSystemAdministratorsGroup")]
        public Input<string>? FileSystemAdministratorsGroup { get; set; }

        [Input("organizationalUnitDistinguishedName")]
        public Input<string>? OrganizationalUnitDistinguishedName { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public StorageVirtualMachineSelfManagedActiveDirectoryConfigurationArgs()
        {
        }
    }
}
