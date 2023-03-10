package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Conf struct {
	Addr       string
	BaseDN     string
	Filter     string
	Attributes []string
	Username   string
	Password   string
}

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Bind(user User) {
	var config Conf
	err := envconfig.Process("LDAP", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("\n%+v\n\n", config)
	conn, err := ldap.DialURL(config.Addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// The “BIND” operation is used to set the authentication state for an LDAP session in which the LDAP client connects to the server.
	err = conn.Bind(user.Username, user.Password)
	if err != nil {
		log.Fatal(err)
	}
}
