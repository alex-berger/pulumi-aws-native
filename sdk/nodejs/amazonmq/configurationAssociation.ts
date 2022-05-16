// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AmazonMQ::ConfigurationAssociation
 *
 * @deprecated ConfigurationAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class ConfigurationAssociation extends pulumi.CustomResource {
    /**
     * Get an existing ConfigurationAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConfigurationAssociation {
        pulumi.log.warn("ConfigurationAssociation is deprecated: ConfigurationAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new ConfigurationAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:amazonmq:ConfigurationAssociation';

    /**
     * Returns true if the given object is an instance of ConfigurationAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigurationAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigurationAssociation.__pulumiType;
    }

    public readonly broker!: pulumi.Output<string>;
    public readonly configuration!: pulumi.Output<outputs.amazonmq.ConfigurationAssociationConfigurationId>;

    /**
     * Create a ConfigurationAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated ConfigurationAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: ConfigurationAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ConfigurationAssociation is deprecated: ConfigurationAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.broker === undefined) && !opts.urn) {
                throw new Error("Missing required property 'broker'");
            }
            if ((!args || args.configuration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configuration'");
            }
            resourceInputs["broker"] = args ? args.broker : undefined;
            resourceInputs["configuration"] = args ? args.configuration : undefined;
        } else {
            resourceInputs["broker"] = undefined /*out*/;
            resourceInputs["configuration"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConfigurationAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConfigurationAssociation resource.
 */
export interface ConfigurationAssociationArgs {
    broker: pulumi.Input<string>;
    configuration: pulumi.Input<inputs.amazonmq.ConfigurationAssociationConfigurationIdArgs>;
}