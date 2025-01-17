// This file is automatically generated. Any manual edits here will be overwritten on the next update.
package cmd

var resourceToEndpoint = map[string]string{
	"cloudflare_account":                                                 "/accounts",
	"cloudflare_account_member":                                          "/accounts/{account_id}/members",
	"cloudflare_account_subscription":                                    "/accounts/{account_id}/subscriptions",
	"cloudflare_account_token":                                           "/accounts/{account_id}/tokens",
	"cloudflare_origin_ca_certificate":                                   "/certificates",
	"cloudflare_user":                                                    "/user",
	"cloudflare_api_token":                                               "/user/tokens",
	"cloudflare_zone":                                                    "/zones",
	"cloudflare_zone_setting":                                            "/zones/{zone_id}/settings/{setting_id}",
	"cloudflare_zone_hold":                                               "/zones/{zone_id}/hold",
	"cloudflare_zone_subscription":                                       "/zones/{identifier}/subscription",
	"cloudflare_load_balancer":                                           "/zones/{zone_id}/load_balancers",
	"cloudflare_load_balancer_monitor":                                   "/accounts/{account_id}/load_balancers/monitors",
	"cloudflare_load_balancer_pool":                                      "/accounts/{account_id}/load_balancers/pools",
	"cloudflare_zone_cache_reserve":                                      "/zones/{zone_id}/cache/cache_reserve",
	"cloudflare_tiered_cache":                                            "/zones/{zone_id}/cache/tiered_cache_smart_topology_enable",
	"cloudflare_zone_cache_variants":                                     "/zones/{zone_id}/cache/variants",
	"cloudflare_regional_tiered_cache":                                   "/zones/{zone_id}/cache/regional_tiered_cache",
	"cloudflare_certificate_pack":                                        "/zones/{zone_id}/ssl/certificate_packs",
	"cloudflare_total_tls":                                               "/zones/{zone_id}/acm/total_tls",
	"cloudflare_argo_smart_routing":                                      "/zones/{zone_id}/argo/smart_routing",
	"cloudflare_argo_tiered_caching":                                     "/zones/{zone_id}/argo/tiered_caching",
	"cloudflare_custom_ssl":                                              "/zones/{zone_id}/custom_certificates",
	"cloudflare_custom_hostname":                                         "/zones/{zone_id}/custom_hostnames",
	"cloudflare_custom_hostname_fallback_origin":                         "/zones/{zone_id}/custom_hostnames/fallback_origin",
	"cloudflare_dns_firewall":                                            "/accounts/{account_id}/dns_firewall",
	"cloudflare_zone_dnssec":                                             "/zones/{zone_id}/dnssec",
	"cloudflare_dns_record":                                              "/zones/{zone_id}/dns_records",
	"cloudflare_dns_zone_transfers_incoming":                             "/zones/{zone_id}/secondary_dns/incoming",
	"cloudflare_dns_zone_transfers_outgoing":                             "/zones/{zone_id}/secondary_dns/outgoing",
	"cloudflare_dns_zone_transfers_acl":                                  "/accounts/{account_id}/secondary_dns/acls",
	"cloudflare_dns_zone_transfers_peer":                                 "/accounts/{account_id}/secondary_dns/peers",
	"cloudflare_dns_zone_transfers_tsig":                                 "/accounts/{account_id}/secondary_dns/tsigs",
	"cloudflare_email_security_block_sender":                             "/accounts/{account_id}/email-security/settings/block_senders",
	"cloudflare_email_security_impersonation_registry":                   "/accounts/{account_id}/email-security/settings/impersonation_registry",
	"cloudflare_email_security_trusted_domains":                          "/accounts/{account_id}/email-security/settings/trusted_domains",
	"cloudflare_email_routing_settings":                                  "/zones/{zone_id}/email/routing",
	"cloudflare_email_routing_dns":                                       "/zones/{zone_id}/email/routing/dns",
	"cloudflare_email_routing_rule":                                      "/zones/{zone_id}/email/routing/rules",
	"cloudflare_email_routing_catch_all":                                 "/zones/{zone_id}/email/routing/rules/catch_all",
	"cloudflare_email_routing_address":                                   "/accounts/{account_id}/email/routing/addresses",
	"cloudflare_filter":                                                  "/zones/{zone_id}/filters",
	"cloudflare_zone_lockdown":                                           "/zones/{zone_id}/firewall/lockdowns",
	"cloudflare_firewall_rule":                                           "/zones/{zone_id}/firewall/rules",
	"cloudflare_access_rule":                                             "/{account_or_zone}/{account_or_zone_id}/firewall/access_rules/rules",
	"cloudflare_user_agent_blocking_rule":                                "/zones/{zone_id}/firewall/ua_rules",
	"cloudflare_healthcheck":                                             "/zones/{zone_id}/healthchecks",
	"cloudflare_keyless_certificate":                                     "/zones/{zone_id}/keyless_certificates",
	"cloudflare_logpush_job":                                             "/{account_or_zone}/{account_or_zone_id}/logpush/jobs",
	"cloudflare_logpull_retention":                                       "/zones/{zone_id}/logs/control/retention/flag",
	"cloudflare_authenticated_origin_pulls_certificate":                  "/zones/{zone_id}/origin_tls_client_auth",
	"cloudflare_authenticated_origin_pulls":                              "/zones/{zone_id}/origin_tls_client_auth/hostnames/{hostname}",
	"cloudflare_page_rule":                                               "/zones/{zone_id}/pagerules",
	"cloudflare_rate_limit":                                              "/zones/{zone_id}/rate_limits",
	"cloudflare_waiting_room":                                            "/zones/{zone_id}/waiting_rooms",
	"cloudflare_waiting_room_event":                                      "/zones/{zone_id}/waiting_rooms/{waiting_room_id}/events",
	"cloudflare_waiting_room_rules":                                      "/zones/{zone_id}/waiting_rooms/{waiting_room_id}/rules",
	"cloudflare_waiting_room_settings":                                   "/zones/{zone_id}/waiting_rooms/settings",
	"cloudflare_web3_hostname":                                           "/zones/{zone_id}/web3/hostnames",
	"cloudflare_workers_script":                                          "/accounts/{account_id}/workers/scripts",
	"cloudflare_workers_script_subdomain":                                "/accounts/{account_id}/workers/scripts/{script_name}/subdomain",
	"cloudflare_workers_cron_trigger":                                    "/accounts/{account_id}/workers/scripts/{script_name}/schedules",
	"cloudflare_workers_deployment":                                      "/accounts/{account_id}/workers/scripts/{script_name}/deployments",
	"cloudflare_workers_custom_domain":                                   "/accounts/{account_id}/workers/domains",
	"cloudflare_workers_kv_namespace":                                    "/accounts/{account_id}/storage/kv/namespaces",
	"cloudflare_workers_kv":                                              "/accounts/{account_id}/storage/kv/namespaces/{namespace_id}/values/{key_name}",
	"cloudflare_queue":                                                   "/accounts/{account_id}/queues",
	"cloudflare_queue_consumer":                                          "/accounts/{account_id}/queues/{queue_id}/consumers",
	"cloudflare_api_shield":                                              "/zones/{zone_id}/api_gateway/configuration",
	"cloudflare_api_shield_discovery_operation":                          "/zones/{zone_id}/api_gateway/discovery/operations",
	"cloudflare_api_shield_operation":                                    "/zones/{zone_id}/api_gateway/operations",
	"cloudflare_api_shield_operation_schema_validation_settings":         "/zones/{zone_id}/api_gateway/operations/{operation_id}/schema_validation",
	"cloudflare_api_shield_schema_validation_settings":                   "/zones/{zone_id}/api_gateway/settings/schema_validation",
	"cloudflare_api_shield_schema":                                       "/zones/{zone_id}/api_gateway/user_schemas",
	"cloudflare_managed_transforms":                                      "/zones/{zone_id}/managed_headers",
	"cloudflare_page_shield_policy":                                      "/zones/{zone_id}/page_shield/policies",
	"cloudflare_ruleset":                                                 "/{account_or_zone}/{account_or_zone_id}/rulesets",
	"cloudflare_url_normalization_settings":                              "/zones/{zone_id}/url_normalization",
	"cloudflare_spectrum_application":                                    "/zones/{zone_id}/spectrum/apps/{app_id}",
	"cloudflare_regional_hostname":                                       "/zones/{zone_id}/addressing/regional_hostnames",
	"cloudflare_address_map":                                             "/accounts/{account_id}/addressing/address_maps",
	"cloudflare_byo_ip_prefix":                                           "/accounts/{account_id}/addressing/prefixes",
	"cloudflare_image":                                                   "/accounts/{account_id}/images/v1",
	"cloudflare_image_variant":                                           "/accounts/{account_id}/images/v1/variants",
	"cloudflare_magic_wan_gre_tunnel":                                    "/accounts/{account_id}/magic/gre_tunnels",
	"cloudflare_magic_wan_ipsec_tunnel":                                  "/accounts/{account_id}/magic/ipsec_tunnels",
	"cloudflare_magic_wan_static_route":                                  "/accounts/{account_id}/magic/routes",
	"cloudflare_magic_transit_site":                                      "/accounts/{account_id}/magic/sites",
	"cloudflare_magic_transit_site_acl":                                  "/accounts/{account_id}/magic/sites/{site_id}/acls",
	"cloudflare_magic_transit_site_lan":                                  "/accounts/{account_id}/magic/sites/{site_id}/lans",
	"cloudflare_magic_transit_site_wan":                                  "/accounts/{account_id}/magic/sites/{site_id}/wans",
	"cloudflare_magic_transit_connector":                                 "/accounts/{account_id}/magic/connectors",
	"cloudflare_magic_network_monitoring_configuration":                  "/accounts/{account_id}/mnm/config",
	"cloudflare_magic_network_monitoring_rule":                           "/accounts/{account_id}/mnm/rules",
	"cloudflare_mtls_certificate":                                        "/accounts/{account_id}/mtls_certificates",
	"cloudflare_pages_project":                                           "/accounts/{account_id}/pages/projects",
	"cloudflare_pages_domain":                                            "/accounts/{account_id}/pages/projects/{project_name}/domains",
	"cloudflare_registrar_domain":                                        "/accounts/{account_id}/registrar/domains",
	"cloudflare_list":                                                    "/accounts/{account_id}/rules/lists",
	"cloudflare_list_item":                                               "/accounts/{account_id}/rules/lists/{list_id}/items",
	"cloudflare_stream":                                                  "/accounts/{account_id}/stream",
	"cloudflare_stream_audio_track":                                      "/accounts/{account_id}/stream/{identifier}/audio",
	"cloudflare_stream_key":                                              "/accounts/{account_id}/stream/keys",
	"cloudflare_stream_live_input":                                       "/accounts/{account_id}/stream/live_inputs",
	"cloudflare_stream_watermark":                                        "/accounts/{account_id}/stream/watermarks",
	"cloudflare_stream_webhook":                                          "/accounts/{account_id}/stream/webhook",
	"cloudflare_stream_caption_language":                                 "/accounts/{account_id}/stream/{identifier}/captions/{language}",
	"cloudflare_stream_download":                                         "/accounts/{account_id}/stream/{identifier}/downloads",
	"cloudflare_notification_policy_webhooks":                            "/accounts/{account_id}/alerting/v3/destinations/webhooks",
	"cloudflare_notification_policy":                                     "/accounts/{account_id}/alerting/v3/policies",
	"cloudflare_d1_database":                                             "/accounts/{account_id}/d1/database",
	"cloudflare_r2_bucket":                                               "/accounts/{account_id}/r2/buckets",
	"cloudflare_r2_custom_domain":                                        "/accounts/{account_id}/r2/buckets/{bucket_name}/domains/custom",
	"cloudflare_r2_managed_domain":                                       "/accounts/{account_id}/r2/buckets/{bucket_name}/domains/managed",
	"cloudflare_workers_for_platforms_dispatch_namespace":                "/accounts/{account_id}/workers/dispatch/namespaces",
	"cloudflare_workers_secret":                                          "/accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/secrets",
	"cloudflare_zero_trust_dex_test":                                     "/accounts/{account_id}/devices/dex_tests",
	"cloudflare_zero_trust_device_managed_networks":                      "/accounts/{account_id}/devices/networks",
	"cloudflare_zero_trust_device_default_profile":                       "/accounts/{account_id}/devices/policy",
	"cloudflare_zero_trust_device_default_profile_local_domain_fallback": "/accounts/{account_id}/devices/policy/fallback_domains",
	"cloudflare_zero_trust_device_default_profile_certificates":          "/zones/{zone_id}/devices/policy/certificates",
	"cloudflare_zero_trust_device_custom_profile":                        "/accounts/{account_id}/devices/policies",
	"cloudflare_zero_trust_device_custom_profile_local_domain_fallback":  "/accounts/{account_id}/devices/policy/{policy_id}/fallback_domains",
	"cloudflare_zero_trust_device_posture_rule":                          "/accounts/{account_id}/devices/posture",
	"cloudflare_zero_trust_device_posture_integration":                   "/accounts/{account_id}/devices/posture/integration",
	"cloudflare_zero_trust_access_identity_provider":                     "/{account_or_zone}/{account_or_zone_id}/access/identity_providers",
	"cloudflare_zero_trust_organization":                                 "/{account_or_zone}/{account_or_zone_id}/access/organizations",
	"cloudflare_zero_trust_access_infrastructure_target":                 "/accounts/{account_id}/infrastructure/targets",
	"cloudflare_zero_trust_access_application":                           "/{account_or_zone}/{account_or_zone_id}/access/apps",
	"cloudflare_zero_trust_access_short_lived_certificate":               "/{account_or_zone}/{account_or_zone_id}/access/apps/ca",
	"cloudflare_zero_trust_access_mtls_certificate":                      "/{account_or_zone}/{account_or_zone_id}/access/certificates",
	"cloudflare_zero_trust_access_mtls_hostname_settings":                "/{account_or_zone}/{account_or_zone_id}/access/certificates/settings",
	"cloudflare_zero_trust_access_group":                                 "/{account_or_zone}/{account_or_zone_id}/access/groups",
	"cloudflare_zero_trust_access_service_token":                         "/{account_or_zone}/{account_or_zone_id}/access/service_tokens",
	"cloudflare_zero_trust_access_key_configuration":                     "/accounts/{account_id}/access/keys",
	"cloudflare_zero_trust_access_custom_page":                           "/accounts/{account_id}/access/custom_pages",
	"cloudflare_zero_trust_access_tag":                                   "/accounts/{account_id}/access/tags",
	"cloudflare_zero_trust_access_policy":                                "/accounts/{account_id}/access/policies",
	"cloudflare_zero_trust_tunnel_cloudflared":                           "/accounts/{account_id}/cfd_tunnel",
	"cloudflare_zero_trust_tunnel_cloudflared_config":                    "/accounts/{account_id}/cfd_tunnel/{tunnel_id}/configurations",
	"cloudflare_zero_trust_dlp_dataset":                                  "/accounts/{account_id}/dlp/datasets",
	"cloudflare_zero_trust_dlp_custom_profile":                           "/accounts/{account_id}/dlp/profiles/custom/{profile_id}",
	"cloudflare_zero_trust_dlp_predefined_profile":                       "/accounts/{account_id}/dlp/profiles/predefined/{profile_id}",
	"cloudflare_zero_trust_dlp_entry":                                    "/accounts/{account_id}/dlp/entries",
	"cloudflare_zero_trust_gateway_categories":                           "/accounts/{account_id}/gateway/categories",
	"cloudflare_zero_trust_gateway_app_types":                            "/accounts/{account_id}/gateway/app_types",
	"cloudflare_zero_trust_gateway_settings":                             "/accounts/{account_id}/gateway/configuration",
	"cloudflare_zero_trust_list":                                         "/accounts/{account_id}/gateway/lists",
	"cloudflare_zero_trust_dns_location":                                 "/accounts/{account_id}/gateway/locations",
	"cloudflare_zero_trust_gateway_proxy_endpoint":                       "/accounts/{account_id}/gateway/proxy_endpoints",
	"cloudflare_zero_trust_gateway_policy":                               "/accounts/{account_id}/gateway/rules",
	"cloudflare_zero_trust_gateway_certificate":                          "/accounts/{account_id}/gateway/certificates",
	"cloudflare_zero_trust_tunnel_cloudflared_route":                     "/accounts/{account_id}/teamnet/routes",
	"cloudflare_zero_trust_tunnel_cloudflared_virtual_network":           "/accounts/{account_id}/teamnet/virtual_networks",
	"cloudflare_zero_trust_risk_behavior":                                "/accounts/{account_id}/zt_risk_scoring/behaviors",
	"cloudflare_zero_trust_risk_scoring_integration":                     "/accounts/{account_id}/zt_risk_scoring/integrations/{integration_id}",
	"cloudflare_turnstile_widget":                                        "/accounts/{account_id}/challenges/widgets",
	"cloudflare_hyperdrive_config":                                       "/accounts/{account_id}/hyperdrive/configs",
	"cloudflare_web_analytics_site":                                      "/accounts/{account_id}/rum/site_info/list",
	"cloudflare_web_analytics_rule":                                      "/accounts/{account_id}/rum/v2/{ruleset_id}/rules",
	"cloudflare_bot_management":                                          "/zones/{zone_id}/bot_management",
	"cloudflare_observatory_scheduled_test":                              "/zones/{zone_id}/speed_api/schedule/{url}",
	"cloudflare_hostname_tls_setting":                                    "/zones/{zone_id}/hostnames/settings/{setting_id}",
	"cloudflare_snippets":                                                "/zones/{zone_id}/snippets",
	"cloudflare_snippet_rules":                                           "/zones/{zone_id}/snippets/snippet_rules",
	"cloudflare_calls_sfu_app":                                           "/accounts/{account_id}/calls/apps",
	"cloudflare_calls_turn_app":                                          "/accounts/{account_id}/calls/turn_keys",
	"cloudflare_cloudforce_one_request_priority":                         "/accounts/{account_identifier}/cloudforce-one/requests/priority/{priority_identifer}",
	"cloudflare_cloudforce_one_request_asset":                            "/accounts/{account_identifier}/cloudforce-one/requests/{request_identifier}/asset/{asset_identifer}",
	"cloudflare_cloud_connector_rules":                                   "/zones/{zone_id}/cloud_connector/rules",
	"cloudflare_leaked_credential_check":                                 "/zones/{zone_id}/leaked-credential-checks",
	"cloudflare_leaked_credential_check_rule":                            "/zones/{zone_id}/leaked-credential-checks/detections",
	"cloudflare_content_scanning_expression":                             "/zones/{zone_id}/content-upload-scan/payloads",
}