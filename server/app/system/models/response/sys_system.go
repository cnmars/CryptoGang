package response

import "golddigger/app/system/models/config"

type SysConfigResponse struct {
	Config config.ConfigCollection `json:"config"`
}
