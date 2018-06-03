package marketsmith

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	regexInfoCell = `\s*<div\s*class="cell">\s*<div\s*class="key">([^<]+)<\/div>\s*<div\s*class="value">([^<]+)<\/div>\s*<\/div>\s*`

	regexIndustryGroup = `\s*<div\s*class="companySymbol">\s*<span\s*class="companyInfCoName">[^<]+<\/span>\(\w+\)\s*\S+\s*([^<]+)<\/div>\s*`

	regexOptions = `\s*<div\s*class="Options">Options\s*([^<]+)<\/div>\s*`

	regexFloatShare = `\s*<div\s*class="companyInfoCenter"\s*[^>]+>\s*<div\s*class="companyInfoLable">\s*<div\s*class="cell">([^<]+)<\/div>\s*<div\s*class="cell">([^<]+)<\/div>\s*<div\s*class="cell">([^<]+)<\/div>\s*<\/div>\s*<div\s*class="companyInfoVlaue">\s*<div\s*class="cell">([^<]+)<\/div>\s*<div\s*class="cell">([^<]+)<\/div>\s*<div\s*class="cell">([^<]+)<\/div>\s*<\/div>\s*<\/div>\s*`

	regexQuarterlyResults = `\s*<div\s*class="cell"\s*style="display:\s*block;">\s*<div\s*class="quarterly\s*\w*">[^<]+<span\s*[^>]+>\s*[^<]*\s*<\/span>\s*<\/div>\s*<div\s*class="eps">\s*([^<]+)\s*<\/div>\s*<div\s*class="epsChg">([^<]+)<\/div>\s*<div\s*class="sales">([^<]+)<\/div>\s*<div\s*class="salesChg\s*\w+">([^<]+)<\/div>\s*<\/div>\s*`

	regexFunds = `\s*<div\s*class="cell">\s*<span>[^<]+<\/span>\s*<span>(\d+)<\/span>\s*<\/div>\s*`
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
