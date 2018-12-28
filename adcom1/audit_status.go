package adcom1

// AuditStatus represents codes used in Audit objects to reflect status or workflow state.
type AuditStatus int

// Codes used in Audit objects to reflect status or workflow state.
//
// Values of 500+ hold endor-specific codes.
const (
	AuditPendingAudit                 AuditStatus = 1 // Pending Audit: An audit has not yet been completed on this ad. A recommendation cannot be made to use this ad, but vendors' policies may override.
	AuditPreApproved                  AuditStatus = 2 // Pre-Approved: An audit has not yet been completed on this ad. Subject to vendors' policies, it can be recommended for use. However, once the audit has been completed, its status will change and it may or may not be approved for continued use.
	AuditApproved                     AuditStatus = 3 // Approved: The audit is complete and the ad is approved for use. Note, however, that some attributes (e.g., adomain, cat, attr, etc.) may have been changed in the process by the auditor.
	AuditDenied                       AuditStatus = 4 // Denied: The audit is complete, but the ad has been found unacceptable in some material aspect and is disapproved for use.
	AuditChangedResubmissionRequested AuditStatus = 5 // Changed; Resubmission Requested: A version of the ad has been detected in use that is materially different from the version that was previously audited, which may result in rejection during use until the ad is resubmitted for audit and approved. Vendors need to communicate offline as to the criteria that constitutes a material change.
	AuditExpired                      AuditStatus = 6 // Expired: The ad has been marked as expired by the vendor. Vendors need to communicate offline as to the expected bidding behavior for ads with this status.
)
