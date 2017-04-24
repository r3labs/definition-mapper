/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

// VirtualMachine ...
type VirtualMachine struct {
	Name                   string                 `json:"name" yaml:"name"`
	Count                  int                    `json:"count" yaml:"count"`
	Size                   string                 `json:"size" yaml:"size"`
	Image                  string                 `json:"image" yaml:"image"`
	Authentication         Authentication         `json:"authentication" yaml:"authentication"`
	StorageOSDisk          StorageOSDisk          `json:"storage_os_disk" yaml:"storage_os_disk"`
	OSProfile              OSProfile              `json:"os_profile" yaml:"os_profile"`
	OSProfileWindowsConfig OSProfileWindowsConfig `json:"os_profile_windows_config" yaml:"os_profile_windows_config"`
	NetworkInterfaces      []string               `json:"network_interfaces" yaml:"network_interfaces"`
	Plan                   struct {
		Name      string `json:"name" yaml:"name"`
		Publisher string `json:"publisher" yaml:"publisher"`
		Product   string `json:"product" yaml:"product"`
	} `json:"plan" yaml:"plan"`
	BootDiagnostics struct {
		Enabled    bool   `json:"enabled" yaml:"enabled"`
		StorageURI string `json:"storage_uri" yaml:"storage_uri"`
	} `json:"boot_diagnostics" yaml:"boot_diagnostics"`
	StorageImageReference struct {
		Publisher string `json:"publisher" yaml:"publisher"`
		Offer     string `json:"offer" yaml:"offer"`
		Sku       string `json:"sku" yaml:"sku"`
		Version   string `json:"version" yaml:"version"`
	} `json:"storage_image_reference" yaml:"storage_image_reference"`
	StorageDataDisk struct {
		Name         string `json:"name" yaml:"name"`
		VhdURI       string `json:"vhd_uri" yaml:"vhd_uri"`
		CreateOption string `json:"create_option" yaml:"create_option"`
		Caching      bool   `json:"caching" yaml:"caching"`
		ImageURI     string `json:"image_uri" yaml:"image_uri"`
		OSType       string `json:"os_type" yaml:"os_type"`
		DiskSizeGB   string `json:"disk_size_gb" yaml:"disk_size_gb"`
	} `json:"storage_data_disk" yaml:"storage_data_disk"`
	//DeleteOSDiskOnTermination    `json:"delete_os_disk_on_termination" yaml:"delete_os_disk_on_termination"`
	//DeleteDataDisksOnTermination `json:"delete_data_disks_on_termination" yaml:"delete_data_disks_on_termination"`
	LicenseType string            `json:"license_type" yaml:"license_type"`
	Tags        map[string]string `json:"tags" yaml:"tags"`
}

// Authentication ...
type Authentication struct {
	AdminUsername                 string            `json:"admin_username" yaml:"admin_username"`
	AdminPassword                 string            `json:"admin_password" yaml:"ssh_keys"`
	SSHKeys                       map[string]string `json:"ssh_keys" yaml:"ssh_keys"`
	DisablePasswordAuthentication bool              `json:"disable_password_authentication" yaml:"disable_password_authentication"`
}

// StorageOSDisk ...
type StorageOSDisk struct {
	Name         string `json:"name" yaml:"name"`
	VHDURI       string `json:"vhd_uri" yaml:"vhd_uri"`
	CreateOption string `json:"create_option" yaml:"create_option"`
	Caching      bool   `json:"caching" yaml:"caching"`
	ImageURI     string `json:"image_uri" yaml:"image_uri"`
	OSType       string `json:"os_type" yaml:"os_type"`
	DiskSizeGB   string `json:"disk_size_gb" yaml:"disk_size_gb"`
}

// OSProfile ...
type OSProfile struct {
	ComputerName string `json:"computer_name" yaml:"computer_name"`
	CustomData   string `json:"custom_data" yaml:"custom_data"`
}

// OSProfileWindowsConfig ...
type OSProfileWindowsConfig struct {
	ProvisionVMAgent         string  `json:"provision_vm_agent" yaml:"provision_vm_agent"`
	EnableAutomaticUpgrades  bool    `json:"enable_automatic_upgrades" yaml:"enable_automatic_upgrades"`
	WinRM                    []WinRM `json:"winrm" yaml:"winrm"`
	AdditionalUnattendConfig struct {
		Pass        string `json:"pass" yaml:"pass"`
		Component   string `json:"component" yaml:"component"`
		SettingName string `json:"setting_name" yaml:"setting_name"`
		Content     string `json:"content" yaml:"content"`
	} `json:"additional_unattend_config" yaml:"additional_unattend_config"`
}

// WinRM ...
type WinRM struct {
	Protocol       string `json:"protocol" yaml:"protocol"`
	CertificateURL string `json:"certificate_url" yaml:"certificate_url"`
}