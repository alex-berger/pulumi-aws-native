// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::IoTTwinMaker::ComponentType
 */
export class ComponentType extends pulumi.CustomResource {
    /**
     * Get an existing ComponentType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ComponentType {
        return new ComponentType(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iottwinmaker:ComponentType';

    /**
     * Returns true if the given object is an instance of ComponentType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComponentType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComponentType.__pulumiType;
    }

    /**
     * The ARN of the component type.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The ID of the component type.
     */
    public readonly componentTypeId!: pulumi.Output<string>;
    /**
     * The date and time when the component type was created.
     */
    public /*out*/ readonly creationDateTime!: pulumi.Output<string>;
    /**
     * The description of the component type.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the parent component type to extend.
     */
    public readonly extendsFrom!: pulumi.Output<string[] | undefined>;
    /**
     * a Map of functions in the component type. Each function's key must be unique to this map.
     */
    public readonly functions!: pulumi.Output<any | undefined>;
    /**
     * A Boolean value that specifies whether the component type is abstract.
     */
    public /*out*/ readonly isAbstract!: pulumi.Output<boolean>;
    /**
     * A Boolean value that specifies whether the component type has a schema initializer and that the schema initializer has run.
     */
    public /*out*/ readonly isSchemaInitialized!: pulumi.Output<boolean>;
    /**
     * A Boolean value that specifies whether an entity can have more than one component of this type.
     */
    public readonly isSingleton!: pulumi.Output<boolean | undefined>;
    /**
     * An map of the property definitions in the component type. Each property definition's key must be unique to this map.
     */
    public readonly propertyDefinitions!: pulumi.Output<any | undefined>;
    /**
     * The current status of the component type.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.iottwinmaker.ComponentTypeStatus>;
    /**
     * A map of key-value pairs to associate with a resource.
     */
    public readonly tags!: pulumi.Output<any | undefined>;
    /**
     * The last date and time when the component type was updated.
     */
    public /*out*/ readonly updateDateTime!: pulumi.Output<string>;
    /**
     * The ID of the workspace that contains the component type.
     */
    public readonly workspaceId!: pulumi.Output<string>;

    /**
     * Create a ComponentType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComponentTypeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.componentTypeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'componentTypeId'");
            }
            if ((!args || args.workspaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceId'");
            }
            resourceInputs["componentTypeId"] = args ? args.componentTypeId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["extendsFrom"] = args ? args.extendsFrom : undefined;
            resourceInputs["functions"] = args ? args.functions : undefined;
            resourceInputs["isSingleton"] = args ? args.isSingleton : undefined;
            resourceInputs["propertyDefinitions"] = args ? args.propertyDefinitions : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["workspaceId"] = args ? args.workspaceId : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["creationDateTime"] = undefined /*out*/;
            resourceInputs["isAbstract"] = undefined /*out*/;
            resourceInputs["isSchemaInitialized"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updateDateTime"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["componentTypeId"] = undefined /*out*/;
            resourceInputs["creationDateTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["extendsFrom"] = undefined /*out*/;
            resourceInputs["functions"] = undefined /*out*/;
            resourceInputs["isAbstract"] = undefined /*out*/;
            resourceInputs["isSchemaInitialized"] = undefined /*out*/;
            resourceInputs["isSingleton"] = undefined /*out*/;
            resourceInputs["propertyDefinitions"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["updateDateTime"] = undefined /*out*/;
            resourceInputs["workspaceId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ComponentType.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ComponentType resource.
 */
export interface ComponentTypeArgs {
    /**
     * The ID of the component type.
     */
    componentTypeId: pulumi.Input<string>;
    /**
     * The description of the component type.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the parent component type to extend.
     */
    extendsFrom?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * a Map of functions in the component type. Each function's key must be unique to this map.
     */
    functions?: any;
    /**
     * A Boolean value that specifies whether an entity can have more than one component of this type.
     */
    isSingleton?: pulumi.Input<boolean>;
    /**
     * An map of the property definitions in the component type. Each property definition's key must be unique to this map.
     */
    propertyDefinitions?: any;
    /**
     * A map of key-value pairs to associate with a resource.
     */
    tags?: any;
    /**
     * The ID of the workspace that contains the component type.
     */
    workspaceId: pulumi.Input<string>;
}