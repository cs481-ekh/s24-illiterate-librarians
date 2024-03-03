package model

type EOSRequest struct {
	EOSPSID string `json:"EOS_p_s_id" binding:"required"`
}