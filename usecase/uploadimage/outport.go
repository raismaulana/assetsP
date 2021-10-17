package uploadimage

import "github.com/raismaulana/assetsP/domain/service"

// Outport of UploadImage
type Outport interface {
	service.GenerateRandomStringService
	service.GetBaseURLRepo
}
