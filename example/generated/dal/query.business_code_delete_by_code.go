// DO NOT EDIT THIS FILE, IT IS GENERATED BY DALC
package dal

import (
	"github.com/pharosnet/dalc/v2"
)

// ************* business_code_delete_by_code *************
const businessCodeDeleteByCodeSQL = "DELETE FROM `applications`.`business_code` WHERE `group` = ? AND `code` "

type BusinessCodeDeleteByCodeRequest struct {
	Group string
}

func BusinessCodeDeleteByCode(ctx dalc.PreparedContext, request *BusinessCodeDeleteByCodeRequest) (affected int64, err error) {

	querySQL := businessCodeDeleteByCodeSQL
	args := dalc.NewArgs()
	args.Arg(request.Group)

	affected, err = dalc.Execute(ctx, querySQL, args)

	return
}
