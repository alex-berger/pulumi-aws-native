// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Backup.Outputs
{

    /// <summary>
    /// Identifies the report template for the report. Reports are built using a report template.
    /// </summary>
    [OutputType]
    public sealed class ReportSettingProperties
    {
        /// <summary>
        /// The Amazon Resource Names (ARNs) of the frameworks a report covers.
        /// </summary>
        public readonly ImmutableArray<string> FrameworkArns;
        /// <summary>
        /// Identifies the report template for the report. Reports are built using a report template. The report templates are: `BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT`
        /// </summary>
        public readonly string ReportTemplate;

        [OutputConstructor]
        private ReportSettingProperties(
            ImmutableArray<string> frameworkArns,

            string reportTemplate)
        {
            FrameworkArns = frameworkArns;
            ReportTemplate = reportTemplate;
        }
    }
}
