package response

import system "golddigger/app/system/models"
type SysAPIResponse struct {
	Api system.SysApi `json:"api"`
}

type SysAPIListResponse struct {
	Apis []system.SysApi `json:"apis"`
}
