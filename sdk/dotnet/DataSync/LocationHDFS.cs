// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync
{
    /// <summary>
    /// Resource schema for AWS::DataSync::LocationHDFS.
    /// </summary>
    [AwsNativeResourceType("aws-native:datasync:LocationHDFS")]
    public partial class LocationHDFS : Pulumi.CustomResource
    {
        /// <summary>
        /// ARN(s) of the agent(s) to use for an HDFS location.
        /// </summary>
        [Output("agentArns")]
        public Output<ImmutableArray<string>> AgentArns { get; private set; } = null!;

        /// <summary>
        /// The authentication mode used to determine identity of user.
        /// </summary>
        [Output("authenticationType")]
        public Output<Pulumi.AwsNative.DataSync.LocationHDFSAuthenticationType> AuthenticationType { get; private set; } = null!;

        /// <summary>
        /// Size of chunks (blocks) in bytes that the data is divided into when stored in the HDFS cluster.
        /// </summary>
        [Output("blockSize")]
        public Output<int?> BlockSize { get; private set; } = null!;

        /// <summary>
        /// The Base64 string representation of the Keytab file.
        /// </summary>
        [Output("kerberosKeytab")]
        public Output<string?> KerberosKeytab { get; private set; } = null!;

        /// <summary>
        /// The string representation of the Krb5Conf file, or the presigned URL to access the Krb5.conf file within an S3 bucket.
        /// </summary>
        [Output("kerberosKrb5Conf")]
        public Output<string?> KerberosKrb5Conf { get; private set; } = null!;

        /// <summary>
        /// The unique identity, or principal, to which Kerberos can assign tickets.
        /// </summary>
        [Output("kerberosPrincipal")]
        public Output<string?> KerberosPrincipal { get; private set; } = null!;

        /// <summary>
        /// The identifier for the Key Management Server where the encryption keys that encrypt data inside HDFS clusters are stored.
        /// </summary>
        [Output("kmsKeyProviderUri")]
        public Output<string?> KmsKeyProviderUri { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the HDFS location.
        /// </summary>
        [Output("locationArn")]
        public Output<string> LocationArn { get; private set; } = null!;

        /// <summary>
        /// The URL of the HDFS location that was described.
        /// </summary>
        [Output("locationUri")]
        public Output<string> LocationUri { get; private set; } = null!;

        /// <summary>
        /// An array of Name Node(s) of the HDFS location.
        /// </summary>
        [Output("nameNodes")]
        public Output<ImmutableArray<Outputs.LocationHDFSNameNode>> NameNodes { get; private set; } = null!;

        [Output("qopConfiguration")]
        public Output<Outputs.LocationHDFSQopConfiguration?> QopConfiguration { get; private set; } = null!;

        /// <summary>
        /// Number of copies of each block that exists inside the HDFS cluster.
        /// </summary>
        [Output("replicationFactor")]
        public Output<int?> ReplicationFactor { get; private set; } = null!;

        /// <summary>
        /// The user name that has read and write permissions on the specified HDFS cluster.
        /// </summary>
        [Output("simpleUser")]
        public Output<string?> SimpleUser { get; private set; } = null!;

        /// <summary>
        /// The subdirectory in HDFS that is used to read data from the HDFS source location or write data to the HDFS destination.
        /// </summary>
        [Output("subdirectory")]
        public Output<string?> Subdirectory { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.LocationHDFSTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a LocationHDFS resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LocationHDFS(string name, LocationHDFSArgs args, CustomResourceOptions? options = null)
            : base("aws-native:datasync:LocationHDFS", name, args ?? new LocationHDFSArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LocationHDFS(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:datasync:LocationHDFS", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LocationHDFS resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LocationHDFS Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LocationHDFS(name, id, options);
        }
    }

    public sealed class LocationHDFSArgs : Pulumi.ResourceArgs
    {
        [Input("agentArns", required: true)]
        private InputList<string>? _agentArns;

        /// <summary>
        /// ARN(s) of the agent(s) to use for an HDFS location.
        /// </summary>
        public InputList<string> AgentArns
        {
            get => _agentArns ?? (_agentArns = new InputList<string>());
            set => _agentArns = value;
        }

        /// <summary>
        /// The authentication mode used to determine identity of user.
        /// </summary>
        [Input("authenticationType", required: true)]
        public Input<Pulumi.AwsNative.DataSync.LocationHDFSAuthenticationType> AuthenticationType { get; set; } = null!;

        /// <summary>
        /// Size of chunks (blocks) in bytes that the data is divided into when stored in the HDFS cluster.
        /// </summary>
        [Input("blockSize")]
        public Input<int>? BlockSize { get; set; }

        /// <summary>
        /// The Base64 string representation of the Keytab file.
        /// </summary>
        [Input("kerberosKeytab")]
        public Input<string>? KerberosKeytab { get; set; }

        /// <summary>
        /// The string representation of the Krb5Conf file, or the presigned URL to access the Krb5.conf file within an S3 bucket.
        /// </summary>
        [Input("kerberosKrb5Conf")]
        public Input<string>? KerberosKrb5Conf { get; set; }

        /// <summary>
        /// The unique identity, or principal, to which Kerberos can assign tickets.
        /// </summary>
        [Input("kerberosPrincipal")]
        public Input<string>? KerberosPrincipal { get; set; }

        /// <summary>
        /// The identifier for the Key Management Server where the encryption keys that encrypt data inside HDFS clusters are stored.
        /// </summary>
        [Input("kmsKeyProviderUri")]
        public Input<string>? KmsKeyProviderUri { get; set; }

        [Input("nameNodes", required: true)]
        private InputList<Inputs.LocationHDFSNameNodeArgs>? _nameNodes;

        /// <summary>
        /// An array of Name Node(s) of the HDFS location.
        /// </summary>
        public InputList<Inputs.LocationHDFSNameNodeArgs> NameNodes
        {
            get => _nameNodes ?? (_nameNodes = new InputList<Inputs.LocationHDFSNameNodeArgs>());
            set => _nameNodes = value;
        }

        [Input("qopConfiguration")]
        public Input<Inputs.LocationHDFSQopConfigurationArgs>? QopConfiguration { get; set; }

        /// <summary>
        /// Number of copies of each block that exists inside the HDFS cluster.
        /// </summary>
        [Input("replicationFactor")]
        public Input<int>? ReplicationFactor { get; set; }

        /// <summary>
        /// The user name that has read and write permissions on the specified HDFS cluster.
        /// </summary>
        [Input("simpleUser")]
        public Input<string>? SimpleUser { get; set; }

        /// <summary>
        /// The subdirectory in HDFS that is used to read data from the HDFS source location or write data to the HDFS destination.
        /// </summary>
        [Input("subdirectory")]
        public Input<string>? Subdirectory { get; set; }

        [Input("tags")]
        private InputList<Inputs.LocationHDFSTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.LocationHDFSTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.LocationHDFSTagArgs>());
            set => _tags = value;
        }

        public LocationHDFSArgs()
        {
        }
    }
}
