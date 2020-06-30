package crawl

import (
	"crawl-data/domain"
	"crawl-data/response"
	"crawl-data/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func StartCrawlerChotot(s *service.Service) error {
	//maximum data is 50
	limit := 50
	page := 0
	for page = 0; page < 10; page++ {
		url := fmt.Sprintf("https://gateway.chotot.com/v1/public/ad-listing?limit=%d&o=%d&page=%d", limit, page*limit, page)
		s.Logger.Info(url)
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		var obj response.AdListing
		err = json.Unmarshal(body, &obj)
		if err != nil {
			return err
		}
		for _, item := range obj.Ads {
			urlProfile := fmt.Sprintf("https://gateway.chotot.com/v1/public/profile/%s", item.AccountOid)
			resp, err := http.Get(urlProfile)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)

			var profile response.Profile
			err = json.Unmarshal(body, &profile)
			if err != nil {
				return err
			}
			s.Logger.Info(item.AccountOid, profile.Phone, profile.Address, profile.FullName, profile.Gender)
			user := domain.User{}
			user.Address = profile.Address
			user.Gender = profile.Gender
			user.Phone = profile.Phone
			user.FullName = profile.FullName
			user.ProfileImage = profile.Avatar
			user.FacebookID = profile.FacebookID
			user.Link = urlProfile
			user.Source = "chotot"
			user.Tag = item.CategoryName
			user.AccountID = item.AccountOid

			s.DB.Table("user").Save(&user)
		}
	}

	// var count int64
	// s.DB.Table("image_front").
	// 	Where("url is NULL").
	// 	Count(&count)

	// s.Logger.Info("Count ", count)

	// for count >= 0 {
	// 	var imageFronts []domain.ImageFront
	// 	err := s.DB.Table("image_front").
	// 		Where("url is NULL").
	// 		Limit(20).
	// 		Find(&imageFronts).Error
	// 	if err != nil {
	// 		return err
	// 	}

	// 	for _, f := range imageFronts {
	// 		s.Logger.Infof("Excute:", f.ImageID)
	// 		dateFormat := f.CreatedTime.Format("2006/01/02")
	// 		url := fmt.Sprintf("%s/%s/front.jpg", dateFormat, f.RequestID)
	// 		f.URL = url

	// 		s.DB.Debug().Table("image_front").
	// 			Where("image_id = ?", f.ImageID).
	// 			Updates(map[string]interface{}{
	// 				"url": f.URL,
	// 			})
	// 	}

	// 	count -= 20
	// }
	return nil
}
