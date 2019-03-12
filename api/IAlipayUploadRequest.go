package api

import (
  "github.com/solarhell/antsdk/utils"
)

type IAlipayUploadRequest interface {
  IAlipayRequest
  GetFileParams() map[string]*utils.FileItem
}
