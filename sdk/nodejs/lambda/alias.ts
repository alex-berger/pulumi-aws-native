// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Lambda::Alias
 *
 * @deprecated Alias is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Alias extends pulumi.CustomResource {
    /**
     * Get an existing Alias resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Alias {
        pulumi.log.warn("Alias is deprecated: Alias is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Alias(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:lambda:Alias';

    /**
     * Returns true if the given object is an instance of Alias.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Alias {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Alias.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public readonly functionName!: pulumi.Output<string>;
    public readonly functionVersion!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly provisionedConcurrencyConfig!: pulumi.Output<outputs.lambda.AliasProvisionedConcurrencyConfiguration | undefined>;
    public readonly routingConfig!: pulumi.Output<outputs.lambda.AliasAliasRoutingConfiguration | undefined>;

    /**
     * Create a Alias resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Alias is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: AliasArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Alias is deprecated: Alias is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.functionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionName'");
            }
            if ((!args || args.functionVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionVersion'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["functionName"] = args ? args.functionName : undefined;
            inputs["functionVersion"] = args ? args.functionVersion : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["provisionedConcurrencyConfig"] = args ? args.provisionedConcurrencyConfig : undefined;
            inputs["routingConfig"] = args ? args.routingConfig : undefined;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["functionName"] = undefined /*out*/;
            inputs["functionVersion"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisionedConcurrencyConfig"] = undefined /*out*/;
            inputs["routingConfig"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Alias.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Alias resource.
 */
export interface AliasArgs {
    description?: pulumi.Input<string>;
    functionName: pulumi.Input<string>;
    functionVersion: pulumi.Input<string>;
    name: pulumi.Input<string>;
    provisionedConcurrencyConfig?: pulumi.Input<inputs.lambda.AliasProvisionedConcurrencyConfigurationArgs>;
    routingConfig?: pulumi.Input<inputs.lambda.AliasAliasRoutingConfigurationArgs>;
}