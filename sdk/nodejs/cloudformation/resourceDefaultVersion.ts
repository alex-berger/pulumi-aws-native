// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The default version of a resource that has been registered in the CloudFormation Registry.
 */
export class ResourceDefaultVersion extends pulumi.CustomResource {
    /**
     * Get an existing ResourceDefaultVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ResourceDefaultVersion {
        return new ResourceDefaultVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cloudformation:ResourceDefaultVersion';

    /**
     * Returns true if the given object is an instance of ResourceDefaultVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourceDefaultVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourceDefaultVersion.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a ResourceDefaultVersion
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of the type being registered.
     *
     * We recommend that type names adhere to the following pattern: company_or_organization::service::type.
     */
    public readonly typeName!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the type version.
     */
    public readonly typeVersionArn!: pulumi.Output<string | undefined>;
    /**
     * The ID of an existing version of the resource to set as the default.
     */
    public readonly versionId!: pulumi.Output<string | undefined>;

    /**
     * Create a ResourceDefaultVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ResourceDefaultVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["typeName"] = args ? args.typeName : undefined;
            resourceInputs["typeVersionArn"] = args ? args.typeVersionArn : undefined;
            resourceInputs["versionId"] = args ? args.versionId : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["typeName"] = undefined /*out*/;
            resourceInputs["typeVersionArn"] = undefined /*out*/;
            resourceInputs["versionId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ResourceDefaultVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ResourceDefaultVersion resource.
 */
export interface ResourceDefaultVersionArgs {
    /**
     * The name of the type being registered.
     *
     * We recommend that type names adhere to the following pattern: company_or_organization::service::type.
     */
    typeName?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the type version.
     */
    typeVersionArn?: pulumi.Input<string>;
    /**
     * The ID of an existing version of the resource to set as the default.
     */
    versionId?: pulumi.Input<string>;
}
