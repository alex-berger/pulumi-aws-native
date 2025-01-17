// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Definition of the AWS::QuickSight::Dashboard Resource Type.
 */
export class Dashboard extends pulumi.CustomResource {
    /**
     * Get an existing Dashboard resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Dashboard {
        return new Dashboard(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:quicksight:Dashboard';

    /**
     * Returns true if the given object is an instance of Dashboard.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Dashboard {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Dashboard.__pulumiType;
    }

    /**
     * <p>The Amazon Resource Name (ARN) of the resource.</p>
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly awsAccountId!: pulumi.Output<string>;
    /**
     * <p>The time that this dataset was created.</p>
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    public readonly dashboardId!: pulumi.Output<string>;
    public readonly dashboardPublishOptions!: pulumi.Output<outputs.quicksight.DashboardPublishOptions | undefined>;
    /**
     * <p>The last time that this dataset was published.</p>
     */
    public /*out*/ readonly lastPublishedTime!: pulumi.Output<string>;
    /**
     * <p>The last time that this dataset was updated.</p>
     */
    public /*out*/ readonly lastUpdatedTime!: pulumi.Output<string>;
    /**
     * <p>The display name of the dashboard.</p>
     */
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly parameters!: pulumi.Output<outputs.quicksight.DashboardParameters | undefined>;
    /**
     * <p>A structure that contains the permissions of the dashboard. You can use this structure
     *             for granting permissions by providing a list of IAM action information for each
     *             principal ARN. </p>
     *
     *         <p>To specify no permissions, omit the permissions list.</p>
     */
    public readonly permissions!: pulumi.Output<outputs.quicksight.DashboardResourcePermission[] | undefined>;
    public readonly sourceEntity!: pulumi.Output<outputs.quicksight.DashboardSourceEntity>;
    /**
     * <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the
     *             dashboard.</p>
     */
    public readonly tags!: pulumi.Output<outputs.quicksight.DashboardTag[] | undefined>;
    /**
     * <p>The Amazon Resource Name (ARN) of the theme that is being used for this dashboard. If
     *             you add a value for this field, it overrides the value that is used in the source
     *             entity. The theme ARN must exist in the same AWS account where you create the
     *             dashboard.</p>
     */
    public readonly themeArn!: pulumi.Output<string | undefined>;
    public /*out*/ readonly version!: pulumi.Output<outputs.quicksight.DashboardVersion>;
    /**
     * <p>A description for the first version of the dashboard being created.</p>
     */
    public readonly versionDescription!: pulumi.Output<string | undefined>;

    /**
     * Create a Dashboard resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DashboardArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.awsAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'awsAccountId'");
            }
            if ((!args || args.dashboardId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dashboardId'");
            }
            if ((!args || args.sourceEntity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceEntity'");
            }
            resourceInputs["awsAccountId"] = args ? args.awsAccountId : undefined;
            resourceInputs["dashboardId"] = args ? args.dashboardId : undefined;
            resourceInputs["dashboardPublishOptions"] = args ? args.dashboardPublishOptions : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["sourceEntity"] = args ? args.sourceEntity : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["themeArn"] = args ? args.themeArn : undefined;
            resourceInputs["versionDescription"] = args ? args.versionDescription : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["lastPublishedTime"] = undefined /*out*/;
            resourceInputs["lastUpdatedTime"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsAccountId"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["dashboardId"] = undefined /*out*/;
            resourceInputs["dashboardPublishOptions"] = undefined /*out*/;
            resourceInputs["lastPublishedTime"] = undefined /*out*/;
            resourceInputs["lastUpdatedTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parameters"] = undefined /*out*/;
            resourceInputs["permissions"] = undefined /*out*/;
            resourceInputs["sourceEntity"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["themeArn"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
            resourceInputs["versionDescription"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Dashboard.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Dashboard resource.
 */
export interface DashboardArgs {
    awsAccountId: pulumi.Input<string>;
    dashboardId: pulumi.Input<string>;
    dashboardPublishOptions?: pulumi.Input<inputs.quicksight.DashboardPublishOptionsArgs>;
    /**
     * <p>The display name of the dashboard.</p>
     */
    name?: pulumi.Input<string>;
    parameters?: pulumi.Input<inputs.quicksight.DashboardParametersArgs>;
    /**
     * <p>A structure that contains the permissions of the dashboard. You can use this structure
     *             for granting permissions by providing a list of IAM action information for each
     *             principal ARN. </p>
     *
     *         <p>To specify no permissions, omit the permissions list.</p>
     */
    permissions?: pulumi.Input<pulumi.Input<inputs.quicksight.DashboardResourcePermissionArgs>[]>;
    sourceEntity: pulumi.Input<inputs.quicksight.DashboardSourceEntityArgs>;
    /**
     * <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the
     *             dashboard.</p>
     */
    tags?: pulumi.Input<pulumi.Input<inputs.quicksight.DashboardTagArgs>[]>;
    /**
     * <p>The Amazon Resource Name (ARN) of the theme that is being used for this dashboard. If
     *             you add a value for this field, it overrides the value that is used in the source
     *             entity. The theme ARN must exist in the same AWS account where you create the
     *             dashboard.</p>
     */
    themeArn?: pulumi.Input<string>;
    /**
     * <p>A description for the first version of the dashboard being created.</p>
     */
    versionDescription?: pulumi.Input<string>;
}
