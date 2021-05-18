// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html
 */
export class Grant extends pulumi.CustomResource {
    /**
     * Get an existing Grant resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Grant {
        return new Grant(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:LicenseManager:Grant';

    /**
     * Returns true if the given object is an instance of Grant.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Grant {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Grant.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-allowedoperations
     */
    public readonly allowedOperations!: pulumi.Output<outputs.LicenseManager.GrantAllowedOperationList | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-clienttoken
     */
    public readonly clientToken!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-filters
     */
    public readonly filters!: pulumi.Output<outputs.LicenseManager.GrantFilterList | undefined>;
    public /*out*/ readonly grantArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantarns
     */
    public readonly grantArns!: pulumi.Output<outputs.LicenseManager.GrantArnList | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantname
     */
    public readonly grantName!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantstatus
     */
    public readonly grantStatus!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantedoperations
     */
    public readonly grantedOperations!: pulumi.Output<outputs.LicenseManager.GrantAllowedOperationList | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-granteeprincipalarn
     */
    public readonly granteePrincipalArn!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-homeregion
     */
    public readonly homeRegion!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-licensearn
     */
    public readonly licenseArn!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-maxresults
     */
    public readonly maxResults!: pulumi.Output<number | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-nexttoken
     */
    public readonly nextToken!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-parentarn
     */
    public readonly parentArn!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-principals
     */
    public readonly principals!: pulumi.Output<outputs.LicenseManager.GrantArnList | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-sourceversion
     */
    public readonly sourceVersion!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-status
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-statusreason
     */
    public readonly statusReason!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-tags
     */
    public readonly tags!: pulumi.Output<outputs.LicenseManager.GrantTagList | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-version
     */
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a Grant resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GrantArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["allowedOperations"] = args ? args.allowedOperations : undefined;
            inputs["clientToken"] = args ? args.clientToken : undefined;
            inputs["filters"] = args ? args.filters : undefined;
            inputs["grantArns"] = args ? args.grantArns : undefined;
            inputs["grantName"] = args ? args.grantName : undefined;
            inputs["grantStatus"] = args ? args.grantStatus : undefined;
            inputs["grantedOperations"] = args ? args.grantedOperations : undefined;
            inputs["granteePrincipalArn"] = args ? args.granteePrincipalArn : undefined;
            inputs["homeRegion"] = args ? args.homeRegion : undefined;
            inputs["licenseArn"] = args ? args.licenseArn : undefined;
            inputs["maxResults"] = args ? args.maxResults : undefined;
            inputs["nextToken"] = args ? args.nextToken : undefined;
            inputs["parentArn"] = args ? args.parentArn : undefined;
            inputs["principals"] = args ? args.principals : undefined;
            inputs["sourceVersion"] = args ? args.sourceVersion : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["statusReason"] = args ? args.statusReason : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["grantArn"] = undefined /*out*/;
        } else {
            inputs["allowedOperations"] = undefined /*out*/;
            inputs["clientToken"] = undefined /*out*/;
            inputs["filters"] = undefined /*out*/;
            inputs["grantArn"] = undefined /*out*/;
            inputs["grantArns"] = undefined /*out*/;
            inputs["grantName"] = undefined /*out*/;
            inputs["grantStatus"] = undefined /*out*/;
            inputs["grantedOperations"] = undefined /*out*/;
            inputs["granteePrincipalArn"] = undefined /*out*/;
            inputs["homeRegion"] = undefined /*out*/;
            inputs["licenseArn"] = undefined /*out*/;
            inputs["maxResults"] = undefined /*out*/;
            inputs["nextToken"] = undefined /*out*/;
            inputs["parentArn"] = undefined /*out*/;
            inputs["principals"] = undefined /*out*/;
            inputs["sourceVersion"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusReason"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Grant.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Grant resource.
 */
export interface GrantArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-allowedoperations
     */
    readonly allowedOperations?: pulumi.Input<inputs.LicenseManager.GrantAllowedOperationListArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-clienttoken
     */
    readonly clientToken?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-filters
     */
    readonly filters?: pulumi.Input<inputs.LicenseManager.GrantFilterListArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantarns
     */
    readonly grantArns?: pulumi.Input<inputs.LicenseManager.GrantArnListArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantname
     */
    readonly grantName?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantstatus
     */
    readonly grantStatus?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantedoperations
     */
    readonly grantedOperations?: pulumi.Input<inputs.LicenseManager.GrantAllowedOperationListArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-granteeprincipalarn
     */
    readonly granteePrincipalArn?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-homeregion
     */
    readonly homeRegion?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-licensearn
     */
    readonly licenseArn?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-maxresults
     */
    readonly maxResults?: pulumi.Input<number>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-nexttoken
     */
    readonly nextToken?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-parentarn
     */
    readonly parentArn?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-principals
     */
    readonly principals?: pulumi.Input<inputs.LicenseManager.GrantArnListArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-sourceversion
     */
    readonly sourceVersion?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-status
     */
    readonly status?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-statusreason
     */
    readonly statusReason?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-tags
     */
    readonly tags?: pulumi.Input<inputs.LicenseManager.GrantTagListArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-version
     */
    readonly version?: pulumi.Input<string>;
}
