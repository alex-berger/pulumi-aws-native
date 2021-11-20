// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Associates a set of DHCP options with a VPC, or associates no DHCP options with the VPC.
 */
export class VPCDHCPOptionsAssociation extends pulumi.CustomResource {
    /**
     * Get an existing VPCDHCPOptionsAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VPCDHCPOptionsAssociation {
        return new VPCDHCPOptionsAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:VPCDHCPOptionsAssociation';

    /**
     * Returns true if the given object is an instance of VPCDHCPOptionsAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VPCDHCPOptionsAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VPCDHCPOptionsAssociation.__pulumiType;
    }

    /**
     * The ID of the DHCP options set, or default to associate no DHCP options with the VPC.
     */
    public readonly dhcpOptionsId!: pulumi.Output<string>;
    /**
     * The ID of the VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a VPCDHCPOptionsAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VPCDHCPOptionsAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.dhcpOptionsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dhcpOptionsId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            inputs["dhcpOptionsId"] = args ? args.dhcpOptionsId : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
        } else {
            inputs["dhcpOptionsId"] = undefined /*out*/;
            inputs["vpcId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VPCDHCPOptionsAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VPCDHCPOptionsAssociation resource.
 */
export interface VPCDHCPOptionsAssociationArgs {
    /**
     * The ID of the DHCP options set, or default to associate no DHCP options with the VPC.
     */
    dhcpOptionsId: pulumi.Input<string>;
    /**
     * The ID of the VPC.
     */
    vpcId: pulumi.Input<string>;
}
