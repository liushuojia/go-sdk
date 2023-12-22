package token

import (
	"github.com/twinj/uuid"
	"time"
)

func NewDouble() *Double {
	return &Double{}
}

type Double struct {
	access  JWT
	refresh JWT
}

func (d *Double) SetData(data any) *Double {
	d.GetAccess().SetData(data)
	d.GetRefresh().SetData(data)
	return d
}
func (d *Double) GetData() any {
	return d.GetAccess().GetData()
}
func (d *Double) GetAccess() *JWT {
	return &d.access
}
func (d *Double) GetRefresh() *JWT {
	return &d.refresh
}
func (d *Double) Generate() (accessToken string, refreshToken string, expire time.Time, err error) {
	if d.GetAccess().GetUUID() == "" {
		d.GetAccess().SetUUID(uuid.NewV4().String())
	}

	if d.GetRefresh().GetUUID() == "" {
		d.GetRefresh().SetUUID(uuid.NewV4().String())
	}
	d.GetRefresh().SetParentUUID(d.GetAccess().GetUUID())

	accessToken, err = d.GetAccess().Generate()
	if err != nil {
		return
	}
	expire = d.GetAccess().GetExpire()

	refreshToken, err = d.GetRefresh().Generate()
	return
}
func (d *Double) Refresh(rfToken string) (accessToken string, refreshToken string, expireTime time.Time, err error) {
	claims, err := d.GetRefresh().Parse(rfToken)
	if err != nil {
		return
	}

	d.GetRefresh().SetClaims(*claims)
	d.GetAccess().SetClaims(*claims).SetUUID(claims.ParentUUID).SetParentUUID("")
	return d.Generate()
}
