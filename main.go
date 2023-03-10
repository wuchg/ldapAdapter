package main

import (
	"ldapAdapter/server"
)

func main() {
	//searchRequest := ldap.NewSearchRequest(
	//	config.BaseDN,
	//	ldap.ScopeWholeSubtree,
	//	ldap.NeverDerefAliases,
	//	0,
	//	0,
	//	false,
	//	config.Filter,
	//	config.Attributes,
	//	nil,
	//)
	//
	//result, err := conn.Search(searchRequest)
	//if err != nil {
	//	log.Fatal(err)
	//}
	// http://www.selfadsi.org/deep-inside/microsoft-sid-attributes.htm
	// https://github.com/owncloud/user_ldap/issues/379
	//for _, entry := range result.Entries {
	//	fmt.Printf("DN=====>%s\n", entry.DN)
	//	attributes := entry.Attributes
	//	for _, attr := range attributes {
	//		fmt.Printf("%s=%s\n", attr.Name, attr.Values)
	//	}
	//	fmt.Printf("\n\n")
	//}

	server.Start()
}
