package service

import (
	"HH_ADMIN/api/v1"
	"HH_ADMIN/internal/consts"
	"HH_ADMIN/internal/dao"
	"HH_ADMIN/internal/model"
	"HH_ADMIN/internal/model/entity"
	"HH_ADMIN/util"
	"HH_ADMIN/utility/rr"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"math"
	"strings"
)

type sFile struct{}

func File() *sFile {
	return &sFile{}
}

func (s *sFile) FileUpload(ctx context.Context, req *api.FileReq) (res *rr.CommonRes, err error) {
	//userId := ctx.Value(consts.CTXUSERID)
	//phone := ctx.Value(consts.CTXPHONE)
	userName := ctx.Value(consts.CTXUSERNAME)

	fileLocation := g.Cfg().MustGet(ctx, "file_upload.location").String()
	urlPrefix := g.Cfg().MustGet(ctx, "file_upload.url_prefix").String()
	fileSize := g.Cfg().MustGet(ctx, "file_upload.size").Int64()
	fileType := g.Cfg().MustGet(ctx, "file_upload.file_type").Array()

	err = util.PathExistOrCreat(fileLocation)
	if err != nil {
		g.Log().Line().Errorf(ctx, "FileUpload err:%s", err.Error())
		return rr.FailedWithMessage("文件上传失败"), err
	}

	uploadRes := make([]model.FileUploadRes, 0)

	// 校验
	for _, uploadFile := range req.Files {
		if util.JudgeFileType(uploadFile.Filename, fileType) {
			return rr.FailedWithMessage("文件 " + uploadFile.Filename + " 类型不符合规范，请重新上传"), err
		}
		if uploadFile.Size > fileSize {
			return rr.FailedWithMessage("文件 " + uploadFile.Filename + " 大小不符合规范，请重新上传"), err
		}
	}

	// 存储
	for _, uploadFile := range req.Files {
		oldFileName := uploadFile.Filename                                                                // 原文件名
		format := strings.Split(uploadFile.Filename, ".")[len(strings.Split(uploadFile.Filename, "."))-1] // 文件后缀名
		randomName := util.GetRandomStringNameNew("", false)                                              // 随机文件名
		newFileName := randomName + "." + format                                                          // 新文件名

		uploadFile.Filename = newFileName
		_, err := uploadFile.Save(fileLocation)
		if err != nil {
			g.Log().Line().Errorf(ctx, "FileUpload err:%s", err.Error())
			return rr.FailedWithMessage("文件上传失败"), err
		}
		// 入库
		id, err := dao.SysFile.Ctx(ctx).InsertAndGetId(entity.SysFile{
			FileOldName:  oldFileName,
			FileName:     newFileName,
			FileType:     format,
			FileSize:     gconv.Int(math.Ceil(float64(uploadFile.Size/1048576))) + 1,
			FileLocation: fileLocation + newFileName,
			FileUrl:      urlPrefix + newFileName,
			CreateTime:   gtime.Now(),
			CreateBy:     gconv.String(userName),
		})
		if err != nil {
			g.Log().Line().Errorf(ctx, "FileUpload err:%s", err.Error())
			return rr.FailedWithMessage("文件上传失败"), err
		}
		uploadRes = append(uploadRes, model.FileUploadRes{
			Id:       id,
			FileName: oldFileName,
			FileUrl:  urlPrefix + newFileName,
		})
	}
	return rr.SuccessWithMessageAndData("上传成功", uploadRes), err
}
