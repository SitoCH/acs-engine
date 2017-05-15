package acsengine

import "fmt"

// AUTOGENERATED FILE - last generated 2017-04-06 11:25:28Z

const (
	// AzureProdFQDNFormat specifies the format for a prod dns name
	AzurePublicProdFQDNFormat = "%s.%s.cloudapp.azure.com"
	//AzureChinaProdFQDNFormat specify the endpoint of Azure China Cloud
	AzureChinaProdFQDNFormat = "%s.%s.cloudapp.chinacloudapi.cn"
)

// AzureLocations provides all azure regions in prod.
// Related powershell to refresh this list:
//   Get-AzureRmLocation | Select-Object -Property Location
var AzureLocations = []string{
	"australiaeast",
	"australiasoutheast",
	"brazilsouth",
	"canadacentral",
	"canadaeast",
	"centralindia",
	"centralus",
	"chinaeast",
	"chinanorth",
	"eastasia",
	"eastus",
	"eastus2",
	"japaneast",
	"japanwest",
	"koreacentral",
	"koreasouth",
	"northcentralus",
	"northeurope",
	"southcentralus",
	"southeastasia",
	"southindia",
	"uksouth",
	"ukwest",
	"westcentralus",
	"westeurope",
	"westindia",
	"westus",
	"westus2",
}

// FormatAzureProdFQDNs constructs all possible Azure prod fqdn
func FormatAzureProdFQDNs(fqdnPrefix string) []string {
	var fqdns []string
	for _, location := range AzureLocations {
		fqdns = append(fqdns, FormatAzureProdFQDN(fqdnPrefix, location))
	}
	return fqdns
}

// FormatAzureProdFQDN constructs an Azure prod fqdn
func FormatAzureProdFQDN(fqdnPrefix string, location string) string {
	FQDNFormat := AzurePublicProdFQDNFormat
	if location == "chinaeast" || location == "chinanorth" {
		FQDNFormat = AzureChinaProdFQDNFormat
	}
	return fmt.Sprintf(FQDNFormat, fqdnPrefix, location)
}

// GetDCOSMasterAllowedSizes returns the master allowed sizes
func GetDCOSMasterAllowedSizes() string {
	return `      "allowedValues": [
        "Standard_A10",
        "Standard_A11",
        "Standard_A2",
        "Standard_A3",
        "Standard_A4",
        "Standard_A5",
        "Standard_A6",
        "Standard_A7",
        "Standard_A8",
        "Standard_A9",
        "Standard_D11",
        "Standard_D11_v2",
        "Standard_D11_v2_Promo",
        "Standard_D12",
        "Standard_D12_v2",
        "Standard_D12_v2_Promo",
        "Standard_D13",
        "Standard_D13_v2",
        "Standard_D13_v2_Promo",
        "Standard_D14",
        "Standard_D14_v2",
        "Standard_D14_v2_Promo",
        "Standard_D15_v2",
        "Standard_D2",
        "Standard_D2_v2",
        "Standard_D2_v2_Promo",
        "Standard_D3",
        "Standard_D3_v2",
        "Standard_D3_v2_Promo",
        "Standard_D4",
        "Standard_D4_v2",
        "Standard_D4_v2_Promo",
        "Standard_D5_v2",
        "Standard_D5_v2_Promo",
        "Standard_DS13",
        "Standard_DS13_v2",
        "Standard_DS13_v2_Promo",
        "Standard_DS14",
        "Standard_DS14_v2",
        "Standard_DS14_v2_Promo",
        "Standard_DS15_v2",
        "Standard_DS5_v2",
        "Standard_DS5_v2_Promo",
        "Standard_F16",
        "Standard_F8",
        "Standard_G1",
        "Standard_G2",
        "Standard_G3",
        "Standard_G4",
        "Standard_G5",
        "Standard_GS2",
        "Standard_GS3",
        "Standard_GS4",
        "Standard_GS5",
        "Standard_H16",
        "Standard_H16m",
        "Standard_H16mr",
        "Standard_H16r",
        "Standard_H8",
        "Standard_H8m",
        "Standard_L16s",
        "Standard_L32s",
        "Standard_L4s",
        "Standard_L8s",
        "Standard_NC12",
        "Standard_NC24",
        "Standard_NC24r",
        "Standard_NC6",
        "Standard_NV12",
        "Standard_NV24",
        "Standard_NV6"
     ],
`
}

// GetMasterAgentAllowedSizes returns the agent allowed sizes
func GetMasterAgentAllowedSizes() string {
	return `      "allowedValues": [
        "Standard_A1",
        "Standard_A1_v2",
        "Standard_A10",
        "Standard_A11",
        "Standard_A2",
        "Standard_A2_v2",
        "Standard_A2m_v2",
        "Standard_A3",
        "Standard_A4",
        "Standard_A4_v2",
        "Standard_A4m_v2",
        "Standard_A5",
        "Standard_A6",
        "Standard_A7",
        "Standard_A8",
        "Standard_A8_v2",
        "Standard_A8m_v2",
        "Standard_A9",
        "Standard_D11",
        "Standard_D11_v2",
        "Standard_D11_v2_Promo",
        "Standard_D12",
        "Standard_D12_v2",
        "Standard_D12_v2_Promo",
        "Standard_D13",
        "Standard_D13_v2",
        "Standard_D13_v2_Promo",
        "Standard_D14",
        "Standard_D14_v2",
        "Standard_D14_v2_Promo",
        "Standard_D15_v2",
        "Standard_D2",
        "Standard_D2_v2",
        "Standard_D2_v2_Promo",
        "Standard_D3",
        "Standard_D3_v2",
        "Standard_D3_v2_Promo",
        "Standard_D4",
        "Standard_D4_v2",
        "Standard_D4_v2_Promo",
        "Standard_D5_v2",
        "Standard_D5_v2_Promo",
        "Standard_DS11",
        "Standard_DS11_v2",
        "Standard_DS11_v2_Promo",
        "Standard_DS12",
        "Standard_DS12_v2",
        "Standard_DS12_v2_Promo",
        "Standard_DS13",
        "Standard_DS13_v2",
        "Standard_DS13_v2_Promo",
        "Standard_DS14",
        "Standard_DS14_v2",
        "Standard_DS14_v2_Promo",
        "Standard_DS15_v2",
        "Standard_DS2",
        "Standard_DS2_v2",
        "Standard_DS2_v2_Promo",
        "Standard_DS3",
        "Standard_DS3_v2",
        "Standard_DS3_v2_Promo",
        "Standard_DS4",
        "Standard_DS4_v2",
        "Standard_DS4_v2_Promo",
        "Standard_DS5_v2",
        "Standard_DS5_v2_Promo",
        "Standard_F1",
        "Standard_F16",
        "Standard_F16s",
        "Standard_F2",
        "Standard_F2s",
        "Standard_F4",
        "Standard_F4s",
        "Standard_F8",
        "Standard_F8s",
        "Standard_G1",
        "Standard_G2",
        "Standard_G3",
        "Standard_G4",
        "Standard_G5",
        "Standard_GS1",
        "Standard_GS2",
        "Standard_GS3",
        "Standard_GS4",
        "Standard_GS5",
        "Standard_H16",
        "Standard_H16m",
        "Standard_H16mr",
        "Standard_H16r",
        "Standard_H8",
        "Standard_H8m",
        "Standard_L16s",
        "Standard_L32s",
        "Standard_L4s",
        "Standard_L8s",
        "Standard_NC12",
        "Standard_NC24",
        "Standard_NC24r",
        "Standard_NC6",
        "Standard_NV12",
        "Standard_NV24",
        "Standard_NV6"
     ],
`
}

// GetKubernetesAgentAllowedSizes returns the allowed sizes for Kubernetes agents
func GetKubernetesAgentAllowedSizes() string {
	return `      "allowedValues": [
        "Standard_A0",
        "Standard_A1",
        "Standard_A1_v2",
        "Standard_A10",
        "Standard_A11",
        "Standard_A2",
        "Standard_A2_v2",
        "Standard_A2m_v2",
        "Standard_A3",
        "Standard_A4",
        "Standard_A4_v2",
        "Standard_A4m_v2",
        "Standard_A5",
        "Standard_A6",
        "Standard_A7",
        "Standard_A8",
        "Standard_A8_v2",
        "Standard_A8m_v2",
        "Standard_A9",
        "Standard_D1",
        "Standard_D1_v2",
        "Standard_D11",
        "Standard_D11_v2",
        "Standard_D11_v2_Promo",
        "Standard_D12",
        "Standard_D12_v2",
        "Standard_D12_v2_Promo",
        "Standard_D13",
        "Standard_D13_v2",
        "Standard_D13_v2_Promo",
        "Standard_D14",
        "Standard_D14_v2",
        "Standard_D14_v2_Promo",
        "Standard_D15_v2",
        "Standard_D2",
        "Standard_D2_v2",
        "Standard_D2_v2_Promo",
        "Standard_D3",
        "Standard_D3_v2",
        "Standard_D3_v2_Promo",
        "Standard_D4",
        "Standard_D4_v2",
        "Standard_D4_v2_Promo",
        "Standard_D5_v2",
        "Standard_D5_v2_Promo",
        "Standard_DS1",
        "Standard_DS1_v2",
        "Standard_DS11",
        "Standard_DS11_v2",
        "Standard_DS11_v2_Promo",
        "Standard_DS12",
        "Standard_DS12_v2",
        "Standard_DS12_v2_Promo",
        "Standard_DS13",
        "Standard_DS13_v2",
        "Standard_DS13_v2_Promo",
        "Standard_DS14",
        "Standard_DS14_v2",
        "Standard_DS14_v2_Promo",
        "Standard_DS15_v2",
        "Standard_DS2",
        "Standard_DS2_v2",
        "Standard_DS2_v2_Promo",
        "Standard_DS3",
        "Standard_DS3_v2",
        "Standard_DS3_v2_Promo",
        "Standard_DS4",
        "Standard_DS4_v2",
        "Standard_DS4_v2_Promo",
        "Standard_DS5_v2",
        "Standard_DS5_v2_Promo",
        "Standard_F1",
        "Standard_F16",
        "Standard_F16s",
        "Standard_F1s",
        "Standard_F2",
        "Standard_F2s",
        "Standard_F4",
        "Standard_F4s",
        "Standard_F8",
        "Standard_F8s",
        "Standard_G1",
        "Standard_G2",
        "Standard_G3",
        "Standard_G4",
        "Standard_G5",
        "Standard_GS1",
        "Standard_GS2",
        "Standard_GS3",
        "Standard_GS4",
        "Standard_GS5",
        "Standard_H16",
        "Standard_H16m",
        "Standard_H16mr",
        "Standard_H16r",
        "Standard_H8",
        "Standard_H8m",
        "Standard_L16s",
        "Standard_L32s",
        "Standard_L4s",
        "Standard_L8s",
        "Standard_NC12",
        "Standard_NC24",
        "Standard_NC24r",
        "Standard_NC6",
        "Standard_NV12",
        "Standard_NV24",
        "Standard_NV6"
     ],
`
}

// GetSizeMap returns the size / storage map
func GetSizeMap() string {
	return `    "vmSizesMap": {
      "Standard_A1": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A1_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A10": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A11": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A2_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A2m_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A3": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A4": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A4_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A4m_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A5": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A6": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A7": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A8": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A8_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A8m_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A9": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D11": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D11_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D11_v2_Promo": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D12": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D12_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D12_v2_Promo": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D13": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D13_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D13_v2_Promo": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D14": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D14_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D14_v2_Promo": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D15_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D2_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D2_v2_Promo": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D3": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D3_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D3_v2_Promo": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D4": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D4_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D4_v2_Promo": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D5_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D5_v2_Promo": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_DS11": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS11_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS11_v2_Promo": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS12": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS12_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS12_v2_Promo": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS13": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS13_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS13_v2_Promo": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS14": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS14_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS14_v2_Promo": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS15_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS1_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS2_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS2_v2_Promo": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS3": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS3_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS3_v2_Promo": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS4": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS4_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS4_v2_Promo": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS5_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS5_v2_Promo": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_F1": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_F16": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_F16s": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_F2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_F2s": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_F4": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_F4s": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_F8": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_F8s": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_G1": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_G2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_G3": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_G4": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_G5": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_GS1": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_GS2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_GS3": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_GS4": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_GS5": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_H16": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_H16m": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_H16mr": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_H16r": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_H8": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_H8m": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_L16s": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_L32s": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_L4s": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_L8s": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_NC12": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_NC24": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_NC24r": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_NC6": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_NV12": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_NV24": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_NV6": {
        "storageAccountType": "Standard_LRS"
      }
    }	
`
}

// GetClassicAllowedSizes returns the classic allowed sizes
func GetClassicAllowedSizes() string {
	return `      "allowedValues": [
        "Standard_A0",
        "Standard_A1",
        "Standard_A1_v2",
        "Standard_A10",
        "Standard_A11",
        "Standard_A2",
        "Standard_A2_v2",
        "Standard_A2m_v2",
        "Standard_A3",
        "Standard_A4",
        "Standard_A4_v2",
        "Standard_A4m_v2",
        "Standard_A5",
        "Standard_A6",
        "Standard_A7",
        "Standard_A8",
        "Standard_A8_v2",
        "Standard_A8m_v2",
        "Standard_A9",
        "Standard_D1",
        "Standard_D1_v2",
        "Standard_D11",
        "Standard_D11_v2",
        "Standard_D11_v2_Promo",
        "Standard_D12",
        "Standard_D12_v2",
        "Standard_D12_v2_Promo",
        "Standard_D13",
        "Standard_D13_v2",
        "Standard_D13_v2_Promo",
        "Standard_D14",
        "Standard_D14_v2",
        "Standard_D14_v2_Promo",
        "Standard_D15_v2",
        "Standard_D2",
        "Standard_D2_v2",
        "Standard_D2_v2_Promo",
        "Standard_D3",
        "Standard_D3_v2",
        "Standard_D3_v2_Promo",
        "Standard_D4",
        "Standard_D4_v2",
        "Standard_D4_v2_Promo",
        "Standard_D5_v2",
        "Standard_D5_v2_Promo",
        "Standard_DS1",
        "Standard_DS1_v2",
        "Standard_DS11",
        "Standard_DS11_v2",
        "Standard_DS11_v2_Promo",
        "Standard_DS12",
        "Standard_DS12_v2",
        "Standard_DS12_v2_Promo",
        "Standard_DS13",
        "Standard_DS13_v2",
        "Standard_DS13_v2_Promo",
        "Standard_DS14",
        "Standard_DS14_v2",
        "Standard_DS14_v2_Promo",
        "Standard_DS15_v2",
        "Standard_DS2",
        "Standard_DS2_v2",
        "Standard_DS2_v2_Promo",
        "Standard_DS3",
        "Standard_DS3_v2",
        "Standard_DS3_v2_Promo",
        "Standard_DS4",
        "Standard_DS4_v2",
        "Standard_DS4_v2_Promo",
        "Standard_DS5_v2",
        "Standard_DS5_v2_Promo",
        "Standard_F1",
        "Standard_F16",
        "Standard_F16s",
        "Standard_F1s",
        "Standard_F2",
        "Standard_F2s",
        "Standard_F4",
        "Standard_F4s",
        "Standard_F8",
        "Standard_F8s",
        "Standard_G1",
        "Standard_G2",
        "Standard_G3",
        "Standard_G4",
        "Standard_G5",
        "Standard_GS1",
        "Standard_GS2",
        "Standard_GS3",
        "Standard_GS4",
        "Standard_GS5",
        "Standard_H16",
        "Standard_H16m",
        "Standard_H16mr",
        "Standard_H16r",
        "Standard_H8",
        "Standard_H8m",
        "Standard_L16s",
        "Standard_L32s",
        "Standard_L4s",
        "Standard_L8s",
        "Standard_NC12",
        "Standard_NC24",
        "Standard_NC24r",
        "Standard_NC6",
        "Standard_NV12",
        "Standard_NV24",
        "Standard_NV6"
     ],
`
}

// GetClassicSizeMap returns the size / storage map
func GetClassicSizeMap() string {
	return `    "vmSizesMap": {
      "Standard_A0": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A1": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A1_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A10": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A11": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A2_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A2m_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A3": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A4": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A4_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A4m_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A5": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A6": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A7": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A8": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A8_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A8m_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_A9": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D1": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D1_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D11": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D11_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D11_v2_Promo": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D12": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D12_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D12_v2_Promo": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D13": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D13_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D13_v2_Promo": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D14": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D14_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D14_v2_Promo": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D15_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D2_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D2_v2_Promo": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D3": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D3_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D3_v2_Promo": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D4": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D4_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D4_v2_Promo": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D5_v2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_D5_v2_Promo": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_DS1": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS1_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS11": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS11_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS11_v2_Promo": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS12": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS12_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS12_v2_Promo": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS13": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS13_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS13_v2_Promo": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS14": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS14_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS14_v2_Promo": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS15_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS2_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS2_v2_Promo": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS3": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS3_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS3_v2_Promo": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS4": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS4_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS4_v2_Promo": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS5_v2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_DS5_v2_Promo": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_F1": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_F16": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_F16s": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_F1s": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_F2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_F2s": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_F4": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_F4s": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_F8": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_F8s": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_G1": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_G2": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_G3": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_G4": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_G5": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_GS1": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_GS2": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_GS3": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_GS4": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_GS5": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_H16": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_H16m": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_H16mr": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_H16r": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_H8": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_H8m": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_L16s": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_L32s": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_L4s": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_L8s": {
        "storageAccountType": "Premium_LRS"
      },
      "Standard_NC12": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_NC24": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_NC24r": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_NC6": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_NV12": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_NV24": {
        "storageAccountType": "Standard_LRS"
      },
      "Standard_NV6": {
        "storageAccountType": "Standard_LRS"
      }
    }	
`
}
