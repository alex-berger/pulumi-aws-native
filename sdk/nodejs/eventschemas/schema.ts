// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EventSchemas::Schema
 *
 * @deprecated Schema is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Schema extends pulumi.CustomResource {
    /**
     * Get an existing Schema resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Schema {
        pulumi.log.warn("Schema is deprecated: Schema is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Schema(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:eventschemas:Schema';

    /**
     * Returns true if the given object is an instance of Schema.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Schema {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Schema.__pulumiType;
    }

    public readonly content!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly registryName!: pulumi.Output<string>;
    public /*out*/ readonly schemaArn!: pulumi.Output<string>;
    public readonly schemaName!: pulumi.Output<string | undefined>;
    public /*out*/ readonly schemaVersion!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.eventschemas.SchemaTagsEntry[] | undefined>;
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Schema resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Schema is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: SchemaArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Schema is deprecated: Schema is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            if ((!args || args.registryName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registryName'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["registryName"] = args ? args.registryName : undefined;
            resourceInputs["schemaName"] = args ? args.schemaName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["schemaArn"] = undefined /*out*/;
            resourceInputs["schemaVersion"] = undefined /*out*/;
        } else {
            resourceInputs["content"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["registryName"] = undefined /*out*/;
            resourceInputs["schemaArn"] = undefined /*out*/;
            resourceInputs["schemaName"] = undefined /*out*/;
            resourceInputs["schemaVersion"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Schema.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Schema resource.
 */
export interface SchemaArgs {
    content: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    registryName: pulumi.Input<string>;
    schemaName?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.eventschemas.SchemaTagsEntryArgs>[]>;
    type: pulumi.Input<string>;
}
