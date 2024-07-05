package qbittorrent

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/bytedance/sonic"
)

type Application interface {
	// Version get application version
	Version() (string, error)
	// WebApiVersion get webapi version
	WebApiVersion() (string, error)
	// BuildInfo get build info
	BuildInfo() (*BuildInfo, error)
	// Shutdown exit application
	Shutdown() error
	// GetPreferences get application preferences
	GetPreferences() (*Preferences, error)
	// SetPreferences set application preferences
	SetPreferences(*Preferences) error
	// DefaultSavePath get default save path
	DefaultSavePath() (string, error)
}

type BuildInfo struct {
	BitNess    int    `json:"bitness,omitempty"`
	Boost      string `json:"boost,omitempty"`
	LibTorrent string `json:"libtorrent,omitempty"`
	Openssl    string `json:"openssl,omitempty"`
	QT         string `json:"qt,omitempty"`
	Zlib       string `json:"zlib,omitempty"`
}

type Preferences struct {
	AddToTopOfQueue                    bool   `json:"add_to_top_of_queue,omitempty"`
	AddTrackers                        string `json:"add_trackers,omitempty"`
	AddTrackersEnabled                 bool   `json:"add_trackers_enabled,omitempty"`
	AltDlLimit                         int    `json:"alt_dl_limit,omitempty"`
	AltUpLimit                         int    `json:"alt_up_limit,omitempty"`
	AlternativeWebuiEnabled            bool   `json:"alternative_webui_enabled,omitempty"`
	AlternativeWebuiPath               string `json:"alternative_webui_path,omitempty"`
	AnnounceIP                         string `json:"announce_ip,omitempty"`
	AnnounceToAllTiers                 bool   `json:"announce_to_all_tiers,omitempty"`
	AnnounceToAllTrackers              bool   `json:"announce_to_all_trackers,omitempty"`
	AnonymousMode                      bool   `json:"anonymous_mode,omitempty"`
	AsyncIoThreads                     int    `json:"async_io_threads,omitempty"`
	AutoDeleteMode                     int    `json:"auto_delete_mode,omitempty"`
	AutoTmmEnabled                     bool   `json:"auto_tmm_enabled,omitempty"`
	AutorunEnabled                     bool   `json:"autorun_enabled,omitempty"`
	AutorunOnTorrentAddedEnabled       bool   `json:"autorun_on_torrent_added_enabled,omitempty"`
	AutorunOnTorrentAddedProgram       string `json:"autorun_on_torrent_added_program,omitempty"`
	AutorunProgram                     string `json:"autorun_program,omitempty"`
	BannedIPs                          string `json:"banned_IPs,omitempty"`
	BdecodeDepthLimit                  int    `json:"bdecode_depth_limit,omitempty"`
	BdecodeTokenLimit                  int    `json:"bdecode_token_limit,omitempty"`
	BittorrentProtocol                 int    `json:"bittorrent_protocol,omitempty"`
	BlockPeersOnPrivilegedPorts        bool   `json:"block_peers_on_privileged_ports,omitempty"`
	BypassAuthSubnetWhitelist          string `json:"bypass_auth_subnet_whitelist,omitempty"`
	BypassAuthSubnetWhitelistEnabled   bool   `json:"bypass_auth_subnet_whitelist_enabled,omitempty"`
	BypassLocalAuth                    bool   `json:"bypass_local_auth,omitempty"`
	CategoryChangedTmmEnabled          bool   `json:"category_changed_tmm_enabled,omitempty"`
	CheckingMemoryUse                  int    `json:"checking_memory_use,omitempty"`
	ConnectionSpeed                    int    `json:"connection_speed,omitempty"`
	CurrentInterfaceAddress            string `json:"current_interface_address,omitempty"`
	CurrentInterfaceName               string `json:"current_interface_name,omitempty"`
	CurrentNetworkInterface            string `json:"current_network_interface,omitempty"`
	Dht                                bool   `json:"dht,omitempty"`
	DiskCache                          int    `json:"disk_cache,omitempty"`
	DiskCacheTTL                       int    `json:"disk_cache_ttl,omitempty"`
	DiskIoReadMode                     int    `json:"disk_io_read_mode,omitempty"`
	DiskIoType                         int    `json:"disk_io_type,omitempty"`
	DiskIoWriteMode                    int    `json:"disk_io_write_mode,omitempty"`
	DiskQueueSize                      int    `json:"disk_queue_size,omitempty"`
	DlLimit                            int    `json:"dl_limit,omitempty"`
	DontCountSlowTorrents              bool   `json:"dont_count_slow_torrents,omitempty"`
	DyndnsDomain                       string `json:"dyndns_domain,omitempty"`
	DyndnsEnabled                      bool   `json:"dyndns_enabled,omitempty"`
	DyndnsPassword                     string `json:"dyndns_password,omitempty"`
	DyndnsService                      int    `json:"dyndns_service,omitempty"`
	DyndnsUsername                     string `json:"dyndns_username,omitempty"`
	EmbeddedTrackerPort                int    `json:"embedded_tracker_port,omitempty"`
	EmbeddedTrackerPortForwarding      bool   `json:"embedded_tracker_port_forwarding,omitempty"`
	EnableCoalesceReadWrite            bool   `json:"enable_coalesce_read_write,omitempty"`
	EnableEmbeddedTracker              bool   `json:"enable_embedded_tracker,omitempty"`
	EnableMultiConnectionsFromSameIP   bool   `json:"enable_multi_connections_from_same_ip,omitempty"`
	EnablePieceExtentAffinity          bool   `json:"enable_piece_extent_affinity,omitempty"`
	EnableUploadSuggestions            bool   `json:"enable_upload_suggestions,omitempty"`
	Encryption                         int    `json:"encryption,omitempty"`
	ExcludedFileNames                  string `json:"excluded_file_names,omitempty"`
	ExcludedFileNamesEnabled           bool   `json:"excluded_file_names_enabled,omitempty"`
	ExportDir                          string `json:"export_dir,omitempty"`
	ExportDirFin                       string `json:"export_dir_fin,omitempty"`
	FileLogAge                         int    `json:"file_log_age,omitempty"`
	FileLogAgeType                     int    `json:"file_log_age_type,omitempty"`
	FileLogBackupEnabled               bool   `json:"file_log_backup_enabled,omitempty"`
	FileLogDeleteOld                   bool   `json:"file_log_delete_old,omitempty"`
	FileLogEnabled                     bool   `json:"file_log_enabled,omitempty"`
	FileLogMaxSize                     int    `json:"file_log_max_size,omitempty"`
	FileLogPath                        string `json:"file_log_path,omitempty"`
	FilePoolSize                       int    `json:"file_pool_size,omitempty"`
	HashingThreads                     int    `json:"hashing_threads,omitempty"`
	I2PAddress                         string `json:"i2p_address,omitempty"`
	I2PEnabled                         bool   `json:"i2p_enabled,omitempty"`
	I2PInboundLength                   int    `json:"i2p_inbound_length,omitempty"`
	I2PInboundQuantity                 int    `json:"i2p_inbound_quantity,omitempty"`
	I2PMixedMode                       bool   `json:"i2p_mixed_mode,omitempty"`
	I2POutboundLength                  int    `json:"i2p_outbound_length,omitempty"`
	I2POutboundQuantity                int    `json:"i2p_outbound_quantity,omitempty"`
	I2PPort                            int    `json:"i2p_port,omitempty"`
	IdnSupportEnabled                  bool   `json:"idn_support_enabled,omitempty"`
	IncompleteFilesExt                 bool   `json:"incomplete_files_ext,omitempty"`
	IPFilterEnabled                    bool   `json:"ip_filter_enabled,omitempty"`
	IPFilterPath                       string `json:"ip_filter_path,omitempty"`
	IPFilterTrackers                   bool   `json:"ip_filter_trackers,omitempty"`
	LimitLanPeers                      bool   `json:"limit_lan_peers,omitempty"`
	LimitTCPOverhead                   bool   `json:"limit_tcp_overhead,omitempty"`
	LimitUtpRate                       bool   `json:"limit_utp_rate,omitempty"`
	ListenPort                         int    `json:"listen_port,omitempty"`
	Locale                             string `json:"locale,omitempty"`
	Lsd                                bool   `json:"lsd,omitempty"`
	MailNotificationAuthEnabled        bool   `json:"mail_notification_auth_enabled,omitempty"`
	MailNotificationEmail              string `json:"mail_notification_email,omitempty"`
	MailNotificationEnabled            bool   `json:"mail_notification_enabled,omitempty"`
	MailNotificationPassword           string `json:"mail_notification_password,omitempty"`
	MailNotificationSender             string `json:"mail_notification_sender,omitempty"`
	MailNotificationSMTP               string `json:"mail_notification_smtp,omitempty"`
	MailNotificationSslEnabled         bool   `json:"mail_notification_ssl_enabled,omitempty"`
	MailNotificationUsername           string `json:"mail_notification_username,omitempty"`
	MaxActiveCheckingTorrents          int    `json:"max_active_checking_torrents,omitempty"`
	MaxActiveDownloads                 int    `json:"max_active_downloads,omitempty"`
	MaxActiveTorrents                  int    `json:"max_active_torrents,omitempty"`
	MaxActiveUploads                   int    `json:"max_active_uploads,omitempty"`
	MaxConcurrentHTTPAnnounces         int    `json:"max_concurrent_http_announces,omitempty"`
	MaxConnec                          int    `json:"max_connec,omitempty"`
	MaxConnecPerTorrent                int    `json:"max_connec_per_torrent,omitempty"`
	MaxInactiveSeedingTime             int    `json:"max_inactive_seeding_time,omitempty"`
	MaxInactiveSeedingTimeEnabled      bool   `json:"max_inactive_seeding_time_enabled,omitempty"`
	MaxRatio                           int    `json:"max_ratio,omitempty"`
	MaxRatioAct                        int    `json:"max_ratio_act,omitempty"`
	MaxRatioEnabled                    bool   `json:"max_ratio_enabled,omitempty"`
	MaxSeedingTime                     int    `json:"max_seeding_time,omitempty"`
	MaxSeedingTimeEnabled              bool   `json:"max_seeding_time_enabled,omitempty"`
	MaxUploads                         int    `json:"max_uploads,omitempty"`
	MaxUploadsPerTorrent               int    `json:"max_uploads_per_torrent,omitempty"`
	MemoryWorkingSetLimit              int    `json:"memory_working_set_limit,omitempty"`
	MergeTrackers                      bool   `json:"merge_trackers,omitempty"`
	OutgoingPortsMax                   int    `json:"outgoing_ports_max,omitempty"`
	OutgoingPortsMin                   int    `json:"outgoing_ports_min,omitempty"`
	PeerTos                            int    `json:"peer_tos,omitempty"`
	PeerTurnover                       int    `json:"peer_turnover,omitempty"`
	PeerTurnoverCutoff                 int    `json:"peer_turnover_cutoff,omitempty"`
	PeerTurnoverInterval               int    `json:"peer_turnover_interval,omitempty"`
	PerformanceWarning                 bool   `json:"performance_warning,omitempty"`
	Pex                                bool   `json:"pex,omitempty"`
	PreallocateAll                     bool   `json:"preallocate_all,omitempty"`
	ProxyAuthEnabled                   bool   `json:"proxy_auth_enabled,omitempty"`
	ProxyBittorrent                    bool   `json:"proxy_bittorrent,omitempty"`
	ProxyHostnameLookup                bool   `json:"proxy_hostname_lookup,omitempty"`
	ProxyIP                            string `json:"proxy_ip,omitempty"`
	ProxyMisc                          bool   `json:"proxy_misc,omitempty"`
	ProxyPassword                      string `json:"proxy_password,omitempty"`
	ProxyPeerConnections               bool   `json:"proxy_peer_connections,omitempty"`
	ProxyPort                          int    `json:"proxy_port,omitempty"`
	ProxyRss                           bool   `json:"proxy_rss,omitempty"`
	ProxyType                          string `json:"proxy_type,omitempty"`
	ProxyUsername                      string `json:"proxy_username,omitempty"`
	QueueingEnabled                    bool   `json:"queueing_enabled,omitempty"`
	RandomPort                         bool   `json:"random_port,omitempty"`
	ReannounceWhenAddressChanged       bool   `json:"reannounce_when_address_changed,omitempty"`
	RecheckCompletedTorrents           bool   `json:"recheck_completed_torrents,omitempty"`
	RefreshInterval                    int    `json:"refresh_interval,omitempty"`
	RequestQueueSize                   int    `json:"request_queue_size,omitempty"`
	ResolvePeerCountries               bool   `json:"resolve_peer_countries,omitempty"`
	ResumeDataStorageType              string `json:"resume_data_storage_type,omitempty"`
	RssAutoDownloadingEnabled          bool   `json:"rss_auto_downloading_enabled,omitempty"`
	RssDownloadRepackProperEpisodes    bool   `json:"rss_download_repack_proper_episodes,omitempty"`
	RssMaxArticlesPerFeed              int    `json:"rss_max_articles_per_feed,omitempty"`
	RssProcessingEnabled               bool   `json:"rss_processing_enabled,omitempty"`
	RssRefreshInterval                 int    `json:"rss_refresh_interval,omitempty"`
	RssSmartEpisodeFilters             string `json:"rss_smart_episode_filters,omitempty"`
	SavePath                           string `json:"save_path,omitempty"`
	SavePathChangedTmmEnabled          bool   `json:"save_path_changed_tmm_enabled,omitempty"`
	SaveResumeDataInterval             int    `json:"save_resume_data_interval,omitempty"`
	ScheduleFromHour                   int    `json:"schedule_from_hour,omitempty"`
	ScheduleFromMin                    int    `json:"schedule_from_min,omitempty"`
	ScheduleToHour                     int    `json:"schedule_to_hour,omitempty"`
	ScheduleToMin                      int    `json:"schedule_to_min,omitempty"`
	SchedulerDays                      int    `json:"scheduler_days,omitempty"`
	SchedulerEnabled                   bool   `json:"scheduler_enabled,omitempty"`
	SendBufferLowWatermark             int    `json:"send_buffer_low_watermark,omitempty"`
	SendBufferWatermark                int    `json:"send_buffer_watermark,omitempty"`
	SendBufferWatermarkFactor          int    `json:"send_buffer_watermark_factor,omitempty"`
	SlowTorrentDlRateThreshold         int    `json:"slow_torrent_dl_rate_threshold,omitempty"`
	SlowTorrentInactiveTimer           int    `json:"slow_torrent_inactive_timer,omitempty"`
	SlowTorrentUlRateThreshold         int    `json:"slow_torrent_ul_rate_threshold,omitempty"`
	SocketBacklogSize                  int    `json:"socket_backlog_size,omitempty"`
	SocketReceiveBufferSize            int    `json:"socket_receive_buffer_size,omitempty"`
	SocketSendBufferSize               int    `json:"socket_send_buffer_size,omitempty"`
	SsrfMitigation                     bool   `json:"ssrf_mitigation,omitempty"`
	StartPausedEnabled                 bool   `json:"start_paused_enabled,omitempty"`
	StopTrackerTimeout                 int    `json:"stop_tracker_timeout,omitempty"`
	TempPath                           string `json:"temp_path,omitempty"`
	TempPathEnabled                    bool   `json:"temp_path_enabled,omitempty"`
	TorrentChangedTmmEnabled           bool   `json:"torrent_changed_tmm_enabled,omitempty"`
	TorrentContentLayout               string `json:"torrent_content_layout,omitempty"`
	TorrentFileSizeLimit               int    `json:"torrent_file_size_limit,omitempty"`
	TorrentStopCondition               string `json:"torrent_stop_condition,omitempty"`
	UpLimit                            int    `json:"up_limit,omitempty"`
	UploadChokingAlgorithm             int    `json:"upload_choking_algorithm,omitempty"`
	UploadSlotsBehavior                int    `json:"upload_slots_behavior,omitempty"`
	Upnp                               bool   `json:"upnp,omitempty"`
	UpnpLeaseDuration                  int    `json:"upnp_lease_duration,omitempty"`
	UseCategoryPathsInManualMode       bool   `json:"use_category_paths_in_manual_mode,omitempty"`
	UseHTTPS                           bool   `json:"use_https,omitempty"`
	UseSubcategories                   bool   `json:"use_subcategories,omitempty"`
	UtpTCPMixedMode                    int    `json:"utp_tcp_mixed_mode,omitempty"`
	ValidateHTTPSTrackerCertificate    bool   `json:"validate_https_tracker_certificate,omitempty"`
	WebUIAddress                       string `json:"web_ui_address,omitempty"`
	WebUIBanDuration                   int    `json:"web_ui_ban_duration,omitempty"`
	WebUIClickjackingProtectionEnabled bool   `json:"web_ui_clickjacking_protection_enabled,omitempty"`
	WebUICsrfProtectionEnabled         bool   `json:"web_ui_csrf_protection_enabled,omitempty"`
	WebUICustomHTTPHeaders             string `json:"web_ui_custom_http_headers,omitempty"`
	WebUIDomainList                    string `json:"web_ui_domain_list,omitempty"`
	WebUIHostHeaderValidationEnabled   bool   `json:"web_ui_host_header_validation_enabled,omitempty"`
	WebUIHTTPSCertPath                 string `json:"web_ui_https_cert_path,omitempty"`
	WebUIHTTPSKeyPath                  string `json:"web_ui_https_key_path,omitempty"`
	WebUIMaxAuthFailCount              int    `json:"web_ui_max_auth_fail_count,omitempty"`
	WebUIPort                          int    `json:"web_ui_port,omitempty"`
	WebUIReverseProxiesList            string `json:"web_ui_reverse_proxies_list,omitempty"`
	WebUIReverseProxyEnabled           bool   `json:"web_ui_reverse_proxy_enabled,omitempty"`
	WebUISecureCookieEnabled           bool   `json:"web_ui_secure_cookie_enabled,omitempty"`
	WebUISessionTimeout                int    `json:"web_ui_session_timeout,omitempty"`
	WebUIUpnp                          bool   `json:"web_ui_upnp,omitempty"`
	WebUIUseCustomHTTPHeadersEnabled   bool   `json:"web_ui_use_custom_http_headers_enabled,omitempty"`
	WebUIUsername                      string `json:"web_ui_username,omitempty"`
}

func (c *client) Version() (string, error) {
	apiUrl := fmt.Sprintf("%s/api/v2/app/version", c.config.Address)

	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return "", err
	}

	if result.code != 200 {
		return "", errors.New("get version failed: " + string(result.body))
	}

	return string(result.body), nil
}

func (c *client) WebApiVersion() (string, error) {
	apiUrl := fmt.Sprintf("%s/api/v2/app/webapiVersion", c.config.Address)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return "", err
	}

	if result.code != 200 {
		return "", errors.New("get version failed: " + string(result.body))
	}

	return string(result.body), nil
}

func (c *client) BuildInfo() (*BuildInfo, error) {
	apiUrl := fmt.Sprintf("%s/api/v2/app/buildInfo", c.config.Address)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return nil, err
	}

	if result.code != 200 {
		return nil, errors.New("get build info failed: " + string(result.body))
	}

	var build = new(BuildInfo)
	if err := sonic.Unmarshal(result.body, build); err != nil {
		return nil, err
	}

	return build, nil
}

func (c *client) Shutdown() error {
	apiUrl := fmt.Sprintf("%s/api/v2/app/shutdown", c.config.Address)
	result, err := c.doRequest(&requestData{
		method: http.MethodPost,
		url:    apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("shutdown application failed: " + string(result.body))
	}

	return nil
}

func (c *client) GetPreferences() (*Preferences, error) {
	apiUrl := fmt.Sprintf("%s/api/v2/app/preferences", c.config.Address)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return nil, err
	}

	if result.code != 200 {
		return nil, errors.New("get preference failed: " + string(result.body))
	}

	var preferences = new(Preferences)
	if err := sonic.Unmarshal(result.body, preferences); err != nil {
		return nil, err
	}

	return preferences, nil
}

func (c *client) SetPreferences(prefs *Preferences) error {
	apiUrl := fmt.Sprintf("%s/api/v2/app/setPreferences", c.config.Address)
	data, err := sonic.Marshal(prefs)
	if err != nil {
		return err
	}
	var formData bytes.Buffer
	formData.Write([]byte("json="))
	formData.Write(data)

	result, err := c.doRequest(&requestData{
		method:      http.MethodPost,
		url:         apiUrl,
		contentType: ContentTypeFormUrlEncoded,
		body:        &formData,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("set preference failed: " + string(result.body))
	}

	return nil
}

func (c *client) DefaultSavePath() (string, error) {
	apiUrl := fmt.Sprintf("%s/api/v2/app/defaultSavePath", c.config.Address)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return "", err
	}

	if result.code != 200 {
		return "", errors.New("get default save path failed: " + string(result.body))
	}

	return string(result.body), nil
}
