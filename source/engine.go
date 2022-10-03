package main

import (
	"app/aliclient"
	"app/config"
	"app/utils"
	"log"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

// run 主运行函数
func run(cfg *config.Config, client *aliclient.Client) {
	domains, err := cfg.GetDoMain()
	if err != nil {
		log.Panic(err)
	}
	var lastIP string
	func() {
		log.Printf("Info: ========================= cron joy begin!======================")
		tempIp, err := utils.Getnetip()
		if err != nil {
			return
		}
		if tempIp == lastIP {
			log.Printf("Info: ip no have change = %s", lastIP)
			return
		}
		log.Printf("Info: begin Update RR To IP:%s", tempIp)
		for _, domain := range domains {
			if len(domain.Host) == 0 {
				log.Printf("Info: domain:%s no have host", domain.Name)
				continue
			}
			for _, host := range domain.Host {
				if host == "" {
					continue
				}
				req := alidns.CreateDescribeDomainRecordsRequest()
				req.DomainName = domain.Name
				req.SearchMode = "EXACT"
				req.KeyWord = host
				response, err := client.DescribeDomainRecords(req)
				if err != nil {
					log.Printf("Error: client.DescribeDomainRecords domain:%s error:%s", domain.Name, err.Error())
					continue
				}
				if len(response.DomainRecords.Record) == 0 {
					// Add RR
					req := alidns.CreateAddDomainRecordRequest()
					req.Value = tempIp
					req.Type = "A"
					req.RR = host
					req.DomainName = domain.Name
					resp, err := client.AddDomainRecord(req)
					if err != nil {
						log.Printf("Error: client.AddDomainRecord domain:%s host:%s add fail:%s", domain.Name, host, err.Error())
						return
					}
					if resp.IsSuccess() {
						log.Printf("Info: domain:%s host:%s add success", domain.Name, host)
						continue
					} else {
						log.Printf("Error: client.AddDomainRecord domain:%s error:%s", domain.Name, resp.GetHttpContentString())
						continue
					}
				} else {
					// Update
					if response.DomainRecords.Record[0].Value == tempIp {
						log.Printf("Info: domain:%s rr:%s no hange change", domain.Name, host)
						continue
					}
					req := alidns.CreateUpdateDomainRecordRequest()
					req.RecordId = response.DomainRecords.Record[0].RecordId
					req.RR = response.DomainRecords.Record[0].RR
					req.Type = response.DomainRecords.Record[0].Type
					req.Value = tempIp
					resp, err := client.UpdateDomainRecord(req)
					if err != nil {
						log.Printf("Error: client.UpdateDomainRecord domain:%s host:%s Update Fail:%s", domain.Name, host, err.Error())
						continue
					}
					if resp.IsSuccess() {
						log.Printf("Info: domain:%s host:%s update success", domain.Name, host)
						continue
					} else {
						log.Printf("Error: client.UpdateDomainRecord domain:%s error:%s", domain.Name, resp.GetHttpContentString())
						continue
					}
				}
			}

		}
		log.Printf("Info: ========================= cron joy done!======================")
		lastIP = tempIp
	}()
}
