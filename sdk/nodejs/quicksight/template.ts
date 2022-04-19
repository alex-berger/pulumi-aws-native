// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Definition of the AWS::QuickSight::Template Resource Type.
 */
export class Template extends pulumi.CustomResource {
    /**
     * Get an existing Template resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Template {
        return new Template(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:quicksight:Template';

    /**
     * Returns true if the given object is an instance of Template.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Template {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Template.__pulumiType;
    }

    /**
     * <p>The Amazon Resource Name (ARN) of the template.</p>
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly awsAccountId!: pulumi.Output<string>;
    /**
     * <p>Time when this was created.</p>
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * <p>Time when this was last updated.</p>
     */
    public /*out*/ readonly lastUpdatedTime!: pulumi.Output<string>;
    /**
     * <p>A display name for the template.</p>
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * <p>A list of resource permissions to be set on the template. </p>
     */
    public readonly permissions!: pulumi.Output<outputs.quicksight.TemplateResourcePermission[] | undefined>;
    public readonly sourceEntity!: pulumi.Output<outputs.quicksight.TemplateSourceEntity>;
    /**
     * <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.</p>
     */
    public readonly tags!: pulumi.Output<outputs.quicksight.TemplateTag[] | undefined>;
    public readonly templateId!: pulumi.Output<string>;
    public /*out*/ readonly version!: pulumi.Output<outputs.quicksight.TemplateVersion>;
    /**
     * <p>A description of the current template version being created. This API operation creates the
     * 			first version of the template. Every time <code>UpdateTemplate</code> is called, a new
     * 			version is created. Each version of the template maintains a description of the version
     * 			in the <code>VersionDescription</code> field.</p>
     */
    public readonly versionDescription!: pulumi.Output<string | undefined>;

    /**
     * Create a Template resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.awsAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'awsAccountId'");
            }
            if ((!args || args.sourceEntity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceEntity'");
            }
            if ((!args || args.templateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateId'");
            }
            resourceInputs["awsAccountId"] = args ? args.awsAccountId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["sourceEntity"] = args ? args.sourceEntity : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["templateId"] = args ? args.templateId : undefined;
            resourceInputs["versionDescription"] = args ? args.versionDescription : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["lastUpdatedTime"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsAccountId"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["lastUpdatedTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["permissions"] = undefined /*out*/;
            resourceInputs["sourceEntity"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["templateId"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
            resourceInputs["versionDescription"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Template.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Template resource.
 */
export interface TemplateArgs {
    awsAccountId: pulumi.Input<string>;
    /**
     * <p>A display name for the template.</p>
     */
    name?: pulumi.Input<string>;
    /**
     * <p>A list of resource permissions to be set on the template. </p>
     */
    permissions?: pulumi.Input<pulumi.Input<inputs.quicksight.TemplateResourcePermissionArgs>[]>;
    sourceEntity: pulumi.Input<inputs.quicksight.TemplateSourceEntityArgs>;
    /**
     * <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.</p>
     */
    tags?: pulumi.Input<pulumi.Input<inputs.quicksight.TemplateTagArgs>[]>;
    templateId: pulumi.Input<string>;
    /**
     * <p>A description of the current template version being created. This API operation creates the
     * 			first version of the template. Every time <code>UpdateTemplate</code> is called, a new
     * 			version is created. Each version of the template maintains a description of the version
     * 			in the <code>VersionDescription</code> field.</p>
     */
    versionDescription?: pulumi.Input<string>;
}
