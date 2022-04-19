// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::EIPAssociation
 *
 * @deprecated EIPAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class EIPAssociation extends pulumi.CustomResource {
    /**
     * Get an existing EIPAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EIPAssociation {
        pulumi.log.warn("EIPAssociation is deprecated: EIPAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new EIPAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:EIPAssociation';

    /**
     * Returns true if the given object is an instance of EIPAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EIPAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EIPAssociation.__pulumiType;
    }

    public readonly allocationId!: pulumi.Output<string | undefined>;
    public readonly eIP!: pulumi.Output<string | undefined>;
    public readonly instanceId!: pulumi.Output<string | undefined>;
    public readonly networkInterfaceId!: pulumi.Output<string | undefined>;
    public readonly privateIpAddress!: pulumi.Output<string | undefined>;

    /**
     * Create a EIPAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated EIPAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: EIPAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("EIPAssociation is deprecated: EIPAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["allocationId"] = args ? args.allocationId : undefined;
            resourceInputs["eIP"] = args ? args.eIP : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["networkInterfaceId"] = args ? args.networkInterfaceId : undefined;
            resourceInputs["privateIpAddress"] = args ? args.privateIpAddress : undefined;
        } else {
            resourceInputs["allocationId"] = undefined /*out*/;
            resourceInputs["eIP"] = undefined /*out*/;
            resourceInputs["instanceId"] = undefined /*out*/;
            resourceInputs["networkInterfaceId"] = undefined /*out*/;
            resourceInputs["privateIpAddress"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EIPAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EIPAssociation resource.
 */
export interface EIPAssociationArgs {
    allocationId?: pulumi.Input<string>;
    eIP?: pulumi.Input<string>;
    instanceId?: pulumi.Input<string>;
    networkInterfaceId?: pulumi.Input<string>;
    privateIpAddress?: pulumi.Input<string>;
}
