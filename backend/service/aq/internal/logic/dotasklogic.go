package logic

import (
	"context"

	"backend/service/aq/internal/svc"
	"backend/service/aq/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DoTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDoTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DoTaskLogic {
	return &DoTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DoTaskLogic) DoTask(req *types.Req) (resp *types.Reply, err error) {
	// todo: add your logic here and delete this line

	return
}
// BoundingBox 定义了一个边界框，其中 (x1, y1) 是左上角坐标，(x2, y2) 是右下角坐标。
type BoundingBox [4]float64

// IoU 计算两个边界框的交并比。
func (bbox1 BoundingBox) IoU(bbox2 BoundingBox) float64 {
	// 计算交集的坐标
	x1 := max(bbox1[0], bbox2[0])
	y1 := max(bbox1[1], bbox2[1])
	x2 := min(bbox1[2], bbox2[2])
	y2 := min(bbox1[3], bbox2[3])

	// 计算交集的面积
	interArea := max(0, x2-x1) * max(0, y2-y1)

	// 计算两个边界框的面积
	bbox1Area := (bbox1[2]-bbox1[0]) * (bbox1[3]-bbox1[1])
	bbox2Area := (bbox2[2]-bbox2[0]) * (bbox2[3]-bbox2[1])

	// 计算并集的面积
	unionArea := bbox1Area + bbox2Area - interArea

	// 计算IoU
	iou := interArea / unionArea
	return iou
}
// max 返回两个数中的最大值。
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// min 返回两个数中的最小值。
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}