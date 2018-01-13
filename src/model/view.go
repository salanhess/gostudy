package model

// api response data model

import (
	"time"
)

type DescVolumeErr struct {
	RequestId string    `json:"requestId"`
	Error     ErrorView `json:"error"`
}

type ErrorView struct {
	Code    uint64      `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

type DescVolume struct {
	RequestId string     `json:"requestId"`
	Result    VolumeView `json:"result"`
}

type VolumeView struct {
	Id             string      `json:"diskId"`
	VolumeName     string      `json:"name"`
	Size           uint64      `json:"diskSizeGB"`
	VolumeTypeName string      `json:"diskType"`
	AZName         string      `json:"az"`
	Status         string      `json:"status"`
	Description    string      `json:"description"`
	SnapshotId     string      `json:"snapshot_id"`
	CreatedAt      time.Time   `json:"create_time"`
	Attachments    interface{} `json:"attachments"`
	Charge         interface{} `json:"charge"`
}
type SnapshotView struct {
	Id           string    `json:"id"`
	TenantId     string    `json:"tenant_id"`
	SnapshotName string    `json:"name"`
	Size         uint64    `json:"size"`
	VolumeId     string    `json:"volume_id"`
	PoolId       string    `json:"pool_id"`
	AZName       string    `json:"az_name"`
	CreatedAt    time.Time `json:"create_time"`
	UpdatedAt    time.Time `json:"update_time"`
	DeletedAt    time.Time `json:"delete_time"`
	Status       string    `json:"status"`
	Description  string    `json:"description"`
}

type QuotaView struct {
	TenantId  string    `json:"tenant_id"`
	Type      string    `json:"type"`
	Used      int       `json:"used_quota"`
	Quota     int       `json:"quota"`
	CreatedAt time.Time `json:"create_time"`
	UpdatedAt time.Time `json:"update_time"`
}

type AttachResultView struct {
	Id string `json:"id"`
}

type AtDeleteView struct {
	TaskId string `json:"task_id"`
}

type SpeedLimitView struct {
	VolumeId string `json:"volume_id"`
	IOPS     string `json:"iops"`
}
