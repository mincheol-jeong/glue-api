package controller

import (
	"Glue-API/utils"
	"errors"
	"github.com/melbahja/goph"
)

func ConnectSSH(host string, keyfile string) (client *goph.Client, err error) {
	auth, err := goph.Key(keyfile, "")
	if err != nil {
		err = errors.Join(err, errors.New("keyfile: "+keyfile))
		utils.FancyHandleError(err)
		return
	}

	client, err = goph.NewUnknown("root", host, auth)
	if err != nil {
		err = errors.Join(err, errors.New("keyfile: "+keyfile))
		utils.FancyHandleError(err)
		return
	}

	return
}
