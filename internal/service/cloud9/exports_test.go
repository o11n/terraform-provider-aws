package cloud9

// Exports for use in tests only.
var (
	ResourceEnvironmentEC2        = resourceEnvironmentEC2
	ResourceEnvironmentMembership = resourceEnvironmentMembership

	FindEnvironmentByID                   = findEnvironmentByID
	FindEnvironmentMembershipByTwoPartKey = findEnvironmentMembershipByTwoPartKey

	EnvironmentMembershipParseResourceID = environmentMembershipParseResourceID
)
