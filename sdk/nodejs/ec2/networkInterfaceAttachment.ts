// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::NetworkInterfaceAttachment
 *
 * @deprecated NetworkInterfaceAttachment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class NetworkInterfaceAttachment extends pulumi.CustomResource {
    /**
     * Get an existing NetworkInterfaceAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NetworkInterfaceAttachment {
        pulumi.log.warn("NetworkInterfaceAttachment is deprecated: NetworkInterfaceAttachment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new NetworkInterfaceAttachment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:NetworkInterfaceAttachment';

    /**
     * Returns true if the given object is an instance of NetworkInterfaceAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkInterfaceAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkInterfaceAttachment.__pulumiType;
    }

    public readonly deleteOnTermination!: pulumi.Output<boolean | undefined>;
    public readonly deviceIndex!: pulumi.Output<string>;
    public readonly instanceId!: pulumi.Output<string>;
    public readonly networkInterfaceId!: pulumi.Output<string>;

    /**
     * Create a NetworkInterfaceAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated NetworkInterfaceAttachment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: NetworkInterfaceAttachmentArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("NetworkInterfaceAttachment is deprecated: NetworkInterfaceAttachment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.deviceIndex === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deviceIndex'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.networkInterfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkInterfaceId'");
            }
            resourceInputs["deleteOnTermination"] = args ? args.deleteOnTermination : undefined;
            resourceInputs["deviceIndex"] = args ? args.deviceIndex : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["networkInterfaceId"] = args ? args.networkInterfaceId : undefined;
        } else {
            resourceInputs["deleteOnTermination"] = undefined /*out*/;
            resourceInputs["deviceIndex"] = undefined /*out*/;
            resourceInputs["instanceId"] = undefined /*out*/;
            resourceInputs["networkInterfaceId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkInterfaceAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a NetworkInterfaceAttachment resource.
 */
export interface NetworkInterfaceAttachmentArgs {
    deleteOnTermination?: pulumi.Input<boolean>;
    deviceIndex: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    networkInterfaceId: pulumi.Input<string>;
}
