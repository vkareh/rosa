package ingress

import (
	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
)

var ValidWildcardPolicies = []string{string(cmv1.WildcardPolicyWildcardsDisallowed),
	string(cmv1.WildcardPolicyWildcardsAllowed)}
var DefaultWildcardPolicy = cmv1.WildcardPolicyWildcardsDisallowed
var ValidNamespaceOwnershipPolicies = []string{string(cmv1.NamespaceOwnershipPolicyStrict),
	string(cmv1.NamespaceOwnershipPolicyInterNamespaceAllowed)}
var DefaultNamespaceOwnershipPolicy = cmv1.NamespaceOwnershipPolicyStrict
