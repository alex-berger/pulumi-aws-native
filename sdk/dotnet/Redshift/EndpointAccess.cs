// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Redshift
{
    /// <summary>
    /// Resource schema for a Redshift-managed VPC endpoint.
    /// </summary>
    [AwsNativeResourceType("aws-native:redshift:EndpointAccess")]
    public partial class EndpointAccess : Pulumi.CustomResource
    {
        /// <summary>
        /// The DNS address of the endpoint.
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account
        /// </summary>
        [Output("clusterIdentifier")]
        public Output<string?> ClusterIdentifier { get; private set; } = null!;

        /// <summary>
        /// The time (UTC) that the endpoint was created.
        /// </summary>
        [Output("endpointCreateTime")]
        public Output<string> EndpointCreateTime { get; private set; } = null!;

        /// <summary>
        /// The name of the endpoint.
        /// </summary>
        [Output("endpointName")]
        public Output<string> EndpointName { get; private set; } = null!;

        /// <summary>
        /// The status of the endpoint.
        /// </summary>
        [Output("endpointStatus")]
        public Output<string> EndpointStatus { get; private set; } = null!;

        /// <summary>
        /// The port number on which the cluster accepts incoming connections.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// The AWS account ID of the owner of the cluster.
        /// </summary>
        [Output("resourceOwner")]
        public Output<string?> ResourceOwner { get; private set; } = null!;

        /// <summary>
        /// The subnet group name where Amazon Redshift chooses to deploy the endpoint.
        /// </summary>
        [Output("subnetGroupName")]
        public Output<string?> SubnetGroupName { get; private set; } = null!;

        /// <summary>
        /// The connection endpoint for connecting to an Amazon Redshift cluster through the proxy.
        /// </summary>
        [Output("vpcEndpoint")]
        public Output<Outputs.VpcEndpointProperties> VpcEndpoint { get; private set; } = null!;

        /// <summary>
        /// A list of vpc security group ids to apply to the created endpoint access.
        /// </summary>
        [Output("vpcSecurityGroupIds")]
        public Output<ImmutableArray<string>> VpcSecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// A list of Virtual Private Cloud (VPC) security groups to be associated with the endpoint.
        /// </summary>
        [Output("vpcSecurityGroups")]
        public Output<ImmutableArray<Outputs.EndpointAccessVpcSecurityGroup>> VpcSecurityGroups { get; private set; } = null!;


        /// <summary>
        /// Create a EndpointAccess resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EndpointAccess(string name, EndpointAccessArgs args, CustomResourceOptions? options = null)
            : base("aws-native:redshift:EndpointAccess", name, args ?? new EndpointAccessArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EndpointAccess(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:redshift:EndpointAccess", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing EndpointAccess resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EndpointAccess Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EndpointAccess(name, id, options);
        }
    }

    public sealed class EndpointAccessArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account
        /// </summary>
        [Input("clusterIdentifier")]
        public Input<string>? ClusterIdentifier { get; set; }

        /// <summary>
        /// The name of the endpoint.
        /// </summary>
        [Input("endpointName", required: true)]
        public Input<string> EndpointName { get; set; } = null!;

        /// <summary>
        /// The AWS account ID of the owner of the cluster.
        /// </summary>
        [Input("resourceOwner")]
        public Input<string>? ResourceOwner { get; set; }

        /// <summary>
        /// The subnet group name where Amazon Redshift chooses to deploy the endpoint.
        /// </summary>
        [Input("subnetGroupName")]
        public Input<string>? SubnetGroupName { get; set; }

        [Input("vpcSecurityGroupIds", required: true)]
        private InputList<string>? _vpcSecurityGroupIds;

        /// <summary>
        /// A list of vpc security group ids to apply to the created endpoint access.
        /// </summary>
        public InputList<string> VpcSecurityGroupIds
        {
            get => _vpcSecurityGroupIds ?? (_vpcSecurityGroupIds = new InputList<string>());
            set => _vpcSecurityGroupIds = value;
        }

        public EndpointAccessArgs()
        {
        }
    }
}
