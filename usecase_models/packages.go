package usecase_models

import "time"

type Package struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Price     int       `json:"price"`
	Limits    Limits    `json:"limits"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type CreatePackageResponse struct {
	PackageId int `json:"package_id"`
}

type Limits struct {
	EndpointLimits   EndpointsLimits  `json:"endpoint_limits"`
	NetCatLimits     NetCatLimits     `json:"net_cat_limits"`
	PingLimits       PingLimits       `json:"ping_limits"`
	PageSpeedLimits  PageSpeedLimits  `json:"page_speed_limits"`
	TraceRouteLimits TraceRouteLimits `json:"trace_route_limits"`
}

type EndpointsLimits struct {
	NumberOfMonitoring int `json:"number_of_monitoring"`
	DurationLimit      int `json:"duration_limit"`
}

type NetCatLimits struct {
	NumberOfMonitoring int `json:"number_of_monitoring"`
	DurationLimit      int `json:"duration_limit"`
}

type PingLimits struct {
	NumberOfMonitoring int `json:"number_of_monitoring"`
	DurationLimit      int `json:"duration_limit"`
}

type PageSpeedLimits struct {
	NumberOfMonitoring int `json:"number_of_monitoring"`
	DurationLimit      int `json:"duration_limit"`
}

type TraceRouteLimits struct {
	NumberOfMonitoring int `json:"number_of_monitoring"`
	DurationLimit      int `json:"duration_limit"`
}
