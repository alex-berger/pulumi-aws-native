// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceCatalogAppRegistry
{
    public static class GetAttributeGroupAssociation
    {
        /// <summary>
        /// Resource Schema for AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation.
        /// </summary>
        public static Task<GetAttributeGroupAssociationResult> InvokeAsync(GetAttributeGroupAssociationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAttributeGroupAssociationResult>("aws-native:servicecatalogappregistry:getAttributeGroupAssociation", args ?? new GetAttributeGroupAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Schema for AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation.
        /// </summary>
        public static Output<GetAttributeGroupAssociationResult> Invoke(GetAttributeGroupAssociationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAttributeGroupAssociationResult>("aws-native:servicecatalogappregistry:getAttributeGroupAssociation", args ?? new GetAttributeGroupAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAttributeGroupAssociationArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetAttributeGroupAssociationArgs()
        {
        }
    }

    public sealed class GetAttributeGroupAssociationInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetAttributeGroupAssociationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAttributeGroupAssociationResult
    {
        /// <summary>
        /// The name or the Id of the Application.
        /// </summary>
        public readonly string? Application;
        public readonly string? ApplicationArn;
        /// <summary>
        /// The name or the Id of the AttributeGroup.
        /// </summary>
        public readonly string? AttributeGroup;
        public readonly string? AttributeGroupArn;
        public readonly string? Id;

        [OutputConstructor]
        private GetAttributeGroupAssociationResult(
            string? application,

            string? applicationArn,

            string? attributeGroup,

            string? attributeGroupArn,

            string? id)
        {
            Application = application;
            ApplicationArn = applicationArn;
            AttributeGroup = attributeGroup;
            AttributeGroupArn = attributeGroupArn;
            Id = id;
        }
    }
}