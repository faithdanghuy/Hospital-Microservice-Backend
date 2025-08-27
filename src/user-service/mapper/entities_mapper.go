package mapper

import (
	"github.com/Hospital-Microservice/user-service/entity"
	"github.com/Hospital-Microservice/user-service/model/res"
)

func TransformUserEntitiesToRes(userEntities []entity.UserEntity) []*res.FilterRes {
	resList := make([]*res.FilterRes, 0, len(userEntities))
	for _, u := range userEntities {
		resList = append(resList, TransformUserEntityToFilterRes(&u))
	}
	return resList
}
