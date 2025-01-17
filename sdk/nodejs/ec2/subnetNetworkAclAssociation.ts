// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::SubnetNetworkAclAssociation
 */
export class SubnetNetworkAclAssociation extends pulumi.CustomResource {
    /**
     * Get an existing SubnetNetworkAclAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SubnetNetworkAclAssociation {
        return new SubnetNetworkAclAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:SubnetNetworkAclAssociation';

    /**
     * Returns true if the given object is an instance of SubnetNetworkAclAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SubnetNetworkAclAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SubnetNetworkAclAssociation.__pulumiType;
    }

    public /*out*/ readonly associationId!: pulumi.Output<string>;
    /**
     * The ID of the network ACL
     */
    public readonly networkAclId!: pulumi.Output<string>;
    /**
     * The ID of the subnet
     */
    public readonly subnetId!: pulumi.Output<string>;

    /**
     * Create a SubnetNetworkAclAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetNetworkAclAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.networkAclId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkAclId'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            resourceInputs["networkAclId"] = args ? args.networkAclId : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["associationId"] = undefined /*out*/;
        } else {
            resourceInputs["associationId"] = undefined /*out*/;
            resourceInputs["networkAclId"] = undefined /*out*/;
            resourceInputs["subnetId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SubnetNetworkAclAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SubnetNetworkAclAssociation resource.
 */
export interface SubnetNetworkAclAssociationArgs {
    /**
     * The ID of the network ACL
     */
    networkAclId: pulumi.Input<string>;
    /**
     * The ID of the subnet
     */
    subnetId: pulumi.Input<string>;
}
